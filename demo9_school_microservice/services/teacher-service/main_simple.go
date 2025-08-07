package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Teacher struct {
	ID           string   `json:"id"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Department   string   `json:"department"`
	Subjects     []string `json:"subjects"`
	Qualification string  `json:"qualification"`
	Experience   int      `json:"experience"`
	Status       string   `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Type         string   `json:"type"`
}

type CouchbaseClient struct {
	Cluster    *gocb.Cluster
	Bucket     *gocb.Bucket
	Collection *gocb.Collection
}

var db *CouchbaseClient

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func connectToCouchbase() (*CouchbaseClient, error) {
	host := getEnv("COUCHBASE_HOST", "localhost")
	username := getEnv("COUCHBASE_USERNAME", "Administrator")
	password := getEnv("COUCHBASE_PASSWORD", "password123")
	bucketName := getEnv("COUCHBASE_BUCKET", "schoolmgmt")
	
	connectionString := fmt.Sprintf("couchbase://%s", host)
	
	logrus.Infof("🔌 Attempting to connect to Couchbase at: %s", connectionString)
	logrus.Infof("📝 Using credentials - User: %s, Bucket: %s", username, bucketName)
	
	// Connect to cluster
	cluster, err := gocb.Connect(connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		TimeoutsConfig: gocb.TimeoutsConfig{
			ConnectTimeout:    10 * time.Second,
			KVTimeout:         5 * time.Second,
			QueryTimeout:      15 * time.Second,
			AnalyticsTimeout:  15 * time.Second,
			SearchTimeout:     15 * time.Second,
			ManagementTimeout: 15 * time.Second,
		},
	})
	if err != nil {
		logrus.Errorf("❌ Failed to connect to Couchbase cluster: %v", err)
		return nil, fmt.Errorf("failed to connect to Couchbase cluster: %w", err)
	}
	
	logrus.Infof("✅ Connected to Couchbase cluster, waiting for readiness...")

	// Wait for cluster to be ready
	err = cluster.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		logrus.Errorf("❌ Cluster not ready: %v", err)
		return nil, fmt.Errorf("cluster not ready: %w", err)
	}
	
	logrus.Infof("✅ Cluster is ready, accessing bucket: %s", bucketName)

	// Get bucket
	bucket := cluster.Bucket(bucketName)
	
	// Wait for bucket to be ready
	err = bucket.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		logrus.Warnf("⚠️ Bucket %s not ready: %v", bucketName, err)
		// Continue anyway - bucket might be created later
	} else {
		logrus.Infof("✅ Bucket %s is ready!", bucketName)
	}

	// Get default collection
	collection := bucket.DefaultCollection()
	logrus.Infof("✅ Connected to default collection")

	logrus.Infof("🎉 Successfully connected to Couchbase bucket: %s", bucketName)

	return &CouchbaseClient{
		Cluster:    cluster,
		Bucket:     bucket,
		Collection: collection,
	}, nil
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"service":  "teacher-service",
		"status":   "healthy",
		"database": "couchbase-connected",
	})
}

func createTeacher(c *gin.Context) {
	var req struct {
		FirstName    string   `json:"firstName" binding:"required"`
		LastName     string   `json:"lastName" binding:"required"`
		Email        string   `json:"email" binding:"required"`
		Phone        string   `json:"phone"`
		Department   string   `json:"department" binding:"required"`
		Subjects     []string `json:"subjects"`
		Qualification string  `json:"qualification"`
		Experience   int      `json:"experience"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher := Teacher{
		ID:           uuid.New().String(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Phone:        req.Phone,
		Department:   req.Department,
		Subjects:     req.Subjects,
		Qualification: req.Qualification,
		Experience:   req.Experience,
		Status:       "active",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Type:         "teacher",
	}

	documentKey := fmt.Sprintf("teacher::%s", teacher.ID)
	logrus.Infof("💾 Creating teacher document with key: %s", documentKey)

	_, err := db.Collection.Insert(documentKey, teacher, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to create teacher: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher"})
		return
	}

	logrus.Infof("✅ Successfully created teacher: %s %s", teacher.FirstName, teacher.LastName)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    teacher,
	})
}

func getTeachers(c *gin.Context) {
	query := `SELECT t.* FROM schoolmgmt t WHERE t.type = "teacher"`
	
	result, err := db.Cluster.Query(query, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to query teachers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query teachers"})
		return
	}
	defer result.Close()

	var teachers []Teacher
	for result.Next() {
		var teacher Teacher
		if err := result.Row(&teacher); err != nil {
			logrus.Errorf("❌ Failed to decode teacher: %v", err)
			continue
		}
		teachers = append(teachers, teacher)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"count":    len(teachers),
			"teachers": teachers,
		},
	})
}

func getTeacher(c *gin.Context) {
	id := c.Param("id")
	documentKey := fmt.Sprintf("teacher::%s", id)

	result, err := db.Collection.Get(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
			return
		}
		logrus.Errorf("❌ Failed to get teacher: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get teacher"})
		return
	}

	var teacher Teacher
	if err := result.Content(&teacher); err != nil {
		logrus.Errorf("❌ Failed to decode teacher: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teacher,
	})
}

func deleteTeacher(c *gin.Context) {
	id := c.Param("id")
	documentKey := fmt.Sprintf("teacher::%s", id)

	_, err := db.Collection.Remove(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
			return
		}
		logrus.Errorf("❌ Failed to delete teacher: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}

	logrus.Infof("✅ Successfully deleted teacher with ID: %s", id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Teacher deleted successfully",
	})
}

func main() {
	// Connect to Couchbase
	var err error
	db, err = connectToCouchbase()
	if err != nil {
		log.Fatalf("Failed to connect to Couchbase: %v", err)
	}
	defer db.Cluster.Close(nil)

	// Setup Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", healthCheck)

	// Teacher endpoints
	router.POST("/api/v1/teachers", createTeacher)
	router.GET("/api/v1/teachers", getTeachers)
	router.GET("/api/v1/teachers/:id", getTeacher)
	router.DELETE("/api/v1/teachers/:id", deleteTeacher)

	// Start server
	log.Println("Starting Teacher Service on port 8082 with Couchbase integration")
	router.Run(":8082")
}

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

type Student struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Grade       string    `json:"grade"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Type        string    `json:"type"`
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

func main() {
	// Setup logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	
	// Connect to Couchbase
	var err error
	db, err = connectToCouchbase()
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Cluster.Close(nil)

	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "student-service",
			"database": "couchbase-connected",
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		v1.POST("/students", createStudent)
		v1.GET("/students", listStudents)
		v1.GET("/students/:id", getStudent)
		v1.PUT("/students/:id", updateStudent)
		v1.DELETE("/students/:id", deleteStudent)
	}

	log.Println("Starting Student Service on port 8081 with Couchbase integration")
	router.Run(":8081")
}

func createStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		logrus.Errorf("❌ Invalid JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set student metadata
	student.ID = uuid.New().String()
	student.Status = "active"
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	student.Type = "student"

	documentKey := fmt.Sprintf("student::%s", student.ID)
	
	logrus.Infof("💾 Inserting student document with key: %s", documentKey)
	_, err := db.Collection.Insert(documentKey, student, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to create student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	logrus.Infof("✅ Successfully created student with ID: %s", student.ID)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    student,
		"message": "Student created successfully",
	})
}

func getStudent(c *gin.Context) {
	id := c.Param("id")
	documentKey := fmt.Sprintf("student::%s", id)
	
	logrus.Infof("🔍 Retrieving student with key: %s", documentKey)
	result, err := db.Collection.Get(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			logrus.Warnf("⚠️ Student not found with ID: %s", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		logrus.Errorf("❌ Failed to get student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get student"})
		return
	}

	var student Student
	err = result.Content(&student)
	if err != nil {
		logrus.Errorf("❌ Failed to decode student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode student"})
		return
	}

	logrus.Infof("✅ Successfully retrieved student: %s %s", student.FirstName, student.LastName)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
	})
}

func updateStudent(c *gin.Context) {
	id := c.Param("id")
	documentKey := fmt.Sprintf("student::%s", id)
	
	// First get the existing student
	result, err := db.Collection.Get(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get student"})
		return
	}

	var student Student
	err = result.Content(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode student"})
		return
	}

	// Update with new data
	var updates Student
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if updates.FirstName != "" {
		student.FirstName = updates.FirstName
	}
	if updates.LastName != "" {
		student.LastName = updates.LastName
	}
	if updates.Email != "" {
		student.Email = updates.Email
	}
	if updates.Grade != "" {
		student.Grade = updates.Grade
	}
	
	student.UpdatedAt = time.Now()

	logrus.Infof("📝 Updating student with key: %s", documentKey)
	_, err = db.Collection.Replace(documentKey, student, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to update student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	logrus.Infof("✅ Successfully updated student: %s %s", student.FirstName, student.LastName)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
		"message": "Student updated successfully",
	})
}

func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	documentKey := fmt.Sprintf("student::%s", id)
	
	logrus.Infof("🗑️ Deleting student with key: %s", documentKey)
	_, err := db.Collection.Remove(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		logrus.Errorf("❌ Failed to delete student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	logrus.Infof("✅ Successfully deleted student with ID: %s", id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Student deleted successfully",
	})
}

func listStudents(c *gin.Context) {
	logrus.Infof("📋 Retrieving all students")
	
	// Use N1QL query to get all students
	query := "SELECT * FROM `schoolmgmt` WHERE type = 'student'"
	
	logrus.Infof("🔍 Executing query: %s", query)
	result, err := db.Cluster.Query(query, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to query students: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query students"})
		return
	}

	var students []Student
	for result.Next() {
		var row struct {
			Schoolmgmt Student `json:"schoolmgmt"`
		}
		err := result.Row(&row)
		if err != nil {
			logrus.Errorf("❌ Failed to decode student row: %v", err)
			continue
		}
		students = append(students, row.Schoolmgmt)
	}

	if err := result.Err(); err != nil {
		logrus.Errorf("❌ Query execution error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query execution failed"})
		return
	}

	logrus.Infof("✅ Successfully retrieved %d students", len(students))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"students": students,
			"count":    len(students),
		},
	})
}

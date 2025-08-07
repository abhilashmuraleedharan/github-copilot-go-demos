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

type Subject struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Credits     int       `json:"credits"`
	Description string    `json:"description"`
	TeacherID   string    `json:"teacherId"`
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

	logrus.Infof(" Attempting to connect to Couchbase at: %s", connectionString)

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
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cluster: %w", err)
	}

	// Open bucket
	bucket := cluster.Bucket(bucketName)

	// Wait for bucket to be ready
	err = bucket.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("bucket not ready: %w", err)
	}

	// Get collection
	collection := bucket.DefaultCollection()

	return &CouchbaseClient{
		Cluster:    cluster,
		Bucket:     bucket,
		Collection: collection,
	}, nil
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"service":  "academic-service",
		"status":   "healthy",
		"database": "couchbase-connected",
	})
}

func createSubject(c *gin.Context) {
	var subject Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject.ID = uuid.New().String()
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()
	subject.Type = "subject"

	// Insert into Couchbase
	_, err := db.Collection.Insert(subject.ID, subject, nil)
	if err != nil {
		logrus.Errorf("Failed to insert subject: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}

	logrus.Infof("Created subject: %s", subject.ID)
	c.JSON(http.StatusCreated, subject)
}

func getSubjects(c *gin.Context) {
	c.JSON(http.StatusOK, []Subject{})
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

	// Subject CRUD endpoints
	router.POST("/subjects", createSubject)
	router.GET("/subjects", getSubjects)

	// Start server
	log.Println("Starting Academic Service on port 8083 with Couchbase integration")
	router.Run(":8083")
}

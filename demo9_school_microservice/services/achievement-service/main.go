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

type Achievement struct {
ID          string    `json:"id"`
Title       string    `json:"title"`
Description string    `json:"description"`
Category    string    `json:"category"`
Points      int       `json:"points"`
StudentID   string    `json:"studentId"`
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
"service":  "achievement-service",
"status":   "healthy",
"database": "couchbase-connected",
})
}

func createAchievement(c *gin.Context) {
var achievement Achievement
if err := c.ShouldBindJSON(&achievement); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

achievement.ID = uuid.New().String()
achievement.CreatedAt = time.Now()
achievement.UpdatedAt = time.Now()
achievement.Type = "achievement"

// Insert into Couchbase
_, err := db.Collection.Insert(achievement.ID, achievement, nil)
if err != nil {
logrus.Errorf("Failed to insert achievement: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create achievement"})
return
}

logrus.Infof("Created achievement: %s", achievement.ID)
c.JSON(http.StatusCreated, achievement)
}

func getAchievements(c *gin.Context) {
query := `SELECT META(a).id, a.* FROM schoolmgmt a WHERE a.type = "achievement"`
results, err := db.Cluster.Query(query, nil)
if err != nil {
logrus.Errorf("Failed to query achievements: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve achievements"})
return
}

var achievements []Achievement
for results.Next() {
var achievement Achievement
if err := results.Row(&achievement); err != nil {
logrus.Errorf("Failed to scan achievement: %v", err)
continue
}
achievements = append(achievements, achievement)
}

if err := results.Err(); err != nil {
logrus.Errorf("Query iteration error: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve achievements"})
return
}

c.JSON(http.StatusOK, achievements)
}

func getAchievement(c *gin.Context) {
id := c.Param("id")

result, err := db.Collection.Get(id, nil)
if err != nil {
logrus.Errorf("Failed to get achievement: %v", err)
c.JSON(http.StatusNotFound, gin.H{"error": "Achievement not found"})
return
}

var achievement Achievement
if err := result.Content(&achievement); err != nil {
logrus.Errorf("Failed to unmarshal achievement: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve achievement"})
return
}

c.JSON(http.StatusOK, achievement)
}

func updateAchievement(c *gin.Context) {
id := c.Param("id")
var updateData Achievement
if err := c.ShouldBindJSON(&updateData); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

// Get existing achievement
result, err := db.Collection.Get(id, nil)
if err != nil {
logrus.Errorf("Failed to get achievement: %v", err)
c.JSON(http.StatusNotFound, gin.H{"error": "Achievement not found"})
return
}

var existingAchievement Achievement
if err := result.Content(&existingAchievement); err != nil {
logrus.Errorf("Failed to unmarshal achievement: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve achievement"})
return
}

// Update fields
if updateData.Title != "" {
existingAchievement.Title = updateData.Title
}
if updateData.Description != "" {
existingAchievement.Description = updateData.Description
}
if updateData.Category != "" {
existingAchievement.Category = updateData.Category
}
if updateData.Points != 0 {
existingAchievement.Points = updateData.Points
}
if updateData.StudentID != "" {
existingAchievement.StudentID = updateData.StudentID
}
if updateData.TeacherID != "" {
existingAchievement.TeacherID = updateData.TeacherID
}
if updateData.Status != "" {
existingAchievement.Status = updateData.Status
}
existingAchievement.UpdatedAt = time.Now()

// Update in Couchbase
_, err = db.Collection.Replace(id, existingAchievement, nil)
if err != nil {
logrus.Errorf("Failed to update achievement: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update achievement"})
return
}

logrus.Infof("Updated achievement: %s", id)
c.JSON(http.StatusOK, existingAchievement)
}

func deleteAchievement(c *gin.Context) {
id := c.Param("id")

// Check if achievement exists
_, err := db.Collection.Get(id, nil)
if err != nil {
logrus.Errorf("Failed to get achievement: %v", err)
c.JSON(http.StatusNotFound, gin.H{"error": "Achievement not found"})
return
}

// Delete from Couchbase
_, err = db.Collection.Remove(id, nil)
if err != nil {
logrus.Errorf("Failed to delete achievement: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete achievement"})
return
}

logrus.Infof("Deleted achievement: %s", id)
c.JSON(http.StatusOK, gin.H{"message": "Achievement deleted successfully"})
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

// Achievement CRUD endpoints
router.POST("/achievements", createAchievement)
router.GET("/achievements", getAchievements)
router.GET("/achievements/:id", getAchievement)
router.PUT("/achievements/:id", updateAchievement)
router.DELETE("/achievements/:id", deleteAchievement)

// Start server
log.Println("Starting Achievement Service on port 8084 with Couchbase integration")
router.Run(":8084")
}

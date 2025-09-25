// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"school-microservice/internal/config"
	"school-microservice/internal/handler"
	"school-microservice/internal/repository"
	"school-microservice/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	appName    = "School Management Microservice"
	appVersion = "v1.0.0"
)

func main() {
	log.Printf("Starting %s %s", appName, appVersion)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Connect to Couchbase
	log.Println("Connecting to Couchbase...")
	db, err := repository.NewCouchbaseDB(&cfg.Couchbase)
	if err != nil {
		log.Fatalf("Failed to connect to Couchbase: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Create indexes for better performance
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	
	log.Println("Creating database indexes...")
	if err := db.CreateIndexes(ctx); err != nil {
		log.Printf("Warning: Failed to create some indexes: %v", err)
	}

	// Initialize repository
	repo := repository.NewRepository(db)

	// Initialize service
	svc := service.NewService(repo)

	// Initialize handlers
	h := handler.NewHandler(svc)

	// Setup router
	router := setupRouter(h)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

// setupRouter configures and returns the HTTP router
func setupRouter(h *handler.Handler) *gin.Engine {
	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Health check endpoint
	router.GET("/health", healthCheck)
	router.GET("/", rootHandler)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Student routes
		students := v1.Group("/students")
		{
			students.POST("", h.Student.CreateStudent)
			students.GET("", h.Student.GetAllStudents)
			students.GET("/:id", h.Student.GetStudent)
			students.PUT("/:id", h.Student.UpdateStudent)
			students.DELETE("/:id", h.Student.DeleteStudent)
			students.GET("/grade/:grade", h.Student.GetStudentsByGrade)
		}

		// Teacher routes (placeholder)
		teachers := v1.Group("/teachers")
		{
			teachers.POST("", h.Teacher.CreateTeacher)
			teachers.GET("/:id", h.Teacher.GetTeacher)
		}

		// Class routes (placeholder)
		classes := v1.Group("/classes")
		{
			classes.POST("", h.Class.CreateClass)
			classes.GET("/:id", h.Class.GetClass)
		}

		// Academic routes (placeholder)
		academics := v1.Group("/academics")
		{
			academics.POST("", h.Academic.CreateAcademic)
			academics.GET("/:id", h.Academic.GetAcademic)
		}

		// Achievement routes (placeholder)
		achievements := v1.Group("/achievements")
		{
			achievements.POST("", h.Achievement.CreateAchievement)
			achievements.GET("/:id", h.Achievement.GetAchievement)
		}
	}

	return router
}

// corsMiddleware adds CORS headers
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// healthCheck returns the health status of the service
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   appName,
		"version":   appVersion,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// rootHandler returns basic service information
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service":     appName,
		"version":     appVersion,
		"description": "A comprehensive school management microservice built with Go and Couchbase",
		"endpoints": gin.H{
			"health":       "/health",
			"api_v1":       "/api/v1",
			"students":     "/api/v1/students",
			"teachers":     "/api/v1/teachers",
			"classes":      "/api/v1/classes",
			"academics":    "/api/v1/academics",
			"achievements": "/api/v1/achievements",
		},
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"schoolmgmt/services/student-service/internal/handlers"
	"schoolmgmt/services/student-service/internal/repository"
	"schoolmgmt/services/student-service/internal/service"
	"schoolmgmt/shared/pkg/config"
	"schoolmgmt/shared/pkg/database"
	"schoolmgmt/shared/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg)
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handlers.NewStudentHandler(studentService)

	// Setup Gin router
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", studentHandler.HealthCheck)

	// API routes
	v1 := router.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.POST("", studentHandler.CreateStudent)
			students.GET("", studentHandler.ListStudents)
			students.GET("/:id", studentHandler.GetStudent)
			students.PUT("/:id", studentHandler.UpdateStudent)
			students.DELETE("/:id", studentHandler.DeleteStudent)
		}
	}

	// Start server
	logrus.Infof("Starting %s on port %s", cfg.ServiceName, cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

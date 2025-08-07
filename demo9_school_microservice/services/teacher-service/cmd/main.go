package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"schoolmgmt/services/teacher-service/internal/handlers"
	"schoolmgmt/services/teacher-service/internal/repository"
	"schoolmgmt/services/teacher-service/internal/service"
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
	teacherRepo := repository.NewTeacherRepository(db)
	teacherService := service.NewTeacherService(teacherRepo)
	teacherHandler := handlers.NewTeacherHandler(teacherService)

	// Setup Gin router
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", teacherHandler.HealthCheck)

	// API routes
	v1 := router.Group("/api/v1")
	{
		teachers := v1.Group("/teachers")
		{
			teachers.POST("", teacherHandler.CreateTeacher)
			teachers.GET("", teacherHandler.ListTeachers)
			teachers.GET("/:id", teacherHandler.GetTeacher)
			teachers.PUT("/:id", teacherHandler.UpdateTeacher)
			teachers.DELETE("/:id", teacherHandler.DeleteTeacher)
		}
	}

	// Start server
	logrus.Infof("Starting %s on port %s", cfg.ServiceName, cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

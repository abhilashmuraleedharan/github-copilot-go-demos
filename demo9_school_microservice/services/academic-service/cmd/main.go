package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"schoolmgmt/services/academic-service/internal/handlers"
	"schoolmgmt/services/academic-service/internal/repository"
	"schoolmgmt/services/academic-service/internal/service"
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

	logrus.Infof("🚀 Starting Academic Service with Couchbase integration")
	logrus.Infof("📝 Configuration - Host: %s, Bucket: %s", cfg.CouchbaseHost, cfg.CouchbaseBucket)

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg)
	if err != nil {
		logrus.Fatalf("❌ Failed to connect to Couchbase: %v", err)
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	logrus.Infof("✅ Couchbase connection established for Academic Service")

	// Initialize repository, service, and handler
	academicRepo := repository.NewAcademicRepository(db)
	academicService := service.NewAcademicService(academicRepo)
	academicHandler := handlers.NewAcademicHandler(academicService)

	logrus.Infof("✅ Academic service components initialized")

	// Setup Gin router
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", academicHandler.HealthCheck)

	// API routes
	v1 := router.Group("/api/v1")
	{
		academics := v1.Group("/academics")
		{
			academics.POST("", academicHandler.CreateAcademic)
			academics.GET("", academicHandler.ListAcademics)
			academics.GET("/:id", academicHandler.GetAcademic)
			academics.PUT("/:id", academicHandler.UpdateAcademic)
			academics.DELETE("/:id", academicHandler.DeleteAcademic)
		}

		classes := v1.Group("/classes")
		{
			classes.POST("", academicHandler.CreateClass)
			classes.GET("", academicHandler.ListClasses)
			classes.GET("/:id", academicHandler.GetClass)
		}
	}

	// Start server
	logrus.Infof("🌐 Starting Academic Service on port %s with Couchbase backend", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		logrus.Fatalf("❌ Failed to start Academic Service: %v", err)
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"schoolmgmt/services/achievement-service/internal/handlers"
	"schoolmgmt/services/achievement-service/internal/repository"
	"schoolmgmt/services/achievement-service/internal/service"
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

	logrus.Infof("🚀 Starting Achievement Service with Couchbase integration")
	logrus.Infof("📝 Configuration - Host: %s, Bucket: %s", cfg.CouchbaseHost, cfg.CouchbaseBucket)

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg)
	if err != nil {
		logrus.Fatalf("❌ Failed to connect to Couchbase: %v", err)
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	logrus.Infof("✅ Couchbase connection established for Achievement Service")

	// Initialize repository, service, and handler
	achievementRepo := repository.NewAchievementRepository(db)
	achievementService := service.NewAchievementService(achievementRepo)
	achievementHandler := handlers.NewAchievementHandler(achievementService)

	logrus.Infof("✅ Achievement service components initialized")

	// Setup Gin router
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", achievementHandler.HealthCheck)

	// API routes
	v1 := router.Group("/api/v1")
	{
		achievements := v1.Group("/achievements")
		{
			achievements.POST("", achievementHandler.CreateAchievement)
			achievements.GET("", achievementHandler.ListAchievements)
			achievements.GET("/:id", achievementHandler.GetAchievement)
			achievements.PUT("/:id", achievementHandler.UpdateAchievement)
			achievements.DELETE("/:id", achievementHandler.DeleteAchievement)
		}

		badges := v1.Group("/badges")
		{
			badges.POST("", achievementHandler.CreateBadge)
			badges.GET("", achievementHandler.ListBadges)
			badges.GET("/:id", achievementHandler.GetBadge)
		}
	}

	// Start server
	logrus.Infof("🌐 Starting Achievement Service on port %s with Couchbase backend", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		logrus.Fatalf("❌ Failed to start Achievement Service: %v", err)
		log.Fatal("Failed to start server:", err)
	}
}

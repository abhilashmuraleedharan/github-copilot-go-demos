package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"schoolmgmt/gateway/internal/handlers"
	"schoolmgmt/shared/pkg/config"
	"schoolmgmt/shared/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Initialize handlers
	gatewayHandler := handlers.NewGatewayHandler()

	// Setup Gin router
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", gatewayHandler.HealthCheck)

	// API Gateway routes
	v1 := router.Group("/api/v1")
	{
		// Student service routes
		students := v1.Group("/students")
		{
			students.POST("", gatewayHandler.ProxyToStudentService)
			students.GET("", gatewayHandler.ProxyToStudentService)
			students.GET("/:id", gatewayHandler.ProxyToStudentService)
			students.PUT("/:id", gatewayHandler.ProxyToStudentService)
			students.DELETE("/:id", gatewayHandler.ProxyToStudentService)
		}

		// Teacher service routes
		teachers := v1.Group("/teachers")
		{
			teachers.POST("", gatewayHandler.ProxyToTeacherService)
			teachers.GET("", gatewayHandler.ProxyToTeacherService)
			teachers.GET("/:id", gatewayHandler.ProxyToTeacherService)
			teachers.PUT("/:id", gatewayHandler.ProxyToTeacherService)
			teachers.DELETE("/:id", gatewayHandler.ProxyToTeacherService)
		}

		// Academic service routes
		academics := v1.Group("/academics")
		{
			academics.POST("", gatewayHandler.ProxyToAcademicService)
			academics.GET("", gatewayHandler.ProxyToAcademicService)
			academics.GET("/:id", gatewayHandler.ProxyToAcademicService)
			academics.PUT("/:id", gatewayHandler.ProxyToAcademicService)
			academics.DELETE("/:id", gatewayHandler.ProxyToAcademicService)
		}

		classes := v1.Group("/classes")
		{
			classes.POST("", gatewayHandler.ProxyToAcademicService)
			classes.GET("", gatewayHandler.ProxyToAcademicService)
			classes.GET("/:id", gatewayHandler.ProxyToAcademicService)
			classes.PUT("/:id", gatewayHandler.ProxyToAcademicService)
			classes.DELETE("/:id", gatewayHandler.ProxyToAcademicService)
		}

		// Achievement service routes
		achievements := v1.Group("/achievements")
		{
			achievements.POST("", gatewayHandler.ProxyToAchievementService)
			achievements.GET("", gatewayHandler.ProxyToAchievementService)
			achievements.GET("/:id", gatewayHandler.ProxyToAchievementService)
			achievements.PUT("/:id", gatewayHandler.ProxyToAchievementService)
			achievements.DELETE("/:id", gatewayHandler.ProxyToAchievementService)
		}

		badges := v1.Group("/badges")
		{
			badges.POST("", gatewayHandler.ProxyToAchievementService)
			badges.GET("", gatewayHandler.ProxyToAchievementService)
			badges.GET("/:id", gatewayHandler.ProxyToAchievementService)
			badges.PUT("/:id", gatewayHandler.ProxyToAchievementService)
			badges.DELETE("/:id", gatewayHandler.ProxyToAchievementService)
		}
	}

	// Start server
	logrus.Infof("Starting API Gateway on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

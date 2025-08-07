package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ServiceConfig struct {
	Name    string
	BaseURL string
}

var services = map[string]ServiceConfig{
	"student":     {Name: "Student Service", BaseURL: "http://student-service:8081"},
	"teacher":     {Name: "Teacher Service", BaseURL: "http://teacher-service:8082"},
	"academic":    {Name: "Academic Service", BaseURL: "http://academic-service:8083"},
	"achievement": {Name: "Achievement Service", BaseURL: "http://achievement-service:8084"},
}

type HealthStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	URL     string `json:"url"`
}

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Gateway health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "api-gateway",
			"version": "1.0.0",
		})
	})

	// Overall system health
	router.GET("/health/system", getSystemHealth)

	// Service discovery
	router.GET("/services", func(c *gin.Context) {
		var serviceList []map[string]string
		for serviceKey, config := range services {
			serviceList = append(serviceList, map[string]string{
				"key":     serviceKey,
				"name":    config.Name,
				"baseUrl": config.BaseURL,
			})
		}
		
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"services": serviceList,
		})
	})

	// API routing
	api := router.Group("/api/v1")
	{
		// Student service routes
		api.Any("/students", proxyRequest("student"))
		api.Any("/students/:id", proxyRequest("student"))
		api.Any("/students/search", proxyRequest("student"))

		// Teacher service routes
		api.Any("/teachers", proxyRequest("teacher"))
		api.Any("/teachers/:id", proxyRequest("teacher"))
		api.Any("/teachers/search", proxyRequest("teacher"))

		// Academic service routes
		api.Any("/academics", proxyRequest("academic"))
		api.Any("/academics/:id", proxyRequest("academic"))
		api.Any("/academics/student/:student_id", proxyRequest("academic"))
		api.Any("/academics/teacher/:teacher_id", proxyRequest("academic"))
		api.Any("/classes", proxyRequest("academic"))
		api.Any("/classes/:id", proxyRequest("academic"))

		// Achievement service routes
		api.Any("/achievements", proxyRequest("achievement"))
		api.Any("/achievements/:id", proxyRequest("achievement"))
		api.Any("/achievements/student/:student_id", proxyRequest("achievement"))
		api.Any("/achievements/teacher/:teacher_id", proxyRequest("achievement"))
		api.Any("/achievements/category/:category", proxyRequest("achievement"))
		api.Any("/achievements/stats", proxyRequest("achievement"))
		api.Any("/awards", proxyRequest("achievement"))
		api.Any("/awards/:id", proxyRequest("achievement"))
		api.Any("/leaderboard", proxyRequest("achievement"))
	}

	// Dashboard endpoints
	dashboard := router.Group("/dashboard")
	{
		dashboard.GET("/stats", getDashboardStats)
		dashboard.GET("/summary", getDashboardSummary)
	}

	log.Println("Starting API Gateway on port 8080")
	router.Run(":8080")
}

func proxyRequest(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		service, exists := services[serviceName]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}

		// Build target URL
		targetURL := service.BaseURL + c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			targetURL += "?" + c.Request.URL.RawQuery
		}

		// Create request
		var body io.Reader
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			body = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequest(c.Request.Method, targetURL, body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Copy headers
		for name, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(name, value)
			}
		}

		// Make request
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":   fmt.Sprintf("Service %s unavailable", service.Name),
				"details": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		for name, values := range resp.Header {
			for _, value := range values {
				c.Header(name, value)
			}
		}

		// Copy response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
	}
}

func getSystemHealth(c *gin.Context) {
	var healthStatuses []HealthStatus
	var healthyCount int

	for _, service := range services {
		healthURL := service.BaseURL + "/health"
		
		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Get(healthURL)
		
		status := HealthStatus{
			Service: service.Name,
			URL:     healthURL,
		}

		if err != nil {
			status.Status = "unhealthy"
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				status.Status = "healthy"
				healthyCount++
			} else {
				status.Status = "unhealthy"
			}
		}

		healthStatuses = append(healthStatuses, status)
	}

	overallStatus := "healthy"
	if healthyCount < len(services) {
		overallStatus = "degraded"
	}
	if healthyCount == 0 {
		overallStatus = "unhealthy"
	}

	c.JSON(http.StatusOK, gin.H{
		"overall_status": overallStatus,
		"healthy_count":  healthyCount,
		"total_services": len(services),
		"services":       healthStatuses,
		"timestamp":      time.Now(),
	})
}

func getDashboardStats(c *gin.Context) {
	stats := make(map[string]interface{})
	
	// Get stats from each service
	for serviceKey, service := range services {
		switch serviceKey {
		case "student":
			if data := fetchServiceData(service.BaseURL + "/api/v1/students"); data != nil {
				stats["students"] = data
			}
		case "teacher":
			if data := fetchServiceData(service.BaseURL + "/api/v1/teachers"); data != nil {
				stats["teachers"] = data
			}
		case "academic":
			if academicData := fetchServiceData(service.BaseURL + "/api/v1/academics"); academicData != nil {
				stats["academics"] = academicData
			}
			if classData := fetchServiceData(service.BaseURL + "/api/v1/classes"); classData != nil {
				stats["classes"] = classData
			}
		case "achievement":
			if achievementData := fetchServiceData(service.BaseURL + "/api/v1/achievements"); achievementData != nil {
				stats["achievements"] = achievementData
			}
			if awardData := fetchServiceData(service.BaseURL + "/api/v1/awards"); awardData != nil {
				stats["awards"] = awardData
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

func getDashboardSummary(c *gin.Context) {
	summary := map[string]int{
		"total_students":     0,
		"total_teachers":     0,
		"total_academics":    0,
		"total_classes":      0,
		"total_achievements": 0,
		"total_awards":       0,
	}

	// Get counts from each service
	for serviceKey, service := range services {
		switch serviceKey {
		case "student":
			if data := fetchServiceData(service.BaseURL + "/api/v1/students"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_students"] = int(count)
							}
						}
					}
				}
			}
		case "teacher":
			if data := fetchServiceData(service.BaseURL + "/api/v1/teachers"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_teachers"] = int(count)
							}
						}
					}
				}
			}
		case "academic":
			if data := fetchServiceData(service.BaseURL + "/api/v1/academics"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_academics"] = int(count)
							}
						}
					}
				}
			}
			if data := fetchServiceData(service.BaseURL + "/api/v1/classes"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_classes"] = int(count)
							}
						}
					}
				}
			}
		case "achievement":
			if data := fetchServiceData(service.BaseURL + "/api/v1/achievements"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_achievements"] = int(count)
							}
						}
					}
				}
			}
			if data := fetchServiceData(service.BaseURL + "/api/v1/awards"); data != nil {
				if dataMap, ok := data.(map[string]interface{}); ok {
					if success, ok := dataMap["success"].(bool); ok && success {
						if dataSection, ok := dataMap["data"].(map[string]interface{}); ok {
							if count, ok := dataSection["count"].(float64); ok {
								summary["total_awards"] = int(count)
							}
						}
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    summary,
	})
}

func fetchServiceData(serviceURL string) interface{} {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(serviceURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return data
}

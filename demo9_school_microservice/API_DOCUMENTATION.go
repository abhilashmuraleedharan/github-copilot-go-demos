// Package handlers provides HTTP handlers for the School Management System API Gateway.
//
// The gateway acts as a central entry point for all microservices in the system,
// providing request routing, service discovery, and unified API access.
//
// Architecture Overview:
//
// The API Gateway follows a proxy pattern where incoming requests are routed
// to appropriate backend microservices based on URL patterns. It supports
// full HTTP method proxying (GET, POST, PUT, DELETE) and maintains request/response
// integrity while providing centralized logging and monitoring.
//
// Service Discovery:
//
// Services are discovered through environment variables or default to localhost
// for development. In production Kubernetes deployments, service URLs are
// automatically configured through Helm templates.
//
// Supported Services:
//   - Student Service (Port 8081): Student data management
//   - Teacher Service (Port 8082): Teacher data management  
//   - Academic Service (Port 8083): Academic records and classes
//   - Achievement Service (Port 8084): Achievements and badges
//
// URL Routing:
//   - /api/v1/students/* -> Student Service
//   - /api/v1/teachers/* -> Teacher Service
//   - /api/v1/academics/* -> Academic Service
//   - /api/v1/classes/* -> Academic Service (classes endpoint)
//   - /api/v1/achievements/* -> Achievement Service
//   - /api/v1/badges/* -> Achievement Service (badges endpoint)
//   - /health -> Gateway health check
package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GatewayHandler handles HTTP requests for the API Gateway service.
//
// The gateway acts as a reverse proxy, routing requests to appropriate
// backend microservices based on URL patterns. It maintains service
// URLs for all backend services and provides health checking capabilities.
//
// Service URLs are configured through environment variables:
//   - STUDENT_SERVICE_URL: Student service endpoint
//   - TEACHER_SERVICE_URL: Teacher service endpoint  
//   - ACADEMIC_SERVICE_URL: Academic service endpoint
//   - ACHIEVEMENT_SERVICE_URL: Achievement service endpoint
//
// If environment variables are not set, defaults to localhost URLs
// for development purposes.
type GatewayHandler struct {
	// studentServiceURL is the base URL for the Student Service
	studentServiceURL string
	
	// teacherServiceURL is the base URL for the Teacher Service
	teacherServiceURL string
	
	// academicServiceURL is the base URL for the Academic Service
	academicServiceURL string
	
	// achievementServiceURL is the base URL for the Achievement Service
	achievementServiceURL string
}

// NewGatewayHandler creates a new GatewayHandler instance with service URLs
// configured from environment variables or sensible defaults.
//
// Environment Variables:
//   - STUDENT_SERVICE_URL: Default "http://localhost:8081"
//   - TEACHER_SERVICE_URL: Default "http://localhost:8082"
//   - ACADEMIC_SERVICE_URL: Default "http://localhost:8083"
//   - ACHIEVEMENT_SERVICE_URL: Default "http://localhost:8084"
//
// Returns:
//   - *GatewayHandler: Configured gateway handler instance
//
// Example:
//   handler := NewGatewayHandler()
//   router.GET("/health", handler.HealthCheck)
func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{
		studentServiceURL:     getEnv("STUDENT_SERVICE_URL", "http://localhost:8081"),
		teacherServiceURL:     getEnv("TEACHER_SERVICE_URL", "http://localhost:8082"),
		academicServiceURL:    getEnv("ACADEMIC_SERVICE_URL", "http://localhost:8083"),
		achievementServiceURL: getEnv("ACHIEVEMENT_SERVICE_URL", "http://localhost:8084"),
	}
}

// getEnv retrieves an environment variable value or returns a default value if not set.
//
// Parameters:
//   - key: The environment variable name to retrieve
//   - defaultValue: The value to return if the environment variable is not set
//
// Returns:
//   - string: The environment variable value or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// HealthCheck provides a health check endpoint for the API Gateway.
//
// This endpoint returns the gateway's operational status and the configured
// URLs for all backend services. It's used for monitoring, load balancer
// health checks, and Kubernetes readiness/liveness probes.
//
// HTTP Method: GET
// Path: /health
//
// Response Format:
//   {
//     "status": "healthy",
//     "service": "api-gateway", 
//     "services": {
//       "student": "http://localhost:8081",
//       "teacher": "http://localhost:8082",
//       "academic": "http://localhost:8083",
//       "achievement": "http://localhost:8084"
//     }
//   }
//
// Status Codes:
//   - 200 OK: Gateway is healthy and operational
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//
// Example:
//   curl http://localhost:8080/health
func (h *GatewayHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "api-gateway",
		"services": gin.H{
			"student":     h.studentServiceURL,
			"teacher":     h.teacherServiceURL,
			"academic":    h.academicServiceURL,
			"achievement": h.achievementServiceURL,
		},
	})
}

// ProxyToStudentService proxies HTTP requests to the Student Service.
//
// Routes all requests matching /api/v1/students/* to the Student Service.
// Supports all HTTP methods (GET, POST, PUT, DELETE) and preserves
// request headers, query parameters, and request body.
//
// HTTP Methods: GET, POST, PUT, DELETE
// Path Pattern: /api/v1/students/*
// Target Service: Student Service (Port 8081)
//
// URL Mapping:
//   /api/v1/students -> {STUDENT_SERVICE_URL}/api/v1/students
//   /api/v1/students/{id} -> {STUDENT_SERVICE_URL}/api/v1/students/{id}
//
// Request Flow:
//   1. Extract request path, query parameters, headers, and body
//   2. Construct target URL with Student Service base URL
//   3. Forward request to Student Service
//   4. Return response with original status code and headers
//
// Error Handling:
//   - 502 Bad Gateway: Student Service is unavailable
//   - 500 Internal Server Error: Failed to create proxy request
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//
// Example:
//   GET /api/v1/students -> Proxied to Student Service
//   POST /api/v1/students -> Proxied to Student Service
func (h *GatewayHandler) ProxyToStudentService(c *gin.Context) {
	h.proxyRequest(c, h.studentServiceURL, "/api/v1/students")
}

// ProxyToTeacherService proxies HTTP requests to the Teacher Service.
//
// Routes all requests matching /api/v1/teachers/* to the Teacher Service.
// Supports all HTTP methods and preserves request integrity.
//
// HTTP Methods: GET, POST, PUT, DELETE
// Path Pattern: /api/v1/teachers/*
// Target Service: Teacher Service (Port 8082)
//
// URL Mapping:
//   /api/v1/teachers -> {TEACHER_SERVICE_URL}/api/v1/teachers
//   /api/v1/teachers/{id} -> {TEACHER_SERVICE_URL}/api/v1/teachers/{id}
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//
// Example:
//   GET /api/v1/teachers -> Proxied to Teacher Service
//   PUT /api/v1/teachers/123 -> Proxied to Teacher Service
func (h *GatewayHandler) ProxyToTeacherService(c *gin.Context) {
	h.proxyRequest(c, h.teacherServiceURL, "/api/v1/teachers")
}

// ProxyToAcademicService proxies HTTP requests to the Academic Service.
//
// Routes requests for both academic records and classes to the Academic Service.
// Automatically determines the correct endpoint based on URL path.
//
// HTTP Methods: GET, POST, PUT, DELETE
// Path Patterns: 
//   - /api/v1/academics/* -> Academic records endpoint
//   - /api/v1/classes/* -> Classes endpoint
// Target Service: Academic Service (Port 8083)
//
// URL Mapping:
//   /api/v1/academics -> {ACADEMIC_SERVICE_URL}/api/v1/academics
//   /api/v1/classes -> {ACADEMIC_SERVICE_URL}/api/v1/classes
//
// The Academic Service handles both academic records and class management
// through different endpoints on the same service instance.
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//
// Example:
//   GET /api/v1/academics -> Academic records
//   POST /api/v1/classes -> Class management
func (h *GatewayHandler) ProxyToAcademicService(c *gin.Context) {
	path := "/api/v1/academics"
	if strings.Contains(c.Request.URL.Path, "/classes") {
		path = "/api/v1/classes"
	}
	h.proxyRequest(c, h.academicServiceURL, path)
}

// ProxyToAchievementService proxies HTTP requests to the Achievement Service.
//
// Routes requests for both achievements and badges to the Achievement Service.
// Automatically determines the correct endpoint based on URL path.
//
// HTTP Methods: GET, POST, PUT, DELETE
// Path Patterns:
//   - /api/v1/achievements/* -> Achievements endpoint
//   - /api/v1/badges/* -> Badges endpoint  
// Target Service: Achievement Service (Port 8084)
//
// URL Mapping:
//   /api/v1/achievements -> {ACHIEVEMENT_SERVICE_URL}/api/v1/achievements
//   /api/v1/badges -> {ACHIEVEMENT_SERVICE_URL}/api/v1/badges
//
// The Achievement Service manages both student achievements and badge
// systems through different endpoints on the same service instance.
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//
// Example:
//   GET /api/v1/achievements -> Student achievements
//   POST /api/v1/badges -> Badge management
func (h *GatewayHandler) ProxyToAchievementService(c *gin.Context) {
	path := "/api/v1/achievements"
	if strings.Contains(c.Request.URL.Path, "/badges") {
		path = "/api/v1/badges"
	}
	h.proxyRequest(c, h.achievementServiceURL, path)
}

// proxyRequest performs the actual HTTP proxying to backend services.
//
// This is a private helper method that handles the low-level details of
// HTTP request proxying, including URL construction, header preservation,
// body forwarding, and response streaming.
//
// Request Processing:
//   1. Constructs target URL by replacing gateway paths with service paths
//   2. Preserves query parameters from original request
//   3. Reads and forwards request body for POST/PUT requests
//   4. Copies all request headers to target request
//   5. Executes HTTP request to target service
//   6. Streams response back to client with original status and headers
//
// Error Handling:
//   - Logs all proxy failures for monitoring and debugging
//   - Returns 500 Internal Server Error for request creation failures
//   - Returns 502 Bad Gateway for service communication failures
//
// Parameters:
//   - c: Gin context for HTTP request/response handling
//   - serviceURL: Base URL of the target service
//   - basePath: API path prefix for the target service
//
// URL Transformation:
//   The method performs multiple string replacements to transform gateway
//   URLs to service-specific URLs. This handles the routing logic for
//   all supported endpoints.
//
// Performance Considerations:
//   - Uses io.Copy for efficient response streaming
//   - Reuses HTTP client instance (could be optimized with connection pooling)
//   - Preserves original request/response headers for compatibility
//
// Security:
//   - Forwards all headers including authentication tokens
//   - Does not modify or inspect request/response bodies
//   - Maintains end-to-end encryption for HTTPS services
func (h *GatewayHandler) proxyRequest(c *gin.Context, serviceURL, basePath string) {
	// Build target URL
	targetURL := fmt.Sprintf("%s%s", serviceURL, strings.Replace(c.Request.URL.Path, "/api/v1/students", basePath, 1))
	targetURL = strings.Replace(targetURL, "/api/v1/teachers", basePath, 1)
	targetURL = strings.Replace(targetURL, "/api/v1/academics", basePath, 1)
	targetURL = strings.Replace(targetURL, "/api/v1/classes", basePath, 1)
	targetURL = strings.Replace(targetURL, "/api/v1/achievements", basePath, 1)
	targetURL = strings.Replace(targetURL, "/api/v1/badges", basePath, 1)

	if c.Request.URL.RawQuery != "" {
		targetURL += "?" + c.Request.URL.RawQuery
	}

	// Read request body
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
	}

	// Create new request
	req, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logrus.Errorf("Failed to create request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy request"})
		return
	}

	// Copy headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Failed to proxy request to %s: %v", targetURL, err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Copy response body
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

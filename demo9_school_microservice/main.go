// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/demo/school-microservice/config"
	"github.com/demo/school-microservice/handlers"
	"github.com/demo/school-microservice/repository"
	"github.com/demo/school-microservice/service"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()
	
	log.Println("Starting School Management Microservice...")
	log.Printf("Server will run on port %s", cfg.Server.Port)
	log.Printf("Connecting to Couchbase at %s", cfg.Couchbase.ConnectionString)
	
	// Connect to Couchbase
	cluster, err := connectToCouchbase(cfg.Couchbase)
	if err != nil {
		log.Fatalf("Failed to connect to Couchbase: %v", err)
	}
	defer cluster.Close(nil)
	
	log.Println("Successfully connected to Couchbase")
	
	// Initialize repository
	baseRepo, err := repository.NewCouchbaseRepository(
		cluster,
		cfg.Couchbase.BucketName,
		"_default",
		"_default",
	)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	
	// Initialize entity-specific repositories
	studentRepo := repository.NewStudentRepository(baseRepo)
	teacherRepo := repository.NewTeacherRepository(baseRepo)
	classRepo := repository.NewClassRepository(baseRepo)
	academicRepo := repository.NewAcademicRepository(baseRepo)
	examRepo := repository.NewExamRepository(baseRepo)
	examResultRepo := repository.NewExamResultRepository(baseRepo)
	achievementRepo := repository.NewAchievementRepository(baseRepo)
	
	// Initialize service layer
	svc := service.NewService(
		studentRepo,
		teacherRepo,
		classRepo,
		academicRepo,
		examRepo,
		examResultRepo,
		achievementRepo,
	)
	
	// Initialize handlers
	handler := handlers.NewHandler(svc)
	
	// Setup router
	router := mux.NewRouter()
	handler.RegisterRoutes(router)
	
	// Add logging middleware
	router.Use(loggingMiddleware)
	
	// Start HTTP server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server listening on %s", addr)
	
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// connectToCouchbase establishes connection to Couchbase cluster
func connectToCouchbase(cfg config.CouchbaseConfig) (*gocb.Cluster, error) {
	// Connect to Couchbase cluster
	cluster, err := gocb.Connect(cfg.ConnectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.Username,
			Password: cfg.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cluster: %w", err)
	}
	
	// Wait for cluster to be ready
	err = cluster.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("cluster not ready: %w", err)
	}
	
	return cluster, nil
}

// loggingMiddleware logs each HTTP request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Log the request
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"school-microservice/pkg/config"
	"school-microservice/pkg/database"
	"school-microservice/services/students/handlers"
	"school-microservice/services/students/repository"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig("students")

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg, "students")
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository and handler
	repo := repository.NewStudentRepository(db)
	handler := handlers.NewStudentHandler(repo)

	// Setup routes
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Students Service is healthy"))
	}).Methods("GET")

	// Student routes
	router.HandleFunc("/students", handler.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", handler.GetStudent).Methods("GET")
	router.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id}", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", handler.DeleteStudent).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler_with_cors := c.Handler(router)

	log.Printf("Students service starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler_with_cors))
}

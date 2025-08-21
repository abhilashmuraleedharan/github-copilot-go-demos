// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"school-microservice/pkg/config"
	"school-microservice/pkg/database"
	"school-microservice/services/teachers/handlers"
	"school-microservice/services/teachers/repository"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig("teachers")

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg, "teachers")
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository and handler
	repo := repository.NewTeacherRepository(db)
	handler := handlers.NewTeacherHandler(repo)

	// Setup routes
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Teachers Service is healthy"))
	}).Methods("GET")

	// Teacher routes
	router.HandleFunc("/teachers", handler.GetTeachers).Methods("GET")
	router.HandleFunc("/teachers/{id}", handler.GetTeacher).Methods("GET")
	router.HandleFunc("/teachers", handler.CreateTeacher).Methods("POST")
	router.HandleFunc("/teachers/{id}", handler.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teachers/{id}", handler.DeleteTeacher).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler_with_cors := c.Handler(router)

	log.Printf("Teachers service starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler_with_cors))
}

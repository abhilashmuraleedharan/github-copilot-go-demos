// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"school-microservice/pkg/config"
	"school-microservice/pkg/database"
	"school-microservice/services/classes/handlers"
	"school-microservice/services/classes/repository"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig("classes")

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg, "classes")
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository and handler
	repo := repository.NewClassRepository(db)
	handler := handlers.NewClassHandler(repo)

	// Setup routes
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Classes Service is healthy"))
	}).Methods("GET")

	// Class routes
	router.HandleFunc("/classes", handler.GetClasses).Methods("GET")
	router.HandleFunc("/classes/{id}", handler.GetClass).Methods("GET")
	router.HandleFunc("/classes", handler.CreateClass).Methods("POST")
	router.HandleFunc("/classes/{id}", handler.UpdateClass).Methods("PUT")
	router.HandleFunc("/classes/{id}", handler.DeleteClass).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler_with_cors := c.Handler(router)

	log.Printf("Classes service starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler_with_cors))
}

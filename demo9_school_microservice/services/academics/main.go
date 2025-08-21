// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"school-microservice/pkg/config"
	"school-microservice/pkg/database"
	"school-microservice/services/academics/handlers"
	"school-microservice/services/academics/repository"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig("academics")

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg, "academics")
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository and handler
	repo := repository.NewAcademicRepository(db)
	handler := handlers.NewAcademicHandler(repo)

	// Setup routes
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Academics Service is healthy"))
	}).Methods("GET")

	// Academic routes
	router.HandleFunc("/academics", handler.GetAcademics).Methods("GET")
	router.HandleFunc("/academics/{id}", handler.GetAcademic).Methods("GET")
	router.HandleFunc("/academics/student/{studentId}", handler.GetAcademicsByStudent).Methods("GET")
	router.HandleFunc("/academics", handler.CreateAcademic).Methods("POST")
	router.HandleFunc("/academics/{id}", handler.UpdateAcademic).Methods("PUT")
	router.HandleFunc("/academics/{id}", handler.DeleteAcademic).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler_with_cors := c.Handler(router)

	log.Printf("Academics service starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler_with_cors))
}

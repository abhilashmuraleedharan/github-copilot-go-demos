// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"school-microservice/pkg/config"
	"school-microservice/pkg/database"
	"school-microservice/services/achievements/handlers"
	"school-microservice/services/achievements/repository"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig("achievements")

	// Connect to Couchbase
	db, err := database.NewCouchbaseClient(cfg, "achievements")
	if err != nil {
		log.Fatal("Failed to connect to Couchbase:", err)
	}
	defer db.Close()

	// Initialize repository and handler
	repo := repository.NewAchievementRepository(db)
	handler := handlers.NewAchievementHandler(repo)

	// Setup routes
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Achievements Service is healthy"))
	}).Methods("GET")

	// Achievement routes
	router.HandleFunc("/achievements", handler.GetAchievements).Methods("GET")
	router.HandleFunc("/achievements/{id}", handler.GetAchievement).Methods("GET")
	router.HandleFunc("/achievements/student/{studentId}", handler.GetAchievementsByStudent).Methods("GET")
	router.HandleFunc("/achievements", handler.CreateAchievement).Methods("POST")
	router.HandleFunc("/achievements/{id}", handler.UpdateAchievement).Methods("PUT")
	router.HandleFunc("/achievements/{id}", handler.DeleteAchievement).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler_with_cors := c.Handler(router)

	log.Printf("Achievements service starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler_with_cors))
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-06
package main

import (
	"log"
	"net/http"
	"github.com/bookapi/internal/api"
)

func main() {
	handler := api.NewHandler()
	log.Println("Starting bookapi server on :8080...")
	if err := http.ListenAndServe(":8080", handler.Router()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

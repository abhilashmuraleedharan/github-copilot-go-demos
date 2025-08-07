// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2023-10-05
package main

import (
    "log"
    "net/http"

    "bookapi/internal/api"
)

func main() {
    // Initialize the HTTP router
    http.HandleFunc("/books", api.HandleBooks)
    http.HandleFunc("/books/", api.HandleBookByID)

    // Start the HTTP server
    port := ":8080"
    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
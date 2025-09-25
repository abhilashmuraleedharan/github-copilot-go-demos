// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2023-10-05
package main

import (
    "log"
    "net/http"

    "bookapi/internal/books"
)

func main() {
    // Initialize a new router
    mux := http.NewServeMux()

    // Define the routes for the REST API
    mux.HandleFunc("/books", books.HandleBooks)
    mux.HandleFunc("/books/", books.HandleBookByID)

    // Start the HTTP server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
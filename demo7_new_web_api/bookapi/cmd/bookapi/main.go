// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2023-10-01
package main

import (
    "log"
    "net/http"

    "bookapi/internal/handlers"
)

func main() {
    // Initialize the HTTP server and routes
    http.HandleFunc("/books", handlers.BooksHandler)
    http.HandleFunc("/books/", handlers.BookHandler)

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
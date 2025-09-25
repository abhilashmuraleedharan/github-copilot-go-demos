// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2023-10-04
package main

import (
    "log"
    "net/http"

    "bookapi/internal/handlers"
)

func main() {
    // Initialize the HTTP server
    http.HandleFunc("/books", handlers.GetBooks)
    http.HandleFunc("/books/", handlers.GetBookByID)
    http.HandleFunc("/books", handlers.CreateBook)
    http.HandleFunc("/books/", handlers.UpdateBook)
    http.HandleFunc("/books/", handlers.DeleteBook)

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
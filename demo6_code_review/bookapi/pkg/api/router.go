// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package api

import (
    "github.com/gorilla/mux"
    "net/http"
    "bookapi/internal/books"
)

// NewRouter initializes a new router for the book API.
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Define the routes
    router.HandleFunc("/books", books.GetBooks).Methods(http.MethodGet)
    router.HandleFunc("/books", books.CreateBook).Methods(http.MethodPost)
    router.HandleFunc("/books/{id}", books.GetBook).Methods(http.MethodGet)
    router.HandleFunc("/books/{id}", books.UpdateBook).Methods(http.MethodPut)
    router.HandleFunc("/books/{id}", books.DeleteBook).Methods(http.MethodDelete)

    return router
}
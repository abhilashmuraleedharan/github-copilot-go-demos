// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-04
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "bookapi/internal/models"
    "bookapi/internal/storage"
)

// BookHandler handles HTTP requests for books.
type BookHandler struct {
    storage storage.BookStorage
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(storage storage.BookStorage) *BookHandler {
    return &BookHandler{storage: storage}
}

// GetBooks retrieves all books.
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
    books := h.storage.GetAllBooks()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// GetBook retrieves a book by ID.
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    book, err := h.storage.GetBookByID(id)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

// CreateBook adds a new book.
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    h.storage.AddBook(book)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// UpdateBook updates an existing book.
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := h.storage.UpdateBook(id, book); err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteBook removes a book by ID.
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    if err := h.storage.DeleteBook(id); err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
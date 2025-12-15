// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package handlers

import (
    "encoding/json"
    "net/http"
    "sync"

    "bookapi/pkg/models"
)

var (
    books      = make(map[string]models.Book)
    booksMutex = &sync.Mutex{}
)

// GetBooks handles GET /books
func GetBooks(w http.ResponseWriter, r *http.Request) {
    booksMutex.Lock()
    defer booksMutex.Unlock()

    var bookList []models.Book
    for _, book := range books {
        bookList = append(bookList, book)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(bookList)
}

// CreateBook handles POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    booksMutex.Lock()
    books[book.ID] = book
    booksMutex.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// GetBook handles GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    booksMutex.Lock()
    book, exists := books[id]
    booksMutex.Unlock()

    if !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    var updatedBook models.Book
    if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    booksMutex.Lock()
    _, exists := books[id]
    if !exists {
        booksMutex.Unlock()
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    updatedBook.ID = id
    books[id] = updatedBook
    booksMutex.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook handles DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    booksMutex.Lock()
    _, exists := books[id]
    if exists {
        delete(books, id)
    }
    booksMutex.Unlock()

    if !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
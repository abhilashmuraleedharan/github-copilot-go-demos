// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package books

import (
    "encoding/json"
    "net/http"
    "strconv"
)

// Book represents the data model for a book in the library.
type Book struct {
    ID            int    `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedYear int    `json:"published_year"`
}

// Store is an in-memory storage for books.
var Store = make(map[int]Book)
var nextID = 1

// GetBooks handles GET /books requests.
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    books := make([]Book, 0, len(Store))
    for _, book := range Store {
        books = append(books, book)
    }
    json.NewEncoder(w).Encode(books)
}

// GetBook handles GET /books/{id} requests.
func GetBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil || Store[id].ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Store[id])
}

// CreateBook handles POST /books requests.
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    book.ID = nextID
    nextID++
    Store[book.ID] = book
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id} requests.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil || Store[id].ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    book.ID = id
    Store[id] = book
    json.NewEncoder(w).Encode(book)
}

// DeleteBook handles DELETE /books/{id} requests.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/books/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil || Store[id].ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    delete(Store, id)
    w.WriteHeader(http.StatusNoContent)
}
// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-04
package books

import (
    "encoding/json"
    "net/http"
    "sync"
)

// Book represents a book in the library.
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// BookStore is an in-memory store for books.
type BookStore struct {
    sync.RWMutex
    books map[string]Book
}

// NewBookStore initializes a new BookStore.
func NewBookStore() *BookStore {
    return &BookStore{
        books: make(map[string]Book),
    }
}

// GetBooks handles GET /books
func (store *BookStore) GetBooks(w http.ResponseWriter, r *http.Request) {
    store.RLock()
    defer store.RUnlock()

    var books []Book
    for _, book := range store.books {
        books = append(books, book)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// AddBook handles POST /books
func (store *BookStore) AddBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    store.Lock()
    store.books[book.ID] = book
    store.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// GetBook handles GET /books/{id}
func (store *BookStore) GetBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    store.RLock()
    book, exists := store.books[id]
    store.RUnlock()

    if !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id}
func (store *BookStore) UpdateBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    store.Lock()
    defer store.Unlock()

    if _, exists := store.books[id]; !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    store.books[id] = book
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

// DeleteBook handles DELETE /books/{id}
func (store *BookStore) DeleteBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/books/"):]

    store.Lock()
    defer store.Unlock()

    if _, exists := store.books[id]; !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    delete(store.books, id)
    w.WriteHeader(http.StatusNoContent)
}
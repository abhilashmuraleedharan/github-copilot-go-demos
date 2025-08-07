// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package store

import (
    "sync"
    "github.com/yourusername/bookapi/internal/model"
)

// MemoryStore is an in-memory store for books.
type MemoryStore struct {
    mu    sync.RWMutex
    books map[string]model.Book
}

// NewMemoryStore creates a new MemoryStore.
func NewMemoryStore() *MemoryStore {
    return &MemoryStore{
        books: make(map[string]model.Book),
    }
}

// AddBook adds a new book to the store.
func (s *MemoryStore) AddBook(book model.Book) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.books[book.ID] = book
}

// GetBook retrieves a book by its ID.
func (s *MemoryStore) GetBook(id string) (model.Book, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    book, exists := s.books[id]
    return book, exists
}

// UpdateBook updates an existing book in the store.
func (s *MemoryStore) UpdateBook(book model.Book) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[book.ID]; exists {
        s.books[book.ID] = book
        return true
    }
    return false
}

// DeleteBook removes a book from the store by its ID.
func (s *MemoryStore) DeleteBook(id string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[id]; exists {
        delete(s.books, id)
        return true
    }
    return false
}

// GetAllBooks retrieves all books from the store.
func (s *MemoryStore) GetAllBooks() []model.Book {
    s.mu.RLock()
    defer s.mu.RUnlock()
    books := make([]model.Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }
    return books
}
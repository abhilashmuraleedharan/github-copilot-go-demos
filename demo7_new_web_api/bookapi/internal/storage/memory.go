// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-04
package storage

import (
    "sync"
    "github.com/yourusername/bookapi/internal/models"
)

// MemoryStorage provides an in-memory storage for books.
type MemoryStorage struct {
    mu    sync.RWMutex
    books map[string]models.Book
}

// NewMemoryStorage initializes a new MemoryStorage.
func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        books: make(map[string]models.Book),
    }
}

// AddBook adds a new book to the storage.
func (s *MemoryStorage) AddBook(book models.Book) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.books[book.ID] = book
}

// GetBook retrieves a book by its ID.
func (s *MemoryStorage) GetBook(id string) (models.Book, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    book, exists := s.books[id]
    return book, exists
}

// UpdateBook updates an existing book in the storage.
func (s *MemoryStorage) UpdateBook(book models.Book) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[book.ID]; exists {
        s.books[book.ID] = book
        return true
    }
    return false
}

// DeleteBook removes a book from the storage.
func (s *MemoryStorage) DeleteBook(id string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[id]; exists {
        delete(s.books, id)
        return true
    }
    return false
}

// GetAllBooks retrieves all books from the storage.
func (s *MemoryStorage) GetAllBooks() []models.Book {
    s.mu.RLock()
    defer s.mu.RUnlock()
    books := make([]models.Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }
    return books
}
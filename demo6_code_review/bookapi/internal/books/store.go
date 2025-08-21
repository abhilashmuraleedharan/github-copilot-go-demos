// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package books

import (
    "sync"
)

// Book represents the data model for a book in the library.
type Book struct {
    ID            string `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedYear int    `json:"published_year"`
}

// Store manages the in-memory storage of books.
type Store struct {
    mu    sync.RWMutex
    books map[string]Book
}

// NewStore creates a new instance of Store.
func NewStore() *Store {
    return &Store{
        books: make(map[string]Book),
    }
}

// AddBook adds a new book to the store.
func (s *Store) AddBook(book Book) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.books[book.ID] = book
}

// GetBook retrieves a book by its ID.
func (s *Store) GetBook(id string) (Book, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    book, exists := s.books[id]
    return book, exists
}

// UpdateBook updates an existing book in the store.
func (s *Store) UpdateBook(book Book) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[book.ID]; exists {
        s.books[book.ID] = book
        return true
    }
    return false
}

// DeleteBook removes a book from the store by its ID.
func (s *Store) DeleteBook(id string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.books[id]; exists {
        delete(s.books, id)
        return true
    }
    return false
}

// GetAllBooks retrieves all books from the store.
func (s *Store) GetAllBooks() []Book {
    s.mu.RLock()
    defer s.mu.RUnlock()
    books := make([]Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }
    return books
}
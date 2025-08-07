// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-06
package storage

import (
	"sync"
	"github.com/bookapi/internal/model"
)

// MemoryStore provides in-memory storage for books.
type MemoryStore struct {
	mu    sync.RWMutex
	books map[int]model.Book
	lastID int
}

// NewMemoryStore creates a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		books: make(map[int]model.Book),
	}
}

// ListBooks returns all books.
func (s *MemoryStore) ListBooks() []model.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	books := make([]model.Book, 0, len(s.books))
	for _, b := range s.books {
		books = append(books, b)
	}
	return books
}

// GetBook returns a book by ID.
func (s *MemoryStore) GetBook(id int) (model.Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[id]
	return b, ok
}

// AddBook adds a new book and returns its ID.
func (s *MemoryStore) AddBook(book model.Book) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastID++
	book.ID = s.lastID
	s.books[book.ID] = book
	return book.ID
}

// UpdateBook updates an existing book.
func (s *MemoryStore) UpdateBook(id int, book model.Book) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return false
	}
	book.ID = id
	s.books[id] = book
	return true
}

// DeleteBook removes a book by ID.
func (s *MemoryStore) DeleteBook(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return false
	}
	delete(s.books, id)
	return true
}

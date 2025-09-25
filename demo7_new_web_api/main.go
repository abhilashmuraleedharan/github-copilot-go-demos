// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// Book represents a book in the library
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	PublishYear int    `json:"publish_year"`
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// BookStore manages the in-memory book storage
type BookStore struct {
	mu    sync.RWMutex
	books map[int]*Book
	nextID int
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// NewBookStore creates a new book store with sample data
func NewBookStore() *BookStore {
	store := &BookStore{
		books:  make(map[int]*Book),
		nextID: 1,
	}
	
	// Add sample books
	sampleBooks := []*Book{
		{Title: "The Go Programming Language", Author: "Alan Donovan", ISBN: "978-0134190440", PublishYear: 2015},
		{Title: "Clean Code", Author: "Robert Martin", ISBN: "978-0132350884", PublishYear: 2008},
	}
	
	for _, book := range sampleBooks {
		store.AddBook(book)
	}
	
	return store
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// AddBook adds a new book to the store
func (bs *BookStore) AddBook(book *Book) *Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	
	book.ID = bs.nextID
	bs.books[book.ID] = book
	bs.nextID++
	
	return book
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// GetBook retrieves a book by ID
func (bs *BookStore) GetBook(id int) (*Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	
	book, exists := bs.books[id]
	return book, exists
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// GetAllBooks returns all books
func (bs *BookStore) GetAllBooks() []*Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	
	books := make([]*Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}
	
	return books
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// UpdateBook updates an existing book
func (bs *BookStore) UpdateBook(id int, updatedBook *Book) (*Book, bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	
	if _, exists := bs.books[id]; !exists {
		return nil, false
	}
	
	updatedBook.ID = id
	bs.books[id] = updatedBook
	
	return updatedBook, true
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// DeleteBook removes a book by ID
func (bs *BookStore) DeleteBook(id int) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	
	if _, exists := bs.books[id]; !exists {
		return false
	}
	
	delete(bs.books, id)
	return true
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// BookAPI handles HTTP requests for book operations
type BookAPI struct {
	store *BookStore
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// NewBookAPI creates a new book API handler
func NewBookAPI(store *BookStore) *BookAPI {
	return &BookAPI{store: store}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// ServeHTTP implements the http.Handler interface
func (api *BookAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case http.MethodGet:
		api.handleGet(w, r)
	case http.MethodPost:
		api.handlePost(w, r)
	case http.MethodPut:
		api.handlePut(w, r)
	case http.MethodDelete:
		api.handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// handleGet handles GET requests for books
func (api *BookAPI) handleGet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/books")
	
	if path == "" || path == "/" {
		// GET /books - return all books
		books := api.store.GetAllBooks()
		json.NewEncoder(w).Encode(books)
		return
	}
	
	// GET /books/{id} - return specific book
	idStr := strings.TrimPrefix(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	book, exists := api.store.GetBook(id)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(book)
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// handlePost handles POST requests to create new books
func (api *BookAPI) handlePost(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	if book.Title == "" || book.Author == "" {
		http.Error(w, "Title and Author are required", http.StatusBadRequest)
		return
	}
	
	createdBook := api.store.AddBook(&book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// handlePut handles PUT requests to update books
func (api *BookAPI) handlePut(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	if book.Title == "" || book.Author == "" {
		http.Error(w, "Title and Author are required", http.StatusBadRequest)
		return
	}
	
	updatedBook, exists := api.store.UpdateBook(id, &book)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(updatedBook)
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// handleDelete handles DELETE requests to remove books
func (api *BookAPI) handleDelete(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	if !api.store.DeleteBook(id) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	store := NewBookStore()
	api := NewBookAPI(store)
	
	http.Handle("/books", api)
	http.Handle("/books/", api)
	
	port := ":8080"
	fmt.Printf("Book API server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /books      - List all books")
	fmt.Println("  POST   /books      - Create a new book")
	fmt.Println("  GET    /books/{id} - Get a specific book")
	fmt.Println("  PUT    /books/{id} - Update a specific book")
	fmt.Println("  DELETE /books/{id} - Delete a specific book")
	
	server := &http.Server{
		Addr:         port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	log.Fatal(server.ListenAndServe())
}

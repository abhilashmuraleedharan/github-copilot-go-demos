// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-06
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bookapi/internal/model"
)

func TestBookAPI_CRUD(t *testing.T) {
	h := NewHandler()
	r := h.Router()

	// Create a book
	book := model.Book{Title: "Go in Action", Author: "William Kennedy", Year: 2015}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// List books
	req = httptest.NewRequest(http.MethodGet, "/books", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var books []model.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(books) != 1 {
		t.Fatalf("expected 1 book, got %d", len(books))
	}

	id := books[0].ID

	// Get book by ID
	req = httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(id), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Update book
	book.Title = "Go in Action Updated"
	body, _ = json.Marshal(book)
	req = httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(id), bytes.NewReader(body))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// Delete book
	req = httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(id), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}
}

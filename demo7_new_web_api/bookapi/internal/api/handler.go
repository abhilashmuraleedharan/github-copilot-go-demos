// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-06
package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"github.com/bookapi/internal/model"
	"github.com/bookapi/internal/storage"
)

// Handler provides HTTP handlers for book API.
type Handler struct {
	store *storage.MemoryStore
}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	return &Handler{store: storage.NewMemoryStore()}
}

// Router returns the HTTP handler for routing.
func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", h.handleBooks)
	mux.HandleFunc("/books/", h.handleBookByID)
	return mux
}

func (h *Handler) handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books := h.store.ListBooks()
		json.NewEncoder(w).Encode(books)
	case http.MethodPost:
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := h.store.AddBook(book)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		book, ok := h.store.GetBook(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(book)
	case http.MethodPut:
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !h.store.UpdateBook(id, book) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if !h.store.DeleteBook(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2023-10-05
// This file contains unit tests for the book handler functions, ensuring that the API endpoints behave as expected.

package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/yourusername/bookapi/internal/handlers"
    "github.com/yourusername/bookapi/internal/models"
)

func TestGetBooks(t *testing.T) {
    req, err := http.NewRequest("GET", "/books", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.GetBooks)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateBook(t *testing.T) {
    book := models.Book{Title: "Test Book", Author: "Author", PublishedYear: 2023}
    jsonData, _ := json.Marshal(book)

    req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.CreateBook)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetBookByID(t *testing.T) {
    req, err := http.NewRequest("GET", "/books/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.GetBookByID)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateBook(t *testing.T) {
    book := models.Book{ID: 1, Title: "Updated Book", Author: "Updated Author", PublishedYear: 2023}
    jsonData, _ := json.Marshal(book)

    req, err := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.UpdateBook)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteBook(t *testing.T) {
    req, err := http.NewRequest("DELETE", "/books/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.DeleteBook)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
}
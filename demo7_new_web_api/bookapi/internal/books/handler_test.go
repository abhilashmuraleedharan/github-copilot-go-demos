// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2023-10-05
package books

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

var testBook = Book{
    ID:     "1",
    Title:  "Test Book",
    Author: "Test Author",
}

func TestGetBooks(t *testing.T) {
    req, err := http.NewRequest("GET", "/books", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetBooks)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var books []Book
    if err := json.Unmarshal(rr.Body.Bytes(), &books); err != nil {
        t.Fatalf("could not unmarshal response: %v", err)
    }

    if len(books) == 0 {
        t.Error("expected books, got none")
    }
}

func TestPostBook(t *testing.T) {
    bookJSON, err := json.Marshal(testBook)
    if err != nil {
        t.Fatal(err)
    }

    req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(PostBook)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }
}

func TestGetBook(t *testing.T) {
    req, err := http.NewRequest("GET", "/books/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetBook)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var book Book
    if err := json.Unmarshal(rr.Body.Bytes(), &book); err != nil {
        t.Fatalf("could not unmarshal response: %v", err)
    }

    if book.ID != testBook.ID {
        t.Errorf("expected book ID %v, got %v", testBook.ID, book.ID)
    }
}

func TestPutBook(t *testing.T) {
    updatedBook := Book{
        ID:     "1",
        Title:  "Updated Book",
        Author: "Updated Author",
    }

    bookJSON, err := json.Marshal(updatedBook)
    if err != nil {
        t.Fatal(err)
    }

    req, err := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookJSON))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(PutBook)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}

func TestDeleteBook(t *testing.T) {
    req, err := http.NewRequest("DELETE", "/books/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(DeleteBook)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNoContent {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusNoContent)
    }
}
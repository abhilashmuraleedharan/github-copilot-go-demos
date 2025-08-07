// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2023-10-05
package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "bookapi/internal/model"
    "bookapi/internal/store"
)

// MockStore is a mock implementation of the Store interface for testing
type MockStore struct {
    mock.Mock
}

func (m *MockStore) Add(book *model.Book) error {
    args := m.Called(book)
    return args.Error(0)
}

func (m *MockStore) GetAll() ([]*model.Book, error) {
    args := m.Called()
    return args.Get(0).([]*model.Book), args.Error(1)
}

func (m *MockStore) GetByID(id string) (*model.Book, error) {
    args := m.Called(id)
    return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockStore) Update(book *model.Book) error {
    args := m.Called(book)
    return args.Error(0)
}

func (m *MockStore) Delete(id string) error {
    args := m.Called(id)
    return args.Error(0)
}

// TestGetBooks tests the GET /books endpoint
func TestGetBooks(t *testing.T) {
    mockStore := new(MockStore)
    mockStore.On("GetAll").Return([]*model.Book{
        {ID: "1", Title: "Book One", Author: "Author A", PublishedYear: 2021},
        {ID: "2", Title: "Book Two", Author: "Author B", PublishedYear: 2022},
    }, nil)

    req, err := http.NewRequest("GET", "/books", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        GetBooks(w, r, mockStore)
    })

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var books []*model.Book
    json.Unmarshal(rr.Body.Bytes(), &books)
    assert.Len(t, books, 2)
}

// TestAddBook tests the POST /books endpoint
func TestAddBook(t *testing.T) {
    mockStore := new(MockStore)
    newBook := &model.Book{ID: "3", Title: "Book Three", Author: "Author C", PublishedYear: 2023}
    mockStore.On("Add", newBook).Return(nil)

    bookJSON, _ := json.Marshal(newBook)
    req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        AddBook(w, r, mockStore)
    })

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
}

// Additional tests for other endpoints would follow a similar pattern...
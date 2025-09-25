// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package books

// Book represents a book in the library.
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   int    `json:"year"`
}

// BookStore is an in-memory store for books.
type BookStore struct {
    books map[string]Book
}

// NewBookStore initializes a new BookStore.
func NewBookStore() *BookStore {
    return &BookStore{
        books: make(map[string]Book),
    }
}

// AddBook adds a new book to the store.
func (s *BookStore) AddBook(book Book) {
    s.books[book.ID] = book
}

// GetBook retrieves a book by its ID.
func (s *BookStore) GetBook(id string) (Book, bool) {
    book, exists := s.books[id]
    return book, exists
}

// UpdateBook updates an existing book in the store.
func (s *BookStore) UpdateBook(book Book) {
    s.books[book.ID] = book
}

// DeleteBook removes a book from the store.
func (s *BookStore) DeleteBook(id string) {
    delete(s.books, id)
}

// GetAllBooks retrieves all books from the store.
func (s *BookStore) GetAllBooks() []Book {
    allBooks := make([]Book, 0, len(s.books))
    for _, book := range s.books {
        allBooks = append(allBooks, book)
    }
    return allBooks
}
# Book API

## Overview
`bookapi` is a simple Go web API service designed to manage books in a library. It provides RESTful endpoints for creating, retrieving, updating, and deleting book records.

## Project Structure
```
bookapi
├── cmd
│   └── bookapi
│       └── main.go        # Entry point of the application
├── internal
│   └── books
│       ├── handler.go     # HTTP handler functions for REST endpoints
│       ├── model.go       # Data structures and methods for book data
│       └── handler_test.go # Unit tests for the handler functions
├── go.mod                  # Module definition and dependencies
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

## Build and Run Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd bookapi
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/bookapi/main.go
   ```

4. **Access the API:**
   The API will be available at `http://localhost:8080`.

## API Endpoints

- **GET /books**: Retrieve a list of all books.
- **POST /books**: Add a new book to the library.
- **GET /books/{id}**: Retrieve a specific book by its ID.
- **PUT /books/{id}**: Update an existing book by its ID.
- **DELETE /books/{id}**: Remove a book from the library by its ID.

## Testing

To run the unit tests, use the following command:
```bash
go test ./internal/books
```

This will execute the tests defined in `handler_test.go` to ensure that the API endpoints function as expected.
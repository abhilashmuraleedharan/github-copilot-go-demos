# Book API

## Overview
The `bookapi` is a simple Go web API service designed to manage books in a library. It provides a RESTful interface for performing CRUD operations on book records.

## Project Structure
```
bookapi
├── cmd
│   └── bookapi
│       └── main.go         # Entry point of the application
├── internal
│   └── handlers
│       └── books.go        # HTTP handler functions for managing books
├── pkg
│   └── models
│       └── book.go         # Data model for a book
├── tests
│   └── handlers_test.go     # Unit tests for the handler functions
├── go.mod                   # Module definition and dependencies
├── go.sum                   # Checksums for module dependencies
└── README.md                # Project documentation
```

## API Endpoints

### GET /books
Retrieves a list of all books in the library.

### POST /books
Adds a new book to the library. The request body should contain the book details in JSON format.

### GET /books/{id}
Retrieves the details of a specific book by its ID.

### PUT /books/{id}
Updates the details of a specific book by its ID. The request body should contain the updated book details in JSON format.

### DELETE /books/{id}
Deletes a specific book from the library by its ID.

## Build and Run Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd bookapi
   ```

2. Build the application:
   ```
   go build -o bookapi ./cmd/bookapi
   ```

3. Run the application:
   ```
   ./bookapi
   ```

4. The API will be available at `http://localhost:8080`.

## Testing

To run the unit tests, use the following command:
```
go test ./tests
```

This will execute all tests defined in the `handlers_test.go` file.
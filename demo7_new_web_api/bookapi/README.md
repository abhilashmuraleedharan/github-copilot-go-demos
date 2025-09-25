# Book API

## Overview
The `bookapi` is a Go web API service designed to manage books in a library. It provides a RESTful interface for performing CRUD operations on book records.

## Features
- REST endpoints for managing books:
  - `GET /books`: Retrieve a list of all books.
  - `POST /books`: Add a new book.
  - `GET /books/{id}`: Retrieve a specific book by ID.
  - `PUT /books/{id}`: Update an existing book by ID.
  - `DELETE /books/{id}`: Remove a book by ID.
- In-memory storage for books, allowing for quick access and manipulation.

## Project Structure
```
bookapi
├── cmd
│   └── bookapi
│       └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── book_handler.go  # HTTP handler functions for managing books
│   ├── models
│   │   └── book.go          # Definition of the Book struct
│   └── storage
│       └── memory.go        # In-memory storage implementation
├── pkg
│   └── api
│       └── response.go      # Utility functions for API responses
├── tests
│   └── book_handler_test.go  # Unit tests for book handler functions
├── go.mod                    # Go module definition
├── go.sum                    # Module dependency checksums
└── README.md                 # Project documentation
```

## Build and Run Instructions
1. Ensure you have Go installed on your machine.
2. Clone the repository:
   ```
   git clone <repository-url>
   cd bookapi
   ```
3. Build the application:
   ```
   go build -o bookapi ./cmd/bookapi
   ```
4. Run the application:
   ```
   ./bookapi
   ```
5. The API will be available at `http://localhost:8080`.

## Testing
To run the unit tests, use the following command:
```
go test ./tests
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.
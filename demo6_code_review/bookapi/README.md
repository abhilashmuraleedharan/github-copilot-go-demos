# Book API

`bookapi` is a simple Go web API service designed to manage books in a library. This service provides RESTful endpoints to perform CRUD operations on book records.

## Requirements

- Go 1.16 or later
- Standard Go packages (net/http)

## Endpoints

The following REST endpoints are available:

- **GET /books**: Retrieve a list of all books.
- **POST /books**: Add a new book to the library.
- **GET /books/{id}**: Retrieve a specific book by its ID.
- **PUT /books/{id}**: Update an existing book by its ID.
- **DELETE /books/{id}**: Remove a book from the library by its ID.

## Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd bookapi
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Running the Application

To run the application, execute the following command:
```
go run cmd/bookapi/main.go
```

The server will start and listen on port 8080 by default. You can access the API at `http://localhost:8080`.

## Testing

To run the unit tests, use the following command:
```
go test ./internal/books
```

This will execute the tests defined in `handler_test.go` to ensure the API behaves as expected.

## Folder Structure

```
bookapi
├── cmd
│   └── bookapi
│       └── main.go         # Entry point of the application
├── internal
│   └── books
│       ├── handler.go      # HTTP handler functions for REST endpoints
│       ├── model.go        # Data model for a book
│       └── store.go        # In-memory storage management
├── pkg
│   └── api
│       └── router.go       # HTTP route setup
├── go.mod                   # Module definition
├── go.sum                   # Dependency checksums
├── README.md                # Project documentation
└── internal
    └── books
        └── handler_test.go  # Unit tests for handler functions
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-06

# bookapi

A simple Go web API service to manage books in a library.

## Build & Run

```sh
# Windows PowerShell
cd demo7_new_web_api/bookapi
# Build
go build -o bookapi.exe
# Run
./bookapi.exe
```

The service runs on port 8080 by default.

## API Endpoints
- `GET    /books`         : List all books
- `POST   /books`         : Add a new book (JSON body)
- `GET    /books/{id}`    : Get book by ID
- `PUT    /books/{id}`    : Update book by ID (JSON body)
- `DELETE /books/{id}`    : Delete book by ID

## Example Book JSON
```
{
  "title": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "year": 2015
}
```

## Run Tests
```sh
# Run all tests
cd demo7_new_web_api/bookapi
# Run tests
 go test ./...
```
=======
# Book API

`bookapi` is a simple Go web API service designed to manage books in a library. This project demonstrates how to create a RESTful API using idiomatic Go and the standard `net/http` package.

## Features

- REST endpoints for managing books:
  - `GET /books` - Retrieve a list of all books
  - `POST /books` - Add a new book
  - `GET /books/{id}` - Retrieve a book by its ID
  - `PUT /books/{id}` - Update a book by its ID
  - `DELETE /books/{id}` - Delete a book by its ID

## Project Structure

```
bookapi
├── cmd
│   └── bookapi
│       └── main.go        # Entry point of the application
├── internal
│   ├── api
│   │   └── handler.go     # HTTP handlers for the REST API
│   ├── model
│   │   └── book.go        # Book model definition
│   └── store
│       └── memory.go      # In-memory store for managing books
├── pkg
│   └── utils.go           # Utility functions
├── go.mod                 # Module definition and dependencies
├── go.sum                 # Dependency checksums
└── internal
    └── api
        └── handler_test.go # Unit tests for the API handlers
```

## Build and Run Instructions

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd bookapi
   ```

2. Build the application:
   ```bash
   go build -o bookapi ./cmd/bookapi
   ```

3. Run the application:
   ```bash
   ./bookapi
   ```

4. The API will be available at `http://localhost:8080`.

## API Usage

You can use tools like `curl` or Postman to interact with the API. Here are some example requests:

- **Get all books:**
  ```bash
  curl -X GET http://localhost:8080/books
  ```

- **Add a new book:**
  ```bash
  curl -X POST http://localhost:8080/books -d '{"title": "Go Programming", "author": "John Doe", "published_year": 2023}'
  ```

- **Get a book by ID:**
  ```bash
  curl -X GET http://localhost:8080/books/1
  ```

- **Update a book:**
  ```bash
  curl -X PUT http://localhost:8080/books/1 -d '{"title": "Go Programming Updated", "author": "John Doe", "published_year": 2023}'
  ```

- **Delete a book:**
  ```bash
  curl -X DELETE http://localhost:8080/books/1
  ```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.


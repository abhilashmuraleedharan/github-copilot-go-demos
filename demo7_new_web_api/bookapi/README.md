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

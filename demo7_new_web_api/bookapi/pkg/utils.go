// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package pkg

import (
    "errors"
    "net/http"
)

// ValidateBook checks if the provided book data is valid.
func ValidateBook(title, author string, publishedYear int) error {
    if title == "" {
        return errors.New("title cannot be empty")
    }
    if author == "" {
        return errors.New("author cannot be empty")
    }
    if publishedYear <= 0 {
        return errors.New("published year must be a positive integer")
    }
    return nil
}

// RespondWithError sends an error response to the client.
func RespondWithError(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"error": "` + message + `"}`))
}

// RespondWithJSON sends a JSON response to the client.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payload)
}
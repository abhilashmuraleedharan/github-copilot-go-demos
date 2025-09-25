// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-04
package api

import (
    "encoding/json"
    "net/http"
)

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response.
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message,omitempty"`
}

// WriteJSON writes a JSON response to the http.ResponseWriter.
func WriteJSON(w http.ResponseWriter, status int, response interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
}

// WriteSuccess writes a success response.
func WriteSuccess(w http.ResponseWriter, data interface{}) {
    response := SuccessResponse{
        Message: "Request successful",
        Data:    data,
    }
    WriteJSON(w, http.StatusOK, response)
}

// WriteError writes an error response.
func WriteError(w http.ResponseWriter, status int, errMsg string) {
    response := ErrorResponse{
        Error:   http.StatusText(status),
        Message: errMsg,
    }
    WriteJSON(w, status, response)
}
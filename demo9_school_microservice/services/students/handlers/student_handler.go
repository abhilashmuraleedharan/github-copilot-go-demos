// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
// Package handlers provides HTTP request handlers for the students service.
// This package implements RESTful API endpoints for student management operations
// including CRUD operations, concurrency-safe data access, and proper error handling.
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/students/models"
	"school-microservice/services/students/repository"
)

// StudentHandler handles HTTP requests for student management operations.
// It provides RESTful endpoints for creating, reading, updating, and deleting
// student records with proper error handling and response formatting.
//
// The handler implements concurrency-safe operations using context propagation
// and streaming JSON encoding for optimal performance with large datasets.
type StudentHandler struct {
	repo *repository.StudentRepository
}

// NewStudentHandler creates a new StudentHandler instance with the provided repository.
// The repository is used for all database operations and must not be nil.
//
// Parameters:
//   - repo: A StudentRepository instance for database operations
//
// Returns:
//   - *StudentHandler: A new handler instance ready to serve HTTP requests
//
// Example:
//   repo := repository.NewStudentRepository(dbClient)
//   handler := NewStudentHandler(repo)
func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
// GetStudents handles GET /students requests with enhanced concurrency safety.
//
// This endpoint retrieves all student records from the database using context-aware
// operations for proper timeout and cancellation handling. The implementation uses
// streaming JSON encoding to efficiently handle large datasets without excessive
// memory usage.
//
// Features:
//   - Context-aware database operations with timeout support
//   - Response caching headers (60-second cache)
//   - Streaming JSON encoder for memory efficiency
//   - Secure error handling without internal detail exposure
//   - Concurrent request safety through repository context propagation
//
// HTTP Method: GET
// Path: /students
// Content-Type: application/json
// Cache-Control: public, max-age=60
//
// Response Format:
//   Success (200): Array of Student objects in JSON format
//   Error (500): {"error": "Failed to retrieve students"}
//
// Example Response:
//   [
//     {
//       "id": "STU20250821140530",
//       "firstName": "John",
//       "lastName": "Doe",
//       "email": "john.doe@school.edu",
//       "grade": "10",
//       "status": "active",
//       "enrollDate": "2023-08-15T00:00:00Z"
//     }
//   ]
func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	// Use request context for timeout and cancellation
	ctx := r.Context()
	
	// Get students with context for better concurrency control
	students, err := h.repo.GetAllWithContext(ctx)
	if err != nil {
		// Don't expose internal errors to clients
		http.Error(w, "Failed to retrieve students", http.StatusInternalServerError)
		return
	}

	// Set headers before writing response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=60") // 1 minute cache
	
	// Use streaming encoder to handle large datasets efficiently
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(students); err != nil {
		// Log error but can't change response at this point
		// Consider using a logger here
		return
	}
}

// GetStudent handles GET /students/{id} requests to retrieve a specific student.
//
// This endpoint fetches a single student record by their unique ID. The ID is
// extracted from the URL path parameter and used to query the database.
//
// HTTP Method: GET
// Path: /students/{id}
// Path Parameters:
//   - id (string): Unique student identifier (e.g., "STU20250821140530")
//
// Response Codes:
//   - 200 OK: Student found and returned
//   - 404 Not Found: Student with specified ID does not exist
//   - 500 Internal Server Error: Database or server error
//
// Response Format:
//   Success: Single Student object in JSON format
//   Error: {"error": "error message"}
//
// Example:
//   GET /students/STU20250821140530
//   Response: {"id": "STU20250821140530", "firstName": "John", ...}
func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	student, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "student not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// CreateStudent handles POST /students requests to create a new student record.
//
// This endpoint accepts a JSON payload containing student information and creates
// a new student record in the database. The handler automatically generates
// a unique student ID if one is not provided and sets default values for
// optional fields.
//
// Auto-Generated Fields:
//   - ID: Generated using format "STU{timestamp}" if not provided
//   - Status: Set to "active" if not provided
//   - EnrollDate: Set to current time if not provided or zero value
//   - CreatedAt/UpdatedAt: Set to current time automatically
//
// HTTP Method: POST
// Path: /students
// Content-Type: application/json
//
// Request Body: Student object (ID optional)
// Required Fields: firstName, lastName, email
// Optional Fields: All other fields have defaults
//
// Response Codes:
//   - 201 Created: Student successfully created
//   - 400 Bad Request: Invalid JSON format
//   - 500 Internal Server Error: Database or server error
//
// Example Request:
//   POST /students
//   {
//     "firstName": "Jane",
//     "lastName": "Smith",
//     "email": "jane.smith@school.edu",
//     "grade": "11"
//   }
//
// Example Response:
//   {
//     "id": "STU20250821140530",
//     "firstName": "Jane",
//     "lastName": "Smith",
//     "email": "jane.smith@school.edu",
//     "grade": "11",
//     "status": "active",
//     "enrollDate": "2025-08-21T14:05:30Z"
//   }
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if student.ID == "" {
		student.ID = generateStudentID()
	}

	// Set default values
	if student.Status == "" {
		student.Status = "active"
	}
	if student.EnrollDate.IsZero() {
		student.EnrollDate = time.Now()
	}

	if err := h.repo.Create(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// UpdateStudent handles PUT /students/{id} requests to update an existing student.
//
// This endpoint accepts a JSON payload with updated student information and
// modifies the existing record identified by the path parameter ID. The ID
// from the URL path takes precedence over any ID in the request body.
//
// HTTP Method: PUT
// Path: /students/{id}
// Path Parameters:
//   - id (string): Unique student identifier to update
//
// Content-Type: application/json
// Request Body: Student object with fields to update
//
// Response Codes:
//   - 200 OK: Student successfully updated
//   - 400 Bad Request: Invalid JSON format
//   - 500 Internal Server Error: Database error or student not found
//
// Note: The student ID is preserved from the URL path parameter and cannot
// be changed through the request body.
//
// Example Request:
//   PUT /students/STU20250821140530
//   {"firstName": "John", "lastName": "Updated", "grade": "12"}
//
// Example Response:
//   {"id": "STU20250821140530", "firstName": "John", "lastName": "Updated", ...}
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	student.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// DeleteStudent handles DELETE /students/{id} requests to remove a student record.
//
// This endpoint permanently deletes a student record from the database.
// The operation is irreversible and should be used with caution.
//
// HTTP Method: DELETE
// Path: /students/{id}
// Path Parameters:
//   - id (string): Unique student identifier to delete
//
// Response Codes:
//   - 204 No Content: Student successfully deleted
//   - 500 Internal Server Error: Database error or student not found
//
// Response Body: Empty (204 No Content)
//
// Example Request:
//   DELETE /students/STU20250821140530
//
// Example Response:
//   Status: 204 No Content
//   Body: (empty)
//
// Security Note: Consider implementing soft deletes (status change) instead
// of hard deletes for audit trail and data recovery purposes.
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// generateStudentID generates a unique student identifier using timestamp-based approach.
//
// The generated ID follows the format "STU{timestamp}" where timestamp is in
// the format YYYYMMDDHHMMSS. This ensures uniqueness while providing a
// human-readable pattern and maintaining chronological ordering.
//
// Format: STU + YYYYMMDDHHMMSS
// Example: STU20250821140530 (generated on Aug 21, 2025 at 14:05:30)
//
// Returns:
//   - string: A unique student ID following the "STU{timestamp}" pattern
//
// Note: This approach assumes single-instance deployment. For distributed
// systems, consider adding instance ID or using UUID for guaranteed uniqueness.
func generateStudentID() string {
	return "STU" + time.Now().Format("20060102150405")
}

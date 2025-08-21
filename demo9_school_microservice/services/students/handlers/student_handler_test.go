// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/students/models"
)

// MockStudentRepository is a mock implementation of the StudentRepository interface
// for testing purposes. It simulates database operations without requiring a real database.
type MockStudentRepository struct {
	students map[string]*models.Student
	calls    map[string]int // Track method calls for verification
}

// NewMockStudentRepository creates a new mock repository instance with predefined test data.
func NewMockStudentRepository() *MockStudentRepository {
	repo := &MockStudentRepository{
		students: make(map[string]*models.Student),
		calls:    make(map[string]int),
	}
	
	// Add test data
	testStudent := &models.Student{
		ID:          "STU20250821140530",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@school.edu",
		DateOfBirth: time.Date(2008, 5, 15, 0, 0, 0, 0, time.UTC),
		Grade:       "10",
		Address:     "123 Main St, Anytown, USA",
		Phone:       "+1-555-0123",
		EnrollDate:  time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.students[testStudent.ID] = testStudent
	
	return repo
}

// GetAllWithContext mocks retrieving all students with context support
func (m *MockStudentRepository) GetAllWithContext(ctx context.Context) ([]models.Student, error) {
	m.calls["GetAllWithContext"]++
	
	// Simulate context cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	
	var students []models.Student
	for _, student := range m.students {
		students = append(students, *student)
	}
	return students, nil
}

// GetAll mocks retrieving all students (legacy method)
func (m *MockStudentRepository) GetAll() ([]models.Student, error) {
	m.calls["GetAll"]++
	return m.GetAllWithContext(context.Background())
}

// GetByID mocks retrieving a student by ID
func (m *MockStudentRepository) GetByID(id string) (*models.Student, error) {
	m.calls["GetByID"]++
	
	student, exists := m.students[id]
	if !exists {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}

// Create mocks creating a new student
func (m *MockStudentRepository) Create(student *models.Student) error {
	m.calls["Create"]++
	
	// Simulate ID generation if not provided
	if student.ID == "" {
		student.ID = "STU" + time.Now().Format("20060102150405")
	}
	
	// Set timestamps
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	
	// Check for duplicate email (business logic simulation)
	for _, existing := range m.students {
		if existing.Email == student.Email {
			return fmt.Errorf("student with email %s already exists", student.Email)
		}
	}
	
	m.students[student.ID] = student
	return nil
}

// Update mocks updating an existing student
func (m *MockStudentRepository) Update(id string, student *models.Student) error {
	m.calls["Update"]++
	
	_, exists := m.students[id]
	if !exists {
		return fmt.Errorf("student not found")
	}
	
	student.UpdatedAt = time.Now()
	m.students[id] = student
	return nil
}

// Delete mocks deleting a student
func (m *MockStudentRepository) Delete(id string) error {
	m.calls["Delete"]++
	
	_, exists := m.students[id]
	if !exists {
		return fmt.Errorf("student not found")
	}
	
	delete(m.students, id)
	return nil
}

// GetCallCount returns the number of times a method was called
func (m *MockStudentRepository) GetCallCount(method string) int {
	return m.calls[method]
}

// Reset clears all call counts and optionally resets data
func (m *MockStudentRepository) Reset() {
	m.calls = make(map[string]int)
}

// TestStudentHandler_GetStudents tests the GetStudents endpoint with various scenarios
func TestStudentHandler_GetStudents(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockStudentRepository)
		expectedStatus int
		expectedCount  int
		contextTimeout time.Duration
	}{
		{
			name: "successful retrieval of students",
			setupMock: func(mock *MockStudentRepository) {
				// Mock already has one student from initialization
			},
			expectedStatus: http.StatusOK,
			expectedCount:  1,
		},
		{
			name: "empty student list",
			setupMock: func(mock *MockStudentRepository) {
				// Clear all students
				mock.students = make(map[string]*models.Student)
			},
			expectedStatus: http.StatusOK,
			expectedCount:  0,
		},
		{
			name: "context cancellation",
			setupMock: func(mock *MockStudentRepository) {
				// Mock will handle context cancellation
			},
			expectedStatus: http.StatusInternalServerError,
			contextTimeout: 1 * time.Nanosecond, // Very short timeout
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockStudentRepository()
			tt.setupMock(mockRepo)
			handler := NewStudentHandler(mockRepo)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/students", nil)
			
			// Add context timeout if specified
			if tt.contextTimeout > 0 {
				ctx, cancel := context.WithTimeout(context.Background(), tt.contextTimeout)
				defer cancel()
				req = req.WithContext(ctx)
				// Give time for context to expire
				time.Sleep(2 * time.Nanosecond)
			}
			
			w := httptest.NewRecorder()

			// Execute
			handler.GetStudents(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var students []models.Student
				if err := json.Unmarshal(w.Body.Bytes(), &students); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if len(students) != tt.expectedCount {
					t.Errorf("expected %d students, got %d", tt.expectedCount, len(students))
				}

				// Verify response headers
				if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", contentType)
				}

				if cacheControl := w.Header().Get("Cache-Control"); cacheControl != "public, max-age=60" {
					t.Errorf("expected Cache-Control public, max-age=60, got %s", cacheControl)
				}
			}

			// Verify repository method was called
			if mockRepo.GetCallCount("GetAllWithContext") != 1 {
				t.Errorf("expected GetAllWithContext to be called once, called %d times", 
					mockRepo.GetCallCount("GetAllWithContext"))
			}
		})
	}
}

// TestStudentHandler_GetStudent tests the GetStudent endpoint
func TestStudentHandler_GetStudent(t *testing.T) {
	tests := []struct {
		name           string
		studentID      string
		setupMock      func(*MockStudentRepository)
		expectedStatus int
		expectedID     string
	}{
		{
			name:           "successful retrieval",
			studentID:      "STU20250821140530",
			setupMock:      func(mock *MockStudentRepository) {}, // Use default test data
			expectedStatus: http.StatusOK,
			expectedID:     "STU20250821140530",
		},
		{
			name:      "student not found",
			studentID: "NONEXISTENT",
			setupMock: func(mock *MockStudentRepository) {},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockStudentRepository()
			tt.setupMock(mockRepo)
			handler := NewStudentHandler(mockRepo)

			// Create request with URL parameters
			req := httptest.NewRequest(http.MethodGet, "/students/"+tt.studentID, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.studentID})
			w := httptest.NewRecorder()

			// Execute
			handler.GetStudent(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var student models.Student
				if err := json.Unmarshal(w.Body.Bytes(), &student); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if student.ID != tt.expectedID {
					t.Errorf("expected student ID %s, got %s", tt.expectedID, student.ID)
				}

				// Verify student data
				if student.FirstName != "John" || student.LastName != "Doe" {
					t.Errorf("unexpected student data: %+v", student)
				}
			}

			// Verify repository method was called
			if mockRepo.GetCallCount("GetByID") != 1 {
				t.Errorf("expected GetByID to be called once, called %d times", 
					mockRepo.GetCallCount("GetByID"))
			}
		})
	}
}

// TestStudentHandler_CreateStudent tests the CreateStudent endpoint
func TestStudentHandler_CreateStudent(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		validateResult func(*testing.T, *httptest.ResponseRecorder, *MockStudentRepository)
	}{
		{
			name: "successful creation with full data",
			requestBody: models.Student{
				FirstName:   "Jane",
				LastName:    "Smith",
				Email:       "jane.smith@school.edu",
				DateOfBirth: time.Date(2007, 3, 20, 0, 0, 0, 0, time.UTC),
				Grade:       "11",
				Address:     "456 Oak Ave, Anytown, USA",
				Phone:       "+1-555-0456",
				Status:      "active",
			},
			expectedStatus: http.StatusCreated,
			validateResult: func(t *testing.T, w *httptest.ResponseRecorder, repo *MockStudentRepository) {
				var student models.Student
				if err := json.Unmarshal(w.Body.Bytes(), &student); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				// Verify auto-generated fields
				if student.ID == "" {
					t.Error("expected ID to be generated")
				}
				if student.EnrollDate.IsZero() {
					t.Error("expected EnrollDate to be set")
				}
				if student.Status != "active" {
					t.Errorf("expected status to be 'active', got %s", student.Status)
				}

				// Verify repository was called
				if repo.GetCallCount("Create") != 1 {
					t.Errorf("expected Create to be called once")
				}
			},
		},
		{
			name: "creation with minimal data",
			requestBody: models.Student{
				FirstName: "Bob",
				LastName:  "Johnson",
				Email:     "bob.johnson@school.edu",
				Grade:     "9",
			},
			expectedStatus: http.StatusCreated,
			validateResult: func(t *testing.T, w *httptest.ResponseRecorder, repo *MockStudentRepository) {
				var student models.Student
				if err := json.Unmarshal(w.Body.Bytes(), &student); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				// Verify defaults were set
				if student.Status != "active" {
					t.Errorf("expected default status 'active', got %s", student.Status)
				}
				if student.EnrollDate.IsZero() {
					t.Error("expected EnrollDate to be set to current time")
				}
			},
		},
		{
			name:           "invalid JSON",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "duplicate email",
			requestBody: models.Student{
				FirstName: "Duplicate",
				LastName:  "User",
				Email:     "john.doe@school.edu", // Same as existing student
				Grade:     "10",
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockStudentRepository()
			handler := NewStudentHandler(mockRepo)

			// Create request body
			var reqBody *bytes.Buffer
			if str, ok := tt.requestBody.(string); ok {
				reqBody = bytes.NewBufferString(str)
			} else {
				bodyBytes, _ := json.Marshal(tt.requestBody)
				reqBody = bytes.NewBuffer(bodyBytes)
			}

			req := httptest.NewRequest(http.MethodPost, "/students", reqBody)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Execute
			handler.CreateStudent(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, w, mockRepo)
			}
		})
	}
}

// TestStudentHandler_UpdateStudent tests the UpdateStudent endpoint
func TestStudentHandler_UpdateStudent(t *testing.T) {
	tests := []struct {
		name           string
		studentID      string
		requestBody    interface{}
		expectedStatus int
		validateResult func(*testing.T, *httptest.ResponseRecorder, *MockStudentRepository)
	}{
		{
			name:      "successful update",
			studentID: "STU20250821140530",
			requestBody: models.Student{
				FirstName: "John",
				LastName:  "Updated",
				Email:     "john.updated@school.edu",
				Grade:     "11",
				Status:    "active",
			},
			expectedStatus: http.StatusOK,
			validateResult: func(t *testing.T, w *httptest.ResponseRecorder, repo *MockStudentRepository) {
				var student models.Student
				if err := json.Unmarshal(w.Body.Bytes(), &student); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if student.LastName != "Updated" {
					t.Errorf("expected last name 'Updated', got %s", student.LastName)
				}
				if student.ID != "STU20250821140530" {
					t.Errorf("expected ID to be preserved as STU20250821140530, got %s", student.ID)
				}
			},
		},
		{
			name:      "student not found",
			studentID: "NONEXISTENT",
			requestBody: models.Student{
				FirstName: "Non",
				LastName:  "Existent",
				Email:     "non.existent@school.edu",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalid JSON",
			studentID:      "STU20250821140530",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockStudentRepository()
			handler := NewStudentHandler(mockRepo)

			// Create request body
			var reqBody *bytes.Buffer
			if str, ok := tt.requestBody.(string); ok {
				reqBody = bytes.NewBufferString(str)
			} else {
				bodyBytes, _ := json.Marshal(tt.requestBody)
				reqBody = bytes.NewBuffer(bodyBytes)
			}

			req := httptest.NewRequest(http.MethodPut, "/students/"+tt.studentID, reqBody)
			req = mux.SetURLVars(req, map[string]string{"id": tt.studentID})
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Execute
			handler.UpdateStudent(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, w, mockRepo)
			}

			// Verify repository method was called for valid requests
			if tt.expectedStatus != http.StatusBadRequest {
				if mockRepo.GetCallCount("Update") != 1 {
					t.Errorf("expected Update to be called once")
				}
			}
		})
	}
}

// TestStudentHandler_DeleteStudent tests the DeleteStudent endpoint
func TestStudentHandler_DeleteStudent(t *testing.T) {
	tests := []struct {
		name           string
		studentID      string
		expectedStatus int
	}{
		{
			name:           "successful deletion",
			studentID:      "STU20250821140530",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "student not found",
			studentID:      "NONEXISTENT",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockStudentRepository()
			handler := NewStudentHandler(mockRepo)

			req := httptest.NewRequest(http.MethodDelete, "/students/"+tt.studentID, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.studentID})
			w := httptest.NewRecorder()

			// Execute
			handler.DeleteStudent(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Verify empty response body for successful deletion
			if tt.expectedStatus == http.StatusNoContent {
				if w.Body.Len() != 0 {
					t.Errorf("expected empty response body, got %s", w.Body.String())
				}
			}

			// Verify repository method was called
			if mockRepo.GetCallCount("Delete") != 1 {
				t.Errorf("expected Delete to be called once")
			}
		})
	}
}

// TestGenerateStudentID tests the student ID generation function
func TestGenerateStudentID(t *testing.T) {
	// Test multiple generations to ensure uniqueness
	ids := make(map[string]bool)
	
	for i := 0; i < 10; i++ {
		id := generateStudentID()
		
		// Check format
		if len(id) != 17 { // STU + 14 digit timestamp
			t.Errorf("expected ID length 17, got %d for ID: %s", len(id), id)
		}
		
		if id[:3] != "STU" {
			t.Errorf("expected ID to start with 'STU', got: %s", id)
		}
		
		// Check uniqueness (though same millisecond might produce duplicates)
		if ids[id] {
			t.Errorf("duplicate ID generated: %s", id)
		}
		ids[id] = true
		
		// Small delay to ensure different timestamps
		time.Sleep(1 * time.Millisecond)
	}
}

// BenchmarkStudentHandler_GetStudents benchmarks the GetStudents endpoint
func BenchmarkStudentHandler_GetStudents(b *testing.B) {
	mockRepo := NewMockStudentRepository()
	
	// Add more test data for realistic benchmarking
	for i := 0; i < 100; i++ {
		student := &models.Student{
			ID:        fmt.Sprintf("STU%d", i),
			FirstName: fmt.Sprintf("Student%d", i),
			LastName:  "Test",
			Email:     fmt.Sprintf("student%d@school.edu", i),
			Grade:     "10",
			Status:    "active",
		}
		mockRepo.students[student.ID] = student
	}
	
	handler := NewStudentHandler(mockRepo)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/students", nil)
		w := httptest.NewRecorder()
		handler.GetStudents(w, req)
	}
}

// TestStudentHandler_Integration tests the complete CRUD workflow
func TestStudentHandler_Integration(t *testing.T) {
	mockRepo := NewMockStudentRepository()
	handler := NewStudentHandler(mockRepo)

	// Test Create
	student := models.Student{
		FirstName: "Integration",
		LastName:  "Test",
		Email:     "integration.test@school.edu",
		Grade:     "12",
	}
	
	bodyBytes, _ := json.Marshal(student)
	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	handler.CreateStudent(w, req)
	
	if w.Code != http.StatusCreated {
		t.Fatalf("failed to create student: %d", w.Code)
	}
	
	var createdStudent models.Student
	json.Unmarshal(w.Body.Bytes(), &createdStudent)
	studentID := createdStudent.ID
	
	// Test Read
	req = httptest.NewRequest(http.MethodGet, "/students/"+studentID, nil)
	req = mux.SetURLVars(req, map[string]string{"id": studentID})
	w = httptest.NewRecorder()
	
	handler.GetStudent(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("failed to get student: %d", w.Code)
	}
	
	// Test Update
	student.LastName = "Updated"
	bodyBytes, _ = json.Marshal(student)
	req = httptest.NewRequest(http.MethodPut, "/students/"+studentID, bytes.NewBuffer(bodyBytes))
	req = mux.SetURLVars(req, map[string]string{"id": studentID})
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	
	handler.UpdateStudent(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("failed to update student: %d", w.Code)
	}
	
	// Test Delete
	req = httptest.NewRequest(http.MethodDelete, "/students/"+studentID, nil)
	req = mux.SetURLVars(req, map[string]string{"id": studentID})
	w = httptest.NewRecorder()
	
	handler.DeleteStudent(w, req)
	
	if w.Code != http.StatusNoContent {
		t.Fatalf("failed to delete student: %d", w.Code)
	}
	
	// Verify deletion
	req = httptest.NewRequest(http.MethodGet, "/students/"+studentID, nil)
	req = mux.SetURLVars(req, map[string]string{"id": studentID})
	w = httptest.NewRecorder()
	
	handler.GetStudent(w, req)
	
	if w.Code != http.StatusNotFound {
		t.Fatalf("student should be deleted but still found: %d", w.Code)
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/demo/school-microservice/models"
	"github.com/gorilla/mux"
)

// MockService is a mock implementation of the service layer for testing
type MockService struct {
	// Student methods
	CreateStudentFunc func(ctx context.Context, student *models.Student) error
	GetStudentFunc    func(ctx context.Context, id string) (*models.Student, error)
	UpdateStudentFunc func(ctx context.Context, student *models.Student) error
	DeleteStudentFunc func(ctx context.Context, id string) error

	// Teacher methods
	CreateTeacherFunc func(ctx context.Context, teacher *models.Teacher) error
	GetTeacherFunc    func(ctx context.Context, id string) (*models.Teacher, error)
	UpdateTeacherFunc func(ctx context.Context, teacher *models.Teacher) error
	DeleteTeacherFunc func(ctx context.Context, id string) error

	// Class methods
	CreateClassFunc func(ctx context.Context, class *models.Class) error
	GetClassFunc    func(ctx context.Context, id string) (*models.Class, error)
	UpdateClassFunc func(ctx context.Context, class *models.Class) error
	DeleteClassFunc func(ctx context.Context, id string) error

	// Achievement methods
	CreateAchievementFunc func(ctx context.Context, achievement *models.Achievement) error
	GetAchievementFunc    func(ctx context.Context, id string) (*models.Achievement, error)
	UpdateAchievementFunc func(ctx context.Context, achievement *models.Achievement) error
	DeleteAchievementFunc func(ctx context.Context, id string) error
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Student mock methods
func (m *MockService) CreateStudent(ctx context.Context, student *models.Student) error {
	if m.CreateStudentFunc != nil {
		return m.CreateStudentFunc(ctx, student)
	}
	return nil
}

func (m *MockService) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	if m.GetStudentFunc != nil {
		return m.GetStudentFunc(ctx, id)
	}
	return &models.Student{ID: id}, nil
}

func (m *MockService) UpdateStudent(ctx context.Context, student *models.Student) error {
	if m.UpdateStudentFunc != nil {
		return m.UpdateStudentFunc(ctx, student)
	}
	return nil
}

func (m *MockService) DeleteStudent(ctx context.Context, id string) error {
	if m.DeleteStudentFunc != nil {
		return m.DeleteStudentFunc(ctx, id)
	}
	return nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Teacher mock methods
func (m *MockService) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	if m.CreateTeacherFunc != nil {
		return m.CreateTeacherFunc(ctx, teacher)
	}
	return nil
}

func (m *MockService) GetTeacher(ctx context.Context, id string) (*models.Teacher, error) {
	if m.GetTeacherFunc != nil {
		return m.GetTeacherFunc(ctx, id)
	}
	return &models.Teacher{ID: id}, nil
}

func (m *MockService) UpdateTeacher(ctx context.Context, teacher *models.Teacher) error {
	if m.UpdateTeacherFunc != nil {
		return m.UpdateTeacherFunc(ctx, teacher)
	}
	return nil
}

func (m *MockService) DeleteTeacher(ctx context.Context, id string) error {
	if m.DeleteTeacherFunc != nil {
		return m.DeleteTeacherFunc(ctx, id)
	}
	return nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Class mock methods
func (m *MockService) CreateClass(ctx context.Context, class *models.Class) error {
	if m.CreateClassFunc != nil {
		return m.CreateClassFunc(ctx, class)
	}
	return nil
}

func (m *MockService) GetClass(ctx context.Context, id string) (*models.Class, error) {
	if m.GetClassFunc != nil {
		return m.GetClassFunc(ctx, id)
	}
	return &models.Class{ID: id}, nil
}

func (m *MockService) UpdateClass(ctx context.Context, class *models.Class) error {
	if m.UpdateClassFunc != nil {
		return m.UpdateClassFunc(ctx, class)
	}
	return nil
}

func (m *MockService) DeleteClass(ctx context.Context, id string) error {
	if m.DeleteClassFunc != nil {
		return m.DeleteClassFunc(ctx, id)
	}
	return nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Achievement mock methods
func (m *MockService) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
	if m.CreateAchievementFunc != nil {
		return m.CreateAchievementFunc(ctx, achievement)
	}
	return nil
}

func (m *MockService) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
	if m.GetAchievementFunc != nil {
		return m.GetAchievementFunc(ctx, id)
	}
	return &models.Achievement{ID: id}, nil
}

func (m *MockService) UpdateAchievement(ctx context.Context, achievement *models.Achievement) error {
	if m.UpdateAchievementFunc != nil {
		return m.UpdateAchievementFunc(ctx, achievement)
	}
	return nil
}

func (m *MockService) DeleteAchievement(ctx context.Context, id string) error {
	if m.DeleteAchievementFunc != nil {
		return m.DeleteAchievementFunc(ctx, id)
	}
	return nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Placeholder methods for other entities (Academic, Exam, ExamResult)
func (m *MockService) CreateAcademic(ctx context.Context, academic *models.Academic) error  { return nil }
func (m *MockService) GetAcademic(ctx context.Context, id string) (*models.Academic, error) { return &models.Academic{ID: id}, nil }
func (m *MockService) UpdateAcademic(ctx context.Context, academic *models.Academic) error  { return nil }
func (m *MockService) DeleteAcademic(ctx context.Context, id string) error                  { return nil }
func (m *MockService) CreateExam(ctx context.Context, exam *models.Exam) error              { return nil }
func (m *MockService) GetExam(ctx context.Context, id string) (*models.Exam, error)         { return &models.Exam{ID: id}, nil }
func (m *MockService) UpdateExam(ctx context.Context, exam *models.Exam) error              { return nil }
func (m *MockService) DeleteExam(ctx context.Context, id string) error                      { return nil }
func (m *MockService) CreateExamResult(ctx context.Context, result *models.ExamResult) error { return nil }
func (m *MockService) GetExamResult(ctx context.Context, id string) (*models.ExamResult, error) { return &models.ExamResult{ID: id}, nil }
func (m *MockService) UpdateExamResult(ctx context.Context, result *models.ExamResult) error { return nil }
func (m *MockService) DeleteExamResult(ctx context.Context, id string) error                 { return nil }

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestHealthCheck tests the health check endpoint
func TestHealthCheck(t *testing.T) {
	mockService := &MockService{}
	handler := NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got '%s'", response["status"])
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestCreateStudent tests student creation with various scenarios
func TestCreateStudent(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		mockSetup      func(*MockService)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "successful student creation",
			requestBody: `{
				"id": "student001",
				"firstName": "Alice",
				"lastName": "Johnson",
				"grade": 10
			}`,
			mockSetup: func(m *MockService) {
				m.CreateStudentFunc = func(ctx context.Context, student *models.Student) error {
					if student.ID != "student001" {
						t.Errorf("Expected ID 'student001', got '%s'", student.ID)
					}
					return nil
				}
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var student models.Student
				if err := json.NewDecoder(w.Body).Decode(&student); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if student.FirstName != "Alice" {
					t.Errorf("Expected firstName 'Alice', got '%s'", student.FirstName)
				}
			},
		},
		{
			name:           "invalid JSON payload",
			requestBody:    `{"invalid": json}`,
			mockSetup:      func(m *MockService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "Invalid request payload" {
					t.Errorf("Expected error message 'Invalid request payload', got '%s'", errResp["error"])
				}
			},
		},
		{
			name: "service layer error",
			requestBody: `{
				"id": "student002",
				"firstName": "Bob",
				"lastName": "Smith"
			}`,
			mockSetup: func(m *MockService) {
				m.CreateStudentFunc = func(ctx context.Context, student *models.Student) error {
					return errors.New("database connection failed")
				}
			},
			expectedStatus: http.StatusInternalServerError,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "database connection failed" {
					t.Errorf("Expected error 'database connection failed', got '%s'", errResp["error"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockService{}
			tt.mockSetup(mockService)
			handler := NewHandler(mockService)

			req := httptest.NewRequest(http.MethodPost, "/api/students", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.CreateStudent(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestGetStudent tests retrieving a student by ID
func TestGetStudent(t *testing.T) {
	tests := []struct {
		name           string
		studentID      string
		mockSetup      func(*MockService)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:      "successful student retrieval",
			studentID: "student001",
			mockSetup: func(m *MockService) {
				m.GetStudentFunc = func(ctx context.Context, id string) (*models.Student, error) {
					return &models.Student{
						ID:        id,
						FirstName: "Alice",
						LastName:  "Johnson",
						Grade:     10,
					}, nil
				}
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var student models.Student
				if err := json.NewDecoder(w.Body).Decode(&student); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if student.ID != "student001" {
					t.Errorf("Expected ID 'student001', got '%s'", student.ID)
				}
			},
		},
		{
			name:      "student not found",
			studentID: "nonexistent",
			mockSetup: func(m *MockService) {
				m.GetStudentFunc = func(ctx context.Context, id string) (*models.Student, error) {
					return nil, errors.New("student not found")
				}
			},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "student not found" {
					t.Errorf("Expected error 'student not found', got '%s'", errResp["error"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockService{}
			tt.mockSetup(mockService)
			handler := NewHandler(mockService)

			req := httptest.NewRequest(http.MethodGet, "/api/students/"+tt.studentID, nil)
			w := httptest.NewRecorder()

			// Use Gorilla Mux to handle path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/students/{id}", handler.GetStudent).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestUpdateTeacher tests teacher update with validation
func TestUpdateTeacher(t *testing.T) {
	tests := []struct {
		name           string
		teacherID      string
		requestBody    string
		mockSetup      func(*MockService)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:      "successful teacher update",
			teacherID: "teacher001",
			requestBody: `{
				"firstName": "John",
				"lastName": "Smith",
				"subject": "Mathematics",
				"email": "john@school.com"
			}`,
			mockSetup: func(m *MockService) {
				m.UpdateTeacherFunc = func(ctx context.Context, teacher *models.Teacher) error {
					if teacher.ID != "teacher001" {
						t.Errorf("Expected ID 'teacher001', got '%s'", teacher.ID)
					}
					return nil
				}
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var teacher models.Teacher
				if err := json.NewDecoder(w.Body).Decode(&teacher); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if teacher.FirstName != "John" {
					t.Errorf("Expected firstName 'John', got '%s'", teacher.FirstName)
				}
			},
		},
		{
			name:      "validation error - missing first name",
			teacherID: "teacher001",
			requestBody: `{
				"lastName": "Smith",
				"subject": "Mathematics",
				"email": "john@school.com"
			}`,
			mockSetup:      func(m *MockService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "first name is required" {
					t.Errorf("Expected error 'first name is required', got '%s'", errResp["error"])
				}
			},
		},
		{
			name:      "validation error - invalid email",
			teacherID: "teacher001",
			requestBody: `{
				"firstName": "John",
				"lastName": "Smith",
				"subject": "Mathematics",
				"email": "invalid-email"
			}`,
			mockSetup:      func(m *MockService) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "invalid email format" {
					t.Errorf("Expected error 'invalid email format', got '%s'", errResp["error"])
				}
			},
		},
		{
			name:      "teacher not found",
			teacherID: "nonexistent",
			requestBody: `{
				"firstName": "John",
				"lastName": "Smith",
				"subject": "Mathematics",
				"email": "john@school.com"
			}`,
			mockSetup: func(m *MockService) {
				m.UpdateTeacherFunc = func(ctx context.Context, teacher *models.Teacher) error {
					return errors.New("teacher not found")
				}
			},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var errResp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if errResp["error"] != "Teacher not found" {
					t.Errorf("Expected error 'Teacher not found', got '%s'", errResp["error"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockService{}
			tt.mockSetup(mockService)
			handler := NewHandler(mockService)

			req := httptest.NewRequest(http.MethodPut, "/api/teachers/"+tt.teacherID, bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/teachers/{id}", handler.UpdateTeacher).Methods(http.MethodPut)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestDeleteTeacher tests teacher deletion
func TestDeleteTeacher(t *testing.T) {
	tests := []struct {
		name           string
		teacherID      string
		mockSetup      func(*MockService)
		expectedStatus int
	}{
		{
			name:      "successful deletion",
			teacherID: "teacher001",
			mockSetup: func(m *MockService) {
				m.DeleteTeacherFunc = func(ctx context.Context, id string) error {
					return nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:      "deletion error",
			teacherID: "teacher001",
			mockSetup: func(m *MockService) {
				m.DeleteTeacherFunc = func(ctx context.Context, id string) error {
					return errors.New("database error")
				}
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockService{}
			tt.mockSetup(mockService)
			handler := NewHandler(mockService)

			req := httptest.NewRequest(http.MethodDelete, "/api/teachers/"+tt.teacherID, nil)
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/teachers/{id}", handler.DeleteTeacher).Methods(http.MethodDelete)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestValidateTeacher tests the teacher validation function
func TestValidateTeacher(t *testing.T) {
	tests := []struct {
		name          string
		teacher       *models.Teacher
		expectedError string
	}{
		{
			name: "valid teacher",
			teacher: &models.Teacher{
				FirstName: "John",
				LastName:  "Smith",
				Subject:   "Mathematics",
				Email:     "john@school.com",
			},
			expectedError: "",
		},
		{
			name: "missing first name",
			teacher: &models.Teacher{
				LastName: "Smith",
				Subject:  "Mathematics",
				Email:    "john@school.com",
			},
			expectedError: "first name is required",
		},
		{
			name: "missing last name",
			teacher: &models.Teacher{
				FirstName: "John",
				Subject:   "Mathematics",
				Email:     "john@school.com",
			},
			expectedError: "last name is required",
		},
		{
			name: "missing subject",
			teacher: &models.Teacher{
				FirstName: "John",
				LastName:  "Smith",
				Email:     "john@school.com",
			},
			expectedError: "subject is required",
		},
		{
			name: "missing email",
			teacher: &models.Teacher{
				FirstName: "John",
				LastName:  "Smith",
				Subject:   "Mathematics",
			},
			expectedError: "email is required",
		},
		{
			name: "invalid email format",
			teacher: &models.Teacher{
				FirstName: "John",
				LastName:  "Smith",
				Subject:   "Mathematics",
				Email:     "invalid-email",
			},
			expectedError: "invalid email format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTeacher(tt.teacher)

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error '%s', got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error '%s', got '%s'", tt.expectedError, err.Error())
				}
			}
		})
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestRespondJSON tests the JSON response helper
func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{"message": "success"}

	respondJSON(w, http.StatusOK, payload)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", response["message"])
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// TestRespondError tests the error response helper
func TestRespondError(t *testing.T) {
	w := httptest.NewRecorder()
	respondError(w, http.StatusBadRequest, "Invalid input")

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var errResp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
		t.Fatalf("Failed to decode error response: %v", err)
	}

	if errResp["error"] != "Invalid input" {
		t.Errorf("Expected error 'Invalid input', got '%s'", errResp["error"])
	}
}

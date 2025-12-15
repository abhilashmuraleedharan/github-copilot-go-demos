// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// Package handlers provides HTTP request handlers for the School Management Microservice REST API.
//
// This package implements RESTful endpoints for managing school entities including students,
// teachers, classes, academics, exams, exam results, and achievements. All handlers follow
// standard HTTP conventions with proper status codes and JSON response formats.
//
// Endpoints:
//   - Health: GET /health
//   - Students: POST, GET, PUT, DELETE /api/students
//   - Teachers: POST, GET, PUT, DELETE /api/teachers
//   - Classes: POST, GET, PUT, DELETE /api/classes
//   - Academics: POST, GET, PUT, DELETE /api/academics
//   - Exams: POST, GET, PUT, DELETE /api/exams
//   - Exam Results: POST, GET, PUT, DELETE /api/exam-results
//   - Achievements: POST, GET, PUT, DELETE /api/achievements
//
// All endpoints accept and return JSON. Error responses follow the format:
//   {"error": "error message"}
//
// Success responses return the entity or a success message with appropriate HTTP status codes:
//   - 200 OK: Successful GET, PUT, DELETE operations
//   - 201 Created: Successful POST operations
//   - 400 Bad Request: Invalid request payload
//   - 404 Not Found: Entity not found
//   - 500 Internal Server Error: Server-side errors
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/demo/school-microservice/models"
	"github.com/demo/school-microservice/service"
	"github.com/gorilla/mux"
)

// Handler holds all HTTP handlers and their dependencies.
// It acts as the entry point for all REST API requests.
type Handler struct {
	service *service.Service
}

// NewHandler creates a new handler instance with the provided service layer.
// The service parameter must not be nil.
//
// Example:
//   svc := service.NewService(repo)
//   handler := handlers.NewHandler(svc)
func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// RegisterRoutes registers all HTTP routes with the provided Gorilla Mux router.
// This method should be called once during application startup to configure all API endpoints.
//
// The router parameter must not be nil. All routes are registered with specific HTTP methods
// using the Methods() constraint to ensure proper REST semantics.
//
// Example:
//   router := mux.NewRouter()
//   handler.RegisterRoutes(router)
//   http.ListenAndServe(":8080", router)
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Health check
	router.HandleFunc("/health", h.HealthCheck).Methods("GET")
	
	// Student routes
	router.HandleFunc("/api/students", h.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", h.GetStudent).Methods("GET")
	router.HandleFunc("/api/students/{id}", h.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", h.DeleteStudent).Methods("DELETE")
	
	// Teacher routes
	router.HandleFunc("/api/teachers", h.CreateTeacher).Methods("POST")
	router.HandleFunc("/api/teachers/{id}", h.GetTeacher).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", h.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/api/teachers/{id}", h.DeleteTeacher).Methods("DELETE")
	
	// Class routes
	router.HandleFunc("/api/classes", h.CreateClass).Methods("POST")
	router.HandleFunc("/api/classes/{id}", h.GetClass).Methods("GET")
	router.HandleFunc("/api/classes/{id}", h.UpdateClass).Methods("PUT")
	router.HandleFunc("/api/classes/{id}", h.DeleteClass).Methods("DELETE")
	
	// Academic routes
	router.HandleFunc("/api/academics", h.CreateAcademic).Methods("POST")
	router.HandleFunc("/api/academics/{id}", h.GetAcademic).Methods("GET")
	router.HandleFunc("/api/academics/{id}", h.UpdateAcademic).Methods("PUT")
	router.HandleFunc("/api/academics/{id}", h.DeleteAcademic).Methods("DELETE")
	
	// Exam routes
	router.HandleFunc("/api/exams", h.CreateExam).Methods("POST")
	router.HandleFunc("/api/exams/{id}", h.GetExam).Methods("GET")
	router.HandleFunc("/api/exams/{id}", h.UpdateExam).Methods("PUT")
	router.HandleFunc("/api/exams/{id}", h.DeleteExam).Methods("DELETE")
	
	// Exam result routes
	router.HandleFunc("/api/exam-results", h.CreateExamResult).Methods("POST")
	router.HandleFunc("/api/exam-results/{id}", h.GetExamResult).Methods("GET")
	router.HandleFunc("/api/exam-results/{id}", h.UpdateExamResult).Methods("PUT")
	router.HandleFunc("/api/exam-results/{id}", h.DeleteExamResult).Methods("DELETE")
	
	// Achievement routes
	router.HandleFunc("/api/achievements", h.CreateAchievement).Methods("POST")
	router.HandleFunc("/api/achievements/{id}", h.GetAchievement).Methods("GET")
	router.HandleFunc("/api/achievements/{id}", h.UpdateAchievement).Methods("PUT")
	router.HandleFunc("/api/achievements/{id}", h.DeleteAchievement).Methods("DELETE")
}

// HealthCheck returns the health status of the service.
// This endpoint is used by load balancers and monitoring systems to verify service availability.
//
// Endpoint: GET /health
//
// Response: 200 OK
//   {"status": "healthy"}
//
// Example:
//   curl http://localhost:8080/health
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

// Student handlers

// CreateStudent creates a new student entity in the system.
// Validates the student data and returns the created student with all fields populated.
//
// Endpoint: POST /api/students
//
// Request Body:
//   {
//     "id": "student001",
//     "firstName": "Alice",
//     "lastName": "Johnson",
//     "grade": 10,
//     "dateOfBirth": "2008-05-15",
//     "email": "alice.j@school.com"
//   }
//
// Response: 201 Created
//   Returns the created student object
//
// Errors:
//   400 Bad Request: Invalid JSON payload
//   500 Internal Server Error: Validation failed or database error
//
// Example:
//   curl -X POST http://localhost:8080/api/students \
//     -H "Content-Type: application/json" \
//     -d '{"id":"student001","firstName":"Alice","lastName":"Johnson","grade":10}'
func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateStudent(r.Context(), &student); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, student)
}

// GetStudent retrieves a student by ID.
//
// Endpoint: GET /api/students/{id}
//
// Path Parameters:
//   id: Student unique identifier
//
// Response: 200 OK
//   Returns the student object
//
// Errors:
//   404 Not Found: Student with the given ID does not exist
//
// Example:
//   curl http://localhost:8080/api/students/student001
func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	student, err := h.service.GetStudent(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, student)
}

// UpdateStudent updates an existing student's information.
// The student ID in the URL path takes precedence over the ID in the request body.
//
// Endpoint: PUT /api/students/{id}
//
// Path Parameters:
//   id: Student unique identifier
//
// Request Body:
//   Complete student object with updated fields
//
// Response: 200 OK
//   Returns the updated student object
//
// Errors:
//   400 Bad Request: Invalid JSON payload
//   500 Internal Server Error: Validation failed or database error
//
// Example:
//   curl -X PUT http://localhost:8080/api/students/student001 \
//     -H "Content-Type: application/json" \
//     -d '{"firstName":"Alice","lastName":"Johnson","grade":11}'
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	student.ID = id
	if err := h.service.UpdateStudent(r.Context(), &student); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, student)
}

// DeleteStudent removes a student from the system.
//
// Endpoint: DELETE /api/students/{id}
//
// Path Parameters:
//   id: Student unique identifier
//
// Response: 200 OK
//   {"message": "Student deleted successfully"}
//
// Errors:
//   500 Internal Server Error: Database error
//
// Example:
//   curl -X DELETE http://localhost:8080/api/students/student001
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteStudent(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Student deleted successfully"})
}

// Teacher handlers

// CreateTeacher creates a new teacher entity in the system.
//
// Endpoint: POST /api/teachers
//
// Request Body:
//   {
//     "id": "teacher001",
//     "firstName": "John",
//     "lastName": "Smith",
//     "subject": "Mathematics",
//     "email": "john.smith@school.com",
//     "hireDate": "2020-08-15"
//   }
//
// Response: 201 Created
//   Returns the created teacher object
//
// Errors:
//   400 Bad Request: Invalid JSON payload
//   500 Internal Server Error: Validation failed or database error
func (h *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateTeacher(r.Context(), &teacher); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, teacher)
}

// GetTeacher retrieves a teacher by ID.
//
// Endpoint: GET /api/teachers/{id}
//
// Response: 200 OK - Returns the teacher object
// Errors: 404 Not Found - Teacher does not exist
func (h *Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	teacher, err := h.service.GetTeacher(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, teacher)
}

// UpdateTeacher updates an existing teacher's information.
// The teacher ID in the URL path takes precedence over the ID in the request body.
//
// Endpoint: PUT /api/teachers/{id}
//
// Path Parameters:
//   id: Teacher unique identifier
//
// Request Body:
//   Complete teacher object with updated fields
//
// Response: 200 OK
//   Returns the updated teacher object
//
// Errors:
//   400 Bad Request: Invalid JSON payload or validation failed
//   404 Not Found: Teacher does not exist
//   500 Internal Server Error: Database error
//
// Example:
//   curl -X PUT http://localhost:8080/api/teachers/teacher001 \
//     -H "Content-Type: application/json" \
//     -d '{"firstName":"John","lastName":"Smith","subject":"Mathematics","email":"john.s@school.com"}'
func (h *Handler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
	// Validate ID parameter is not empty
	if id == "" {
		respondError(w, http.StatusBadRequest, "Teacher ID is required")
		return
	}
	
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	teacher.ID = id
	
	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
	// Validate teacher data before updating
	if err := validateTeacher(&teacher); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
	// Handle different error types from service layer
	if err := h.service.UpdateTeacher(r.Context(), &teacher); err != nil {
		// Check if error is "not found" type
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") {
			respondError(w, http.StatusNotFound, "Teacher not found")
			return
		}
		// Check if error is validation type
		if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "validation") {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		// Default to internal server error
		respondError(w, http.StatusInternalServerError, "Failed to update teacher")
		return
	}
	
	respondJSON(w, http.StatusOK, teacher)
}

// DeleteTeacher removes a teacher from the system.
//
// Endpoint: DELETE /api/teachers/{id}
//
// Response: 200 OK - {"message": "Teacher deleted successfully"}
func (h *Handler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteTeacher(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Teacher deleted successfully"})
}

// Class handlers

// CreateClass creates a new class with teacher assignment validation.
// The service layer verifies that the assigned teacher exists before creating the class.
//
// Endpoint: POST /api/classes
//
// Request Body:
//   {
//     "id": "class001",
//     "name": "Algebra I",
//     "teacherID": "teacher001",
//     "capacity": 30
//   }
//
// Response: 201 Created
// Errors: 400 Bad Request, 500 Internal Server Error (if teacher not found)
func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateClass(r.Context(), &class); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, class)
}

// GetClass retrieves a class by ID.
//
// Endpoint: GET /api/classes/{id}
//
// Response: 200 OK - Returns the class object
// Errors: 404 Not Found
func (h *Handler) GetClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	class, err := h.service.GetClass(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, class)
}

// UpdateClass updates an existing class's information.
//
// Endpoint: PUT /api/classes/{id}
//
// Response: 200 OK
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	class.ID = id
	if err := h.service.UpdateClass(r.Context(), &class); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, class)
}

// DeleteClass removes a class from the system.
//
// Endpoint: DELETE /api/classes/{id}
//
// Response: 200 OK
func (h *Handler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteClass(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Class deleted successfully"})
}

// Academic handlers

// CreateAcademic creates a new academic enrollment linking a student to a class.
//
// Endpoint: POST /api/academics
//
// Response: 201 Created
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) CreateAcademic(w http.ResponseWriter, r *http.Request) {
	var academic models.Academic
	if err := json.NewDecoder(r.Body).Decode(&academic); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateAcademic(r.Context(), &academic); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, academic)
}

// GetAcademic retrieves an academic enrollment by ID.
//
// Endpoint: GET /api/academics/{id}
//
// Response: 200 OK
// Errors: 404 Not Found
func (h *Handler) GetAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	academic, err := h.service.GetAcademic(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, academic)
}

// UpdateAcademic updates an existing academic enrollment.
//
// Endpoint: PUT /api/academics/{id}
//
// Response: 200 OK
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) UpdateAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var academic models.Academic
	if err := json.NewDecoder(r.Body).Decode(&academic); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	academic.ID = id
	if err := h.service.UpdateAcademic(r.Context(), &academic); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, academic)
}

// DeleteAcademic removes an academic enrollment from the system.
//
// Endpoint: DELETE /api/academics/{id}
//
// Response: 200 OK
func (h *Handler) DeleteAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteAcademic(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Academic enrollment deleted successfully"})
}

// Exam handlers

// CreateExam creates a new exam with scoring parameters.
//
// Endpoint: POST /api/exams
//
// Request Body:
//   {
//     "id": "exam001",
//     "name": "Midterm Exam",
//     "classID": "class001",
//     "examDate": "2025-06-15",
//     "totalPoints": 100
//   }
//
// Response: 201 Created
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) CreateExam(w http.ResponseWriter, r *http.Request) {
	var exam models.Exam
	if err := json.NewDecoder(r.Body).Decode(&exam); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateExam(r.Context(), &exam); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, exam)
}

// GetExam retrieves an exam by ID.
//
// Endpoint: GET /api/exams/{id}
//
// Response: 200 OK
// Errors: 404 Not Found
func (h *Handler) GetExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	exam, err := h.service.GetExam(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, exam)
}

// UpdateExam updates an existing exam's information.
//
// Endpoint: PUT /api/exams/{id}
//
// Response: 200 OK
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) UpdateExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var exam models.Exam
	if err := json.NewDecoder(r.Body).Decode(&exam); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	exam.ID = id
	if err := h.service.UpdateExam(r.Context(), &exam); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, exam)
}

// DeleteExam removes an exam from the system.
//
// Endpoint: DELETE /api/exams/{id}
//
// Response: 200 OK
func (h *Handler) DeleteExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteExam(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Exam deleted successfully"})
}

// ExamResult handlers

// CreateExamResult creates a new exam result with automatic grade calculation.
// The service layer calculates the letter grade (A, B, C, D, F) based on the percentage score.
//
// Endpoint: POST /api/exam-results
//
// Request Body:
//   {
//     "id": "result001",
//     "examID": "exam001",
//     "studentID": "student001",
//     "score": 87,
//     "totalPoints": 100
//   }
//
// Response: 201 Created (with auto-calculated "grade": "B")
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) CreateExamResult(w http.ResponseWriter, r *http.Request) {
	var result models.ExamResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateExamResult(r.Context(), &result); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, result)
}

// GetExamResult retrieves an exam result by ID.
//
// Endpoint: GET /api/exam-results/{id}
//
// Response: 200 OK
// Errors: 404 Not Found
func (h *Handler) GetExamResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	result, err := h.service.GetExamResult(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, result)
}

// UpdateExamResult updates an existing exam result with grade recalculation.
//
// Endpoint: PUT /api/exam-results/{id}
//
// Response: 200 OK
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) UpdateExamResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var result models.ExamResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	result.ID = id
	if err := h.service.UpdateExamResult(r.Context(), &result); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, result)
}

// DeleteExamResult removes an exam result from the system.
//
// Endpoint: DELETE /api/exam-results/{id}
//
// Response: 200 OK
func (h *Handler) DeleteExamResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteExamResult(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Exam result deleted successfully"})
}

// Achievement handlers

// CreateAchievement creates a new achievement award for a student.
//
// Endpoint: POST /api/achievements
//
// Request Body:
//   {
//     "id": "achievement001",
//     "studentID": "student001",
//     "title": "Honor Roll",
//     "description": "Academic excellence",
//     "dateAwarded": "2025-06-01"
//   }
//
// Response: 201 Created
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	if err := h.service.CreateAchievement(r.Context(), &achievement); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusCreated, achievement)
}

// GetAchievement retrieves an achievement by ID.
//
// Endpoint: GET /api/achievements/{id}
//
// Response: 200 OK
// Errors: 404 Not Found
func (h *Handler) GetAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	achievement, err := h.service.GetAchievement(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, achievement)
}

// UpdateAchievement updates an existing achievement's information.
//
// Endpoint: PUT /api/achievements/{id}
//
// Response: 200 OK
// Errors: 400 Bad Request, 500 Internal Server Error
func (h *Handler) UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	
	achievement.ID = id
	if err := h.service.UpdateAchievement(r.Context(), &achievement); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, achievement)
}

// DeleteAchievement removes an achievement from the system.
//
// Endpoint: DELETE /api/achievements/{id}
//
// Response: 200 OK
func (h *Handler) DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if err := h.service.DeleteAchievement(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"message": "Achievement deleted successfully"})
}

// Helper functions

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// validateTeacher validates teacher data before creation or update.
// Returns an error if required fields are missing or invalid.
func validateTeacher(teacher *models.Teacher) error {
	if teacher.FirstName == "" {
		return errors.New("first name is required")
	}
	if teacher.LastName == "" {
		return errors.New("last name is required")
	}
	if teacher.Subject == "" {
		return errors.New("subject is required")
	}
	if teacher.Email == "" {
		return errors.New("email is required")
	}
	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
	// Simple email validation using standard library
	if !strings.Contains(teacher.Email, "@") || !strings.Contains(teacher.Email, ".") {
		return errors.New("invalid email format")
	}
	return nil
}

// respondJSON sends a JSON response with the specified HTTP status code.
// It marshals the payload to JSON and sets the appropriate Content-Type header.
// If marshaling fails, it returns a 500 Internal Server Error.
//
// Parameters:
//   w: HTTP response writer
//   status: HTTP status code (e.g., 200, 201, 404)
//   payload: Data to be marshaled to JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// respondError sends an error response in JSON format.
// Error responses follow the standard format: {"error": "message"}
//
// Parameters:
//   w: HTTP response writer
//   code: HTTP error status code (e.g., 400, 404, 500)
//   message: Error message to return to the client
//
// Example:
//   respondError(w, http.StatusBadRequest, "Invalid request payload")
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/teachers/models"
	"school-microservice/services/teachers/repository"
)

// TeacherHandler handles HTTP requests for teachers
type TeacherHandler struct {
	repo *repository.TeacherRepository
}

// NewTeacherHandler creates a new teacher handler
func NewTeacherHandler(repo *repository.TeacherRepository) *TeacherHandler {
	return &TeacherHandler{repo: repo}
}

// GetTeachers handles GET /teachers
func (h *TeacherHandler) GetTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

// GetTeacher handles GET /teachers/{id}
func (h *TeacherHandler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	teacher, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "teacher not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

// CreateTeacher handles POST /teachers
func (h *TeacherHandler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if teacher.ID == "" {
		teacher.ID = generateTeacherID()
	}

	// Set default values
	if teacher.Status == "" {
		teacher.Status = "active"
	}
	if teacher.HireDate.IsZero() {
		teacher.HireDate = time.Now()
	}

	if err := h.repo.Create(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

// UpdateTeacher handles PUT /teachers/{id}
func (h *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teacher.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

// DeleteTeacher handles DELETE /teachers/{id}
func (h *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// generateTeacherID generates a unique teacher ID
func generateTeacherID() string {
	return "TCH" + time.Now().Format("20060102150405")
}

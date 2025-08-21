// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/academics/models"
	"school-microservice/services/academics/repository"
)

// AcademicHandler handles HTTP requests for academics
type AcademicHandler struct {
	repo *repository.AcademicRepository
}

// NewAcademicHandler creates a new academic handler
func NewAcademicHandler(repo *repository.AcademicRepository) *AcademicHandler {
	return &AcademicHandler{repo: repo}
}

// GetAcademics handles GET /academics
func (h *AcademicHandler) GetAcademics(w http.ResponseWriter, r *http.Request) {
	academics, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(academics)
}

// GetAcademic handles GET /academics/{id}
func (h *AcademicHandler) GetAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	academic, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "academic record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(academic)
}

// GetAcademicsByStudent handles GET /academics/student/{studentId}
func (h *AcademicHandler) GetAcademicsByStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]

	academics, err := h.repo.GetByStudentID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(academics)
}

// CreateAcademic handles POST /academics
func (h *AcademicHandler) CreateAcademic(w http.ResponseWriter, r *http.Request) {
	var academic models.Academic
	if err := json.NewDecoder(r.Body).Decode(&academic); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if academic.ID == "" {
		academic.ID = generateAcademicID()
	}

	// Set default values
	if academic.ExamDate.IsZero() {
		academic.ExamDate = time.Now()
	}
	if academic.Year == 0 {
		academic.Year = time.Now().Year()
	}

	if err := h.repo.Create(&academic); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(academic)
}

// UpdateAcademic handles PUT /academics/{id}
func (h *AcademicHandler) UpdateAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var academic models.Academic
	if err := json.NewDecoder(r.Body).Decode(&academic); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &academic); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	academic.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(academic)
}

// DeleteAcademic handles DELETE /academics/{id}
func (h *AcademicHandler) DeleteAcademic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// generateAcademicID generates a unique academic record ID
func generateAcademicID() string {
	return "ACD" + time.Now().Format("20060102150405")
}

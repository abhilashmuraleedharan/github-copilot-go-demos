// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/classes/models"
	"school-microservice/services/classes/repository"
)

// ClassHandler handles HTTP requests for classes
type ClassHandler struct {
	repo *repository.ClassRepository
}

// NewClassHandler creates a new class handler
func NewClassHandler(repo *repository.ClassRepository) *ClassHandler {
	return &ClassHandler{repo: repo}
}

// GetClasses handles GET /classes
func (h *ClassHandler) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

// GetClass handles GET /classes/{id}
func (h *ClassHandler) GetClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	class, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "class not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

// CreateClass handles POST /classes
func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if class.ID == "" {
		class.ID = generateClassID()
	}

	// Set default values
	if class.Status == "" {
		class.Status = "active"
	}
	if class.StudentIDs == nil {
		class.StudentIDs = []string{}
	}
	if class.Year == 0 {
		class.Year = time.Now().Year()
	}

	if err := h.repo.Create(&class); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

// UpdateClass handles PUT /classes/{id}
func (h *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &class); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	class.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

// DeleteClass handles DELETE /classes/{id}
func (h *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// generateClassID generates a unique class ID
func generateClassID() string {
	return "CLS" + time.Now().Format("20060102150405")
}

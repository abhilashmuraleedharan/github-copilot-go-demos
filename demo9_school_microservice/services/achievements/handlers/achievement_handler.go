// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"school-microservice/services/achievements/models"
	"school-microservice/services/achievements/repository"
)

// AchievementHandler handles HTTP requests for achievements
type AchievementHandler struct {
	repo *repository.AchievementRepository
}

// NewAchievementHandler creates a new achievement handler
func NewAchievementHandler(repo *repository.AchievementRepository) *AchievementHandler {
	return &AchievementHandler{repo: repo}
}

// GetAchievements handles GET /achievements
func (h *AchievementHandler) GetAchievements(w http.ResponseWriter, r *http.Request) {
	achievements, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}

// GetAchievement handles GET /achievements/{id}
func (h *AchievementHandler) GetAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	achievement, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "achievement not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}

// GetAchievementsByStudent handles GET /achievements/student/{studentId}
func (h *AchievementHandler) GetAchievementsByStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]

	achievements, err := h.repo.GetByStudentID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}

// CreateAchievement handles POST /achievements
func (h *AchievementHandler) CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if achievement.ID == "" {
		achievement.ID = generateAchievementID()
	}

	// Set default values
	if achievement.Status == "" {
		achievement.Status = "pending"
	}
	if achievement.AwardDate.IsZero() {
		achievement.AwardDate = time.Now()
	}

	if err := h.repo.Create(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(achievement)
}

// UpdateAchievement handles PUT /achievements/{id}
func (h *AchievementHandler) UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &achievement); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	achievement.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}

// DeleteAchievement handles DELETE /achievements/{id}
func (h *AchievementHandler) DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// generateAchievementID generates a unique achievement ID
func generateAchievementID() string {
	return "ACH" + time.Now().Format("20060102150405")
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/database"
	"school-microservice/services/achievements/models"
)

// AchievementRepository handles achievement data operations
type AchievementRepository struct {
	db *database.CouchbaseClient
}

// NewAchievementRepository creates a new achievement repository
func NewAchievementRepository(db *database.CouchbaseClient) *AchievementRepository {
	return &AchievementRepository{db: db}
}

// GetAll retrieves all achievements
func (r *AchievementRepository) GetAll() ([]models.Achievement, error) {
	query := `SELECT META().id, * FROM school WHERE type = "achievement"`
	
	results, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query achievements: %w", err)
	}
	defer results.Close()

	var achievements []models.Achievement
	for results.Next() {
		var row struct {
			ID          string              `json:"id"`
			Achievement models.Achievement `json:"achievement"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Achievement.ID = row.ID
		achievements = append(achievements, row.Achievement)
	}

	return achievements, nil
}

// GetByID retrieves an achievement by ID
func (r *AchievementRepository) GetByID(id string) (*models.Achievement, error) {
	result, err := r.db.Collection.Get(id, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("achievement not found")
		}
		return nil, fmt.Errorf("failed to get achievement: %w", err)
	}

	var achievement models.Achievement
	if err := result.Content(&achievement); err != nil {
		return nil, fmt.Errorf("failed to decode achievement: %w", err)
	}
	achievement.ID = id

	return &achievement, nil
}

// GetByStudentID retrieves achievements by student ID
func (r *AchievementRepository) GetByStudentID(studentID string) ([]models.Achievement, error) {
	query := `SELECT META().id, * FROM school WHERE type = "achievement" AND studentId = $1`
	
	results, err := r.db.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{studentID},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query achievements by student: %w", err)
	}
	defer results.Close()

	var achievements []models.Achievement
	for results.Next() {
		var row struct {
			ID          string              `json:"id"`
			Achievement models.Achievement `json:"achievement"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Achievement.ID = row.ID
		achievements = append(achievements, row.Achievement)
	}

	return achievements, nil
}

// Create creates a new achievement
func (r *AchievementRepository) Create(achievement *models.Achievement) error {
	achievement.CreatedAt = time.Now()
	achievement.UpdatedAt = time.Now()
	
	// Set default points based on level if not provided
	if achievement.Points == 0 {
		achievement.Points = calculatePoints(achievement.Level)
	}
	
	// Add type field for querying
	achievementData := map[string]interface{}{
		"type":        "achievement",
		"studentId":   achievement.StudentID,
		"title":       achievement.Title,
		"description": achievement.Description,
		"category":    achievement.Category,
		"level":       achievement.Level,
		"awardedBy":   achievement.AwardedBy,
		"awardDate":   achievement.AwardDate,
		"points":      achievement.Points,
		"certificate": achievement.Certificate,
		"status":      achievement.Status,
		"teacherId":   achievement.TeacherID,
		"comments":    achievement.Comments,
		"createdAt":   achievement.CreatedAt,
		"updatedAt":   achievement.UpdatedAt,
	}

	_, err := r.db.Collection.Insert(achievement.ID, achievementData, nil)
	if err != nil {
		return fmt.Errorf("failed to create achievement: %w", err)
	}

	return nil
}

// Update updates an existing achievement
func (r *AchievementRepository) Update(id string, achievement *models.Achievement) error {
	achievement.UpdatedAt = time.Now()
	
	achievementData := map[string]interface{}{
		"type":        "achievement",
		"studentId":   achievement.StudentID,
		"title":       achievement.Title,
		"description": achievement.Description,
		"category":    achievement.Category,
		"level":       achievement.Level,
		"awardedBy":   achievement.AwardedBy,
		"awardDate":   achievement.AwardDate,
		"points":      achievement.Points,
		"certificate": achievement.Certificate,
		"status":      achievement.Status,
		"teacherId":   achievement.TeacherID,
		"comments":    achievement.Comments,
		"updatedAt":   achievement.UpdatedAt,
	}

	_, err := r.db.Collection.Replace(id, achievementData, nil)
	if err != nil {
		return fmt.Errorf("failed to update achievement: %w", err)
	}

	return nil
}

// Delete deletes an achievement
func (r *AchievementRepository) Delete(id string) error {
	_, err := r.db.Collection.Remove(id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete achievement: %w", err)
	}

	return nil
}

// calculatePoints calculates default points based on achievement level
func calculatePoints(level string) int {
	switch level {
	case "school":
		return 10
	case "district":
		return 25
	case "state":
		return 50
	case "national":
		return 100
	case "international":
		return 200
	default:
		return 5
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/database"
	"school-microservice/services/academics/models"
)

// AcademicRepository handles academic data operations
type AcademicRepository struct {
	db *database.CouchbaseClient
}

// NewAcademicRepository creates a new academic repository
func NewAcademicRepository(db *database.CouchbaseClient) *AcademicRepository {
	return &AcademicRepository{db: db}
}

// GetAll retrieves all academic records
func (r *AcademicRepository) GetAll() ([]models.Academic, error) {
	query := `SELECT META().id, * FROM school WHERE type = "academic"`
	
	results, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query academics: %w", err)
	}
	defer results.Close()

	var academics []models.Academic
	for results.Next() {
		var row struct {
			ID       string           `json:"id"`
			Academic models.Academic `json:"academic"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Academic.ID = row.ID
		academics = append(academics, row.Academic)
	}

	return academics, nil
}

// GetByID retrieves an academic record by ID
func (r *AcademicRepository) GetByID(id string) (*models.Academic, error) {
	result, err := r.db.Collection.Get(id, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("academic record not found")
		}
		return nil, fmt.Errorf("failed to get academic record: %w", err)
	}

	var academic models.Academic
	if err := result.Content(&academic); err != nil {
		return nil, fmt.Errorf("failed to decode academic record: %w", err)
	}
	academic.ID = id

	return &academic, nil
}

// GetByStudentID retrieves academic records by student ID
func (r *AcademicRepository) GetByStudentID(studentID string) ([]models.Academic, error) {
	query := `SELECT META().id, * FROM school WHERE type = "academic" AND studentId = $1`
	
	results, err := r.db.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{studentID},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query academics by student: %w", err)
	}
	defer results.Close()

	var academics []models.Academic
	for results.Next() {
		var row struct {
			ID       string           `json:"id"`
			Academic models.Academic `json:"academic"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Academic.ID = row.ID
		academics = append(academics, row.Academic)
	}

	return academics, nil
}

// Create creates a new academic record
func (r *AcademicRepository) Create(academic *models.Academic) error {
	academic.CreatedAt = time.Now()
	academic.UpdatedAt = time.Now()
	
	// Calculate grade based on score percentage
	if academic.Grade == "" {
		academic.Grade = calculateGrade(academic.Score, academic.MaxScore)
	}
	
	// Add type field for querying
	academicData := map[string]interface{}{
		"type":      "academic",
		"studentId": academic.StudentID,
		"classId":   academic.ClassID,
		"subject":   academic.Subject,
		"examType":  academic.ExamType,
		"score":     academic.Score,
		"maxScore":  academic.MaxScore,
		"grade":     academic.Grade,
		"examDate":  academic.ExamDate,
		"semester":  academic.Semester,
		"year":      academic.Year,
		"teacherId": academic.TeacherID,
		"comments":  academic.Comments,
		"createdAt": academic.CreatedAt,
		"updatedAt": academic.UpdatedAt,
	}

	_, err := r.db.Collection.Insert(academic.ID, academicData, nil)
	if err != nil {
		return fmt.Errorf("failed to create academic record: %w", err)
	}

	return nil
}

// Update updates an existing academic record
func (r *AcademicRepository) Update(id string, academic *models.Academic) error {
	academic.UpdatedAt = time.Now()
	
	// Recalculate grade if score changed
	if academic.Grade == "" {
		academic.Grade = calculateGrade(academic.Score, academic.MaxScore)
	}
	
	academicData := map[string]interface{}{
		"type":      "academic",
		"studentId": academic.StudentID,
		"classId":   academic.ClassID,
		"subject":   academic.Subject,
		"examType":  academic.ExamType,
		"score":     academic.Score,
		"maxScore":  academic.MaxScore,
		"grade":     academic.Grade,
		"examDate":  academic.ExamDate,
		"semester":  academic.Semester,
		"year":      academic.Year,
		"teacherId": academic.TeacherID,
		"comments":  academic.Comments,
		"updatedAt": academic.UpdatedAt,
	}

	_, err := r.db.Collection.Replace(id, academicData, nil)
	if err != nil {
		return fmt.Errorf("failed to update academic record: %w", err)
	}

	return nil
}

// Delete deletes an academic record
func (r *AcademicRepository) Delete(id string) error {
	_, err := r.db.Collection.Remove(id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete academic record: %w", err)
	}

	return nil
}

// calculateGrade calculates letter grade based on score percentage
func calculateGrade(score, maxScore float64) string {
	if maxScore == 0 {
		return "F"
	}
	
	percentage := (score / maxScore) * 100
	
	switch {
	case percentage >= 90:
		return "A"
	case percentage >= 80:
		return "B"
	case percentage >= 70:
		return "C"
	case percentage >= 60:
		return "D"
	default:
		return "F"
	}
}

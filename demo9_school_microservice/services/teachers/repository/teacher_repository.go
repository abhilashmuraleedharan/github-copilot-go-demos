// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/database"
	"school-microservice/services/teachers/models"
)

// TeacherRepository handles teacher data operations
type TeacherRepository struct {
	db *database.CouchbaseClient
}

// NewTeacherRepository creates a new teacher repository
func NewTeacherRepository(db *database.CouchbaseClient) *TeacherRepository {
	return &TeacherRepository{db: db}
}

// GetAll retrieves all teachers
func (r *TeacherRepository) GetAll() ([]models.Teacher, error) {
	query := `SELECT META().id, * FROM school WHERE type = "teacher"`
	
	results, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query teachers: %w", err)
	}
	defer results.Close()

	var teachers []models.Teacher
	for results.Next() {
		var row struct {
			ID      string         `json:"id"`
			Teacher models.Teacher `json:"teacher"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Teacher.ID = row.ID
		teachers = append(teachers, row.Teacher)
	}

	return teachers, nil
}

// GetByID retrieves a teacher by ID
func (r *TeacherRepository) GetByID(id string) (*models.Teacher, error) {
	result, err := r.db.Collection.Get(id, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to get teacher: %w", err)
	}

	var teacher models.Teacher
	if err := result.Content(&teacher); err != nil {
		return nil, fmt.Errorf("failed to decode teacher: %w", err)
	}
	teacher.ID = id

	return &teacher, nil
}

// Create creates a new teacher
func (r *TeacherRepository) Create(teacher *models.Teacher) error {
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()
	
	// Add type field for querying
	teacherData := map[string]interface{}{
		"type":       "teacher",
		"firstName":  teacher.FirstName,
		"lastName":   teacher.LastName,
		"email":      teacher.Email,
		"phone":      teacher.Phone,
		"department": teacher.Department,
		"subject":    teacher.Subject,
		"hireDate":   teacher.HireDate,
		"salary":     teacher.Salary,
		"address":    teacher.Address,
		"status":     teacher.Status,
		"createdAt":  teacher.CreatedAt,
		"updatedAt":  teacher.UpdatedAt,
	}

	_, err := r.db.Collection.Insert(teacher.ID, teacherData, nil)
	if err != nil {
		return fmt.Errorf("failed to create teacher: %w", err)
	}

	return nil
}

// Update updates an existing teacher
func (r *TeacherRepository) Update(id string, teacher *models.Teacher) error {
	teacher.UpdatedAt = time.Now()
	
	teacherData := map[string]interface{}{
		"type":       "teacher",
		"firstName":  teacher.FirstName,
		"lastName":   teacher.LastName,
		"email":      teacher.Email,
		"phone":      teacher.Phone,
		"department": teacher.Department,
		"subject":    teacher.Subject,
		"hireDate":   teacher.HireDate,
		"salary":     teacher.Salary,
		"address":    teacher.Address,
		"status":     teacher.Status,
		"updatedAt":  teacher.UpdatedAt,
	}

	_, err := r.db.Collection.Replace(id, teacherData, nil)
	if err != nil {
		return fmt.Errorf("failed to update teacher: %w", err)
	}

	return nil
}

// Delete deletes a teacher
func (r *TeacherRepository) Delete(id string) error {
	_, err := r.db.Collection.Remove(id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	return nil
}

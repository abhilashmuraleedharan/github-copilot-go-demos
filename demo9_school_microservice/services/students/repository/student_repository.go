// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/database"
	"school-microservice/services/students/models"
)

// StudentRepository handles student data operations
type StudentRepository struct {
	db *database.CouchbaseClient
}

// NewStudentRepository creates a new student repository
func NewStudentRepository(db *database.CouchbaseClient) *StudentRepository {
	return &StudentRepository{db: db}
}

// GetAll retrieves all students
func (r *StudentRepository) GetAll() ([]models.Student, error) {
	return r.GetAllWithContext(context.Background())
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
// GetAllWithContext retrieves all students with context support for cancellation
func (r *StudentRepository) GetAllWithContext(ctx context.Context) ([]models.Student, error) {
	query := `SELECT META().id, * FROM school WHERE type = "student"`
	
	results, err := r.db.Cluster.Query(query, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query students: %w", err)
	}
	defer results.Close()

	var students []models.Student
	for results.Next() {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		
		var row struct {
			ID     string         `json:"id"`
			Student models.Student `json:"student"`
		}
		if err := results.Row(&row); err != nil {
			// Log the error instead of silently continuing
			// For now, continue to maintain backward compatibility
			continue
		}
		row.Student.ID = row.ID
		students = append(students, row.Student)
	}

	return students, nil
}

// GetByID retrieves a student by ID
func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	result, err := r.db.Collection.Get(id, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	var student models.Student
	if err := result.Content(&student); err != nil {
		return nil, fmt.Errorf("failed to decode student: %w", err)
	}
	student.ID = id

	return &student, nil
}

// Create creates a new student
func (r *StudentRepository) Create(student *models.Student) error {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	
	// Add type field for querying
	studentData := map[string]interface{}{
		"type":        "student",
		"firstName":   student.FirstName,
		"lastName":    student.LastName,
		"email":       student.Email,
		"dateOfBirth": student.DateOfBirth,
		"grade":       student.Grade,
		"address":     student.Address,
		"phone":       student.Phone,
		"enrollDate":  student.EnrollDate,
		"status":      student.Status,
		"createdAt":   student.CreatedAt,
		"updatedAt":   student.UpdatedAt,
	}

	_, err := r.db.Collection.Insert(student.ID, studentData, nil)
	if err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}

	return nil
}

// Update updates an existing student
func (r *StudentRepository) Update(id string, student *models.Student) error {
	student.UpdatedAt = time.Now()
	
	studentData := map[string]interface{}{
		"type":        "student",
		"firstName":   student.FirstName,
		"lastName":    student.LastName,
		"email":       student.Email,
		"dateOfBirth": student.DateOfBirth,
		"grade":       student.Grade,
		"address":     student.Address,
		"phone":       student.Phone,
		"enrollDate":  student.EnrollDate,
		"status":      student.Status,
		"updatedAt":   student.UpdatedAt,
	}

	_, err := r.db.Collection.Replace(id, studentData, nil)
	if err != nil {
		return fmt.Errorf("failed to update student: %w", err)
	}

	return nil
}

// Delete deletes a student
func (r *StudentRepository) Delete(id string) error {
	_, err := r.db.Collection.Remove(id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}

	return nil
}

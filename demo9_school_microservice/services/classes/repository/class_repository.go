// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/database"
	"school-microservice/services/classes/models"
)

// ClassRepository handles class data operations
type ClassRepository struct {
	db *database.CouchbaseClient
}

// NewClassRepository creates a new class repository
func NewClassRepository(db *database.CouchbaseClient) *ClassRepository {
	return &ClassRepository{db: db}
}

// GetAll retrieves all classes
func (r *ClassRepository) GetAll() ([]models.Class, error) {
	query := `SELECT META().id, * FROM school WHERE type = "class"`
	
	results, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query classes: %w", err)
	}
	defer results.Close()

	var classes []models.Class
	for results.Next() {
		var row struct {
			ID    string        `json:"id"`
			Class models.Class `json:"class"`
		}
		if err := results.Row(&row); err != nil {
			continue
		}
		row.Class.ID = row.ID
		classes = append(classes, row.Class)
	}

	return classes, nil
}

// GetByID retrieves a class by ID
func (r *ClassRepository) GetByID(id string) (*models.Class, error) {
	result, err := r.db.Collection.Get(id, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("class not found")
		}
		return nil, fmt.Errorf("failed to get class: %w", err)
	}

	var class models.Class
	if err := result.Content(&class); err != nil {
		return nil, fmt.Errorf("failed to decode class: %w", err)
	}
	class.ID = id

	return &class, nil
}

// Create creates a new class
func (r *ClassRepository) Create(class *models.Class) error {
	class.CreatedAt = time.Now()
	class.UpdatedAt = time.Now()
	
	// Add type field for querying
	classData := map[string]interface{}{
		"type":        "class",
		"name":        class.Name,
		"subject":     class.Subject,
		"teacherId":   class.TeacherID,
		"grade":       class.Grade,
		"room":        class.Room,
		"schedule":    class.Schedule,
		"maxStudents": class.MaxStudents,
		"studentIds":  class.StudentIDs,
		"semester":    class.Semester,
		"year":        class.Year,
		"status":      class.Status,
		"createdAt":   class.CreatedAt,
		"updatedAt":   class.UpdatedAt,
	}

	_, err := r.db.Collection.Insert(class.ID, classData, nil)
	if err != nil {
		return fmt.Errorf("failed to create class: %w", err)
	}

	return nil
}

// Update updates an existing class
func (r *ClassRepository) Update(id string, class *models.Class) error {
	class.UpdatedAt = time.Now()
	
	classData := map[string]interface{}{
		"type":        "class",
		"name":        class.Name,
		"subject":     class.Subject,
		"teacherId":   class.TeacherID,
		"grade":       class.Grade,
		"room":        class.Room,
		"schedule":    class.Schedule,
		"maxStudents": class.MaxStudents,
		"studentIds":  class.StudentIDs,
		"semester":    class.Semester,
		"year":        class.Year,
		"status":      class.Status,
		"updatedAt":   class.UpdatedAt,
	}

	_, err := r.db.Collection.Replace(id, classData, nil)
	if err != nil {
		return fmt.Errorf("failed to update class: %w", err)
	}

	return nil
}

// Delete deletes a class
func (r *ClassRepository) Delete(id string) error {
	_, err := r.db.Collection.Remove(id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete class: %w", err)
	}

	return nil
}

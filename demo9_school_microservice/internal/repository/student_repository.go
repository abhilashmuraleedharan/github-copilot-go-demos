// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package repository

import (
	"context"
	"fmt"
	"school-microservice/internal/models"
	"time"

	"github.com/couchbase/gocb/v2"
)

// studentRepository implements StudentRepository interface
type studentRepository struct {
	db *CouchbaseDB
}

// NewStudentRepository creates a new student repository
func NewStudentRepository(db *CouchbaseDB) StudentRepository {
	return &studentRepository{db: db}
}

// Create creates a new student
func (r *studentRepository) Create(ctx context.Context, student *models.Student) error {
	// Generate ID if not provided
	if student.ID == "" {
		student.ID = generateID("student")
	}

	// Set timestamps
	now := time.Now()
	student.CreatedAt = now
	student.UpdatedAt = now

	// Add document type for N1QL queries
	doc := map[string]interface{}{
		"type":          "student",
		"id":            student.ID,
		"first_name":    student.FirstName,
		"last_name":     student.LastName,
		"email":         student.Email,
		"date_of_birth": student.DateOfBirth,
		"grade":         student.Grade,
		"address":       student.Address,
		"phone":         student.Phone,
		"parent_name":   student.ParentName,
		"parent_phone":  student.ParentPhone,
		"created_at":    student.CreatedAt,
		"updated_at":    student.UpdatedAt,
	}

	return r.db.Insert(ctx, student.ID, doc)
}

// GetByID retrieves a student by ID
func (r *studentRepository) GetByID(ctx context.Context, id string) (*models.Student, error) {
	var doc map[string]interface{}
	err := r.db.Get(ctx, id, &doc)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			return nil, fmt.Errorf("student not found")
		}
		return nil, err
	}

	student, err := mapToStudent(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to map document to student: %w", err)
	}

	return student, nil
}

// GetAll retrieves all students with pagination
func (r *studentRepository) GetAll(ctx context.Context, page, pageSize int) ([]*models.Student, int, error) {
	offset := (page - 1) * pageSize

	// Count query
	countQuery := "SELECT COUNT(*) as count FROM `school` WHERE type = 'student'"
	countResult, err := r.db.Query(ctx, countQuery, nil)
	if err != nil {
		return nil, 0, err
	}

	var countRow struct {
		Count int `json:"count"`
	}
	err = countResult.One(&countRow)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf("SELECT * FROM `school` WHERE type = 'student' ORDER BY created_at DESC LIMIT %d OFFSET %d", pageSize, offset)
	result, err := r.db.Query(ctx, query, nil)
	if err != nil {
		return nil, 0, err
	}

	var students []*models.Student
	for result.Next() {
		var doc map[string]interface{}
		err := result.Row(&doc)
		if err != nil {
			continue
		}

		student, err := mapToStudent(doc["school"].(map[string]interface{}))
		if err != nil {
			continue
		}

		students = append(students, student)
	}

	return students, countRow.Count, nil
}

// Update updates a student
func (r *studentRepository) Update(ctx context.Context, id string, student *models.Student) error {
	// Check if student exists
	exists, err := r.db.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("student not found")
	}

	// Set update timestamp
	student.UpdatedAt = time.Now()
	student.ID = id

	// Update document
	doc := map[string]interface{}{
		"type":          "student",
		"id":            student.ID,
		"first_name":    student.FirstName,
		"last_name":     student.LastName,
		"email":         student.Email,
		"date_of_birth": student.DateOfBirth,
		"grade":         student.Grade,
		"address":       student.Address,
		"phone":         student.Phone,
		"parent_name":   student.ParentName,
		"parent_phone":  student.ParentPhone,
		"created_at":    student.CreatedAt,
		"updated_at":    student.UpdatedAt,
	}

	return r.db.Upsert(ctx, id, doc)
}

// Delete deletes a student
func (r *studentRepository) Delete(ctx context.Context, id string) error {
	return r.db.Remove(ctx, id)
}

// GetByGrade retrieves students by grade
func (r *studentRepository) GetByGrade(ctx context.Context, grade string) ([]*models.Student, error) {
	query := "SELECT * FROM `school` WHERE type = 'student' AND grade = $1 ORDER BY last_name, first_name"
	options := &gocb.QueryOptions{
		PositionalParameters: []interface{}{grade},
	}

	result, err := r.db.Query(ctx, query, options)
	if err != nil {
		return nil, err
	}

	var students []*models.Student
	for result.Next() {
		var doc map[string]interface{}
		err := result.Row(&doc)
		if err != nil {
			continue
		}

		student, err := mapToStudent(doc["school"].(map[string]interface{}))
		if err != nil {
			continue
		}

		students = append(students, student)
	}

	return students, nil
}

// GetByEmail retrieves a student by email
func (r *studentRepository) GetByEmail(ctx context.Context, email string) (*models.Student, error) {
	query := "SELECT * FROM `school` WHERE type = 'student' AND email = $1"
	options := &gocb.QueryOptions{
		PositionalParameters: []interface{}{email},
	}

	result, err := r.db.Query(ctx, query, options)
	if err != nil {
		return nil, err
	}

	var doc map[string]interface{}
	err = result.One(&doc)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, fmt.Errorf("student not found")
		}
		return nil, err
	}

	student, err := mapToStudent(doc["school"].(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("failed to map document to student: %w", err)
	}

	return student, nil
}

// mapToStudent maps a document to Student struct
func mapToStudent(doc map[string]interface{}) (*models.Student, error) {
	student := &models.Student{}

	if id, ok := doc["id"].(string); ok {
		student.ID = id
	}
	if firstName, ok := doc["first_name"].(string); ok {
		student.FirstName = firstName
	}
	if lastName, ok := doc["last_name"].(string); ok {
		student.LastName = lastName
	}
	if email, ok := doc["email"].(string); ok {
		student.Email = email
	}
	if grade, ok := doc["grade"].(string); ok {
		student.Grade = grade
	}
	if address, ok := doc["address"].(string); ok {
		student.Address = address
	}
	if phone, ok := doc["phone"].(string); ok {
		student.Phone = phone
	}
	if parentName, ok := doc["parent_name"].(string); ok {
		student.ParentName = parentName
	}
	if parentPhone, ok := doc["parent_phone"].(string); ok {
		student.ParentPhone = parentPhone
	}

	// Handle date fields
	if dob, ok := doc["date_of_birth"].(string); ok {
		if t, err := time.Parse(time.RFC3339, dob); err == nil {
			student.DateOfBirth = t
		}
	}
	if createdAt, ok := doc["created_at"].(string); ok {
		if t, err := time.Parse(time.RFC3339, createdAt); err == nil {
			student.CreatedAt = t
		}
	}
	if updatedAt, ok := doc["updated_at"].(string); ok {
		if t, err := time.Parse(time.RFC3339, updatedAt); err == nil {
			student.UpdatedAt = t
		}
	}

	return student, nil
}
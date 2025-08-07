package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"schoolmgmt/services/student-service/internal/models"
	"schoolmgmt/shared/pkg/database"
)

type StudentRepository struct {
	db *database.CouchbaseClient
}

func NewStudentRepository(db *database.CouchbaseClient) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(req *models.CreateStudentRequest) (*models.Student, error) {
	logrus.Infof("🔄 Creating new student: %s %s", req.FirstName, req.LastName)
	
	student := &models.Student{
		ID:          uuid.New().String(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Grade:       req.Grade,
		Address:     req.Address,
		Phone:       req.Phone,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		EnrollDate:  time.Now(),
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Type:        "student",
	}

	documentKey := fmt.Sprintf("student::%s", student.ID)
	
	logrus.Infof("💾 Inserting student document with key: %s", documentKey)
	_, err := r.db.Collection.Insert(documentKey, student, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to create student: %v", err)
		return nil, fmt.Errorf("failed to create student: %w", err)
	}

	logrus.Infof("✅ Successfully created student with ID: %s", student.ID)
	return student, nil
}

func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	documentKey := fmt.Sprintf("student::%s", id)
	
	logrus.Infof("🔍 Retrieving student with key: %s", documentKey)
	result, err := r.db.Collection.Get(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			logrus.Warnf("⚠️ Student not found with ID: %s", id)
			return nil, fmt.Errorf("student not found")
		}
		logrus.Errorf("❌ Failed to get student: %v", err)
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	var student models.Student
	err = result.Content(&student)
	if err != nil {
		logrus.Errorf("❌ Failed to decode student: %v", err)
		return nil, fmt.Errorf("failed to decode student: %w", err)
	}

	logrus.Infof("✅ Successfully retrieved student: %s %s", student.FirstName, student.LastName)
	return &student, nil
}

func (r *StudentRepository) Update(id string, req *models.UpdateStudentRequest) (*models.Student, error) {
	logrus.Infof("🔄 Updating student with ID: %s", id)
	
	student, err := r.GetByID(id)
	if err != nil {
		logrus.Errorf("❌ Failed to get student for update: %v", err)
		return nil, err
	}

	logrus.Infof("📝 Applying updates to student: %s %s", student.FirstName, student.LastName)

	// Update fields if provided
	if req.FirstName != nil {
		logrus.Infof("   → First name: %s → %s", student.FirstName, *req.FirstName)
		student.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		logrus.Infof("   → Last name: %s → %s", student.LastName, *req.LastName)
		student.LastName = *req.LastName
	}
	if req.Email != nil {
		logrus.Infof("   → Email: %s → %s", student.Email, *req.Email)
		student.Email = *req.Email
	}
	if req.DateOfBirth != nil {
		student.DateOfBirth = *req.DateOfBirth
	}
	if req.Grade != nil {
		logrus.Infof("   → Grade: %s → %s", student.Grade, *req.Grade)
		student.Grade = *req.Grade
	}
	if req.Address != nil {
		student.Address = *req.Address
	}
	if req.Phone != nil {
		student.Phone = *req.Phone
	}
	if req.ParentName != nil {
		student.ParentName = *req.ParentName
	}
	if req.ParentPhone != nil {
		student.ParentPhone = *req.ParentPhone
	}
	if req.Status != nil {
		logrus.Infof("   → Status: %s → %s", student.Status, *req.Status)
		student.Status = *req.Status
	}

	student.UpdatedAt = time.Now()

	documentKey := fmt.Sprintf("student::%s", id)
	logrus.Infof("💾 Updating student document with key: %s", documentKey)
	
	_, err = r.db.Collection.Replace(documentKey, student, nil)
	if err != nil {
		logrus.Errorf("❌ Failed to update student in Couchbase: %v", err)
		return nil, fmt.Errorf("failed to update student: %w", err)
	}

	logrus.Infof("✅ Successfully updated student: %s %s (ID: %s)", student.FirstName, student.LastName, student.ID)
	return student, nil
}

func (r *StudentRepository) Delete(id string) error {
	documentKey := fmt.Sprintf("student::%s", id)
	
	logrus.Infof("🗑️ Deleting student with key: %s", documentKey)
	
	_, err := r.db.Collection.Remove(documentKey, nil)
	if err != nil {
		if err == gocb.ErrDocumentNotFound {
			logrus.Warnf("⚠️ Student not found for deletion: %s", id)
			return fmt.Errorf("student not found")
		}
		logrus.Errorf("❌ Failed to delete student: %v", err)
		return fmt.Errorf("failed to delete student: %w", err)
	}

	logrus.Infof("✅ Successfully deleted student with ID: %s", id)
	return nil
}

func (r *StudentRepository) List(limit, offset int) ([]*models.Student, error) {
	logrus.Infof("📋 Listing students with limit=%d, offset=%d", limit, offset)
	
	query := `SELECT s.* FROM students s WHERE s.type = "student" ORDER BY s.created_at DESC LIMIT $1 OFFSET $2`
	
	logrus.Infof("🔍 Executing N1QL query: %s", query)
	
	result, err := r.db.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{limit, offset},
	})
	if err != nil {
		logrus.Errorf("❌ Failed to execute students query: %v", err)
		return nil, fmt.Errorf("failed to query students: %w", err)
	}
	defer result.Close()

	var students []*models.Student
	for result.Next() {
		var student models.Student
		err := result.Row(&student)
		if err != nil {
			logrus.Warnf("⚠️ Failed to decode student row, skipping: %v", err)
			continue // Skip invalid documents
		}
		students = append(students, &student)
	}

	logrus.Infof("✅ Successfully retrieved %d students", len(students))
	return students, nil
}

func (r *StudentRepository) GetByEmail(email string) (*models.Student, error) {
	query := `SELECT s.* FROM students s WHERE s.email = $1 AND s.type = "student" LIMIT 1`
	
	result, err := r.db.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{email},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query student by email: %w", err)
	}
	defer result.Close()

	if !result.Next() {
		return nil, fmt.Errorf("student not found")
	}

	var student models.Student
	err = result.Row(&student)
	if err != nil {
		return nil, fmt.Errorf("failed to decode student: %w", err)
	}

	return &student, nil
}

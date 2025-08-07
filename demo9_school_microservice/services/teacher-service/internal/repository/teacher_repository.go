package repository

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
)

type Teacher struct {
	ID           string   `json:"id"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Department   string   `json:"department"`
	Subjects     []string `json:"subjects"`
	Qualification string  `json:"qualification"`
	Experience   int      `json:"experience"`
	Status       string   `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Type         string   `json:"type"` // For document type identification
}

type TeacherRepository struct {
	collection *gocb.Collection
	cluster    *gocb.Cluster
	logger     *logrus.Logger
}

func NewTeacherRepository(cluster *gocb.Cluster, collection *gocb.Collection, logger *logrus.Logger) *TeacherRepository {
	return &TeacherRepository{
		collection: collection,
		cluster:    cluster,
		logger:     logger,
	}
}

func (r *TeacherRepository) Create(teacher *Teacher) error {
	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_create",
		"teacher_id": teacher.ID,
		"email":      teacher.Email,
	}).Info("Creating new teacher in Couchbase")

	// Set metadata
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()
	teacher.Type = "teacher"

	// Generate ID if not provided
	if teacher.ID == "" {
		teacher.ID = fmt.Sprintf("teacher_%d", time.Now().UnixNano())
		r.logger.WithField("generated_id", teacher.ID).Info("Generated new teacher ID")
	}

	documentKey := fmt.Sprintf("teacher::%s", teacher.ID)
	
	_, err := r.collection.Insert(documentKey, teacher, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "teacher_create",
			"teacher_id":   teacher.ID,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to create teacher in Couchbase")
		return fmt.Errorf("failed to create teacher: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "teacher_create",
		"teacher_id":   teacher.ID,
		"document_key": documentKey,
	}).Info("Successfully created teacher in Couchbase")

	return nil
}

func (r *TeacherRepository) GetByID(id string) (*Teacher, error) {
	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_get_by_id",
		"teacher_id": id,
	}).Info("Retrieving teacher from Couchbase")

	documentKey := fmt.Sprintf("teacher::%s", id)
	
	result, err := r.collection.Get(documentKey, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "teacher_get_by_id",
			"teacher_id":   id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to retrieve teacher from Couchbase")
		return nil, fmt.Errorf("failed to get teacher: %w", err)
	}

	var teacher Teacher
	if err := result.Content(&teacher); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "teacher_get_by_id",
			"teacher_id":   id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to decode teacher document")
		return nil, fmt.Errorf("failed to decode teacher: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "teacher_get_by_id",
		"teacher_id":   id,
		"document_key": documentKey,
	}).Info("Successfully retrieved teacher from Couchbase")

	return &teacher, nil
}

func (r *TeacherRepository) Update(id string, teacher *Teacher) error {
	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_update",
		"teacher_id": id,
	}).Info("Updating teacher in Couchbase")

	// Ensure the ID matches
	teacher.ID = id
	teacher.UpdatedAt = time.Now()
	teacher.Type = "teacher"

	documentKey := fmt.Sprintf("teacher::%s", id)
	
	_, err := r.collection.Replace(documentKey, teacher, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "teacher_update",
			"teacher_id":   id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to update teacher in Couchbase")
		return fmt.Errorf("failed to update teacher: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "teacher_update",
		"teacher_id":   id,
		"document_key": documentKey,
	}).Info("Successfully updated teacher in Couchbase")

	return nil
}

func (r *TeacherRepository) Delete(id string) error {
	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_delete",
		"teacher_id": id,
	}).Info("Deleting teacher from Couchbase")

	documentKey := fmt.Sprintf("teacher::%s", id)
	
	_, err := r.collection.Remove(documentKey, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":    "teacher_delete",
			"teacher_id":   id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to delete teacher from Couchbase")
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":    "teacher_delete",
		"teacher_id":   id,
		"document_key": documentKey,
	}).Info("Successfully deleted teacher from Couchbase")

	return nil
}

func (r *TeacherRepository) List(limit, offset int) ([]*Teacher, error) {
	r.logger.WithFields(logrus.Fields{
		"operation": "teacher_list",
		"limit":     limit,
		"offset":    offset,
	}).Info("Listing teachers from Couchbase")

	if limit == 0 {
		limit = 50 // Default limit
	}

	query := fmt.Sprintf(
		`SELECT t.* FROM schoolmgmt t WHERE t.type = "teacher" LIMIT %d OFFSET %d`,
		limit, offset,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for teachers")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "teacher_list",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for teachers")
		return nil, fmt.Errorf("failed to query teachers: %w", err)
	}
	defer result.Close()

	var teachers []*Teacher
	for result.Next() {
		var teacher Teacher
		if err := result.Row(&teacher); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation": "teacher_list",
				"error":     err.Error(),
			}).Error("Failed to decode teacher row")
			continue
		}
		teachers = append(teachers, &teacher)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "teacher_list",
			"error":     err.Error(),
		}).Error("Error during teacher query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation": "teacher_list",
		"count":     len(teachers),
		"limit":     limit,
		"offset":    offset,
	}).Info("Successfully retrieved teachers from Couchbase")

	return teachers, nil
}

func (r *TeacherRepository) GetByDepartment(department string) ([]*Teacher, error) {
	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_get_by_department",
		"department": department,
	}).Info("Retrieving teachers by department from Couchbase")

	query := fmt.Sprintf(
		`SELECT t.* FROM schoolmgmt t WHERE t.type = "teacher" AND t.department = "%s"`,
		department,
	)

	r.logger.WithField("query", query).Debug("Executing N1QL query for teachers by department")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":  "teacher_get_by_department",
			"department": department,
			"query":      query,
			"error":      err.Error(),
		}).Error("Failed to execute N1QL query for teachers by department")
		return nil, fmt.Errorf("failed to query teachers by department: %w", err)
	}
	defer result.Close()

	var teachers []*Teacher
	for result.Next() {
		var teacher Teacher
		if err := result.Row(&teacher); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation":  "teacher_get_by_department",
				"department": department,
				"error":      err.Error(),
			}).Error("Failed to decode teacher row")
			continue
		}
		teachers = append(teachers, &teacher)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation":  "teacher_get_by_department",
			"department": department,
			"error":      err.Error(),
		}).Error("Error during teacher query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation":  "teacher_get_by_department",
		"department": department,
		"count":      len(teachers),
	}).Info("Successfully retrieved teachers by department from Couchbase")

	return teachers, nil
}

func (r *TeacherRepository) GetActiveTeachers() ([]*Teacher, error) {
	r.logger.WithField("operation", "teacher_get_active").Info("Retrieving active teachers from Couchbase")

	query := `SELECT t.* FROM schoolmgmt t WHERE t.type = "teacher" AND t.status = "active"`

	r.logger.WithField("query", query).Debug("Executing N1QL query for active teachers")

	result, err := r.cluster.Query(query, nil)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "teacher_get_active",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for active teachers")
		return nil, fmt.Errorf("failed to query active teachers: %w", err)
	}
	defer result.Close()

	var teachers []*Teacher
	for result.Next() {
		var teacher Teacher
		if err := result.Row(&teacher); err != nil {
			r.logger.WithFields(logrus.Fields{
				"operation": "teacher_get_active",
				"error":     err.Error(),
			}).Error("Failed to decode teacher row")
			continue
		}
		teachers = append(teachers, &teacher)
	}

	if err := result.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"operation": "teacher_get_active",
			"error":     err.Error(),
		}).Error("Error during active teachers query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	r.logger.WithFields(logrus.Fields{
		"operation": "teacher_get_active",
		"count":     len(teachers),
	}).Info("Successfully retrieved active teachers from Couchbase")

	return teachers, nil
}

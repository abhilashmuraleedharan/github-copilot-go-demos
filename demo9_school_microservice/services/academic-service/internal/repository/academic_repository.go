package repository

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"schoolmgmt/shared/pkg/database"
)

type Academic struct {
	ID            string    `json:"id"`
	StudentID     string    `json:"student_id"`
	TeacherID     string    `json:"teacher_id"`
	Subject       string    `json:"subject"`
	Grade         string    `json:"grade"`
	Semester      string    `json:"semester"`
	AcademicYear  string    `json:"academic_year"`
	ExamType      string    `json:"exam_type"`
	ExamDate      time.Time `json:"exam_date"`
	MaxMarks      float64   `json:"max_marks"`
	ObtainedMarks float64   `json:"obtained_marks"`
	Percentage    float64   `json:"percentage"`
	Status        string    `json:"status"`
	Remarks       string    `json:"remarks"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Type          string    `json:"type"` // For document type identification
}

type Class struct {
	ID           string    `json:"id"`
	ClassName    string    `json:"class_name"`
	Grade        string    `json:"grade"`
	Section      string    `json:"section"`
	TeacherID    string    `json:"teacher_id"`
	Subject      string    `json:"subject"`
	AcademicYear string    `json:"academic_year"`
	Semester     string    `json:"semester"`
	StudentIDs   []string  `json:"student_ids"`
	MaxCapacity  int       `json:"max_capacity"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Type         string    `json:"type"` // For document type identification
}

type AcademicRepository struct {
	db *database.CouchbaseClient
}

func NewAcademicRepository(db *database.CouchbaseClient) *AcademicRepository {
	return &AcademicRepository{
		db: db,
	}
}

// Academic CRUD operations
func (r *AcademicRepository) CreateAcademic(academic *Academic) error {
	logrus.WithFields(logrus.Fields{
		"operation":   "academic_create",
		"academic_id": academic.ID,
		"student_id":  academic.StudentID,
		"subject":     academic.Subject,
	}).Info("Creating new academic record in Couchbase")

	// Set metadata
	academic.CreatedAt = time.Now()
	academic.UpdatedAt = time.Now()
	academic.Type = "academic"

	// Generate ID if not provided
	if academic.ID == "" {
		academic.ID = fmt.Sprintf("academic_%d", time.Now().UnixNano())
		logrus.WithField("generated_id", academic.ID).Info("Generated new academic ID")
	}

	// Calculate percentage if not provided
	if academic.Percentage == 0 && academic.MaxMarks > 0 && academic.ObtainedMarks >= 0 {
		academic.Percentage = (academic.ObtainedMarks / academic.MaxMarks) * 100
	}

	documentKey := fmt.Sprintf("academic::%s", academic.ID)
	
	_, err := r.db.Collection.Insert(documentKey, academic, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "academic_create",
			"academic_id":  academic.ID,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to create academic record in Couchbase")
		return fmt.Errorf("failed to create academic record: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "academic_create",
		"academic_id":  academic.ID,
		"document_key": documentKey,
	}).Info("Successfully created academic record in Couchbase")

	return nil
}

func (r *AcademicRepository) GetAcademicByID(id string) (*Academic, error) {
	logrus.WithFields(logrus.Fields{
		"operation":   "academic_get_by_id",
		"academic_id": id,
	}).Info("Retrieving academic record from Couchbase")

	documentKey := fmt.Sprintf("academic::%s", id)
	
	result, err := r.db.Collection.Get(documentKey, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "academic_get_by_id",
			"academic_id":  id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to retrieve academic record from Couchbase")
		return nil, fmt.Errorf("failed to get academic record: %w", err)
	}

	var academic Academic
	if err := result.Content(&academic); err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "academic_get_by_id",
			"academic_id":  id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to decode academic document")
		return nil, fmt.Errorf("failed to decode academic record: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "academic_get_by_id",
		"academic_id":  id,
		"document_key": documentKey,
	}).Info("Successfully retrieved academic record from Couchbase")

	return &academic, nil
}

func (r *AcademicRepository) UpdateAcademic(id string, academic *Academic) error {
	logrus.WithFields(logrus.Fields{
		"operation":   "academic_update",
		"academic_id": id,
	}).Info("Updating academic record in Couchbase")

	// Ensure the ID matches
	academic.ID = id
	academic.UpdatedAt = time.Now()
	academic.Type = "academic"

	// Recalculate percentage if needed
	if academic.MaxMarks > 0 && academic.ObtainedMarks >= 0 {
		academic.Percentage = (academic.ObtainedMarks / academic.MaxMarks) * 100
	}

	documentKey := fmt.Sprintf("academic::%s", id)
	
	_, err := r.db.Collection.Replace(documentKey, academic, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "academic_update",
			"academic_id":  id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to update academic record in Couchbase")
		return fmt.Errorf("failed to update academic record: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "academic_update",
		"academic_id":  id,
		"document_key": documentKey,
	}).Info("Successfully updated academic record in Couchbase")

	return nil
}

func (r *AcademicRepository) DeleteAcademic(id string) error {
	logrus.WithFields(logrus.Fields{
		"operation":   "academic_delete",
		"academic_id": id,
	}).Info("Deleting academic record from Couchbase")

	documentKey := fmt.Sprintf("academic::%s", id)
	
	_, err := r.db.Collection.Remove(documentKey, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "academic_delete",
			"academic_id":  id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to delete academic record from Couchbase")
		return fmt.Errorf("failed to delete academic record: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "academic_delete",
		"academic_id":  id,
		"document_key": documentKey,
	}).Info("Successfully deleted academic record from Couchbase")

	return nil
}

func (r *AcademicRepository) ListAcademics(limit, offset int) ([]*Academic, error) {
	logrus.WithFields(logrus.Fields{
		"operation": "academic_list",
		"limit":     limit,
		"offset":    offset,
	}).Info("Listing academic records from Couchbase")

	if limit == 0 {
		limit = 50 // Default limit
	}

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "academic" LIMIT %d OFFSET %d`,
		limit, offset,
	)

	logrus.WithField("query", query).Debug("Executing N1QL query for academics")

	result, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation": "academic_list",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for academics")
		return nil, fmt.Errorf("failed to query academics: %w", err)
	}
	defer result.Close()

	var academics []*Academic
	for result.Next() {
		var academic Academic
		if err := result.Row(&academic); err != nil {
			logrus.WithFields(logrus.Fields{
				"operation": "academic_list",
				"error":     err.Error(),
			}).Error("Failed to decode academic row")
			continue
		}
		academics = append(academics, &academic)
	}

	if err := result.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"operation": "academic_list",
			"error":     err.Error(),
		}).Error("Error during academic query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation": "academic_list",
		"count":     len(academics),
		"limit":     limit,
		"offset":    offset,
	}).Info("Successfully retrieved academic records from Couchbase")

	return academics, nil
}

func (r *AcademicRepository) GetAcademicsByStudent(studentID string) ([]*Academic, error) {
	logrus.WithFields(logrus.Fields{
		"operation":  "academic_get_by_student",
		"student_id": studentID,
	}).Info("Retrieving academic records by student from Couchbase")

	query := fmt.Sprintf(
		`SELECT a.* FROM schoolmgmt a WHERE a.type = "academic" AND a.student_id = "%s"`,
		studentID,
	)

	logrus.WithField("query", query).Debug("Executing N1QL query for academics by student")

	result, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":  "academic_get_by_student",
			"student_id": studentID,
			"query":      query,
			"error":      err.Error(),
		}).Error("Failed to execute N1QL query for academics by student")
		return nil, fmt.Errorf("failed to query academics by student: %w", err)
	}
	defer result.Close()

	var academics []*Academic
	for result.Next() {
		var academic Academic
		if err := result.Row(&academic); err != nil {
			logrus.WithFields(logrus.Fields{
				"operation":  "academic_get_by_student",
				"student_id": studentID,
				"error":      err.Error(),
			}).Error("Failed to decode academic row")
			continue
		}
		academics = append(academics, &academic)
	}

	if err := result.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":  "academic_get_by_student",
			"student_id": studentID,
			"error":      err.Error(),
		}).Error("Error during academic query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":  "academic_get_by_student",
		"student_id": studentID,
		"count":      len(academics),
	}).Info("Successfully retrieved academic records by student from Couchbase")

	return academics, nil
}

// Class CRUD operations
func (r *AcademicRepository) CreateClass(class *Class) error {
	logrus.WithFields(logrus.Fields{
		"operation":  "class_create",
		"class_id":   class.ID,
		"class_name": class.ClassName,
	}).Info("Creating new class in Couchbase")

	// Set metadata
	class.CreatedAt = time.Now()
	class.UpdatedAt = time.Now()
	class.Type = "class"

	// Generate ID if not provided
	if class.ID == "" {
		class.ID = fmt.Sprintf("class_%d", time.Now().UnixNano())
		logrus.WithField("generated_id", class.ID).Info("Generated new class ID")
	}

	documentKey := fmt.Sprintf("class::%s", class.ID)
	
	_, err := r.db.Collection.Insert(documentKey, class, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "class_create",
			"class_id":     class.ID,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to create class in Couchbase")
		return fmt.Errorf("failed to create class: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "class_create",
		"class_id":     class.ID,
		"document_key": documentKey,
	}).Info("Successfully created class in Couchbase")

	return nil
}

func (r *AcademicRepository) GetClassByID(id string) (*Class, error) {
	logrus.WithFields(logrus.Fields{
		"operation": "class_get_by_id",
		"class_id":  id,
	}).Info("Retrieving class from Couchbase")

	documentKey := fmt.Sprintf("class::%s", id)
	
	result, err := r.db.Collection.Get(documentKey, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "class_get_by_id",
			"class_id":     id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to retrieve class from Couchbase")
		return nil, fmt.Errorf("failed to get class: %w", err)
	}

	var class Class
	if err := result.Content(&class); err != nil {
		logrus.WithFields(logrus.Fields{
			"operation":    "class_get_by_id",
			"class_id":     id,
			"document_key": documentKey,
			"error":        err.Error(),
		}).Error("Failed to decode class document")
		return nil, fmt.Errorf("failed to decode class: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation":    "class_get_by_id",
		"class_id":     id,
		"document_key": documentKey,
	}).Info("Successfully retrieved class from Couchbase")

	return &class, nil
}

func (r *AcademicRepository) ListClasses(limit, offset int) ([]*Class, error) {
	logrus.WithFields(logrus.Fields{
		"operation": "class_list",
		"limit":     limit,
		"offset":    offset,
	}).Info("Listing classes from Couchbase")

	if limit == 0 {
		limit = 50 // Default limit
	}

	query := fmt.Sprintf(
		`SELECT c.* FROM schoolmgmt c WHERE c.type = "class" LIMIT %d OFFSET %d`,
		limit, offset,
	)

	logrus.WithField("query", query).Debug("Executing N1QL query for classes")

	result, err := r.db.Cluster.Query(query, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"operation": "class_list",
			"query":     query,
			"error":     err.Error(),
		}).Error("Failed to execute N1QL query for classes")
		return nil, fmt.Errorf("failed to query classes: %w", err)
	}
	defer result.Close()

	var classes []*Class
	for result.Next() {
		var class Class
		if err := result.Row(&class); err != nil {
			logrus.WithFields(logrus.Fields{
				"operation": "class_list",
				"error":     err.Error(),
			}).Error("Failed to decode class row")
			continue
		}
		classes = append(classes, &class)
	}

	if err := result.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"operation": "class_list",
			"error":     err.Error(),
		}).Error("Error during class query iteration")
		return nil, fmt.Errorf("query iteration error: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"operation": "class_list",
		"count":     len(classes),
		"limit":     limit,
		"offset":    offset,
	}).Info("Successfully retrieved classes from Couchbase")

	return classes, nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/demo/school-microservice/models"
)

// Repository provides generic CRUD operations
type Repository interface {
	Create(ctx context.Context, id string, doc interface{}) error
	Get(ctx context.Context, id string, doc interface{}) error
	Update(ctx context.Context, id string, doc interface{}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, docType string, limit int) ([]interface{}, error)
}

// CouchbaseRepository implements Repository using Couchbase
type CouchbaseRepository struct {
	collection *gocb.Collection
	cluster    *gocb.Cluster
	bucket     *gocb.Bucket
}

// NewCouchbaseRepository creates a new Couchbase repository
func NewCouchbaseRepository(cluster *gocb.Cluster, bucketName, scopeName, collectionName string) (*CouchbaseRepository, error) {
	bucket := cluster.Bucket(bucketName)
	
	// Wait for bucket to be ready
	err := bucket.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for bucket: %w", err)
	}

	collection := bucket.Scope(scopeName).Collection(collectionName)
	
	return &CouchbaseRepository{
		collection: collection,
		cluster:    cluster,
		bucket:     bucket,
	}, nil
}

// Create inserts a new document
func (r *CouchbaseRepository) Create(ctx context.Context, id string, doc interface{}) error {
	_, err := r.collection.Insert(id, doc, &gocb.InsertOptions{})
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	return nil
}

// Get retrieves a document by ID
func (r *CouchbaseRepository) Get(ctx context.Context, id string, doc interface{}) error {
	result, err := r.collection.Get(id, &gocb.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get document: %w", err)
	}
	
	err = result.Content(doc)
	if err != nil {
		return fmt.Errorf("failed to decode document: %w", err)
	}
	
	return nil
}

// Update modifies an existing document
func (r *CouchbaseRepository) Update(ctx context.Context, id string, doc interface{}) error {
	_, err := r.collection.Replace(id, doc, &gocb.ReplaceOptions{})
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}
	return nil
}

// Delete removes a document
func (r *CouchbaseRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.Remove(id, &gocb.RemoveOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	return nil
}

// List retrieves all documents of a specific type
func (r *CouchbaseRepository) List(ctx context.Context, docType string, limit int) ([]interface{}, error) {
	// For simplicity, we'll use a basic query
	// In production, you'd want proper indexing and pagination
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE type = $1 LIMIT $2", r.bucket.Name())
	
	results, err := r.cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{docType, limit},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query documents: %w", err)
	}
	defer results.Close()
	
	var docs []interface{}
	for results.Next() {
		var row map[string]interface{}
		err := results.Row(&row)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row: %w", err)
		}
		docs = append(docs, row)
	}
	
	if err := results.Err(); err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	
	return docs, nil
}

// StudentRepository provides student-specific operations
type StudentRepository struct {
	repo *CouchbaseRepository
}

// NewStudentRepository creates a new student repository
func NewStudentRepository(repo *CouchbaseRepository) *StudentRepository {
	return &StudentRepository{repo: repo}
}

// CreateStudent creates a new student
func (r *StudentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	student.Type = "student"
	return r.repo.Create(ctx, student.ID, student)
}

// GetStudent retrieves a student by ID
func (r *StudentRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	err := r.repo.Get(ctx, id, &student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// UpdateStudent updates a student
func (r *StudentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	student.Type = "student"
	return r.repo.Update(ctx, student.ID, student)
}

// DeleteStudent deletes a student
func (r *StudentRepository) DeleteStudent(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// TeacherRepository provides teacher-specific operations
type TeacherRepository struct {
	repo *CouchbaseRepository
}

// NewTeacherRepository creates a new teacher repository
func NewTeacherRepository(repo *CouchbaseRepository) *TeacherRepository {
	return &TeacherRepository{repo: repo}
}

// CreateTeacher creates a new teacher
func (r *TeacherRepository) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	teacher.Type = "teacher"
	return r.repo.Create(ctx, teacher.ID, teacher)
}

// GetTeacher retrieves a teacher by ID
func (r *TeacherRepository) GetTeacher(ctx context.Context, id string) (*models.Teacher, error) {
	var teacher models.Teacher
	err := r.repo.Get(ctx, id, &teacher)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

// UpdateTeacher updates a teacher
func (r *TeacherRepository) UpdateTeacher(ctx context.Context, teacher *models.Teacher) error {
	teacher.Type = "teacher"
	return r.repo.Update(ctx, teacher.ID, teacher)
}

// DeleteTeacher deletes a teacher
func (r *TeacherRepository) DeleteTeacher(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// ClassRepository provides class-specific operations
type ClassRepository struct {
	repo *CouchbaseRepository
}

// NewClassRepository creates a new class repository
func NewClassRepository(repo *CouchbaseRepository) *ClassRepository {
	return &ClassRepository{repo: repo}
}

// CreateClass creates a new class
func (r *ClassRepository) CreateClass(ctx context.Context, class *models.Class) error {
	class.Type = "class"
	return r.repo.Create(ctx, class.ID, class)
}

// GetClass retrieves a class by ID
func (r *ClassRepository) GetClass(ctx context.Context, id string) (*models.Class, error) {
	var class models.Class
	err := r.repo.Get(ctx, id, &class)
	if err != nil {
		return nil, err
	}
	return &class, nil
}

// UpdateClass updates a class
func (r *ClassRepository) UpdateClass(ctx context.Context, class *models.Class) error {
	class.Type = "class"
	return r.repo.Update(ctx, class.ID, class)
}

// DeleteClass deletes a class
func (r *ClassRepository) DeleteClass(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// AcademicRepository provides academic enrollment operations
type AcademicRepository struct {
	repo *CouchbaseRepository
}

// NewAcademicRepository creates a new academic repository
func NewAcademicRepository(repo *CouchbaseRepository) *AcademicRepository {
	return &AcademicRepository{repo: repo}
}

// CreateAcademic creates a new academic enrollment
func (r *AcademicRepository) CreateAcademic(ctx context.Context, academic *models.Academic) error {
	academic.Type = "academic"
	return r.repo.Create(ctx, academic.ID, academic)
}

// GetAcademic retrieves an academic enrollment by ID
func (r *AcademicRepository) GetAcademic(ctx context.Context, id string) (*models.Academic, error) {
	var academic models.Academic
	err := r.repo.Get(ctx, id, &academic)
	if err != nil {
		return nil, err
	}
	return &academic, nil
}

// UpdateAcademic updates an academic enrollment
func (r *AcademicRepository) UpdateAcademic(ctx context.Context, academic *models.Academic) error {
	academic.Type = "academic"
	return r.repo.Update(ctx, academic.ID, academic)
}

// DeleteAcademic deletes an academic enrollment
func (r *AcademicRepository) DeleteAcademic(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// ExamRepository provides exam operations
type ExamRepository struct {
	repo *CouchbaseRepository
}

// NewExamRepository creates a new exam repository
func NewExamRepository(repo *CouchbaseRepository) *ExamRepository {
	return &ExamRepository{repo: repo}
}

// CreateExam creates a new exam
func (r *ExamRepository) CreateExam(ctx context.Context, exam *models.Exam) error {
	exam.Type = "exam"
	return r.repo.Create(ctx, exam.ID, exam)
}

// GetExam retrieves an exam by ID
func (r *ExamRepository) GetExam(ctx context.Context, id string) (*models.Exam, error) {
	var exam models.Exam
	err := r.repo.Get(ctx, id, &exam)
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

// UpdateExam updates an exam
func (r *ExamRepository) UpdateExam(ctx context.Context, exam *models.Exam) error {
	exam.Type = "exam"
	return r.repo.Update(ctx, exam.ID, exam)
}

// DeleteExam deletes an exam
func (r *ExamRepository) DeleteExam(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// ExamResultRepository provides exam result operations
type ExamResultRepository struct {
	repo *CouchbaseRepository
}

// NewExamResultRepository creates a new exam result repository
func NewExamResultRepository(repo *CouchbaseRepository) *ExamResultRepository {
	return &ExamResultRepository{repo: repo}
}

// CreateExamResult creates a new exam result
func (r *ExamResultRepository) CreateExamResult(ctx context.Context, result *models.ExamResult) error {
	result.Type = "examResult"
	return r.repo.Create(ctx, result.ID, result)
}

// GetExamResult retrieves an exam result by ID
func (r *ExamResultRepository) GetExamResult(ctx context.Context, id string) (*models.ExamResult, error) {
	var result models.ExamResult
	err := r.repo.Get(ctx, id, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateExamResult updates an exam result
func (r *ExamResultRepository) UpdateExamResult(ctx context.Context, result *models.ExamResult) error {
	result.Type = "examResult"
	return r.repo.Update(ctx, result.ID, result)
}

// DeleteExamResult deletes an exam result
func (r *ExamResultRepository) DeleteExamResult(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

// AchievementRepository provides achievement operations
type AchievementRepository struct {
	repo *CouchbaseRepository
}

// NewAchievementRepository creates a new achievement repository
func NewAchievementRepository(repo *CouchbaseRepository) *AchievementRepository {
	return &AchievementRepository{repo: repo}
}

// CreateAchievement creates a new achievement
func (r *AchievementRepository) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
	achievement.Type = "achievement"
	return r.repo.Create(ctx, achievement.ID, achievement)
}

// GetAchievement retrieves an achievement by ID
func (r *AchievementRepository) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
	var achievement models.Achievement
	err := r.repo.Get(ctx, id, &achievement)
	if err != nil {
		return nil, err
	}
	return &achievement, nil
}

// UpdateAchievement updates an achievement
func (r *AchievementRepository) UpdateAchievement(ctx context.Context, achievement *models.Achievement) error {
	achievement.Type = "achievement"
	return r.repo.Update(ctx, achievement.ID, achievement)
}

// DeleteAchievement deletes an achievement
func (r *AchievementRepository) DeleteAchievement(ctx context.Context, id string) error {
	return r.repo.Delete(ctx, id)
}

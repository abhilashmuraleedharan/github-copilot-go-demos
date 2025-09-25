// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"school-microservice/internal/models"
	"school-microservice/internal/repository"
	"strings"
	"time"
)

// StudentService defines the interface for student business logic
type StudentService interface {
	CreateStudent(ctx context.Context, req *models.CreateStudentRequest) (*models.Student, error)
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	GetAllStudents(ctx context.Context, page, pageSize int) ([]*models.Student, int, error)
	UpdateStudent(ctx context.Context, id string, req *models.CreateStudentRequest) (*models.Student, error)
	DeleteStudent(ctx context.Context, id string) error
	GetStudentsByGrade(ctx context.Context, grade string) ([]*models.Student, error)
	GetStudentByEmail(ctx context.Context, email string) (*models.Student, error)
}

// studentService implements StudentService interface
type studentService struct {
	repo *repository.Repository
}

// NewStudentService creates a new student service
func NewStudentService(repo *repository.Repository) StudentService {
	return &studentService{repo: repo}
}

// CreateStudent creates a new student with validation
func (s *studentService) CreateStudent(ctx context.Context, req *models.CreateStudentRequest) (*models.Student, error) {
	// Validate request
	if err := s.validateStudentRequest(req); err != nil {
		return nil, err
	}

	// Check if email already exists
	existingStudent, err := s.repo.Student.GetByEmail(ctx, req.Email)
	if err == nil && existingStudent != nil {
		return nil, fmt.Errorf("student with email %s already exists", req.Email)
	}

	// Create student model
	student := &models.Student{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Grade:       req.Grade,
		Address:     req.Address,
		Phone:       req.Phone,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
	}

	// Save to repository
	if err := s.repo.Student.Create(ctx, student); err != nil {
		return nil, fmt.Errorf("failed to create student: %w", err)
	}

	return student, nil
}

// GetStudent retrieves a student by ID
func (s *studentService) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("student ID is required")
	}

	return s.repo.Student.GetByID(ctx, id)
}

// GetAllStudents retrieves all students with pagination
func (s *studentService) GetAllStudents(ctx context.Context, page, pageSize int) ([]*models.Student, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.repo.Student.GetAll(ctx, page, pageSize)
}

// UpdateStudent updates a student
func (s *studentService) UpdateStudent(ctx context.Context, id string, req *models.CreateStudentRequest) (*models.Student, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("student ID is required")
	}

	// Validate request
	if err := s.validateStudentRequest(req); err != nil {
		return nil, err
	}

	// Check if student exists
	existingStudent, err := s.repo.Student.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("student not found: %w", err)
	}

	// Check if email is being changed and if new email already exists
	if req.Email != existingStudent.Email {
		emailStudent, err := s.repo.Student.GetByEmail(ctx, req.Email)
		if err == nil && emailStudent != nil {
			return nil, fmt.Errorf("student with email %s already exists", req.Email)
		}
	}

	// Update student model
	student := &models.Student{
		ID:          id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Grade:       req.Grade,
		Address:     req.Address,
		Phone:       req.Phone,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		CreatedAt:   existingStudent.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	// Update in repository
	if err := s.repo.Student.Update(ctx, id, student); err != nil {
		return nil, fmt.Errorf("failed to update student: %w", err)
	}

	return student, nil
}

// DeleteStudent deletes a student
func (s *studentService) DeleteStudent(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("student ID is required")
	}

	// Check if student exists
	_, err := s.repo.Student.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("student not found: %w", err)
	}

	// Delete student
	if err := s.repo.Student.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}

	return nil
}

// GetStudentsByGrade retrieves students by grade
func (s *studentService) GetStudentsByGrade(ctx context.Context, grade string) ([]*models.Student, error) {
	if strings.TrimSpace(grade) == "" {
		return nil, errors.New("grade is required")
	}

	return s.repo.Student.GetByGrade(ctx, grade)
}

// GetStudentByEmail retrieves a student by email
func (s *studentService) GetStudentByEmail(ctx context.Context, email string) (*models.Student, error) {
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}

	return s.repo.Student.GetByEmail(ctx, email)
}

// validateStudentRequest validates student creation/update request
func (s *studentService) validateStudentRequest(req *models.CreateStudentRequest) error {
	// Required fields
	if strings.TrimSpace(req.FirstName) == "" {
		return errors.New("first name is required")
	}
	if strings.TrimSpace(req.LastName) == "" {
		return errors.New("last name is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Grade) == "" {
		return errors.New("grade is required")
	}

	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	// Date of birth validation
	if req.DateOfBirth.IsZero() {
		return errors.New("date of birth is required")
	}
	if req.DateOfBirth.After(time.Now()) {
		return errors.New("date of birth cannot be in the future")
	}

	// Age validation (must be between 3 and 25 years old)
	age := time.Since(req.DateOfBirth).Hours() / (24 * 365.25)
	if age < 3 || age > 25 {
		return errors.New("student age must be between 3 and 25 years")
	}

	// Grade validation
	validGrades := []string{"KG", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	isValidGrade := false
	for _, grade := range validGrades {
		if req.Grade == grade {
			isValidGrade = true
			break
		}
	}
	if !isValidGrade {
		return errors.New("invalid grade: must be KG or 1-12")
	}

	// Phone validation (optional but if provided should be valid)
	if req.Phone != "" && len(req.Phone) < 10 {
		return errors.New("phone number must be at least 10 digits")
	}
	if req.ParentPhone != "" && len(req.ParentPhone) < 10 {
		return errors.New("parent phone number must be at least 10 digits")
	}

	return nil
}

// Service aggregates all business logic services
type Service struct {
	Student     StudentService
	Teacher     TeacherService
	Class       ClassService
	Academic    AcademicService
	Achievement AchievementService
}

// NewService creates a new service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Student:     NewStudentService(repo),
		Teacher:     NewTeacherService(repo),
		Class:       NewClassService(repo),
		Academic:    NewAcademicService(repo),
		Achievement: NewAchievementService(repo),
	}
}

// Placeholder service interfaces and implementations
type TeacherService interface {
	CreateTeacher(ctx context.Context, req *models.CreateTeacherRequest) (*models.Teacher, error)
	GetTeacher(ctx context.Context, id string) (*models.Teacher, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, req *models.CreateClassRequest) (*models.Class, error)
	GetClass(ctx context.Context, id string) (*models.Class, error)
}

type AcademicService interface {
	CreateAcademic(ctx context.Context, req *models.CreateAcademicRequest) (*models.Academic, error)
	GetAcademic(ctx context.Context, id string) (*models.Academic, error)
}

type AchievementService interface {
	CreateAchievement(ctx context.Context, req *models.CreateAchievementRequest) (*models.Achievement, error)
	GetAchievement(ctx context.Context, id string) (*models.Achievement, error)
}

// Placeholder implementations
func NewTeacherService(repo *repository.Repository) TeacherService {
	return &teacherService{repo: repo}
}

func NewClassService(repo *repository.Repository) ClassService {
	return &classService{repo: repo}
}

func NewAcademicService(repo *repository.Repository) AcademicService {
	return &academicService{repo: repo}
}

func NewAchievementService(repo *repository.Repository) AchievementService {
	return &achievementService{repo: repo}
}

type teacherService struct{ repo *repository.Repository }
type classService struct{ repo *repository.Repository }
type academicService struct{ repo *repository.Repository }
type achievementService struct{ repo *repository.Repository }

func (s *teacherService) CreateTeacher(ctx context.Context, req *models.CreateTeacherRequest) (*models.Teacher, error) {
	return &models.Teacher{}, nil
}
func (s *teacherService) GetTeacher(ctx context.Context, id string) (*models.Teacher, error) {
	return &models.Teacher{}, nil
}
func (s *classService) CreateClass(ctx context.Context, req *models.CreateClassRequest) (*models.Class, error) {
	return &models.Class{}, nil
}
func (s *classService) GetClass(ctx context.Context, id string) (*models.Class, error) {
	return &models.Class{}, nil
}
func (s *academicService) CreateAcademic(ctx context.Context, req *models.CreateAcademicRequest) (*models.Academic, error) {
	return &models.Academic{}, nil
}
func (s *academicService) GetAcademic(ctx context.Context, id string) (*models.Academic, error) {
	return &models.Academic{}, nil
}
func (s *achievementService) CreateAchievement(ctx context.Context, req *models.CreateAchievementRequest) (*models.Achievement, error) {
	return &models.Achievement{}, nil
}
func (s *achievementService) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
	return &models.Achievement{}, nil
}
// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/demo/school-microservice/models"
	"github.com/demo/school-microservice/repository"
)

// Service provides business logic operations
type Service struct {
	studentRepo     *repository.StudentRepository
	teacherRepo     *repository.TeacherRepository
	classRepo       *repository.ClassRepository
	academicRepo    *repository.AcademicRepository
	examRepo        *repository.ExamRepository
	examResultRepo  *repository.ExamResultRepository
	achievementRepo *repository.AchievementRepository
}

// NewService creates a new service instance
func NewService(
	studentRepo *repository.StudentRepository,
	teacherRepo *repository.TeacherRepository,
	classRepo *repository.ClassRepository,
	academicRepo *repository.AcademicRepository,
	examRepo *repository.ExamRepository,
	examResultRepo *repository.ExamResultRepository,
	achievementRepo *repository.AchievementRepository,
) *Service {
	return &Service{
		studentRepo:     studentRepo,
		teacherRepo:     teacherRepo,
		classRepo:       classRepo,
		academicRepo:    academicRepo,
		examRepo:        examRepo,
		examResultRepo:  examResultRepo,
		achievementRepo: achievementRepo,
	}
}

// Student operations

// CreateStudent creates a new student with validation
func (s *Service) CreateStudent(ctx context.Context, student *models.Student) error {
	if student.ID == "" {
		return errors.New("student ID is required")
	}
	if student.FirstName == "" || student.LastName == "" {
		return errors.New("student name is required")
	}
	if student.Email == "" {
		return errors.New("student email is required")
	}
	
	return s.studentRepo.CreateStudent(ctx, student)
}

// GetStudent retrieves a student by ID
func (s *Service) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	if id == "" {
		return nil, errors.New("student ID is required")
	}
	return s.studentRepo.GetStudent(ctx, id)
}

// UpdateStudent updates a student with validation
func (s *Service) UpdateStudent(ctx context.Context, student *models.Student) error {
	if student.ID == "" {
		return errors.New("student ID is required")
	}
	
	// Check if student exists
	_, err := s.studentRepo.GetStudent(ctx, student.ID)
	if err != nil {
		return fmt.Errorf("student not found: %w", err)
	}
	
	return s.studentRepo.UpdateStudent(ctx, student)
}

// DeleteStudent deletes a student
func (s *Service) DeleteStudent(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("student ID is required")
	}
	return s.studentRepo.DeleteStudent(ctx, id)
}

// Teacher operations

// CreateTeacher creates a new teacher with validation
func (s *Service) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	if teacher.ID == "" {
		return errors.New("teacher ID is required")
	}
	if teacher.FirstName == "" || teacher.LastName == "" {
		return errors.New("teacher name is required")
	}
	if teacher.Email == "" {
		return errors.New("teacher email is required")
	}
	
	return s.teacherRepo.CreateTeacher(ctx, teacher)
}

// GetTeacher retrieves a teacher by ID
func (s *Service) GetTeacher(ctx context.Context, id string) (*models.Teacher, error) {
	if id == "" {
		return nil, errors.New("teacher ID is required")
	}
	return s.teacherRepo.GetTeacher(ctx, id)
}

// UpdateTeacher updates a teacher with validation
func (s *Service) UpdateTeacher(ctx context.Context, teacher *models.Teacher) error {
	if teacher.ID == "" {
		return errors.New("teacher ID is required")
	}
	
	_, err := s.teacherRepo.GetTeacher(ctx, teacher.ID)
	if err != nil {
		return fmt.Errorf("teacher not found: %w", err)
	}
	
	return s.teacherRepo.UpdateTeacher(ctx, teacher)
}

// DeleteTeacher deletes a teacher
func (s *Service) DeleteTeacher(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("teacher ID is required")
	}
	return s.teacherRepo.DeleteTeacher(ctx, id)
}

// Class operations

// CreateClass creates a new class with validation
func (s *Service) CreateClass(ctx context.Context, class *models.Class) error {
	if class.ID == "" {
		return errors.New("class ID is required")
	}
	if class.Name == "" {
		return errors.New("class name is required")
	}
	if class.TeacherID == "" {
		return errors.New("teacher ID is required")
	}
	
	// Verify teacher exists
	_, err := s.teacherRepo.GetTeacher(ctx, class.TeacherID)
	if err != nil {
		return fmt.Errorf("teacher not found: %w", err)
	}
	
	return s.classRepo.CreateClass(ctx, class)
}

// GetClass retrieves a class by ID
func (s *Service) GetClass(ctx context.Context, id string) (*models.Class, error) {
	if id == "" {
		return nil, errors.New("class ID is required")
	}
	return s.classRepo.GetClass(ctx, id)
}

// UpdateClass updates a class with validation
func (s *Service) UpdateClass(ctx context.Context, class *models.Class) error {
	if class.ID == "" {
		return errors.New("class ID is required")
	}
	
	_, err := s.classRepo.GetClass(ctx, class.ID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}
	
	return s.classRepo.UpdateClass(ctx, class)
}

// DeleteClass deletes a class
func (s *Service) DeleteClass(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("class ID is required")
	}
	return s.classRepo.DeleteClass(ctx, id)
}

// Academic operations

// CreateAcademic creates a new academic enrollment with validation
func (s *Service) CreateAcademic(ctx context.Context, academic *models.Academic) error {
	if academic.ID == "" {
		return errors.New("academic ID is required")
	}
	if academic.StudentID == "" {
		return errors.New("student ID is required")
	}
	if academic.ClassID == "" {
		return errors.New("class ID is required")
	}
	
	// Verify student exists
	_, err := s.studentRepo.GetStudent(ctx, academic.StudentID)
	if err != nil {
		return fmt.Errorf("student not found: %w", err)
	}
	
	// Verify class exists
	_, err = s.classRepo.GetClass(ctx, academic.ClassID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}
	
	if academic.EnrollmentDate.IsZero() {
		academic.EnrollmentDate = time.Now()
	}
	if academic.Status == "" {
		academic.Status = "active"
	}
	
	return s.academicRepo.CreateAcademic(ctx, academic)
}

// GetAcademic retrieves an academic enrollment by ID
func (s *Service) GetAcademic(ctx context.Context, id string) (*models.Academic, error) {
	if id == "" {
		return nil, errors.New("academic ID is required")
	}
	return s.academicRepo.GetAcademic(ctx, id)
}

// UpdateAcademic updates an academic enrollment
func (s *Service) UpdateAcademic(ctx context.Context, academic *models.Academic) error {
	if academic.ID == "" {
		return errors.New("academic ID is required")
	}
	
	_, err := s.academicRepo.GetAcademic(ctx, academic.ID)
	if err != nil {
		return fmt.Errorf("academic enrollment not found: %w", err)
	}
	
	return s.academicRepo.UpdateAcademic(ctx, academic)
}

// DeleteAcademic deletes an academic enrollment
func (s *Service) DeleteAcademic(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("academic ID is required")
	}
	return s.academicRepo.DeleteAcademic(ctx, id)
}

// Exam operations

// CreateExam creates a new exam with validation
func (s *Service) CreateExam(ctx context.Context, exam *models.Exam) error {
	if exam.ID == "" {
		return errors.New("exam ID is required")
	}
	if exam.ClassID == "" {
		return errors.New("class ID is required")
	}
	if exam.Title == "" {
		return errors.New("exam title is required")
	}
	
	// Verify class exists
	_, err := s.classRepo.GetClass(ctx, exam.ClassID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}
	
	return s.examRepo.CreateExam(ctx, exam)
}

// GetExam retrieves an exam by ID
func (s *Service) GetExam(ctx context.Context, id string) (*models.Exam, error) {
	if id == "" {
		return nil, errors.New("exam ID is required")
	}
	return s.examRepo.GetExam(ctx, id)
}

// UpdateExam updates an exam
func (s *Service) UpdateExam(ctx context.Context, exam *models.Exam) error {
	if exam.ID == "" {
		return errors.New("exam ID is required")
	}
	
	_, err := s.examRepo.GetExam(ctx, exam.ID)
	if err != nil {
		return fmt.Errorf("exam not found: %w", err)
	}
	
	return s.examRepo.UpdateExam(ctx, exam)
}

// DeleteExam deletes an exam
func (s *Service) DeleteExam(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("exam ID is required")
	}
	return s.examRepo.DeleteExam(ctx, id)
}

// ExamResult operations

// CreateExamResult creates a new exam result with validation
func (s *Service) CreateExamResult(ctx context.Context, result *models.ExamResult) error {
	if result.ID == "" {
		return errors.New("exam result ID is required")
	}
	if result.ExamID == "" {
		return errors.New("exam ID is required")
	}
	if result.StudentID == "" {
		return errors.New("student ID is required")
	}
	
	// Verify exam exists
	exam, err := s.examRepo.GetExam(ctx, result.ExamID)
	if err != nil {
		return fmt.Errorf("exam not found: %w", err)
	}
	
	// Verify student exists
	_, err = s.studentRepo.GetStudent(ctx, result.StudentID)
	if err != nil {
		return fmt.Errorf("student not found: %w", err)
	}
	
	// Validate score
	if result.Score < 0 || result.Score > exam.MaxScore {
		return fmt.Errorf("invalid score: must be between 0 and %d", exam.MaxScore)
	}
	
	// Calculate grade based on percentage
	percentage := float64(result.Score) / float64(exam.MaxScore) * 100
	result.Grade = calculateGrade(percentage)
	
	if result.TakenDate.IsZero() {
		result.TakenDate = time.Now()
	}
	
	return s.examResultRepo.CreateExamResult(ctx, result)
}

// GetExamResult retrieves an exam result by ID
func (s *Service) GetExamResult(ctx context.Context, id string) (*models.ExamResult, error) {
	if id == "" {
		return nil, errors.New("exam result ID is required")
	}
	return s.examResultRepo.GetExamResult(ctx, id)
}

// UpdateExamResult updates an exam result
func (s *Service) UpdateExamResult(ctx context.Context, result *models.ExamResult) error {
	if result.ID == "" {
		return errors.New("exam result ID is required")
	}
	
	_, err := s.examResultRepo.GetExamResult(ctx, result.ID)
	if err != nil {
		return fmt.Errorf("exam result not found: %w", err)
	}
	
	return s.examResultRepo.UpdateExamResult(ctx, result)
}

// DeleteExamResult deletes an exam result
func (s *Service) DeleteExamResult(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("exam result ID is required")
	}
	return s.examResultRepo.DeleteExamResult(ctx, id)
}

// Achievement operations

// CreateAchievement creates a new achievement with validation
func (s *Service) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
	if achievement.ID == "" {
		return errors.New("achievement ID is required")
	}
	if achievement.StudentID == "" {
		return errors.New("student ID is required")
	}
	if achievement.Title == "" {
		return errors.New("achievement title is required")
	}
	
	// Verify student exists
	_, err := s.studentRepo.GetStudent(ctx, achievement.StudentID)
	if err != nil {
		return fmt.Errorf("student not found: %w", err)
	}
	
	if achievement.AwardDate.IsZero() {
		achievement.AwardDate = time.Now()
	}
	
	return s.achievementRepo.CreateAchievement(ctx, achievement)
}

// GetAchievement retrieves an achievement by ID
func (s *Service) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
	if id == "" {
		return nil, errors.New("achievement ID is required")
	}
	return s.achievementRepo.GetAchievement(ctx, id)
}

// UpdateAchievement updates an achievement
func (s *Service) UpdateAchievement(ctx context.Context, achievement *models.Achievement) error {
	if achievement.ID == "" {
		return errors.New("achievement ID is required")
	}
	
	_, err := s.achievementRepo.GetAchievement(ctx, achievement.ID)
	if err != nil {
		return fmt.Errorf("achievement not found: %w", err)
	}
	
	return s.achievementRepo.UpdateAchievement(ctx, achievement)
}

// DeleteAchievement deletes an achievement
func (s *Service) DeleteAchievement(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("achievement ID is required")
	}
	return s.achievementRepo.DeleteAchievement(ctx, id)
}

// Helper functions

// calculateGrade calculates letter grade from percentage
func calculateGrade(percentage float64) string {
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

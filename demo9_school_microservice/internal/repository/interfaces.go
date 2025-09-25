// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package repository

import (
	"context"
	"school-microservice/internal/models"
)

// StudentRepository defines the interface for student data operations
type StudentRepository interface {
	Create(ctx context.Context, student *models.Student) error
	GetByID(ctx context.Context, id string) (*models.Student, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*models.Student, int, error)
	Update(ctx context.Context, id string, student *models.Student) error
	Delete(ctx context.Context, id string) error
	GetByGrade(ctx context.Context, grade string) ([]*models.Student, error)
	GetByEmail(ctx context.Context, email string) (*models.Student, error)
}

// TeacherRepository defines the interface for teacher data operations
type TeacherRepository interface {
	Create(ctx context.Context, teacher *models.Teacher) error
	GetByID(ctx context.Context, id string) (*models.Teacher, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*models.Teacher, int, error)
	Update(ctx context.Context, id string, teacher *models.Teacher) error
	Delete(ctx context.Context, id string) error
	GetByDepartment(ctx context.Context, department string) ([]*models.Teacher, error)
	GetBySubject(ctx context.Context, subject string) ([]*models.Teacher, error)
	GetByEmail(ctx context.Context, email string) (*models.Teacher, error)
}

// ClassRepository defines the interface for class data operations
type ClassRepository interface {
	Create(ctx context.Context, class *models.Class) error
	GetByID(ctx context.Context, id string) (*models.Class, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*models.Class, int, error)
	Update(ctx context.Context, id string, class *models.Class) error
	Delete(ctx context.Context, id string) error
	GetByTeacherID(ctx context.Context, teacherID string) ([]*models.Class, error)
	GetByGrade(ctx context.Context, grade string) ([]*models.Class, error)
	GetBySubject(ctx context.Context, subject string) ([]*models.Class, error)
}

// AcademicRepository defines the interface for academic record operations
type AcademicRepository interface {
	Create(ctx context.Context, academic *models.Academic) error
	GetByID(ctx context.Context, id string) (*models.Academic, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*models.Academic, int, error)
	Update(ctx context.Context, id string, academic *models.Academic) error
	Delete(ctx context.Context, id string) error
	GetByStudentID(ctx context.Context, studentID string) ([]*models.Academic, error)
	GetByClassID(ctx context.Context, classID string) ([]*models.Academic, error)
	GetByStudentIDAndSubject(ctx context.Context, studentID, subject string) ([]*models.Academic, error)
}

// AchievementRepository defines the interface for achievement operations
type AchievementRepository interface {
	Create(ctx context.Context, achievement *models.Achievement) error
	GetByID(ctx context.Context, id string) (*models.Achievement, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*models.Achievement, int, error)
	Update(ctx context.Context, id string, achievement *models.Achievement) error
	Delete(ctx context.Context, id string) error
	GetByStudentID(ctx context.Context, studentID string) ([]*models.Achievement, error)
	GetByCategory(ctx context.Context, category string) ([]*models.Achievement, error)
	GetByLevel(ctx context.Context, level string) ([]*models.Achievement, error)
}

// StudentClassRepository defines the interface for student-class relationship operations
type StudentClassRepository interface {
	EnrollStudent(ctx context.Context, studentClass *models.StudentClass) error
	UnenrollStudent(ctx context.Context, studentID, classID string) error
	GetStudentsByClassID(ctx context.Context, classID string) ([]*models.Student, error)
	GetClassesByStudentID(ctx context.Context, studentID string) ([]*models.Class, error)
	IsStudentEnrolled(ctx context.Context, studentID, classID string) (bool, error)
}

// Repository aggregates all repository interfaces
type Repository struct {
	Student      StudentRepository
	Teacher      TeacherRepository
	Class        ClassRepository
	Academic     AcademicRepository
	Achievement  AchievementRepository
	StudentClass StudentClassRepository
}
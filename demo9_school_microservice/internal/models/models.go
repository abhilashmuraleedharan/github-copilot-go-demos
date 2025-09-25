// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package models

import (
	"time"
)

// Student represents a student in the school system
type Student struct {
	ID          string    `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Grade       string    `json:"grade" db:"grade"`
	Address     string    `json:"address" db:"address"`
	Phone       string    `json:"phone" db:"phone"`
	ParentName  string    `json:"parent_name" db:"parent_name"`
	ParentPhone string    `json:"parent_phone" db:"parent_phone"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Teacher represents a teacher in the school system
type Teacher struct {
	ID          string    `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Department  string    `json:"department" db:"department"`
	Subject     string    `json:"subject" db:"subject"`
	Experience  int       `json:"experience" db:"experience"`
	Salary      float64   `json:"salary" db:"salary"`
	HireDate    time.Time `json:"hire_date" db:"hire_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Class represents a class in the school system
type Class struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Subject     string    `json:"subject" db:"subject"`
	TeacherID   string    `json:"teacher_id" db:"teacher_id"`
	Grade       string    `json:"grade" db:"grade"`
	Room        string    `json:"room" db:"room"`
	Schedule    string    `json:"schedule" db:"schedule"`
	Capacity    int       `json:"capacity" db:"capacity"`
	Enrolled    int       `json:"enrolled" db:"enrolled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Academic represents academic records and exams
type Academic struct {
	ID         string    `json:"id" db:"id"`
	StudentID  string    `json:"student_id" db:"student_id"`
	ClassID    string    `json:"class_id" db:"class_id"`
	ExamType   string    `json:"exam_type" db:"exam_type"` // midterm, final, quiz, assignment
	Subject    string    `json:"subject" db:"subject"`
	MaxMarks   float64   `json:"max_marks" db:"max_marks"`
	ObtMarks   float64   `json:"obt_marks" db:"obt_marks"`
	Grade      string    `json:"grade" db:"grade"`
	ExamDate   time.Time `json:"exam_date" db:"exam_date"`
	Remarks    string    `json:"remarks" db:"remarks"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Achievement represents student achievements and awards
type Achievement struct {
	ID          string    `json:"id" db:"id"`
	StudentID   string    `json:"student_id" db:"student_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Category    string    `json:"category" db:"category"` // academic, sports, cultural, behavior
	Level       string    `json:"level" db:"level"`       // school, district, state, national
	Date        time.Time `json:"date" db:"date"`
	AwardedBy   string    `json:"awarded_by" db:"awarded_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// StudentClass represents the many-to-many relationship between students and classes
type StudentClass struct {
	StudentID  string    `json:"student_id" db:"student_id"`
	ClassID    string    `json:"class_id" db:"class_id"`
	EnrolledAt time.Time `json:"enrolled_at" db:"enrolled_at"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	Error      string      `json:"error,omitempty"`
}

// CreateStudentRequest represents the request payload for creating a student
type CreateStudentRequest struct {
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Grade       string    `json:"grade" binding:"required"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	ParentName  string    `json:"parent_name"`
	ParentPhone string    `json:"parent_phone"`
}

// CreateTeacherRequest represents the request payload for creating a teacher
type CreateTeacherRequest struct {
	FirstName  string    `json:"first_name" binding:"required"`
	LastName   string    `json:"last_name" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Phone      string    `json:"phone"`
	Department string    `json:"department" binding:"required"`
	Subject    string    `json:"subject" binding:"required"`
	Experience int       `json:"experience"`
	Salary     float64   `json:"salary"`
	HireDate   time.Time `json:"hire_date"`
}

// CreateClassRequest represents the request payload for creating a class
type CreateClassRequest struct {
	Name      string `json:"name" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	TeacherID string `json:"teacher_id" binding:"required"`
	Grade     string `json:"grade" binding:"required"`
	Room      string `json:"room"`
	Schedule  string `json:"schedule"`
	Capacity  int    `json:"capacity"`
}

// CreateAcademicRequest represents the request payload for creating academic records
type CreateAcademicRequest struct {
	StudentID string    `json:"student_id" binding:"required"`
	ClassID   string    `json:"class_id" binding:"required"`
	ExamType  string    `json:"exam_type" binding:"required"`
	Subject   string    `json:"subject" binding:"required"`
	MaxMarks  float64   `json:"max_marks" binding:"required"`
	ObtMarks  float64   `json:"obt_marks" binding:"required"`
	ExamDate  time.Time `json:"exam_date" binding:"required"`
	Remarks   string    `json:"remarks"`
}

// CreateAchievementRequest represents the request payload for creating achievements
type CreateAchievementRequest struct {
	StudentID   string    `json:"student_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Category    string    `json:"category" binding:"required"`
	Level       string    `json:"level" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	AwardedBy   string    `json:"awarded_by"`
}
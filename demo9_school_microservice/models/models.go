// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package models

import "time"

// Student represents a student entity
type Student struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
	Grade       string    `json:"grade"`
	Type        string    `json:"type"` // Document type for Couchbase
}

// Teacher represents a teacher entity
type Teacher struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Subject     string    `json:"subject"`
	HireDate    time.Time `json:"hireDate"`
	Type        string    `json:"type"`
}

// Class represents a class entity
type Class struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Subject     string   `json:"subject"`
	TeacherID   string   `json:"teacherId"`
	Schedule    string   `json:"schedule"`
	MaxStudents int      `json:"maxStudents"`
	Type        string   `json:"type"`
}

// Academic represents academic enrollment linking students to classes
type Academic struct {
	ID             string    `json:"id"`
	StudentID      string    `json:"studentId"`
	ClassID        string    `json:"classId"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
	Status         string    `json:"status"` // active, completed, dropped
	Type           string    `json:"type"`
}

// Exam represents an exam for a class
type Exam struct {
	ID          string    `json:"id"`
	ClassID     string    `json:"classId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExamDate    time.Time `json:"examDate"`
	MaxScore    int       `json:"maxScore"`
	Type        string    `json:"type"`
}

// ExamResult represents a student's exam result
type ExamResult struct {
	ID        string    `json:"id"`
	ExamID    string    `json:"examId"`
	StudentID string    `json:"studentId"`
	Score     int       `json:"score"`
	Grade     string    `json:"grade"`
	TakenDate time.Time `json:"takenDate"`
	Type      string    `json:"type"`
}

// Achievement represents a student achievement or award
type Achievement struct {
	ID          string    `json:"id"`
	StudentID   string    `json:"studentId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AwardDate   time.Time `json:"awardDate"`
	Category    string    `json:"category"` // academic, sports, arts, leadership
	Type        string    `json:"type"`
}

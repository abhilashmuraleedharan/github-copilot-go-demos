// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package models

import "time"

// Academic represents an academic record in the school system
type Academic struct {
	ID         string    `json:"id" couchbase:"id"`
	StudentID  string    `json:"studentId" couchbase:"studentId"`
	ClassID    string    `json:"classId" couchbase:"classId"`
	Subject    string    `json:"subject" couchbase:"subject"`
	ExamType   string    `json:"examType" couchbase:"examType"` // midterm, final, quiz, assignment
	Score      float64   `json:"score" couchbase:"score"`
	MaxScore   float64   `json:"maxScore" couchbase:"maxScore"`
	Grade      string    `json:"grade" couchbase:"grade"` // A, B, C, D, F
	ExamDate   time.Time `json:"examDate" couchbase:"examDate"`
	Semester   string    `json:"semester" couchbase:"semester"`
	Year       int       `json:"year" couchbase:"year"`
	TeacherID  string    `json:"teacherId" couchbase:"teacherId"`
	Comments   string    `json:"comments" couchbase:"comments"`
	CreatedAt  time.Time `json:"createdAt" couchbase:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" couchbase:"updatedAt"`
}

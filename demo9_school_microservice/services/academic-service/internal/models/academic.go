package models

import (
	"time"
)

type Academic struct {
	ID          string    `json:"id" couchbase:"id"`
	StudentID   string    `json:"student_id" binding:"required" couchbase:"student_id"`
	TeacherID   string    `json:"teacher_id" binding:"required" couchbase:"teacher_id"`
	Subject     string    `json:"subject" binding:"required" couchbase:"subject"`
	Grade       string    `json:"grade" binding:"required" couchbase:"grade"`
	Semester    string    `json:"semester" binding:"required" couchbase:"semester"`
	AcademicYear string   `json:"academic_year" binding:"required" couchbase:"academic_year"`
	ExamType    string    `json:"exam_type" binding:"required" couchbase:"exam_type"` // midterm, final, quiz, assignment
	ExamDate    time.Time `json:"exam_date" couchbase:"exam_date"`
	MaxMarks    float64   `json:"max_marks" binding:"required" couchbase:"max_marks"`
	ObtainedMarks float64 `json:"obtained_marks" binding:"required" couchbase:"obtained_marks"`
	Percentage  float64   `json:"percentage" couchbase:"percentage"`
	Status      string    `json:"status" couchbase:"status"` // pass, fail, pending
	Remarks     string    `json:"remarks" couchbase:"remarks"`
	CreatedAt   time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" couchbase:"updated_at"`
	Type        string    `json:"type" couchbase:"type"` // Document type for Couchbase
}

type Class struct {
	ID           string    `json:"id" couchbase:"id"`
	ClassName    string    `json:"class_name" binding:"required" couchbase:"class_name"`
	Grade        string    `json:"grade" binding:"required" couchbase:"grade"`
	Section      string    `json:"section" couchbase:"section"`
	TeacherID    string    `json:"teacher_id" binding:"required" couchbase:"teacher_id"`
	Subject      string    `json:"subject" binding:"required" couchbase:"subject"`
	AcademicYear string    `json:"academic_year" binding:"required" couchbase:"academic_year"`
	Semester     string    `json:"semester" binding:"required" couchbase:"semester"`
	StudentIDs   []string  `json:"student_ids" couchbase:"student_ids"`
	Schedule     Schedule  `json:"schedule" couchbase:"schedule"`
	MaxCapacity  int       `json:"max_capacity" couchbase:"max_capacity"`
	Status       string    `json:"status" couchbase:"status"` // active, inactive, completed
	CreatedAt    time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" couchbase:"updated_at"`
	Type         string    `json:"type" couchbase:"type"`
}

type Schedule struct {
	DayOfWeek string `json:"day_of_week" couchbase:"day_of_week"` // Monday, Tuesday, etc.
	StartTime string `json:"start_time" couchbase:"start_time"`   // HH:MM format
	EndTime   string `json:"end_time" couchbase:"end_time"`       // HH:MM format
	Room      string `json:"room" couchbase:"room"`
}

type CreateAcademicRequest struct {
	StudentID     string    `json:"student_id" binding:"required"`
	TeacherID     string    `json:"teacher_id" binding:"required"`
	Subject       string    `json:"subject" binding:"required"`
	Grade         string    `json:"grade" binding:"required"`
	Semester      string    `json:"semester" binding:"required"`
	AcademicYear  string    `json:"academic_year" binding:"required"`
	ExamType      string    `json:"exam_type" binding:"required"`
	ExamDate      time.Time `json:"exam_date"`
	MaxMarks      float64   `json:"max_marks" binding:"required"`
	ObtainedMarks float64   `json:"obtained_marks" binding:"required"`
	Remarks       string    `json:"remarks"`
}

type CreateClassRequest struct {
	ClassName    string   `json:"class_name" binding:"required"`
	Grade        string   `json:"grade" binding:"required"`
	Section      string   `json:"section"`
	TeacherID    string   `json:"teacher_id" binding:"required"`
	Subject      string   `json:"subject" binding:"required"`
	AcademicYear string   `json:"academic_year" binding:"required"`
	Semester     string   `json:"semester" binding:"required"`
	StudentIDs   []string `json:"student_ids"`
	Schedule     Schedule `json:"schedule"`
	MaxCapacity  int      `json:"max_capacity"`
}

type UpdateAcademicRequest struct {
	StudentID     *string    `json:"student_id,omitempty"`
	TeacherID     *string    `json:"teacher_id,omitempty"`
	Subject       *string    `json:"subject,omitempty"`
	Grade         *string    `json:"grade,omitempty"`
	Semester      *string    `json:"semester,omitempty"`
	AcademicYear  *string    `json:"academic_year,omitempty"`
	ExamType      *string    `json:"exam_type,omitempty"`
	ExamDate      *time.Time `json:"exam_date,omitempty"`
	MaxMarks      *float64   `json:"max_marks,omitempty"`
	ObtainedMarks *float64   `json:"obtained_marks,omitempty"`
	Remarks       *string    `json:"remarks,omitempty"`
	Status        *string    `json:"status,omitempty"`
}

type StudentGradeReport struct {
	StudentID    string             `json:"student_id"`
	AcademicYear string             `json:"academic_year"`
	Semester     string             `json:"semester"`
	Subjects     []SubjectGrade     `json:"subjects"`
	OverallGPA   float64            `json:"overall_gpa"`
	TotalMarks   float64            `json:"total_marks"`
	MaxMarks     float64            `json:"max_marks"`
	Percentage   float64            `json:"percentage"`
	Status       string             `json:"status"`
}

type SubjectGrade struct {
	Subject       string  `json:"subject"`
	TeacherID     string  `json:"teacher_id"`
	TotalMarks    float64 `json:"total_marks"`
	MaxMarks      float64 `json:"max_marks"`
	Percentage    float64 `json:"percentage"`
	Grade         string  `json:"grade"`
	Status        string  `json:"status"`
}

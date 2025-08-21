// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package models

import "time"

// Class represents a class in the school system
type Class struct {
	ID          string    `json:"id" couchbase:"id"`
	Name        string    `json:"name" couchbase:"name"`
	Subject     string    `json:"subject" couchbase:"subject"`
	TeacherID   string    `json:"teacherId" couchbase:"teacherId"`
	Grade       string    `json:"grade" couchbase:"grade"`
	Room        string    `json:"room" couchbase:"room"`
	Schedule    string    `json:"schedule" couchbase:"schedule"` // e.g., "Mon,Wed,Fri 9:00-10:00"
	MaxStudents int       `json:"maxStudents" couchbase:"maxStudents"`
	StudentIDs  []string  `json:"studentIds" couchbase:"studentIds"`
	Semester    string    `json:"semester" couchbase:"semester"`
	Year        int       `json:"year" couchbase:"year"`
	Status      string    `json:"status" couchbase:"status"` // active, inactive, completed
	CreatedAt   time.Time `json:"createdAt" couchbase:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" couchbase:"updatedAt"`
}

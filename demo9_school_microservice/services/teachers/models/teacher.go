// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package models

import "time"

// Teacher represents a teacher in the school system
type Teacher struct {
	ID          string    `json:"id" couchbase:"id"`
	FirstName   string    `json:"firstName" couchbase:"firstName"`
	LastName    string    `json:"lastName" couchbase:"lastName"`
	Email       string    `json:"email" couchbase:"email"`
	Phone       string    `json:"phone" couchbase:"phone"`
	Department  string    `json:"department" couchbase:"department"`
	Subject     string    `json:"subject" couchbase:"subject"`
	HireDate    time.Time `json:"hireDate" couchbase:"hireDate"`
	Salary      float64   `json:"salary" couchbase:"salary"`
	Address     string    `json:"address" couchbase:"address"`
	Status      string    `json:"status" couchbase:"status"` // active, inactive, retired
	CreatedAt   time.Time `json:"createdAt" couchbase:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" couchbase:"updatedAt"`
}

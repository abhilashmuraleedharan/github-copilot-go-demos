// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package models

import "time"

// Achievement represents an achievement in the school system
type Achievement struct {
	ID          string    `json:"id" couchbase:"id"`
	StudentID   string    `json:"studentId" couchbase:"studentId"`
	Title       string    `json:"title" couchbase:"title"`
	Description string    `json:"description" couchbase:"description"`
	Category    string    `json:"category" couchbase:"category"` // academic, sports, arts, community, leadership
	Level       string    `json:"level" couchbase:"level"`       // school, district, state, national, international
	AwardedBy   string    `json:"awardedBy" couchbase:"awardedBy"`
	AwardDate   time.Time `json:"awardDate" couchbase:"awardDate"`
	Points      int       `json:"points" couchbase:"points"` // achievement points
	Certificate string    `json:"certificate" couchbase:"certificate"` // URL to certificate
	Status      string    `json:"status" couchbase:"status"` // pending, approved, revoked
	TeacherID   string    `json:"teacherId" couchbase:"teacherId"` // nominating teacher
	Comments    string    `json:"comments" couchbase:"comments"`
	CreatedAt   time.Time `json:"createdAt" couchbase:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" couchbase:"updatedAt"`
}

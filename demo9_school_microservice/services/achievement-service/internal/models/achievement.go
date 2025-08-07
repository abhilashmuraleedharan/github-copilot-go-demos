package models

import (
	"time"
)

type Achievement struct {
	ID          string    `json:"id" couchbase:"id"`
	StudentID   string    `json:"student_id" binding:"required" couchbase:"student_id"`
	TeacherID   string    `json:"teacher_id" couchbase:"teacher_id"` // Optional, for teacher-assigned achievements
	Title       string    `json:"title" binding:"required" couchbase:"title"`
	Description string    `json:"description" couchbase:"description"`
	Category    string    `json:"category" binding:"required" couchbase:"category"` // academic, sports, arts, behavior, leadership
	Level       string    `json:"level" binding:"required" couchbase:"level"`       // school, district, state, national, international
	AwardedDate time.Time `json:"awarded_date" binding:"required" couchbase:"awarded_date"`
	AcademicYear string   `json:"academic_year" couchbase:"academic_year"`
	Semester    string    `json:"semester" couchbase:"semester"`
	Points      int       `json:"points" couchbase:"points"` // Achievement points for gamification
	Certificate string    `json:"certificate" couchbase:"certificate"` // URL or path to certificate file
	Verified    bool      `json:"verified" couchbase:"verified"`
	VerifiedBy  string    `json:"verified_by" couchbase:"verified_by"` // Admin/Teacher ID who verified
	CreatedAt   time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" couchbase:"updated_at"`
	Type        string    `json:"type" couchbase:"type"` // Document type for Couchbase
}

type Badge struct {
	ID          string    `json:"id" couchbase:"id"`
	StudentID   string    `json:"student_id" binding:"required" couchbase:"student_id"`
	BadgeName   string    `json:"badge_name" binding:"required" couchbase:"badge_name"`
	BadgeType   string    `json:"badge_type" binding:"required" couchbase:"badge_type"` // attendance, performance, behavior, participation
	Criteria    string    `json:"criteria" couchbase:"criteria"`
	EarnedDate  time.Time `json:"earned_date" couchbase:"earned_date"`
	ValidUntil  time.Time `json:"valid_until" couchbase:"valid_until"` // For temporary badges
	BadgeIcon   string    `json:"badge_icon" couchbase:"badge_icon"`   // URL to badge icon
	Points      int       `json:"points" couchbase:"points"`
	IsActive    bool      `json:"is_active" couchbase:"is_active"`
	CreatedAt   time.Time `json:"created_at" couchbase:"created_at"`
	Type        string    `json:"type" couchbase:"type"`
}

type Competition struct {
	ID           string    `json:"id" couchbase:"id"`
	Name         string    `json:"name" binding:"required" couchbase:"name"`
	Description  string    `json:"description" couchbase:"description"`
	Category     string    `json:"category" binding:"required" couchbase:"category"`
	Level        string    `json:"level" binding:"required" couchbase:"level"`
	StartDate    time.Time `json:"start_date" binding:"required" couchbase:"start_date"`
	EndDate      time.Time `json:"end_date" binding:"required" couchbase:"end_date"`
	Participants []string  `json:"participants" couchbase:"participants"` // Student IDs
	Winners      []Winner  `json:"winners" couchbase:"winners"`
	Prizes       []Prize   `json:"prizes" couchbase:"prizes"`
	Status       string    `json:"status" couchbase:"status"` // upcoming, ongoing, completed, cancelled
	CreatedBy    string    `json:"created_by" couchbase:"created_by"`     // Teacher/Admin ID
	CreatedAt    time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" couchbase:"updated_at"`
	Type         string    `json:"type" couchbase:"type"`
}

type Winner struct {
	StudentID string `json:"student_id" couchbase:"student_id"`
	Position  int    `json:"position" couchbase:"position"` // 1st, 2nd, 3rd place
	Points    int    `json:"points" couchbase:"points"`
}

type Prize struct {
	Position    int    `json:"position" couchbase:"position"`
	Description string `json:"description" couchbase:"description"`
	Value       string `json:"value" couchbase:"value"`
}

// Request/Response DTOs
type CreateAchievementRequest struct {
	StudentID    string    `json:"student_id" binding:"required"`
	TeacherID    string    `json:"teacher_id"`
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	Category     string    `json:"category" binding:"required"`
	Level        string    `json:"level" binding:"required"`
	AwardedDate  time.Time `json:"awarded_date" binding:"required"`
	AcademicYear string    `json:"academic_year"`
	Semester     string    `json:"semester"`
	Points       int       `json:"points"`
	Certificate  string    `json:"certificate"`
}

type CreateBadgeRequest struct {
	StudentID  string    `json:"student_id" binding:"required"`
	BadgeName  string    `json:"badge_name" binding:"required"`
	BadgeType  string    `json:"badge_type" binding:"required"`
	Criteria   string    `json:"criteria"`
	ValidUntil time.Time `json:"valid_until"`
	BadgeIcon  string    `json:"badge_icon"`
	Points     int       `json:"points"`
}

type CreateCompetitionRequest struct {
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description"`
	Category     string    `json:"category" binding:"required"`
	Level        string    `json:"level" binding:"required"`
	StartDate    time.Time `json:"start_date" binding:"required"`
	EndDate      time.Time `json:"end_date" binding:"required"`
	Participants []string  `json:"participants"`
	Prizes       []Prize   `json:"prizes"`
	CreatedBy    string    `json:"created_by"`
}

type UpdateAchievementRequest struct {
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Category     *string    `json:"category,omitempty"`
	Level        *string    `json:"level,omitempty"`
	AwardedDate  *time.Time `json:"awarded_date,omitempty"`
	AcademicYear *string    `json:"academic_year,omitempty"`
	Semester     *string    `json:"semester,omitempty"`
	Points       *int       `json:"points,omitempty"`
	Certificate  *string    `json:"certificate,omitempty"`
	Verified     *bool      `json:"verified,omitempty"`
	VerifiedBy   *string    `json:"verified_by,omitempty"`
}

type StudentAchievementSummary struct {
	StudentID        string              `json:"student_id"`
	TotalAchievements int                `json:"total_achievements"`
	TotalPoints      int                 `json:"total_points"`
	TotalBadges      int                 `json:"total_badges"`
	Categories       map[string]int      `json:"categories"` // category -> count
	Levels           map[string]int      `json:"levels"`     // level -> count
	RecentAchievements []Achievement     `json:"recent_achievements"`
	ActiveBadges     []Badge             `json:"active_badges"`
}

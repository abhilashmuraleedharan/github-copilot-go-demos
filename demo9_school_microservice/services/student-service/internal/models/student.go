package models

import (
	"time"
)

type Student struct {
	ID          string    `json:"id" couchbase:"id"`
	FirstName   string    `json:"first_name" binding:"required" couchbase:"first_name"`
	LastName    string    `json:"last_name" binding:"required" couchbase:"last_name"`
	Email       string    `json:"email" binding:"required,email" couchbase:"email"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" couchbase:"date_of_birth"`
	Grade       string    `json:"grade" binding:"required" couchbase:"grade"`
	Address     Address   `json:"address" couchbase:"address"`
	Phone       string    `json:"phone" couchbase:"phone"`
	ParentName  string    `json:"parent_name" couchbase:"parent_name"`
	ParentPhone string    `json:"parent_phone" couchbase:"parent_phone"`
	EnrollDate  time.Time `json:"enroll_date" couchbase:"enroll_date"`
	Status      string    `json:"status" couchbase:"status"` // active, inactive, graduated
	CreatedAt   time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" couchbase:"updated_at"`
	Type        string    `json:"type" couchbase:"type"` // Document type for Couchbase
}

type Address struct {
	Street  string `json:"street" couchbase:"street"`
	City    string `json:"city" couchbase:"city"`
	State   string `json:"state" couchbase:"state"`
	ZipCode string `json:"zip_code" couchbase:"zip_code"`
	Country string `json:"country" couchbase:"country"`
}

type CreateStudentRequest struct {
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Grade       string    `json:"grade" binding:"required"`
	Address     Address   `json:"address"`
	Phone       string    `json:"phone"`
	ParentName  string    `json:"parent_name"`
	ParentPhone string    `json:"parent_phone"`
}

type UpdateStudentRequest struct {
	FirstName   *string  `json:"first_name,omitempty"`
	LastName    *string  `json:"last_name,omitempty"`
	Email       *string  `json:"email,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Grade       *string  `json:"grade,omitempty"`
	Address     *Address `json:"address,omitempty"`
	Phone       *string  `json:"phone,omitempty"`
	ParentName  *string  `json:"parent_name,omitempty"`
	ParentPhone *string  `json:"parent_phone,omitempty"`
	Status      *string  `json:"status,omitempty"`
}

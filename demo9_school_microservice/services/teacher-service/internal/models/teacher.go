package models

import (
	"time"
)

type Teacher struct {
	ID              string    `json:"id" couchbase:"id"`
	FirstName       string    `json:"first_name" binding:"required" couchbase:"first_name"`
	LastName        string    `json:"last_name" binding:"required" couchbase:"last_name"`
	Email           string    `json:"email" binding:"required,email" couchbase:"email"`
	Phone           string    `json:"phone" couchbase:"phone"`
	DateOfBirth     time.Time `json:"date_of_birth" couchbase:"date_of_birth"`
	HireDate        time.Time `json:"hire_date" couchbase:"hire_date"`
	Department      string    `json:"department" binding:"required" couchbase:"department"`
	Subject         []string  `json:"subjects" couchbase:"subjects"`
	Qualification   string    `json:"qualification" couchbase:"qualification"`
	Experience      int       `json:"experience" couchbase:"experience"` // Years of experience
	Salary          float64   `json:"salary" couchbase:"salary"`
	Address         Address   `json:"address" couchbase:"address"`
	EmergencyContact EmergencyContact `json:"emergency_contact" couchbase:"emergency_contact"`
	Status          string    `json:"status" couchbase:"status"` // active, inactive, on_leave
	CreatedAt       time.Time `json:"created_at" couchbase:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" couchbase:"updated_at"`
	Type            string    `json:"type" couchbase:"type"` // Document type for Couchbase
}

type Address struct {
	Street  string `json:"street" couchbase:"street"`
	City    string `json:"city" couchbase:"city"`
	State   string `json:"state" couchbase:"state"`
	ZipCode string `json:"zip_code" couchbase:"zip_code"`
	Country string `json:"country" couchbase:"country"`
}

type EmergencyContact struct {
	Name         string `json:"name" couchbase:"name"`
	Relationship string `json:"relationship" couchbase:"relationship"`
	Phone        string `json:"phone" couchbase:"phone"`
}

type CreateTeacherRequest struct {
	FirstName        string           `json:"first_name" binding:"required"`
	LastName         string           `json:"last_name" binding:"required"`
	Email            string           `json:"email" binding:"required,email"`
	Phone            string           `json:"phone"`
	DateOfBirth      time.Time        `json:"date_of_birth"`
	Department       string           `json:"department" binding:"required"`
	Subject          []string         `json:"subjects"`
	Qualification    string           `json:"qualification"`
	Experience       int              `json:"experience"`
	Salary           float64          `json:"salary"`
	Address          Address          `json:"address"`
	EmergencyContact EmergencyContact `json:"emergency_contact"`
}

type UpdateTeacherRequest struct {
	FirstName        *string           `json:"first_name,omitempty"`
	LastName         *string           `json:"last_name,omitempty"`
	Email            *string           `json:"email,omitempty"`
	Phone            *string           `json:"phone,omitempty"`
	DateOfBirth      *time.Time        `json:"date_of_birth,omitempty"`
	Department       *string           `json:"department,omitempty"`
	Subject          *[]string         `json:"subjects,omitempty"`
	Qualification    *string           `json:"qualification,omitempty"`
	Experience       *int              `json:"experience,omitempty"`
	Salary           *float64          `json:"salary,omitempty"`
	Address          *Address          `json:"address,omitempty"`
	EmergencyContact *EmergencyContact `json:"emergency_contact,omitempty"`
	Status           *string           `json:"status,omitempty"`
}

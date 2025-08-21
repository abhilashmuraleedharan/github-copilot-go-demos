// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
// Package models defines the data structures and domain models for the students service.
// This package contains the core business entities and their JSON/database mappings
// for student management operations.
package models

import "time"

// Student represents a student entity in the school management system.
//
// A Student contains all the essential information needed to manage a student's
// enrollment, personal details, and academic status within the school system.
// The struct uses both JSON and Couchbase tags for proper serialization and
// database storage.
//
// ID Generation:
// Student IDs are automatically generated using the format "STU{timestamp}"
// if not provided during creation. This ensures uniqueness while maintaining
// a readable pattern.
//
// Status Values:
//   - "active": Currently enrolled and attending
//   - "inactive": Temporarily not attending (e.g., leave of absence)
//   - "graduated": Completed studies and graduated
//
// Example:
//   student := &Student{
//       FirstName: "John",
//       LastName:  "Doe",
//       Email:     "john.doe@school.edu",
//       Grade:     "10",
//       Status:    "active",
//   }
type Student struct {
	// ID is the unique identifier for the student (auto-generated: STU{timestamp})
	ID string `json:"id" couchbase:"id"`
	
	// FirstName is the student's given name (required field)
	FirstName string `json:"firstName" couchbase:"firstName"`
	
	// LastName is the student's family name (required field)
	LastName string `json:"lastName" couchbase:"lastName"`
	
	// Email is the student's email address (required, should be unique)
	Email string `json:"email" couchbase:"email"`
	
	// DateOfBirth is the student's birth date for age calculation and verification
	DateOfBirth time.Time `json:"dateOfBirth" couchbase:"dateOfBirth"`
	
	// Grade represents the current grade level (e.g., "9", "10", "11", "12")
	Grade string `json:"grade" couchbase:"grade"`
	
	// Address is the student's home address for contact purposes
	Address string `json:"address" couchbase:"address"`
	
	// Phone is the primary contact phone number
	Phone string `json:"phone" couchbase:"phone"`
	
	// EnrollDate is when the student first enrolled (auto-set if not provided)
	EnrollDate time.Time `json:"enrollDate" couchbase:"enrollDate"`
	
	// Status indicates the student's current enrollment status
	// Valid values: "active", "inactive", "graduated"
	Status string `json:"status" couchbase:"status"`
	
	// CreatedAt is the timestamp when this record was created (auto-set)
	CreatedAt time.Time `json:"createdAt" couchbase:"createdAt"`
	
	// UpdatedAt is the timestamp when this record was last modified (auto-updated)
	UpdatedAt time.Time `json:"updatedAt" couchbase:"updatedAt"`
}

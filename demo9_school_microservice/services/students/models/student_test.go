// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
package models

import (
	"encoding/json"
	"testing"
	"time"
)

// TestStudent_JSONSerialization tests JSON marshaling and unmarshaling of Student struct
func TestStudent_JSONSerialization(t *testing.T) {
	// Create a test student with all fields populated
	originalStudent := Student{
		ID:          "STU20250821140530",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@school.edu",
		DateOfBirth: time.Date(2008, 5, 15, 0, 0, 0, 0, time.UTC),
		Grade:       "10",
		Address:     "123 Main St, Anytown, USA",
		Phone:       "+1-555-0123",
		EnrollDate:  time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
		Status:      "active",
		CreatedAt:   time.Date(2025, 8, 21, 14, 5, 30, 0, time.UTC),
		UpdatedAt:   time.Date(2025, 8, 21, 14, 5, 30, 0, time.UTC),
	}

	// Test marshaling to JSON
	jsonData, err := json.Marshal(originalStudent)
	if err != nil {
		t.Fatalf("failed to marshal student to JSON: %v", err)
	}

	// Verify JSON contains expected fields
	jsonStr := string(jsonData)
	expectedFields := []string{
		`"id":"STU20250821140530"`,
		`"firstName":"John"`,
		`"lastName":"Doe"`,
		`"email":"john.doe@school.edu"`,
		`"grade":"10"`,
		`"status":"active"`,
	}

	for _, field := range expectedFields {
		if !contains(jsonStr, field) {
			t.Errorf("JSON missing expected field: %s\nActual JSON: %s", field, jsonStr)
		}
	}

	// Test unmarshaling from JSON
	var unmarshaledStudent Student
	err = json.Unmarshal(jsonData, &unmarshaledStudent)
	if err != nil {
		t.Fatalf("failed to unmarshal student from JSON: %v", err)
	}

	// Verify all fields were unmarshaled correctly
	if unmarshaledStudent.ID != originalStudent.ID {
		t.Errorf("ID mismatch: expected %s, got %s", originalStudent.ID, unmarshaledStudent.ID)
	}
	if unmarshaledStudent.FirstName != originalStudent.FirstName {
		t.Errorf("FirstName mismatch: expected %s, got %s", originalStudent.FirstName, unmarshaledStudent.FirstName)
	}
	if unmarshaledStudent.LastName != originalStudent.LastName {
		t.Errorf("LastName mismatch: expected %s, got %s", originalStudent.LastName, unmarshaledStudent.LastName)
	}
	if unmarshaledStudent.Email != originalStudent.Email {
		t.Errorf("Email mismatch: expected %s, got %s", originalStudent.Email, unmarshaledStudent.Email)
	}
	if !unmarshaledStudent.DateOfBirth.Equal(originalStudent.DateOfBirth) {
		t.Errorf("DateOfBirth mismatch: expected %v, got %v", originalStudent.DateOfBirth, unmarshaledStudent.DateOfBirth)
	}
	if unmarshaledStudent.Grade != originalStudent.Grade {
		t.Errorf("Grade mismatch: expected %s, got %s", originalStudent.Grade, unmarshaledStudent.Grade)
	}
	if unmarshaledStudent.Address != originalStudent.Address {
		t.Errorf("Address mismatch: expected %s, got %s", originalStudent.Address, unmarshaledStudent.Address)
	}
	if unmarshaledStudent.Phone != originalStudent.Phone {
		t.Errorf("Phone mismatch: expected %s, got %s", originalStudent.Phone, unmarshaledStudent.Phone)
	}
	if !unmarshaledStudent.EnrollDate.Equal(originalStudent.EnrollDate) {
		t.Errorf("EnrollDate mismatch: expected %v, got %v", originalStudent.EnrollDate, unmarshaledStudent.EnrollDate)
	}
	if unmarshaledStudent.Status != originalStudent.Status {
		t.Errorf("Status mismatch: expected %s, got %s", originalStudent.Status, unmarshaledStudent.Status)
	}
	if !unmarshaledStudent.CreatedAt.Equal(originalStudent.CreatedAt) {
		t.Errorf("CreatedAt mismatch: expected %v, got %v", originalStudent.CreatedAt, unmarshaledStudent.CreatedAt)
	}
	if !unmarshaledStudent.UpdatedAt.Equal(originalStudent.UpdatedAt) {
		t.Errorf("UpdatedAt mismatch: expected %v, got %v", originalStudent.UpdatedAt, unmarshaledStudent.UpdatedAt)
	}
}

// TestStudent_JSONWithNilValues tests JSON handling with nil/zero values
func TestStudent_JSONWithNilValues(t *testing.T) {
	// Create student with minimal fields
	student := Student{
		ID:        "STU20250821140530",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@school.edu",
		// Other fields left as zero values
	}

	// Test marshaling
	jsonData, err := json.Marshal(student)
	if err != nil {
		t.Fatalf("failed to marshal student with nil values: %v", err)
	}

	// Test unmarshaling
	var unmarshaledStudent Student
	err = json.Unmarshal(jsonData, &unmarshaledStudent)
	if err != nil {
		t.Fatalf("failed to unmarshal student with nil values: %v", err)
	}

	// Verify required fields are preserved
	if unmarshaledStudent.ID != student.ID {
		t.Errorf("ID not preserved: expected %s, got %s", student.ID, unmarshaledStudent.ID)
	}
	if unmarshaledStudent.FirstName != student.FirstName {
		t.Errorf("FirstName not preserved: expected %s, got %s", student.FirstName, unmarshaledStudent.FirstName)
	}

	// Verify zero values are handled correctly
	if !unmarshaledStudent.DateOfBirth.IsZero() {
		t.Errorf("expected DateOfBirth to be zero value, got %v", unmarshaledStudent.DateOfBirth)
	}
	if !unmarshaledStudent.EnrollDate.IsZero() {
		t.Errorf("expected EnrollDate to be zero value, got %v", unmarshaledStudent.EnrollDate)
	}
}

// TestStudent_JSONFieldTags tests that JSON field tags are correctly applied
func TestStudent_JSONFieldTags(t *testing.T) {
	student := Student{
		ID:        "STU20250821140530",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@school.edu",
		Grade:     "10",
		Status:    "active",
	}

	jsonData, err := json.Marshal(student)
	if err != nil {
		t.Fatalf("failed to marshal student: %v", err)
	}

	jsonStr := string(jsonData)

	// Test that JSON uses camelCase field names as specified in tags
	expectedJSONFields := map[string]string{
		"id":          `"id":"STU20250821140530"`,
		"firstName":   `"firstName":"John"`,
		"lastName":    `"lastName":"Doe"`,
		"email":       `"email":"john.doe@school.edu"`,
		"dateOfBirth": `"dateOfBirth":"0001-01-01T00:00:00Z"`, // Zero time
		"grade":       `"grade":"10"`,
		"address":     `"address":""`,
		"phone":       `"phone":""`,
		"enrollDate":  `"enrollDate":"0001-01-01T00:00:00Z"`, // Zero time
		"status":      `"status":"active"`,
		"createdAt":   `"createdAt":"0001-01-01T00:00:00Z"`,   // Zero time
		"updatedAt":   `"updatedAt":"0001-01-01T00:00:00Z"`,   // Zero time
	}

	for fieldName, expectedJSON := range expectedJSONFields {
		if !contains(jsonStr, expectedJSON) {
			t.Errorf("JSON missing or incorrect field %s: expected %s\nActual JSON: %s", 
				fieldName, expectedJSON, jsonStr)
		}
	}

	// Verify it doesn't contain Go struct field names (PascalCase)
	unexpectedFields := []string{
		"ID:", "FirstName:", "LastName:", "Email:", "DateOfBirth:",
		"Grade:", "Address:", "Phone:", "EnrollDate:", "Status:",
		"CreatedAt:", "UpdatedAt:",
	}

	for _, field := range unexpectedFields {
		if contains(jsonStr, field) {
			t.Errorf("JSON should not contain Go field name %s\nActual JSON: %s", field, jsonStr)
		}
	}
}

// TestStudent_ValidationScenarios tests various student data validation scenarios
func TestStudent_ValidationScenarios(t *testing.T) {
	tests := []struct {
		name            string
		student         Student
		expectedValid   bool
		validationNotes string
	}{
		{
			name: "valid complete student",
			student: Student{
				ID:          "STU20250821140530",
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "john.doe@school.edu",
				DateOfBirth: time.Date(2008, 5, 15, 0, 0, 0, 0, time.UTC),
				Grade:       "10",
				Address:     "123 Main St",
				Phone:       "+1-555-0123",
				EnrollDate:  time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
				Status:      "active",
			},
			expectedValid:   true,
			validationNotes: "All fields properly populated",
		},
		{
			name: "valid minimal student",
			student: Student{
				ID:        "STU20250821140530",
				FirstName: "Jane",
				LastName:  "Smith",
				Email:     "jane.smith@school.edu",
				Status:    "active",
			},
			expectedValid:   true,
			validationNotes: "Minimal required fields only",
		},
		{
			name: "student with future date of birth",
			student: Student{
				ID:          "STU20250821140530",
				FirstName:   "Future",
				LastName:    "Student",
				Email:       "future@school.edu",
				DateOfBirth: time.Now().Add(24 * time.Hour), // Tomorrow
				Status:      "active",
			},
			expectedValid:   false,
			validationNotes: "Date of birth should not be in the future",
		},
		{
			name: "student with empty required fields",
			student: Student{
				ID:     "",
				Email:  "",
				Status: "active",
			},
			expectedValid:   false,
			validationNotes: "Missing required fields: ID, FirstName, LastName, Email",
		},
		{
			name: "student with invalid status",
			student: Student{
				ID:        "STU20250821140530",
				FirstName: "Invalid",
				LastName:  "Status",
				Email:     "invalid@school.edu",
				Status:    "unknown_status",
			},
			expectedValid:   false,
			validationNotes: "Status should be one of: active, inactive, graduated",
		},
		{
			name: "student with invalid grade",
			student: Student{
				ID:        "STU20250821140530",
				FirstName: "Invalid",
				LastName:  "Grade",
				Email:     "invalid.grade@school.edu",
				Grade:     "13", // Assuming grades are 9-12
				Status:    "active",
			},
			expectedValid:   false,
			validationNotes: "Grade should be valid school grade level",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Perform basic validation checks
			isValid := validateStudent(tt.student)
			
			if isValid != tt.expectedValid {
				t.Errorf("validation result mismatch: expected %v, got %v (%s)", 
					tt.expectedValid, isValid, tt.validationNotes)
			}
		})
	}
}

// TestStudent_StatusValues tests valid status values
func TestStudent_StatusValues(t *testing.T) {
	validStatuses := []string{"active", "inactive", "graduated"}
	invalidStatuses := []string{"", "pending", "suspended", "unknown", "ACTIVE", "Active"}

	for _, status := range validStatuses {
		student := Student{
			ID:        "STU20250821140530",
			FirstName: "Test",
			LastName:  "Student",
			Email:     "test@school.edu",
			Status:    status,
		}

		if !isValidStatus(student.Status) {
			t.Errorf("expected status '%s' to be valid", status)
		}
	}

	for _, status := range invalidStatuses {
		if isValidStatus(status) {
			t.Errorf("expected status '%s' to be invalid", status)
		}
	}
}

// TestStudent_GradeValues tests valid grade values
func TestStudent_GradeValues(t *testing.T) {
	validGrades := []string{"9", "10", "11", "12", "K", "1", "2", "3", "4", "5", "6", "7", "8"}
	invalidGrades := []string{"13", "0", "-1", "grade9", "ninth", ""}

	for _, grade := range validGrades {
		if !isValidGrade(grade) {
			t.Errorf("expected grade '%s' to be valid", grade)
		}
	}

	for _, grade := range invalidGrades {
		if isValidGrade(grade) {
			t.Errorf("expected grade '%s' to be invalid", grade)
		}
	}
}

// TestStudent_EmailValidation tests email format validation
func TestStudent_EmailValidation(t *testing.T) {
	validEmails := []string{
		"student@school.edu",
		"john.doe@school.edu",
		"jane_smith@school.edu",
		"test123@school.edu",
		"student+parent@school.edu",
	}

	invalidEmails := []string{
		"",
		"not-an-email",
		"@school.edu",
		"student@",
		"student space@school.edu",
		"student@school",
		"student.school.edu",
	}

	for _, email := range validEmails {
		if !isValidEmail(email) {
			t.Errorf("expected email '%s' to be valid", email)
		}
	}

	for _, email := range invalidEmails {
		if isValidEmail(email) {
			t.Errorf("expected email '%s' to be invalid", email)
		}
	}
}

// TestStudent_IDFormat tests student ID format validation
func TestStudent_IDFormat(t *testing.T) {
	validIDs := []string{
		"STU20250821140530",
		"STU20241201120000",
		"STU20230915080000",
	}

	invalidIDs := []string{
		"",
		"STU",
		"20250821140530",
		"STU2025",
		"STUDENT20250821140530",
		"stu20250821140530",
		"STU20250821140530X",
		"STU202508211405",
	}

	for _, id := range validIDs {
		if !isValidStudentID(id) {
			t.Errorf("expected ID '%s' to be valid", id)
		}
	}

	for _, id := range invalidIDs {
		if isValidStudentID(id) {
			t.Errorf("expected ID '%s' to be invalid", id)
		}
	}
}

// TestStudent_TimestampHandling tests timestamp field handling
func TestStudent_TimestampHandling(t *testing.T) {
	now := time.Now()
	
	student := Student{
		ID:        "STU20250821140530",
		FirstName: "Time",
		LastName:  "Test",
		Email:     "time.test@school.edu",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Test that timestamps are preserved in JSON serialization
	jsonData, err := json.Marshal(student)
	if err != nil {
		t.Fatalf("failed to marshal student: %v", err)
	}

	var unmarshaledStudent Student
	err = json.Unmarshal(jsonData, &unmarshaledStudent)
	if err != nil {
		t.Fatalf("failed to unmarshal student: %v", err)
	}

	// Verify timestamps are preserved (within reasonable precision)
	if !almostEqual(unmarshaledStudent.CreatedAt, student.CreatedAt) {
		t.Errorf("CreatedAt not preserved: expected %v, got %v", 
			student.CreatedAt, unmarshaledStudent.CreatedAt)
	}

	if !almostEqual(unmarshaledStudent.UpdatedAt, student.UpdatedAt) {
		t.Errorf("UpdatedAt not preserved: expected %v, got %v", 
			student.UpdatedAt, unmarshaledStudent.UpdatedAt)
	}
}

// Helper functions for validation tests

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    contains(s[1:], substr) || 
		    (len(s) > 0 && s[:len(substr)] == substr))
}

// validateStudent performs basic student validation
func validateStudent(s Student) bool {
	// Check required fields
	if s.ID == "" || s.FirstName == "" || s.LastName == "" || s.Email == "" {
		return false
	}

	// Check ID format
	if !isValidStudentID(s.ID) {
		return false
	}

	// Check email format
	if !isValidEmail(s.Email) {
		return false
	}

	// Check status
	if !isValidStatus(s.Status) {
		return false
	}

	// Check grade if provided
	if s.Grade != "" && !isValidGrade(s.Grade) {
		return false
	}

	// Check date of birth is not in the future
	if !s.DateOfBirth.IsZero() && s.DateOfBirth.After(time.Now()) {
		return false
	}

	return true
}

// isValidStatus checks if a status value is valid
func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"active":    true,
		"inactive":  true,
		"graduated": true,
	}
	return validStatuses[status]
}

// isValidGrade checks if a grade value is valid
func isValidGrade(grade string) bool {
	validGrades := map[string]bool{
		"K": true, "1": true, "2": true, "3": true, "4": true,
		"5": true, "6": true, "7": true, "8": true, "9": true,
		"10": true, "11": true, "12": true,
	}
	return validGrades[grade]
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// Simple email validation - contains @ and .
	hasAt := false
	hasDot := false
	for i, c := range email {
		if c == '@' {
			if hasAt || i == 0 || i == len(email)-1 {
				return false
			}
			hasAt = true
		} else if c == '.' && hasAt {
			hasDot = true
		} else if c == ' ' {
			return false
		}
	}
	return hasAt && hasDot
}

// isValidStudentID checks if a student ID follows the correct format
func isValidStudentID(id string) bool {
	// Expected format: STU + 14 digit timestamp (YYYYMMDDHHMMSS)
	if len(id) != 17 {
		return false
	}
	if id[:3] != "STU" {
		return false
	}
	// Check that the rest are digits
	for _, c := range id[3:] {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// almostEqual checks if two times are equal within a reasonable precision
func almostEqual(t1, t2 time.Time) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff < time.Second // Allow up to 1 second difference
}

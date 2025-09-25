// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-25
package emailutils

import "testing"

// TestIsValidEmail tests the IsValidEmail function with various email addresses.
func TestIsValidEmail(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name    string
		email   string
		valid   bool
	}{
		// Valid email addresses
		{"simple valid", "user@example.com", true},
		{"valid with dot", "first.last@example.co", true},
		{"valid with plus", "user+tag@example.org", true},
		// Edge case: sub-domain
		{"valid subdomain", "user@mail.example.com", true},
		// Invalid email addresses
		{"missing at", "userexample.com", false},
		{"missing domain", "user@.com", false},
		{"invalid chars", "user@exa!mple.com", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidEmail(tc.email)
			if got != tc.valid {
				t.Errorf("IsValidEmail(%q) = %v; want %v", tc.email, got, tc.valid)
			}
		})
	}
}

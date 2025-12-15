// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package emailutils

import (
	"testing"
)

// TestIsValidEmail tests the IsValidEmail function with various valid and invalid email addresses.
func TestIsValidEmail(t *testing.T) {
	// Table-driven test cases for different email scenarios
	testCases := []struct {
		email    string
		want    bool
		comment string
	}{
		// Valid email addresses
		{"user@example.com", true, "simple valid email"},
		{"user.name+tag@domain.co", true, "valid email with plus and dot"},
		{"user@mail.example.com", true, "valid email with sub-domain (edge case)"},

		// Invalid email addresses
		{"user@.com", false, "missing domain name"},
		{"user@domain", false, "missing TLD"},
		{"plainaddress", false, "missing @ and domain"},
	}

	for _, tc := range testCases {
		got := IsValidEmail(tc.email)
		if got != tc.want {
			t.Errorf("IsValidEmail(%q) = %v; want %v (%s)", tc.email, got, tc.want, tc.comment)
		}
	}
}

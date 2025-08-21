// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package emailutils

import (
	"testing"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
// TestIsValidEmail tests the IsValidEmail function with various email addresses.
// It uses table-driven tests for clarity and maintainability.
func TestIsValidEmail(t *testing.T) {
	// Define a list of test cases with input and expected output.
	testCases := []struct {
		email    string
		isValid  bool
		desc     string // Description for clarity in test output
	}{
		// Valid email addresses
		{"test@example.com", true, "simple valid email"},
		{"user.name+tag@example.co.uk", true, "valid email with plus and dot"},
		{"user@mail.example.com", true, "valid email with sub-domain (edge case)"},

		// Invalid email addresses
		{"plainaddress", false, "missing @ and domain"},
		{"user@.com", false, "domain starts with dot"},
		{"@example.com", false, "missing local part"},
	}

	for _, tc := range testCases {
		got := IsValidEmail(tc.email)
		if got != tc.isValid {
			t.Errorf("%s: IsValidEmail(%q) = %v; want %v", tc.desc, tc.email, got, tc.isValid)
		}
	}
}

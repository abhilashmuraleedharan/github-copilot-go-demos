// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-08
package emailutils

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		// Valid emails
		{"user@example.com", true},
		{"john.doe@company.org", true},
		{"user@mail.example.com", true}, // sub-domain edge case

		// Invalid emails
		{"plainaddress", false},
		{"user@.com", false},
		{"@example.com", false},
	}

	for _, tt := range tests {
		got := IsValidEmail(tt.email)
		if got != tt.expected {
			t.Errorf("IsValidEmail(%q) = %v; want %v", tt.email, got, tt.expected)
		}
	}
}

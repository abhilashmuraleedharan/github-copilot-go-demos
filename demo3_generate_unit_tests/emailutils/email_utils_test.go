// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
package emailutils

import (
    "testing"
)

// TestIsValidEmail tests the IsValidEmail function with various email addresses.
func TestIsValidEmail(t *testing.T) {
    // Table-driven test cases for IsValidEmail
    tests := []struct {
        email    string
        expected bool
        desc     string
    }{
        // Valid email addresses
        {"user@example.com", true, "simple valid email"},
        {"john.doe@company.org", true, "valid email with dot in local part"},
        {"alice+bob@service.co", true, "valid email with plus sign"},
        {"user@mail.example.com", true, "valid email with sub-domain"}, // edge case

        // Invalid email addresses
        {"plainaddress", false, "missing @ and domain"},
        {"@no-local-part.com", false, "missing local part"},
        {"user@.com", false, "domain starts with dot"},
    }

    for _, tc := range tests {
        result := IsValidEmail(tc.email)
        if result != tc.expected {
            t.Errorf("IsValidEmail(%q) = %v; want %v (%s)", tc.email, result, tc.expected, tc.desc)
        }
    }
}
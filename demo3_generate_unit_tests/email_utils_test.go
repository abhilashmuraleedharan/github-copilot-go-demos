package emailutils

import "testing"

func TestIsValidEmail(t *testing.T) {
    testCases := []struct {
        name     string
        email    string
        expected bool
    }{
        {"valid_simple", "user@example.com", true},
        {"valid_dot", "john.doe@example.co.uk", true},
        {"valid_subdomain", "user@mail.example.com", true},
        {"invalid_missing_at", "userexample.com", false},
        {"invalid_double_at", "user@@example.com", false},
        {"invalid_trailing_dot", "user@example.", false},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := IsValidEmail(tc.email)
            if result != tc.expected {
                t.Errorf("IsValidEmail(%q) = %v; want %v", tc.email, result, tc.expected)
            }
        })
    }
}

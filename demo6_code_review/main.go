package main

import (
	"fmt"
	"regexp"
	"strings"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
var emailPattern = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// IsValidEmail validates whether the given email address matches a valid email format.
// Returns true if the email is valid, false otherwise.
func IsValidEmail(email string) bool {
	return emailPattern.MatchString(email)
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// GetEmailDomain extracts the domain part from an email address.
// Returns an error if the email is invalid or doesn't contain "@".
func GetEmailDomain(email string) (string, error) {
	idx := strings.LastIndex(email, "@")
	if idx == -1 {
		return "", fmt.Errorf("invalid email: missing @")
	}
	return email[idx+1:], nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// GetEmailLocalPart extracts the local part (before "@") from an email address.
// Returns an error if the email doesn't contain "@".
func GetEmailLocalPart(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email format")
	}
	return parts[0], nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
// MaskEmail masks an email address, showing only the first visibleChars
// characters of the local part. Returns an error if the email is invalid
// or if visibleChars is out of range.
func MaskEmail(email string, visibleChars int) (string, error) {
	if !IsValidEmail(email) {
		return "", fmt.Errorf("invalid email format")
	}
	parts := strings.Split(email, "@")
	localPart := parts[0]
	domain := parts[1]
	
	if visibleChars < 0 || visibleChars > len(localPart) {
		return "", fmt.Errorf("visibleChars must be between 0 and %d", len(localPart))
	}
	
	masked := localPart[:visibleChars] + strings.Repeat("*", len(localPart)-visibleChars)
	return masked + "@" + domain, nil
}

func main() {
	email := "john.doe@example.com"
	masked, err := MaskEmail(email, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(masked)
}

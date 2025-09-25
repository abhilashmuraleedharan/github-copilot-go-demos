package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// Use anchors for full string matching instead of partial matching
var emailPattern = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// IsValidEmail validates email format using full string matching
func IsValidEmail(address string) bool {
	return emailPattern.MatchString(address)
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// GetDomain extracts the domain part from a valid email address
func GetDomain(addr string) (string, error) {
	if !IsValidEmail(addr) {
		return "", errors.New("invalid email address")
	}
	atIndex := strings.LastIndex(addr, "@")
	if atIndex == -1 || atIndex == len(addr)-1 {
		return "", errors.New("invalid email format")
	}
	return addr[atIndex+1:], nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// GetLocalPart extracts the local part from a valid email address
func GetLocalPart(addr string) (string, error) {
	if !IsValidEmail(addr) {
		return "", errors.New("invalid email address")
	}
	atIndex := strings.Index(addr, "@")
	if atIndex == -1 || atIndex == 0 {
		return "", errors.New("invalid email format")
	}
	return addr[:atIndex], nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
// MaskEmail masks an email showing only the specified number of characters from the local part
func MaskEmail(email string, showChars int) (string, error) {
	if !IsValidEmail(email) {
		return "", errors.New("invalid email address")
	}

	if showChars < 0 {
		return "", errors.New("showChars must be non-negative")
	}

	localPart, err := GetLocalPart(email)
	if err != nil {
		return "", err
	}

	domain, err := GetDomain(email)
	if err != nil {
		return "", err
	}

	// Handle case where showChars exceeds local part length
	if showChars >= len(localPart) {
		return email, nil // Return original if nothing to mask
	}

	masked := localPart[:showChars] + strings.Repeat("*", len(localPart)-showChars)
	return masked + "@" + domain, nil
}

func main() {
	email := "john.doe@example.com"

	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
	// Demonstrate proper error handling
	masked, err := MaskEmail(email, 2)
	if err != nil {
		fmt.Printf("Error masking email: %v\n", err)
		return
	}
	fmt.Printf("Original: %s\nMasked: %s\n", email, masked)

	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-27
	// Demonstrate domain extraction with error handling
	domain, err := GetDomain(email)
	if err != nil {
		fmt.Printf("Error getting domain: %v\n", err)
		return
	}
	fmt.Printf("Domain: %s\n", domain)
}

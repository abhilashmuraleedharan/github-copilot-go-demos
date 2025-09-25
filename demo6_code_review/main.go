package main

import (
	"fmt"
	"regexp"
	"strings"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
var emailPattern = regexp.MustCompile(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`)


// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2025-09-24
// IsValidEmail checks if the provided address matches a basic email pattern.
func IsValidEmailAddress(address string) bool {
	return emailPattern.MatchString(address)
}


// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2025-09-24
// GetEmailDomain returns the domain part of an email address.
// Returns an empty string if '@' is missing or at the end.
func GetEmailDomain(addr string) string {
	idx := strings.LastIndex(addr, "@")
	if idx == -1 || idx == len(addr)-1 {
		return ""
	}
	return addr[idx+1:]
}


// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2025-09-24
// GetEmailLocalPart returns the local part of an email address.
// Returns an empty string if '@' is missing.
func GetEmailLocalPart(addr string) string {
	idx := strings.Index(addr, "@")
	if idx == -1 {
		return ""
	}
	return addr[:idx]
}


// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2025-09-24
// MaskEmailLocalPart masks the local part of an email, showing only the first 'show' characters.
// Returns the original email if invalid or if 'show' is out of bounds.
func MaskEmailLocalPart(e string, show int) string {
	if !IsValidEmailAddress(e) {
		return e
	}
	parts := strings.SplitN(e, "@", 2)
	if len(parts) != 2 {
		return e
	}
	lp := parts[0]
	dom := parts[1]
	if show < 0 || show > len(lp) {
		return e
	}
	masked := lp[:show] + strings.Repeat("*", len(lp)-show)
	return masked + "@" + dom
}


func main() {
	email := "john.doe@example.com"
	masked := MaskEmailLocalPart(email, 2)
	fmt.Println(masked)
}

//Generate
package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// IsValidEmail checks if the provided email address is valid.
// Returns true if valid, false otherwise.
func IsValidEmail(address string) bool {
	return emailPattern.MatchString(address)
}

// DomainFromEmail extracts the domain part from an email address.
// Returns the domain and an error if the email is invalid or missing '@'.
func DomainFromEmail(addr string) (string, error) {
	at := strings.LastIndex(addr, "@")
	if at == -1 {
		return "", errors.New("missing '@' in email address")
	}
	domain := addr[at+1:]
	if domain == "" {
		return "", errors.New("domain part is empty")
	}
	return domain, nil
}

// LocalPartFromEmail extracts the local part from an email address.
// Returns the local part and an error if the email is invalid or missing '@'.
func LocalPartFromEmail(addr string) (string, error) {
	at := strings.LastIndex(addr, "@")
	if at == -1 {
		return "", errors.New("missing '@' in email address")
	}
	local := addr[:at]
	if local == "" {
		return "", errors.New("local part is empty")
	}
	return local, nil
}

// MaskEmail masks an email so only 'show' chars of the local part remain visible.
// Returns the masked email or an error describing the issue.
func MaskEmail(e string, show int) (string, error) {
	if !IsValidEmail(e) {
		return "", errors.New("invalid email address format")
	}
	local, err := LocalPartFromEmail(e)
	if err != nil {
		return "", err
	}
	domain, err := DomainFromEmail(e)
	if err != nil {
		return "", err
	}
	if show < 0 || show > len(local) {
		return "", fmt.Errorf("'show' parameter (%d) out of bounds for local part length %d", show, len(local))
	}
	masked := local[:show] + strings.Repeat("*", len(local)-show)
	return masked + "@" + domain, nil
}

func main() {
	email := "john.doe@example.com"
	masked, err := MaskEmail(email, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(masked)
	}
}

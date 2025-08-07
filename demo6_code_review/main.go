package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const minLocalPartLength = 1

// Improved regex to match the entire email string.
var emailPattern = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// IsValid returns true if the address is a valid email.
func IsValid(address string) bool {
	return emailPattern.MatchString(address)
}

// GetDomain returns the domain part of an email address.
// Returns an error if the address is invalid or missing "@".
func GetDomain(addr string) (string, error) {
	at := strings.LastIndex(addr, "@")
	if at == -1 || at == len(addr)-1 {
		return "", errors.New("invalid email address: missing or misplaced '@'")
	}
	return addr[at+1:], nil
}

// LocalPart returns the local part of an email address.
// Returns an error if the address is invalid or missing "@".
func LocalPart(addr string) (string, error) {
	at := strings.Index(addr, "@")
	if at == -1 {
		return "", errors.New("invalid email address: missing '@'")
	}
	return addr[:at], nil
}

// MaskedEmail masks an email so only 'show' chars of the local part remain visible.
// Returns an error if the email is invalid or show is out of range.
func MaskedEmail(e string, show int) (string, error) {
	if !IsValid(e) {
		return "", errors.New("invalid email address")
	}
	parts := strings.SplitN(e, "@", 2)
	localPart := parts[0]
	domain := parts[1]
	if show < 0 || show > len(localPart) {
		return "", errors.New("show parameter out of range")
	}
	masked := localPart[:show] + strings.Repeat("*", len(localPart)-show)
	return masked + "@" + domain, nil
}

func main() {
	email := "john.doe@example.com"
	masked, err := MaskedEmail(email, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(masked)
}

package main

import (
	"fmt"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`)

// NOTE: Chose a shorter regex but lost some precision
func IsValid(address string) bool {
	// TODO REVIEW: should we use MatchString or full match equivalent?
	return emailPattern.MatchString(address)
}

func GetDomain(addr string) string {
	// Returns everything after the last "@"
	return addr[strings.LastIndex(addr, "@")+1:]
}

func LocalPart(addr string) string {
	return strings.Split(addr, "@")[0]
}

func MaskedEmail(e string, show int) string {
	// Mask an email so only *show* chars of the local part remain visible
	if !IsValid(e) {
		return e
	}
	parts := strings.Split(e, "@")
	lp := parts[0]
	dom := parts[1]
	masked := lp[:show] + strings.Repeat("*", len(lp)-show)
	return masked + "@" + dom
}

func main() {
	email := "john.doe@example.com"
	masked := MaskedEmail(email, 2)
	fmt.Println(masked)
}

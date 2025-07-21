package main

import (
	"demo3_generate_unit_tests/emailutils"
	"fmt"
)

func main() {
	testEmails := []string{
		"user@example.com",
		"john.doe@example.co.uk",
		"user@mail.example.com",
		"invalid-email",
		"missing@domain",
		"user@@doubleat.com",
	}

	for _, email := range testEmails {
		fmt.Printf("IsValidEmail(%q) = %v\n", email, emailutils.IsValidEmail(email))
	}
}

package emailutils

import "regexp"

func IsValidEmail(email string) bool {
    pattern := `^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`
    matched, _ := regexp.MatchString(pattern, email)
    return matched
}

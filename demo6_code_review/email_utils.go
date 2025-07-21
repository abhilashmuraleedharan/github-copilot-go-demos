package emailutils

import (
    "regexp"
    "strings"
)

var emailPattern = regexp.MustCompile(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`)

func IsValid(address string) bool {
    return emailPattern.MatchString(address)
}

func GetDomain(addr string) string {
    return addr[strings.LastIndex(addr, "@")+1:]
}

func LocalPart(addr string) string {
    return strings.Split(addr, "@")[0]
}

func MaskedEmail(e string, show int) string {
    if !IsValid(e) {
        return e
    }
    parts := strings.Split(e, "@")
    lp := parts[0]
    dom := parts[1]
    masked := lp[:show] + strings.Repeat("*", len(lp)-show)
    return masked + "@" + dom
}

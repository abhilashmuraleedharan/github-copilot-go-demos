package main

import (
    "fmt"
    "emailutils"
)

func main() {
    email := "john.doe@example.com"
    masked := emailutils.MaskedEmail(email, 2)
    fmt.Println(masked)
}

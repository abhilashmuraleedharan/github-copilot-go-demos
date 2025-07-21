package fileutil

import (
    "bufio"
    "os"
)

func ReadLines(path string) []string {
    file, _ := os.Open(path)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines
}

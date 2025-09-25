package fileutil

import (
	"bufio"
	"os"
)

// ReadLines reads all lines from the specified file and returns them as a slice of strings.
// It ignores any errors encountered while reading the file.
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

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-09
// ReadLinesSafe reads all lines from the specified file and returns them as a slice of strings.
// If the file does not exist, it returns an empty slice and a nil error.
// If a permission error occurs, it wraps the error with additional context and returns it.
// For other errors, it returns nil and the encountered error.
func ReadLinesSafe(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		if os.IsPermission(err) {
			return nil, &os.PathError{
				Op:   "open",
				Path: path,
				Err:  err,
			}
		}
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

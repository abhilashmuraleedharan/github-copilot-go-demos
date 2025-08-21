package fileutil

import (
	"bufio"
	"fmt"
	"os"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
// ReadLinesSafe reads all lines from the file at the given path.
// If the file does not exist, it returns an empty slice and no error.
// If a permission error occurs, it wraps the error with additional context and returns it.
// For other errors, it returns the error as is.
func ReadLinesSafe(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("permission denied opening file %q: %w", path, err)
		}
		return nil, err
	}
	defer func() {
		cerr := file.Close()
		if cerr != nil {
			// Optionally log or handle close error
		}
	}()

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

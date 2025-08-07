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

// ReadLinesSafe reads lines from a file, returning an empty slice if the file is not found.
// Permission errors are wrapped with context and returned.
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
	if scanErr := scanner.Err(); scanErr != nil {
		return nil, scanErr
	}
	return lines, nil
}

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

// ReadLinesSafe reads all lines from the specified file path.
//
// If the file does not exist, it returns an empty slice and no error.
// If a permission error occurs, it wraps the error with additional context and returns it.
// For other errors, it returns the error as is.
func ReadLinesSafe(path string) ([]string, error) {
       file, err := os.Open(path)
       if err != nil {
	       if os.IsNotExist(err) {
		       // File not found: return empty slice, no error
		       return []string{}, nil
	       }
	       if os.IsPermission(err) {
		       // Permission error: wrap with context
		       return nil, fmt.Errorf("permission denied opening %q: %w", path, err)
	       }
	       // Other errors
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

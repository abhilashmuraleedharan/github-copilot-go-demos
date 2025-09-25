package main

import (
	"fmt"
	"strings"
)

// cleanText removes punctuation and converts text to lowercase
func cleanText(line string) string {
	line = strings.ToLower(line)
	line = strings.ReplaceAll(line, ",", "")
	line = strings.ReplaceAll(line, ".", "")
	return line
}

// countTokens increments the count for each token in the provided map
func countTokens(tokens []string, counts map[string]int) {
	for _, tok := range tokens {
		counts[tok]++
	}
}

func ProcessTranscript(lines []string) map[string]int {
	counts := make(map[string]int)
	for _, line := range lines {
		cleanLine := cleanText(line)
		tokens := strings.Fields(cleanLine)
		countTokens(tokens, counts)
	}

	var mostCommon string
	var maxCount int
	for word, count := range counts {
		if count > maxCount {
			mostCommon = word
			maxCount = count
		}
	}
	fmt.Printf("Most common word → %s (%d)\n", mostCommon, maxCount)
	return counts
}

func main() {
	lines := []string{
		"Hello, how are you?",
		"I am fine. How are you doing?",
		"I am doing well. Thank you!",
	}

	wordCounts := ProcessTranscript(lines)
	fmt.Println(wordCounts)
}

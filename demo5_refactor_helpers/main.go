package main

import (
	"fmt"
	"strings"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-09
func cleanText(line string) string {
	line = strings.ToLower(line)
	line = strings.ReplaceAll(line, ",", "")
	line = strings.ReplaceAll(line, ".", "")
	return line
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-09
func countTokens(tokens []string, counts map[string]int) {
	for _, tok := range tokens {
		counts[tok]++
	}
}

func ProcessTranscript(lines []string) map[string]int {
	counts := make(map[string]int)
	for _, line := range lines {
		line = cleanText(line)
		tokens := strings.Fields(line)
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

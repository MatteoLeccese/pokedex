package main

import "strings"

// cleanInput processes the input string by trimming spaces and splitting into words
func cleanInput(text string) []string {
	// Trim spaces from start and end
	trimmed := strings.TrimSpace(text)

	// String on lowercase
	lower := strings.ToLower(trimmed)

	return strings.Fields(lower)
}

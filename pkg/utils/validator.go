package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidGitHubUsername validates GitHub username format
func IsValidGitHubUsername(username string) bool {
	if len(username) < 1 || len(username) > 39 {
		return false
	}

	// GitHub usernames can only contain alphanumeric characters and hyphens
	// Cannot start or end with hyphen, and cannot have consecutive hyphens
	validPattern := regexp.MustCompile(`^[a-zA-Z\d](?:[a-zA-Z\d]|-(?=[a-zA-Z\d])){0,38}$`)
	return validPattern.MatchString(username)
}

// TruncateString safely truncates string with ellipsis
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}

	if maxLength < 3 {
		return s[:maxLength]
	}

	return s[:maxLength-3] + "..."
}

// FormatNumber formats large numbers with commas
func FormatNumber(n int) string {
	s := fmt.Sprintf("%d", n)

	// Add commas every 3 digits
	var result strings.Builder
	for i, char := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteByte(',')
		}
		result.WriteRune(char)
	}

	return result.String()
}

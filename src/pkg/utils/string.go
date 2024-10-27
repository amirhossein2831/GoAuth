package utils

import (
	"regexp"
	"strings"
)

// CamelToSnake converts a camel case string to snake case.
func CamelToSnake(s string) string {
	// Use a regular expression to find positions where a lowercase is followed by an uppercase or where two uppercase letters are followed by a lowercase.
	re := regexp.MustCompile("([a-z0-9])([A-Z])|([A-Z])([A-Z][a-z])")
	// Replace matches with an underscore and the lowercase version of the second part.
	snake := re.ReplaceAllString(s, "${1}${3}_${2}${4}")
	// Convert the whole string to lowercase.
	return strings.ToLower(snake)
}

package utils

import (
	"crypto/rand"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ParamDoesNotExit = errors.New("path param does not exist")
var ParamShouldBeNumeric = errors.New("path param should be numeric")

// CamelToSnake converts a camel case string to snake case.
func CamelToSnake(s string) string {
	// Use a regular expression to find positions where a lowercase is followed by an uppercase or where two uppercase letters are followed by a lowercase.
	re := regexp.MustCompile("([a-z0-9])([A-Z])|([A-Z])([A-Z][a-z])")
	// Replace matches with an underscore and the lowercase version of the second part.
	snake := re.ReplaceAllString(s, "${1}${3}_${2}${4}")
	// Convert the whole string to lowercase.
	return strings.ToLower(snake)
}

// GenerateSalt generates a new salt of the given length.
func GenerateSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GetID(param string) (uint, error) {
	if param == "" {
		return 0, ParamDoesNotExit
	}
	println("param:", param)
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, ParamShouldBeNumeric
	}

	return uint(id), nil
}

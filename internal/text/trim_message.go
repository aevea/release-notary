package text

import (
	"regexp"
	"strings"
)

var messageRegex = regexp.MustCompile("^.*\n")

// TrimMessage returns only the first line of commit message
func TrimMessage(message string) string {
	match := messageRegex.FindString(message)

	match = strings.Replace(match, "\n", "", 1)

	return match
}

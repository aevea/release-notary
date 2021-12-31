package text

import (
	"strings"
)

// TrimMessage returns only the first line of commit message
func TrimMessage(message string) string {
	match := strings.Split(message, "\n")

	final := strings.TrimRight(match[0], " ")

	return final
}

package text

import (
	"strings"
)

// BuildHistory takes commit messages and builds a complete list
func BuildHistory(messages []string) string {
	builder := strings.Builder{}

	for i := 0; i < len(messages); i++ {
		builder.WriteString("- ")
		builder.WriteString(messages[i])
		builder.WriteString("\n")
	}

	return builder.String()
}

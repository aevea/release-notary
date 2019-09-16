package text

import (
	"regexp"
	"strings"
)

var (
	expectedFormatRegex = regexp.MustCompile(`^(?P<type>\S+?)(?P<scope>\(\S+\)?)?!?:\s(?P<message>.+)\n`)
)

// ParseCommitMessage creates a slice of Commits that contain information about category and scope parsed from commit message
func ParseCommitMessage(commitMessage string) Commit {
	match := expectedFormatRegex.FindStringSubmatch(commitMessage)
	if len(match) > 0 {
		// in case of no scope
		category := match[1]
		message := match[2]
		scope := ""

		// if 2nd match is found then scope is present
		if match[3] != "" {
			message = match[3]
			scope = match[2]

			scope = strings.Replace(scope, "(", "", 1)
			scope = strings.Replace(scope, ")", "", 1)
		}

		return Commit{Category: category, Heading: message, Scope: scope}
	}

	return Commit{Category: "other", Heading: commitMessage, Scope: ""}
}

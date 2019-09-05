package text

import (
	"regexp"
	"strings"
)

var (
	expectedFormatRegex = regexp.MustCompile(`^(?P<type>\S+?)(?P<scope>\(\S+\)?)?!?:\s(?P<message>.+)$`)
)

// CategorisedCommit is a parsed commit that contains information about category, scope and heading
type CategorisedCommit struct {
	Category string
	Heading  string
	Scope    string
}

// CategoriseCommits creeates a slice of CategorizedCommits that contain information about category and scope parsed from commit message
func CategoriseCommits(commitMessages []string) []CategorisedCommit {
	var categorisedCommits []CategorisedCommit

	for _, message := range commitMessages {
		match := expectedFormatRegex.FindStringSubmatch(message)
		if match != nil {
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

			categorisedCommits = append(categorisedCommits, CategorisedCommit{Category: category, Heading: message, Scope: scope})
		} else {
			categorisedCommits = append(categorisedCommits, CategorisedCommit{Category: "other", Heading: message, Scope: ""})
		}
	}

	return categorisedCommits
}

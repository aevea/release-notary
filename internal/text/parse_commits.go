package text

import (
	"regexp"
	"strings"
)

var (
	expectedFormatRegex = regexp.MustCompile(`(?s)^(?P<category>\S+?)?(?P<scope>\(\S+\))?(?P<breaking>!?)?: (?P<heading>[^\n\r]+)?([\n\r]{2}(?P<body>.*))?`)
)

// ParseCommitMessage creates a slice of Commits that contain information about category and scope parsed from commit message
func ParseCommitMessage(commitMessage string) Commit {
	match := expectedFormatRegex.FindStringSubmatch(commitMessage)

	if len(match) > 0 {
		result := make(map[string]string)

		for i, name := range expectedFormatRegex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		category := result["category"]
		scope := result["scope"]
		heading := result["heading"]
		body := result["body"]

		scope = strings.Replace(scope, "(", "", 1)
		scope = strings.Replace(scope, ")", "", 1)

		return Commit{Category: category, Heading: heading, Scope: scope, Body: body}
	}

	return Commit{Category: "other", Heading: commitMessage, Scope: ""}
}

package text

// Sections are sections used by release notes. Please check expected-output.md
type Sections struct {
	Features []string
	Chores   []string
	Bugs     []string
	Others   []string
}

// SplitSections accepts categorised commits and further organises them into separate sections for release notes
func SplitSections(categorisedCommits []Commit) Sections {
	categoryMappings := map[string]string{
		"feat":        "feat",
		"chore":       "chore",
		"improvement": "chore",
		"bug":         "bug",
		"other":       "other",
		"fix": 		   "bug",
	}

	var features []string
	var chores []string
	var bugs []string
	var other []string

	for _, commit := range categorisedCommits {
		if categoryMappings[commit.Category] == "feat" {
			features = append(features, commit.Heading)
		}

		if categoryMappings[commit.Category] == "bug" {
			bugs = append(bugs, commit.Heading)
		}

		if categoryMappings[commit.Category] == "chore" {
			chores = append(chores, commit.Heading)
		}

		if categoryMappings[commit.Category] == "other" || categoryMappings[commit.Category] == "" {
			other = append(other, commit.Heading)
		}
	}

	return Sections{
		Features: features,
		Chores:   chores,
		Bugs:     bugs,
		Others:   other,
	}
}

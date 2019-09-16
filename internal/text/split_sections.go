package text

// Sections are sections used by release notes. Please check expected-output.md
type Sections struct {
	Features []Commit
	Chores   []Commit
	Bugs     []Commit
	Others   []Commit
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

	var features []Commit
	var chores []Commit
	var bugs []Commit
	var other []Commit

	for _, commit := range categorisedCommits {
		if categoryMappings[commit.Category] == "feat" {
			features = append(features, commit)
		}

		if categoryMappings[commit.Category] == "bug" {
			bugs = append(bugs, commit)
		}

		if categoryMappings[commit.Category] == "chore" {
			chores = append(chores, commit)
		}

		if categoryMappings[commit.Category] == "other" || categoryMappings[commit.Category] == "" {
			other = append(other, commit)
		}
	}

	return Sections{
		Features: features,
		Chores:   chores,
		Bugs:     bugs,
		Others:   other,
	}
}

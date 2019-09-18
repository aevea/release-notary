// Package release holds the common Release struct shared across all other packages
package release

// Release holds all the important metadata and message of a release
type Release struct {
	ID   int
	Tag  string
	Name string
	// Original message of a release
	Message string
	// Release Notes to update release with
	ReleaseNotes string
}

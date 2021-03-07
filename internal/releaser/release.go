package releaser

import (
	"strings"

	"github.com/apex/log"
)

// Release pulls latest tag and its text, then it appends passed release notes
func (r *Releaser) Release(releaseNotes string) error {
	latestRelease, err := r.service.LatestRelease()

	if err != nil {
		return err
	}

	if strings.Contains(latestRelease.Message, releaseNotes) {
		log.Warn("Release already contains these release notes")
		return nil
	}

	builder := strings.Builder{}
	builder.WriteString(latestRelease.Message)
	builder.WriteString(releaseNotes)

	latestRelease.ReleaseNotes = builder.String()

	err = r.service.Publish(latestRelease)

	if err != nil {
		return err
	}

	return nil
}

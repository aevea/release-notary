package releaser

import "strings"

// Release pulls latest tag and its text, then it appends passed release notes
func (r *Releaser) Release(releaseNotes string) error {
	latestRelease, err := r.service.LatestRelease()

	if err != nil {
		return err
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

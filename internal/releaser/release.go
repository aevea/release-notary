package releaser

import "strings"

// Release pulls latest tag and its text, then it appends passed release notes
func (r *Releaser) Release(releaseNotes string) error {
	latestRelease, err := r.service.LatestRelease()

	if err != nil {
		return err
	}

	builder := strings.Builder{}
	builder.WriteString(latestRelease)
	builder.WriteString(releaseNotes)

	err = r.service.Publish(builder.String(), "")

	if err != nil {
		return err
	}

	return nil
}

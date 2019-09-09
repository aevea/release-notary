package releaser

// Service describes the basic usage needed from a service such as Github
type Service interface {
	LatestRelease() (string, error)
	Publish(string, string) error
}

// Releaser holds all functionality related to releasing.
type Releaser struct {
	service Service
}

// Options are used to initialize releaser
type Options struct {
	Provider Provider
	DryRun   bool
}

// CreateReleaser initializes an instance of Releaser
func CreateReleaser(options Options) *Releaser {
	if options.Provider == "github" {
		return &Releaser{}
	}

	return &Releaser{}
}

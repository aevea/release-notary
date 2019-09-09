package releaser

// Provider refers to git service provider such as Gitlab/Github
type Provider string

const (
	// Github sets Releaser to use Github specific commands
	Github Provider = "github"
	// Gitlab sets Releaser to use Gitlab specific commands
	Gitlab Provider = "gitlab"
)

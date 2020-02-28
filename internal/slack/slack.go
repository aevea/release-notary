package slack

// Slack is the struct holding all the methods to work with the Slack integration.
type Slack struct {
	WebHookURL string
}

type GitRemoter interface {
	GetRemoteURL() string
	Host() string
	Project() string
}

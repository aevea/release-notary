package slack

// MockRemote serves to mock the GitRemoter interface
type MockRemote struct {
}

func (r MockRemote) Host() string {
	return "example.com"
}

func (r MockRemote) Project() string {
	return "some/thing"
}

func (r MockRemote) GetRemoteURL() string {
	return "https://example.com/some/thing"
}

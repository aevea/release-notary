package gitlab

import (
	"net/http"
	"time"
)

func createClient(token string) *http.Client {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	rt := attachHeader(client.Transport)
	rt.Set("PRIVATE-TOKEN", token)
	client.Transport = rt

	return &client
}

type withHeader struct {
	http.Header
	rt http.RoundTripper
}

func attachHeader(rt http.RoundTripper) withHeader {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return withHeader{Header: make(http.Header), rt: rt}
}

func (h withHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.Header {
		req.Header[k] = v
	}

	return h.rt.RoundTrip(req)
}

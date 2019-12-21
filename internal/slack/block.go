package slack

type content struct {
	// expected type is "mrkdwn"
	Type string `json:"type,omitempty"`
	// markdown compliant message
	Text string `json:"text,omitempty"`
}

// Block holds the different blocks uses in the Slack block API
// Hmm... omitempty doesn't omit zero structs https://github.com/golang/go/issues/11939
type Block struct {
	Type     string    `json:"type"`
	Section  content   `json:"text,omitempty"`
	Elements []content `json:"elements,omitempty"`
}

// WebhookMessage is the specific structure that Slack uses for the Webhook API
type WebhookMessage struct {
	Blocks []Block `json:"blocks"`
}

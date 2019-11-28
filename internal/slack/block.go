package slack

type content struct {
	// expected type is "mrkdwn"
	Type string `json:"type"`
	// markdown compliant message
	Text string `json:"text"`
}

// Block holds the different blocks uses in the Slack block API, due to some unfortunate naming json:"text" actually contains subsections.
type Block struct {
	// expected type is "section"
	Type    string  `json:"type"`
	Section content `json:"text"`
}

// WebhookMessage is the specific structure that Slack uses for the Webhook API
type WebhookMessage struct {
	Blocks []Block `json:"blocks"`
}

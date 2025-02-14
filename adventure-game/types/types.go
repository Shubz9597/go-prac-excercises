package types

type OptionKey struct {
	Text string `json: "text"`
	Arc  string `json: "arc"`
}

type Stories struct {
	Title   string      `json: "title"`
	Story   []string    `json: "story"`
	Options []OptionKey `json: "options"`
}

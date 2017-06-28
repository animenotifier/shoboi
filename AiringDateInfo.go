package shoboi

// AiringDateInfo is the airing information for a single episode on a single channel.
type AiringDateInfo struct {
	PID           string      `json:"PID"`
	TID           string      `json:"TID"`
	ChannelID     string      `json:"ChID"`
	ChannelName   string      `json:"ChName"`
	ChannelEPGURL string      `json:"ChEPGURL"`
	EpisodeNumber string      `json:"Count"`
	StartTime     string      `json:"StTime"`
	EndTime       string      `json:"EdTime"`
	SubTitle2     string      `json:"SubTitle2"`
	ProgComment   string      `json:"ProgComment"`
	ConfFlag      interface{} `json:"ConfFlag"`
}

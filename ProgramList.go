package shoboi

// ProgramList lists all programs where the given episode is aired.
type ProgramList struct {
	Programs map[string]AiringDateInfo `json:"Programs"`
}

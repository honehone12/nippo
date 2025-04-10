package schema

type Calls struct {
	MorningCalls []MorningCall `json:"morning_calls"`
	EveningCalls []EveningCall `json:"evening_calls"`
}

type Photos struct {
	MorningAlc string `json:"morning_alc"`
	EveningAlc string `json:"evening_alc"`
	Meter      string `json:"meter"`
}

package schema

type Calls struct {
	MorningCalls []MorningCall `json:"morning_calls"`
	EveningCalls []EveningCall `json:"evening_calls"`
}

type DailyReportPrint struct {
	DailyReport *DailyReport `json:"daily_report"`
	MorningCall *MorningCall `json:"morning_call"`
	EveningCall *EveningCall `json:"evening_call"`

	Locations []Location `json:"locations"`
	Waitings  []Waiting  `json:"waitings"`
	Loadings  []Loading  `json:"loadings"`
	Resings   []Resting  `json:"restings"`
}

type Photos struct {
	MorningAlc []byte `json:"morning_alc"`
	EveningAlc []byte `json:"evening_alc"`
	Meter      []byte `json:"meter"`
}

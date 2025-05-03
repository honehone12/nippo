package schema

import "time"

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

type UserMini struct {
	Id        []uint8   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      CleanText `json:"name"`
}

type Users struct {
	Admin      bool       `json:"admin"`
	Users      []UserMini `json:"users"`
	Invitables []UserMini `json:"invitables"`
}

type AuthOrg struct {
	OrgId []uint8 `json:"org_id"`
}

type Invited struct {
	UserEmail CleanText `json:"user_email"`
}

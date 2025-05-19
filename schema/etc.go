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

	Sites     []Site     `json:"sites"`
	Locations []Location `json:"locations"`
	Waitings  []Waiting  `json:"waitings"`
	Loadings  []Loading  `json:"loadings"`
	Resings   []Resting  `json:"restings"`
}

type Photos struct {
	MorningAlc []byte `json:"morning_alc"`
	EveningAlc []byte `json:"evening_alc"`
	MorningMtr []byte `json:"morning_mtr"`
	EveningMtr []byte `json:"evening_mtr"`
}

type UserMini struct {
	Id        []uint8   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      CleanText `json:"name"`
}

type DailyReportMini struct {
	Id        []uint8   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfo struct {
	User       UserMini `json:"user"`
	Invitable  bool     `json:"invitable"`
	Promotable bool     `json:"promotable"`
}

type Users struct {
	Admin bool       `json:"admin"`
	Users []UserInfo `json:"users"`
}

type AuthOrg struct {
	OrgId []uint8 `json:"org_id"`
}

type Invited struct {
	UserEmail CleanText `json:"user_email"`
}

type Promoted struct {
	SubAdmin bool `json:"sub_admin"`
}

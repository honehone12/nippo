package schema

import (
	"time"

	"gorm.io/gorm"
)

const (
	ID_LEN             = 16
	ID_STR_LEN         = 36
	LINE_ID_LEN        = 33
	MAX_SHORT_TEXT_LEN = 100
	MAX_LONG_TEXT_LEN  = 5000
)

type Org struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name   CleanText `gorm:"not null;size:100" json:"name"`
	Email  CleanText `gorm:"unique;not null;size:100" json:"email"`
	LineId CleanText `gorm:"unique;type:char(33)" json:"line_id"`
	Users  []User    `gorm:"foreignKey:OrgId;references:Id" json:"-"`
}

type User struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name             CleanText `gorm:"size:100" json:"name"`
	SubAdmin         bool      `json:"sub_admin"`
	Email            CleanText `gorm:"unique;size:100;index" json:"email"`
	LineId           CleanText `gorm:"unique;not null;type:char(33);index" json:"line_id"`
	CurrentOperation []uint8   `gorm:"type:binary(16)" json:"-"`
	CurrentDuty      []uint8   `gorm:"type:binary(16)" json:"-"`
	NextState        StateIdx  `json:"-"`

	MorningCalls []MorningCall `gorm:"foreignKey:UserId;references:Id" json:"-"`
	EveningCalls []EveningCall `gorm:"foreignKey:UserId;references:Id" json:"-"`
	DailyReports []DailyReport `gorm:"foreignKey:UserId;references:Id" json:"-"`

	OrgId []uint8 `gorm:"not null;type:binary(16)" json:"org_id"`
}

type MorningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null;index" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Caller               CleanText  `gorm:"size:100" json:"caller"`
	Name                 CleanText  `gorm:"size:100" json:"name"`
	LicenseExpiration    time.Time  `gorm:"default:null" json:"license_expiration"`
	CarNumber            CleanText  `gorm:"size:100" json:"car_number"`
	InspectionExpiration time.Time  `gorm:"default:null" json:"inspection_expiration"`
	Method               CallMethod `json:"method"`
	UsingAlcChecker      bool       `json:"using_alc_checker"`
	AlcCheck             bool       `json:"alc_check"`
	AlcPhoto             string     `gorm:"size:100" json:"alc_photo"`
	HealthCheck          bool       `json:"health_check"`
	CarCheck             bool       `json:"car_check"`
	Note                 CleanText  `gorm:"size:5000" json:"note"`

	UserId []uint8 `gorm:"not null;type:binary(16)" json:"user_id"`
}

type EveningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null;index" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Caller          CleanText  `gorm:"size:100" json:"caller"`
	Name            CleanText  `gorm:"size:100" json:"name"`
	CarNumber       CleanText  `gorm:"size:100" json:"car_number"`
	Method          CallMethod `json:"method"`
	UsingAlcChecker bool       `json:"using_alc_checker"`
	AlcCheck        bool       `json:"alc_check"`
	AlcPhoto        string     `gorm:"size:100" json:"alc_photo"`
	Note            CleanText  `gorm:"size:5000" json:"note"`

	UserId []uint8 `gorm:"not null;type:binary(16)" json:"user_id"`
}

type DailyReport struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null;index" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	MorningCallId []uint8 `gorm:"not null;type:binary(16)" json:"morning_call_id"`
	EveningCallId []uint8 `gorm:"type:binary(16)" json:"evening_call_id"`

	Name      CleanText `gorm:"size:100" json:"name"`
	CarNumber CleanText `gorm:"size:100" json:"car_number"`

	MorningMeter      uint64    `json:"morning_meter"`
	MorningMeterPhoto string    `gorm:"size:100" json:"morning_meter_photo"`
	EveningMeter      uint64    `json:"evening_meter"`
	EveningMeterPhoto string    `gorm:"size:100" json:"evening_meter_photo"`
	NonDutyDistance   uint64    `json:"non_duty_distance"`
	Trouble           CleanText `gorm:"size:5000" json:"trouble"`
	Note              CleanText `gorm:"size:5000" json:"note"`

	Sites     []Site     `gorm:"foreignKey:DailyReportId;references:Id" json:"-"`
	Locations []Location `gorm:"foreignKey:DailyReportId;references:Id" json:"-"`
	Waitings  []Waiting  `gorm:"foreignKey:DailyReportId;references:Id" json:"-"`
	Loadings  []Loading  `gorm:"foreignKey:DailyReportId;references:Id" json:"-"`
	Restings  []Resting  `gorm:"foreignKey:DailyReportId;references:Id" json:"-"`

	UserId []uint8 `gorm:"not null;type:binary(16)" json:"user_id"`
}

type Site struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name      CleanText `gorm:"not null;size:100" json:"name"`
	ShortNote CleanText `gorm:"size:100" json:"short_note"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)" json:"daily_report_id"`
}

type Location struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Label     CleanText `gorm:"size:100" json:"label"`
	Address   CleanText `gorm:"size:100" json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ShortNote CleanText `gorm:"size:100" json:"short_note"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)" json:"daily_report_id"`
}

type Waiting struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Label     CleanText `gorm:"size:100" json:"label"`
	Address   CleanText `gorm:"size:100" json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Note      CleanText `gorm:"size:5000" json:"note"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)" json:"daily_report_id"`
}

type Loading struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Label         CleanText `gorm:"size:100" json:"label"`
	Address       CleanText `gorm:"size:100" json:"address"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	ShippingCheck bool      `json:"shipping_check"`
	Note          CleanText `gorm:"size:5000" json:"note"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)" json:"daily_report_id"`
}

type Resting struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Label     CleanText `gorm:"size:100" json:"label"`
	Address   CleanText `gorm:"size:100" json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ShortNote CleanText `gorm:"size:100" json:"short_note"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)" json:"daily_report_id"`
}

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
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name   string `gorm:"not null;size:100"`
	LineId string `gorm:"type:char(33);index"`
	Users  []User `gorm:"foreignKey:OrgId;references:Id"`
}

type User struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string  `gorm:"not null;size:100"`
	LineId      string  `gorm:"unique;not null;type:char(33);index"`
	CarNumber   string  `gorm:"size:100"`
	OrgId       []uint8 `gorm:"not null;type:binary(16)"`
	CurrentCall []uint8 `gorm:"type:binary(16)"`
	CurrentDuty []uint8 `gorm:"type:binary(16)"`
	NextState   StateIdx

	MorningCalls []MorningCall `gorm:"foreignKey:UserId;references:Id"`
	EveningCalls []EveningCall `gorm:"foreignKey:UserId;references:Id"`
	DailyReports []DailyReport `gorm:"foreignKey:UserId;references:Id"`
}

type MorningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Caller          string `gorm:"size:100"`
	Name            string `gorm:"size:100"`
	CarNumber       string `gorm:"size:100"`
	Method          CallMethod
	UsingAlcChecker bool
	AlcCheck        bool
	Photo           string `gorm:"size:100"`
	HealthCheck     bool
	CarCheck        bool
	Note            string `gorm:"size:5000"`

	UserId []uint8 `gorm:"not null;type:binary(16)"`
}

type EveningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Caller          string `gorm:"size:100"`
	Name            string `gorm:"size:100"`
	CarNumber       string `gorm:"size:100"`
	Method          CallMethod
	UsingAlcChecker bool
	AlcCheck        bool
	Photo           string `gorm:"size:100"`
	Note            string `gorm:"size:5000"`

	UserId []uint8 `gorm:"not null;type:binary(16)"`
}

type DailyReport struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MorningCallId []uint8 `gorm:"not null;type:binary(16)"`
	EveningCallId []uint8 `gorm:"type:binary(16)"`

	Name      string `gorm:"size:100"`
	CarNumber string `gorm:"size:100"`

	MeterPhoto string `gorm:"size:5000"`
	Troubles   string `gorm:"size:5000"`
	Note       string `gorm:"size:5000"`

	Locations []Location `gorm:"foreignKey:DailyReportId;references:Id"`
	Waitings  []Waiting  `gorm:"foreignKey:DailyReportId;references:Id"`
	Loadings  []Loading  `gorm:"foreignKey:DailyReportId;references:Id"`
	Restings  []Resting  `gorm:"foreignKey:DailyReportId;references:Id"`

	UserId []uint8 `gorm:"not null;type:binary(16)"`
}

type Location struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Label     string `gorm:"size:100"`
	Address   string `gorm:"size:100"`
	Latitude  float64
	Longitude float64

	DailyReportId []uint8 `gorm:"not null;type:binary(16)"`
}

type Waiting struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Label     string `gorm:"size:100"`
	Address   string `gorm:"size:100"`
	Latitude  float64
	Longitude float64
	Note      string `gorm:"size:5000"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)"`
}

type Loading struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Label         string `gorm:"size:100"`
	Address       string `gorm:"size:100"`
	Latitude      float64
	Longitude     float64
	ShippingCheck bool
	Note          string `gorm:"size:5000"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)"`
}

type Resting struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Label     string `gorm:"size:100"`
	Address   string `gorm:"size:100"`
	Latitude  float64
	Longitude float64
	Note      string `gorm:"size:5000"`

	DailyReportId []uint8 `gorm:"not null;type:binary(16)"`
}

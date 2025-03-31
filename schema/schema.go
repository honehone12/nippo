package schema

import (
	"time"

	"gorm.io/gorm"
)

const (
	ID_LEN       = 16
	ID_STR_LEN   = 36
	MAX_NAME_LEN = 50
	LINE_ID_LEN  = 33
)

type Org struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name   string `gorm:"not null;size:50"`
	LineId string `gorm:"unique;type:char(33);index"`
	Users  []User `gorm:"foreignKey:OrgId;references:Id"`
}

type User struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name      string   `gorm:"not null;size:50"`
	LineId    string   `gorm:"unique;not null;type:char(33);index"`
	CarNumber string   `gorm:"size:50"`
	OrgId     []uint8  `gorm:"not null;type:binary(16)"`
	NextState StateIdx `gorm:"default:4294967295"`

	MorningCalls []MorningCall `gorm:"foreignKey:UserId;references:Id"`
	EveningCalls []EveningCall `gorm:"foreignKey:UserId;references:Id"`
	DailyReports []DailyReport `gorm:"foreignKey:UserId;references:Id"`
}

type MorningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Caller          string `gorm:"size:50"`
	Name            string `gorm:"size:50"`
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

	Caller          string `gorm:"size:50"`
	Name            string `gorm:"size:50"`
	CarNumber       string `gorm:"size:100"`
	Method          CallMethod
	UsingAlcChecker bool
	AlcCheck        bool
	Photo           string `gorm:"size:100"`
	Note            string `gorm:"size:5000"`

	UserId []uint8 `gorm:"not null;type:binary(16)"`
}

type DailyReport struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	StartAt       time.Time `gorm:"not null;index"`
	MorningCallId []uint8   `gorm:"type:binary(16)"`
	EndAt         time.Time
	EveningCallId []uint8 `gorm:"type:binary(16)"`

	Name      string `gorm:"size:50"`
	CarNumber string `gorm:"size:50"`

	Pickup    uint32
	Deliverry uint32
	Odo       float64
	Trip      float64
	Rests     string `gorm:"size:2000"`
	Highway   string `gorm:"size:2000"`
	Troubles  string `gorm:"size:2000"`
	Note      string `gorm:"size:5000"`

	Locations []Location
	Standbys  []Standby
	Loadings  []Loading

	UserId []uint8 `gorm:"not null;type:binary(16)"`
}

type Location struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Address   string `gorm:"size:100"`
	Latitude  float64
	Longitude float64
}

type Standby struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	StartAt time.Time `gorm:"not null;index"`
	Place   string    `gorm:"size:100"`
	EndAt   time.Time `gorm:"not null;index"`
	Note    string    `gorm:"size:5000"`
}

type Loading struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	StartAt       time.Time `gorm:"not null;index"`
	Place         string    `gorm:"size:100"`
	ShippingCheck bool
	EndAt         time.Time `gorm:"not null;index"`
	Note          string    `gorm:"size:5000"`
}

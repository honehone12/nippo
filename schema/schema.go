package schema

import (
	"nippo/message/state"
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
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"size:50"`
	LineId        string `gorm:"unique;type:char(33);index"`
	AdminLineId   string `gorm:"unique;type:char(33)"`
	LastOperation *state.OrgOperation
	Users         []User `gorm:"foreignKey:OrgId;references:Id"`
}

type User struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"not null;size:50"`
	LineId        string `gorm:"unique;not null;type:char(33);index"`
	LastOperation *state.UserOperation
	OrgId         []uint8       `gorm:"not null;type:binary(16)"`
	MorningCalls  []MorningCall `gorm:"foreignKey:UserId;references:Id"`
	EveningCalls  []EveningCall `gorm:"foreignKey:UserId;references:Id"`
	DailyReports  []DailyReport `gorm:"foreignKey:UserId;references:Id"`
}

type MorningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm:"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type EveningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm:"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type DailyReport struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm:"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

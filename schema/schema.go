package schema

import (
	"nippo/operation"
	"time"

	"gorm.io/gorm"
)

const (
	ID_LEN       int = 16
	MAX_NAME_LEN int = 50
	LINE_ID_LEN  int = 33
)

type Org struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"not null;size:50"`
	LineId        string `gorm:"unique;not null;type:char(33);index"`
	AdminLineId   string `gorm:"unique;not null;type:char(33)"`
	LastOperation *operation.OrgOperation
	Users         []User
}

type User struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"not null;size:50"`
	LineId        string `gorm:"unique;not null;type:char(33);index"`
	LastOperation *operation.UserOperation
	OrgId         []uint8 `gorm;"not null;type:binary(16)"`
	MoringCalls   []MorningCall
	EveningCalls  []EveningCall
	DailyReports  []DailyReport
}

type MorningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm;"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type EveningCall struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm;"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type DailyReport struct {
	Id        []uint8        `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId []uint8   `gorm;"not null;type:binary(16)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

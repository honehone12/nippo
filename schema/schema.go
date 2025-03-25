package schema

import (
	"nippo/operation"
	"time"

	"gorm.io/gorm"
)

const (
	ID_LEN       int = 36
	MAX_NAME_LEN int = 50
	LINE_ID_LEN  int = 33
)

type Org struct {
	Id        string         `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"not null;size:50"`
	LineId        string `gorm:"unique;not null;type:char(33);index"`
	AdminLineId   string `gorm:"unique;not null;type:char(33)"`
	LastOperation *operation.OrgOperation
	Users         []User `gorm:"foreignKey:OrgId;references:Id"`
}

type User struct {
	Id        string         `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"not null;size:50"`
	LineId        string `gorm:"unique;not null;type:char(33);index"`
	LastOperation *operation.UserOperation
	OrgId         string        `gorm:"not null;type:varchar(36)"`
	MorningCalls  []MorningCall `gorm:"foreignKey:UserId;references:Id"`
	EveningCalls  []EveningCall `gorm:"foreignKey:UserId;references:Id"`
	DailyReports  []DailyReport `gorm:"foreignKey:UserId;references:Id"`
}

type MorningCall struct {
	Id        string         `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId string    `gorm:"not null;type:varchar(36)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type EveningCall struct {
	Id        string         `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId string    `gorm:"not null;type:varchar(36)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

type DailyReport struct {
	Id        string         `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId string    `gorm:"not null;type:varchar(36)"`
	Date   time.Time `gorm:"not null;type:date;index"`
}

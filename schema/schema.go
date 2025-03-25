package schema

import (
	"time"

	"gorm.io/gorm"
)

type Org struct {
	Id        []uint8 `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	Name  string `gorm:"not null;size:25"`
	Admin User
	Users []User
}

type User struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	Name         string `gorm:"not null; size:50"`
	OrgId        uint64
	MoringCalls  []MorningCall
	EveningCalls []EveningCall
	DailyReports []DailyReport
}

type MorningCall struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	UserId uint64
	Date   time.Time `gorm:"not null;type:date;index"`
}

type EveningCall struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	UserId uint64
	Date   time.Time `gorm:"not null;type:date;index"`
}

type DailyReport struct {
	Id        []uint8   `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	UserId uint64
	Date   time.Time `gorm:"not null;type:date;index"`
}

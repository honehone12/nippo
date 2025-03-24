package schema

import (
	"time"

	"gorm.io/gorm"
)

type Org struct {
	ID        uint `gorm:"primaryKey;default:auto_random()"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	Name  string `gorm:"not null;size:50"`
	Admin User
	Users []User
}

type User struct {
	ID        uint `gorm:"primaryKey;default:auto_random()"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	Name    string `gorm:"not null; size:50"`
	OrgId   uint
	Calls   []Call
	Reports []Report
}

type Call struct {
	ID        uint `gorm:"primaryKey;default:auto_random()"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	UserId    uint
	Date      time.Time `gorm:"type:date;index:date"`
	IsMorning bool
}

type Report struct {
	ID        uint `gorm:"primaryKey;default:auto_random()"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	UserId uint
	Date   time.Time `gorm:"type:date;index:date"`
}

package schema

import "gorm.io/gorm"

type Org struct {
	gorm.Model
}

type Admin struct {
	gorm.Model
}

type User struct {
	gorm.Model
}

type Call struct {
	gorm.Model
}

type Report struct {
	gorm.Model
}

package database

import (
	"nippo/schema"
)

type Database interface {
	CreateOrg(name string) ([]uint8, error)
	ReadOrg(idStr string) (*schema.Org, error)
	UpdateOrgLineId(idStr string, newLineId string) error

	CreateUser(orgIdStr string, lineId string) ([]uint8, error)
	ReadUser(lineId string) (*schema.User, error)
	UpdateUserName(id []uint8, name string) error
	UpdateUserCarNumber(id []uint8, carNumber string) error
	UpdateUserNextState(id []uint8, nextState schema.StateIdx) error
	UpdateUserCurrentDuty(id []uint8, duty []uint8) error

	CreateMorningCall(userId []uint8) ([]uint8, error)

	CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error)

	CreateEveningCall(userId []uint8) ([]uint8, error)
}

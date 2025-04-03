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
	UpdateUserCurrentCall(id []uint8, call []uint8) error
	UpdateUserCurrentDuty(id []uint8, duty []uint8) error
	UpdateUserNextState(id []uint8, nextState schema.StateIdx) error

	CreateMorningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateMoCallHealth(userId []uint8) error
	UpdateMoCallCarCheck(userId []uint8) error
	UpdateMoCallNote(userId []uint8, note string) error

	CreateEveningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateEvCallNote(userId []uint8, note string) error

	CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error)
}

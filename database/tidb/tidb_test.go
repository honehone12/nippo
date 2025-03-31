package tidb

import (
	"bytes"
	"errors"
	"nippo/database"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var _ database.Database = &TiDb{}

func initTest() (database.Database, error) {
	godotenv.Load("migration/.env")

	tiUri := os.Getenv("TIDB_URI")
	if len(tiUri) == 0 {
		return nil, errors.New("env TIDB_URI is not set")
	}

	return Open(tiUri)
}

func TestUserAndOrg(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	orgName := "テスト　企業"
	bId, err := db.CreateOrg(orgName)
	if err != nil {
		t.Fatal(err)
	}

	orgId, err := uuid.FromBytes(bId)
	if err != nil {
		t.Fatal(err)
	}

	orgLineId := "U9198cf6745fc0d808977bdb0b9f22995"
	err = db.UpdateOrgLineId(orgId.String(), orgLineId)
	if err != nil {
		t.Fatal(err)
	}

	userName := "テスト　ユーザー太郎"
	userLineId := "U8189cf6745fc0d808977bdb0b9f22885"
	userId, err := db.CreateUser(
		orgId.String(),
		userName,
		userLineId,
	)
	if err != nil {
		t.Fatal(err)
	}

	carNumber := "く００９１２ー２３"
	err = db.UpdateUserCarNumber(userLineId, carNumber)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserNextState(userLineId, 2)
	if err != nil {
		t.Fatal(err)
	}

	readOrg, err := db.ReadOrg(orgId.String())
	if err != nil {
		t.Fatal(err)
	}
	if readOrg.LineId != orgLineId || readOrg.Name != orgName {
		t.Fatal("read wrong data")
	}

	readUser, err := db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.Id, userId) || readUser.Name != userName {
		t.Fatal("read wrong user")
	}
	if readUser.CarNumber != carNumber || readUser.NextState != 2 {
		t.Fatal("read wrong user")
	}
}

func TestCallAndReport(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	userLineId := "U8189cf6745fc0d808977bdb0b9f22885"
	morningId, err := db.CreateMorningCall(userLineId)
	if err != nil {
		t.Fatal(err)
	}

	reportId, err := db.CreateDailyReport(userLineId, morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentDuty(userLineId, reportId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateEveningCall(userLineId)
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(readUser.CurrentDuty, reportId) {
		t.Fatal("read wrong user")
	}
}

package test

import (
	"bytes"
	"errors"
	"nippo/database"
	"nippo/database/tidb"

	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var _ database.Database = &tidb.TiDb{}

func initTest() (database.Database, error) {
	godotenv.Load("../bin/.env")

	tiUri := os.Getenv("DB_URI")
	if len(tiUri) == 0 {
		return nil, errors.New("env DB_URI is not set")
	}

	return tidb.Open(tiUri)
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

	userLineId := "U8189cf6745fc0d808977bdb0b9f22885"
	userId, err := db.CreateUser(orgId.String(), userLineId)
	if err != nil {
		t.Fatal(err)
	}

	userName := "テスト　ユーザー太郎"
	err = db.UpdateUserName(userId, userName)
	if err != nil {
		t.Fatal(err)
	}

	carNumber := "て００９１２ー２３"
	err = db.UpdateUserCarNumber(userId, carNumber)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserNextState(userId, 2)
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
	user, err := db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}

	photo := "user-id-009.jpg"
	morningId, err := db.CreateMorningCall(user.Id, photo)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentCall(user.Id, morningId)
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentCall, morningId) {
		t.Fatal("read wrong user")
	}

	reportId, err := db.CreateDailyReport(user.Id, morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentDuty(user.Id, reportId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallCarCheck(user.Id)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallCarCheck(user.Id)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallNote(user.Id, "a: \nb: ")
	if err != nil {
		t.Fatal(err)
	}

	eveningId, err := db.CreateEveningCall(user.Id)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentCall(user.Id, eveningId)
	if err != nil {
		t.Fatal(err)
	}

	readUser, err = db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentCall, eveningId) {
		t.Fatal("read wrong user")
	}
	if !bytes.Equal(readUser.CurrentDuty, reportId) {
		t.Fatal("read wrong user")
	}

	err = db.UpdateUserCurrentCall(user.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentDuty(user.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	readUser, err = db.ReadUser(userLineId)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentCall, nil) {
		t.Fatal("read wrong user")
	}
	if !bytes.Equal(readUser.CurrentDuty, nil) {
		t.Fatal("read wrong user")
	}
}

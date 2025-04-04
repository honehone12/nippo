package test

import (
	"bytes"
	"nippo/database"
	"nippo/database/tidb"

	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const ORG_NAME = "テスト　企業"
const ORG_LINE_ID = "U9198cf6745fc0d808977bdb0b9f22995"
const USER_LINE_ID = "U8189cf6745fc0d808977bdb0b9f22885"
const USER_NAME = "テスト　ユーザー太郎"
const CAR_NUMBER = "て００９１２ー２３"
const ALC_PHOTO = "user-id-009.jpg"
const MTR_PHOTO = "user-id-018.jpg"

var db database.Database = &tidb.TiDb{}
var orgIdStr string
var userId []uint8
var morningId []uint8
var reportId []uint8

func TestOpenDb(t *testing.T) {
	godotenv.Load("../bin/.env")

	tiUri := os.Getenv("DB_URI")
	if len(tiUri) == 0 {
		t.Fatal("env DB_URI is not set")
	}

	var err error
	db, err = tidb.Open(tiUri)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateOrg(t *testing.T) {
	bId, err := db.CreateOrg(ORG_NAME)
	if err != nil {
		t.Fatal(err)
	}

	orgId, err := uuid.FromBytes(bId)
	if err != nil {
		t.Fatal(err)
	}
	orgIdStr = orgId.String()

	err = db.UpdateOrgLineId(orgIdStr, ORG_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserSetting(t *testing.T) {
	var err error
	userId, err = db.CreateUser(orgIdStr, USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserName(userId, USER_NAME)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCarNumber(userId, CAR_NUMBER)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserNextState(userId, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func ReadOrgAndUser(t *testing.T) {
	readOrg, err := db.ReadOrg(orgIdStr)
	if err != nil {
		t.Fatal(err)
	}
	if readOrg.LineId != ORG_LINE_ID || readOrg.Name != ORG_NAME {
		t.Fatal("read wrong data")
	}

	readUser, err := db.ReadUser(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.Id, userId) || readUser.Name != USER_NAME {
		t.Fatal("read wrong user")
	}
	if readUser.CarNumber != CAR_NUMBER || readUser.NextState != 2 {
		t.Fatal("read wrong user")
	}
}

func TestMorningCall(t *testing.T) {
	var err error
	morningId, err = db.CreateMorningCall(userId, ALC_PHOTO)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentCall(userId, morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallHealth(morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallCarCheck(morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateMoCallNote(morningId, "【ラベル】\n無し\n【ラベル】\n無し")
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentCall, morningId) {
		t.Fatal("read wrong user")
	}
}

func TestDailyReport(t *testing.T) {
	var err error
	reportId, err = db.CreateDailyReport(userId, morningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentDuty(userId, reportId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateDReportMtrPhoto(reportId, MTR_PHOTO)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateDReportTrouble(reportId, "【ラベル】\n特になし")
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateDReportNote(reportId, "【ラベル】\n無し")
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentDuty, reportId) {
		t.Fatal("read wrong user")
	}
}

func TestEveningCall(t *testing.T) {
	eveningId, err := db.CreateEveningCall(userId, ALC_PHOTO)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentCall(userId, eveningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateDReportEvCall(reportId, eveningId)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateEvCallNote(eveningId, "【ラベル】\n無し\n【ラベル】\n無し")
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentCall, eveningId) {
		t.Fatal("read wrong user")
	}
}

func TestOnEnd(t *testing.T) {
	err := db.UpdateUserCurrentCall(userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateUserCurrentDuty(userId, nil)
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser(USER_LINE_ID)
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

package test

import (
	"bytes"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/timeconv"
	"time"

	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const EMAIL = "test@test.com"
const ORG_NAME = "テスト　企業"
const ORG_LINE_ID = "U9198cf6745fc0d808977bdb0b9f22995"
const USER_LINE_ID = "U8189cf6745fc0d808977bdb0b9f22885"
const USER_NAME = "テスト　ユーザー太郎"
const CAR_NUMBER = "て００９１２ー２３"
const ALC_PHOTO = "user-id-009.jpg"
const MTR_PHOTO = "user-id-018.jpg"

var db database.Database = tidb.TiDb{}
var orgId []uint8
var userId []uint8
var morningId []uint8
var reportId []uint8
var eveningId []uint8
var loadId []uint8
var waitId []uint8
var restId []uint8
var locId []uint8

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
	err := db.Tx(func(tx database.Transaction) error {
		bId, err := tx.CreateOrg(ORG_NAME, EMAIL)
		if err != nil {
			return err
		}

		orgUuid, err := uuid.FromBytes(bId)
		if err != nil {
			return err
		}
		orgId = orgUuid[:]

		return tx.UpdateOrgLineId(orgId, ORG_LINE_ID)
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserSetting(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		userId, err = tx.CreateUser(orgId, USER_LINE_ID)
		if err != nil {
			return err
		}

		err = tx.UpdateUserName(userId, USER_NAME)
		if err != nil {
			return err
		}

		err = tx.UpdateUserEmail(userId, EMAIL)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCarNumber(userId, CAR_NUMBER)
		if err != nil {
			return err
		}

		return tx.UpdateUserNextState(userId, 9)
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMorningCall(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		morningId, err = tx.CreateMorningCall(userId, USER_NAME, CAR_NUMBER, ALC_PHOTO)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, morningId)
		if err != nil {
			return err
		}

		err = tx.UpdateMCallHealth(morningId)
		if err != nil {
			return err
		}

		err = tx.UpdateMCallCarCheck(morningId)
		if err != nil {
			return err
		}

		return tx.UpdateMCallNote(morningId, "【ラベル】\n無し\n【ラベル】\n無し")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, morningId) {
		t.Fatal("read wrong user")
	}
}

func TestDailyReport(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		reportId, err = tx.CreateDailyReport(userId, USER_NAME, CAR_NUMBER, morningId)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentDuty(userId, reportId)
		if err != nil {
			return err
		}

		err = tx.UpdateDReportMtrPhoto(reportId, MTR_PHOTO)
		if err != nil {
			return err
		}

		err = tx.UpdateDReportTrouble(reportId, "【ラベル】\n特になし")
		if err != nil {
			return err
		}

		return tx.UpdateDReportNote(reportId, "【ラベル】\n無し")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentDuty, reportId) {
		t.Fatal("read wrong user")
	}
}

func TestEveningCall(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		eveningId, err = tx.CreateEveningCall(userId, USER_NAME, CAR_NUMBER, ALC_PHOTO)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, eveningId)
		if err != nil {
			return err
		}

		err = tx.UpdateDReportECall(reportId, eveningId)
		if err != nil {
			return err
		}

		return tx.UpdateECallNote(eveningId, "【ラベル】\n無し\n【ラベル】\n無し")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, eveningId) {
		t.Fatal("read wrong user")
	}
}

func TestLocation(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		locId, err = tx.CreateLocation(
			reportId,
			"倉庫",
			"東京都港区倉庫町９９－９９",
			99.99,
			88.88,
		)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, locId)
		if err != nil {
			return err
		}

		return tx.UpdateLocNote(locId, "【説明】集荷")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, locId) {
		t.Fatal("read wrong user")
	}
}

func TestResting(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		restId, err = tx.CreateResting(
			reportId,
			"倉庫",
			"東京都港区倉庫町９９－９９",
			99.99,
			88.88,
		)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, restId)
		if err != nil {
			return err
		}

		return tx.UpdateRestNote(restId, "【睡眠】無し")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, restId) {
		t.Fatal("read wrong user")
	}
}

func TestWaiting(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		waitId, err = tx.CreateWaiting(
			reportId,
			"倉庫",
			"東京都港区倉庫町９９－９９",
			99.99,
			88.88,
		)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, waitId)
		if err != nil {
			return err
		}

		return tx.UpdateWaitNote(waitId, "【荷主指定時刻】１５時\n【備考】無し\n")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, waitId) {
		t.Fatal("read wrong user")
	}
}

func TestLoading(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		var err error
		loadId, err = tx.CreateLoading(
			reportId,
			"倉庫",
			"東京都港区倉庫町９９－９９",
			99.99,
			88.88,
		)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentOpe(userId, loadId)
		if err != nil {
			return err
		}

		err = tx.UpdateLoadShipCheck(loadId, true)
		if err != nil {
			return err
		}

		return tx.UpdateLoadNote(loadId, "【内容】積み込み\n【備考】無し\n")
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, loadId) {
		t.Fatal("read wrong user")
	}
}

func TestReadFromOrg(t *testing.T) {
	readOrg, err := db.ReadOrg(orgId)
	if err != nil {
		t.Fatal(err)
	}
	if readOrg.LineId != ORG_LINE_ID || readOrg.Name != ORG_NAME {
		t.Fatal("read wrong data")
	}

	{
		readLineId, err := db.ReadOrgLineId(orgId)
		if err != nil {
			t.Fatal(err)
		}
		if string(readLineId) != ORG_LINE_ID {
			t.Fatal("read wrong org")
		}
	}

	{
		readEmail, err := db.ReadOrgEmail(orgId)
		if err != nil {
			t.Fatal(err)
		}
		if string(readEmail) != EMAIL {
			t.Fatal(err)
		}
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.Id, userId) || readUser.Name != USER_NAME {
		t.Fatal("read wrong user")
	}
	if readUser.Email != EMAIL {
		t.Fatal("read wrong user")
	}
	if readUser.CarNumber != CAR_NUMBER || readUser.NextState != 9 {
		t.Fatal("read wrong user")
	}

	{
		readOrgId, err := db.ReadUserOrgId(userId)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(readOrgId, orgId) {
			t.Fatal("read wrong org id")
		}
	}

	{
		readOrgId, err := db.ReadUserOrgId_Email(EMAIL)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(readOrgId, orgId) {
			t.Fatal("read wrong org id")
		}
	}

	{
		readEmail, err := db.ReadUserEmail(userId)
		if err != nil {
			t.Fatal(err)
		}
		if string(readEmail) != EMAIL {
			t.Fatal("read wrong org email")
		}
	}

	users, err := db.ReadUsersMini(orgId)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 1 {
		t.Fatal("read wrong users")
	}
	if !bytes.Equal(users[0].Id, userId) {
		t.Fatal("read wrong users")
	}

	now := time.Now() // this code does not care TZs other than JST
	from, to := timeconv.UtcRangeFromJstMonth(now.Year(), now.Month())

	mCall, err := db.ReadMorningCall(morningId)
	if err != nil {
		t.Fatal(err)
	}
	if mCall.AlcPhoto != ALC_PHOTO {
		t.Fatal("read wrong morning call")
	}

	mCalls, err := db.ReadMorningCalls(userId, from, to)
	if err != nil {
		t.Fatal(err)
	}

	if len(mCalls) != 1 {
		t.Fatal("read wrong morning calls")
	}
	if !bytes.Equal(mCalls[0].Id, morningId) {
		t.Fatal("read wrong morning calls")
	}

	eCall, err := db.ReadEveningCall(eveningId)
	if err != nil {
		t.Fatal(err)
	}
	if eCall.AlcPhoto != ALC_PHOTO {
		t.Fatal("read wrong evening call")
	}

	eCalls, err := db.ReadEveningCalls(userId, from, to)
	if err != nil {
		t.Fatal(err)
	}

	if len(eCalls) != 1 {
		t.Fatal("read wrong evening calls")
	}
	if !bytes.Equal(eCalls[0].Id, eveningId) {
		t.Fatal("read wrong evening calls")
	}

	report, err := db.ReadDailyReport(reportId)
	if err != nil {
		t.Fatal(err)
	}
	if report.MeterPhoto != MTR_PHOTO {
		t.Fatal("read wrong dialy report")
	}

	reports, err := db.ReadDailyReports(userId, from, to)
	if err != nil {
		t.Fatal(err)
	}

	if len(reports) != 1 {
		t.Fatal("read wrong daily reports")
	}
	if !bytes.Equal(reports[0].Id, reportId) {
		t.Fatal("read wrong daily reports")
	}
}

func TestReadFromReport(t *testing.T) {
	locs, err := db.ReadLocations(reportId)
	if err != nil {
		t.Fatal(err)
	}

	if len(locs) != 1 {
		t.Fatal("read wrong locations")
	}
	if !bytes.Equal(locs[0].Id, locId) {
		t.Fatal("read wrong locations")
	}

	rests, err := db.ReadRestings(reportId)
	if err != nil {
		t.Fatal(err)
	}

	if len(rests) != 1 {
		t.Fatal("read wrong restings")
	}
	if !bytes.Equal(rests[0].Id, restId) {
		t.Fatal("read wrong restings")
	}

	waits, err := db.ReadWaitings(reportId)
	if err != nil {
		t.Fatal(err)
	}

	if len(waits) != 1 {
		t.Fatal("read wrong waitings")
	}
	if !bytes.Equal(waits[0].Id, waitId) {
		t.Fatal("read wrong waitings")
	}

	loads, err := db.ReadLoadings(reportId)
	if err != nil {
		t.Fatal(err)
	}

	if len(loads) != 1 {
		t.Fatal("read wrong loadings")
	}
	if !bytes.Equal(loads[0].Id, loadId) {
		t.Fatal("read wrong loadings")
	}
}

func TestOnEnd(t *testing.T) {
	err := db.Tx(func(tx database.Transaction) error {
		err := tx.UpdateUserCurrentOpe(userId, nil)
		if err != nil {
			return err
		}

		return tx.UpdateUserCurrentDuty(userId, nil)
	})
	if err != nil {
		t.Fatal(err)
	}

	readUser, err := db.ReadUser_Line(USER_LINE_ID)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.CurrentOperation, nil) {
		t.Fatal("read wrong user")
	}
	if !bytes.Equal(readUser.CurrentDuty, nil) {
		t.Fatal("read wrong user")
	}
}

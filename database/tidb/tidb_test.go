package tidb

import (
	"bytes"
	"errors"
	"nippo/database"
	"os"
	"testing"
	"time"

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

	org, err := db.CreateOrg("テスト　企業")
	if err != nil {
		t.Fatal(err)
	}

	orgId, err := uuid.FromBytes(org.Id)
	if err != nil {
		t.Fatal(err)
	}

	err = db.UpdateOrgLineId(orgId.String(), "U9198cf6745fc0d808977bdb0b9f22995")
	if err != nil {
		t.Fatal(err)
	}

	user, err := db.CreateUser(
		orgId.String(),
		"テスト　ユーザー太郎",
		"U8189cf6745fc0d808977bdb0b9f22885",
	)
	if err != nil {
		t.Fatal(err)
	}

	readOrg, err := db.ReadOrg(orgId.String())
	if err != nil {
		t.Fatal(err)
	}
	if readOrg.LineId != "U9198cf6745fc0d808977bdb0b9f22995" || readOrg.Name != org.Name {
		t.Fatal("read wrong data")
	}

	readUser, err := db.ReadUser("U8189cf6745fc0d808977bdb0b9f22885")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readUser.Id, user.Id) || readUser.Name != user.Name {
		t.Fatal("read wrong user")
	}
}

func TestMorningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	_, err = db.CreateMorningCall("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEveningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	_, err = db.CreateEveningCall("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDailyReport(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	_, err = db.CreateDailyReport("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}
}

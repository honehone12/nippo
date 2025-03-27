package tidb

import (
	"bytes"
	"errors"
	"fmt"
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

func TestOrg(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	org, err := db.CreateOrg(
		uuid.MustParse("b205621d-615f-41a7-9b5b-794d18737d3d"),
		"テスト　企業",
		"U9198cf6745fc0d808977bdb0b9f22995",
	)
	if err != nil {
		t.Fatal(err)
	}

	read, err := db.ReadOrg("b205621d-615f-41a7-9b5b-794d18737d3d")
	if err != nil {
		t.Fatal(err)
	}
	if read.LineId != org.LineId || read.Name != org.Name {
		t.Fatal("read wrong data")
	}
}

func TestUser(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	user, err := db.CreateUser(
		"b205621d-615f-41a7-9b5b-794d18737d3d",
		"テスト　ユーザー太郎",
		"U8189cf6745fc0d808977bdb0b9f22885",
	)
	if err != nil {
		t.Fatal(err)
	}

	read, err := db.ReadUser("U8189cf6745fc0d808977bdb0b9f22885")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(read.Id, user.Id) || read.Name != user.Name {
		t.Fatal("read wrong user")
	}
}

func TestMorningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	call, err := db.CreateMorningCall("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", call)
}

func TestEveningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	call, err := db.CreateEveningCall("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", call)
}

func TestDailyReport(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	report, err := db.CreateDailyReport("U8189cf6745fc0d808977bdb0b9f22885", today)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", report)
}

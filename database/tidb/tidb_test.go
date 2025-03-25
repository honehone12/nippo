package tidb

import (
	"errors"
	"fmt"
	"nippo/database"
	"os"
	"testing"
	"time"

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
		"テスト　オーガニゼーション会社",
		"U8189cf6745fc0d808977bdb0b9f22995",
		"U8189cf6745fc0d808977bdb0b9f22885",
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", org)
}

func TestUser(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	user, err := db.CreateUser(
		"テスト　ユーザー太郎",
		"U8189cf6745fc0d808977bdb0b9f22885",
		"U8189cf6745fc0d808977bdb0b9f22995",
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", user)
}

func TestMorningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	call, err := db.CreateMorningCall(today, "U8189cf6745fc0d808977bdb0b9f22885")
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
	call, err := db.CreateEveningCall(today, "U8189cf6745fc0d808977bdb0b9f22885")
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
	report, err := db.CreateDailyReport(today, "U8189cf6745fc0d808977bdb0b9f22885")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", report)
}

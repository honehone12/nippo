package tidb

import (
	"errors"
	"nippo/db"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var _ db.Db = &TiDb{}

func initTest() (db.Db, error) {
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

	if err := db.CreateOrg("test org"); err != nil {
		t.Fatal(err)
	}
}

func TestUser(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	if err := db.CreateUser("test user"); err != nil {
		t.Fatal(err)
	}
}

func TestMorningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	if err := db.CreateMorningCall(today); err != nil {
		t.Fatal(err)
	}
}

func TestEveningCall(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	if err := db.CreateEveningCall(today); err != nil {
		t.Fatal(err)
	}
}

func TestDailyReport(t *testing.T) {
	db, err := initTest()
	if err != nil {
		t.Fatal(err)
	}

	today := time.Now().Local()
	if err := db.CreateDailyReport(today); err != nil {
		t.Fatal(err)
	}
}

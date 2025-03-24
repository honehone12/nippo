package tidb

import (
	"nippo/schema"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestTiDb(t *testing.T) {
	godotenv.Load("migration/.env")

	tiUri := os.Getenv("TIDB_URI")
	if len(tiUri) == 0 {
		t.Fatal("env TIDB_URI is not set")
	}

	ti, err := Open(tiUri)
	if err != nil {
		t.Fatal(err)
	}

	org := schema.Org{
		Name: "test",
	}
	res := ti.orm.Create(&org)
	if res.Error != nil {
		t.Fatal(res.Error)
	}
}

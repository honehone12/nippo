package main

import (
	"errors"
	"fmt"
	"nippo/database/tidb"
	"os"

	"github.com/joho/godotenv"
)

func migrate() error {
	dbUri := os.Getenv("DB_URI")
	if len(dbUri) == 0 {
		return errors.New("env DB_URI is not set")
	}

	db, err := tidb.Open(dbUri)
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(); err != nil {
		return err
	}

	return nil
}

func main() {
	godotenv.Load("../.env")

	if err := migrate(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done")
}

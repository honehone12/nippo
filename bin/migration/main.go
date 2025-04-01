package main

import (
	"errors"
	"nippo/database/tidb"
	"nippo/log"
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
	log.FmtPlain()

	if err := migrate(); err != nil {
		log.L.Error(err)
		return
	}

	log.L.Info("done")
}

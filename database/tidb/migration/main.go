package main

import (
	"errors"
	"nippo/database/tidb"
	"nippo/log"
	"os"

	"github.com/joho/godotenv"
)

func migrate() error {
	tiUri := os.Getenv("TIDB_URI")
	if len(tiUri) == 0 {
		return errors.New("env TIDB_URI is not set")
	}

	ti, err := tidb.Open(tiUri)
	if err != nil {
		return err
	}

	if err := ti.AutoMigrate(); err != nil {
		return err
	}

	return nil
}

func init() {
	godotenv.Load()
	log.FmtPlain()
}

func main() {
	if err := migrate(); err != nil {
		log.L.Error(err)
		return
	}

	log.L.Info("done")
}

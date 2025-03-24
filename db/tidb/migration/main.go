package main

import (
	"nippo/db/tidb"
	"nippo/log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log.FmtDefault()
}

func main() {
	tiUri := os.Getenv("TIDB_URI")
	if len(tiUri) == 0 {
		panic("env TIDB_URI is not set")
	}

	ti, err := tidb.Open(tiUri)
	if err != nil {
		panic(err)
	}

	ti.AutoMigrate()
	log.L.Info("done")
}

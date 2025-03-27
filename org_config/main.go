package main

import (
	"flag"
	"fmt"
	"nippo/database"
	"nippo/database/tidb"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	name := flag.String("name", "", "name")
	lineId := flag.String("line-id", "", "line id")
	flag.Parse()
	if len(*name) == 0 {
		fmt.Println("name is empty")
		return
	}
	if len(*lineId) == 0 {
		fmt.Println("line-id is empty")
		return
	}

	godotenv.Load()
	dbUri := os.Getenv("DB_URI")
	if len(dbUri) == 0 {
		fmt.Println("DB_URI is empty")
		return
	}

	var db database.Database
	var err error
	db, err = tidb.Open(dbUri)
	if err != nil {
		fmt.Println(err)
		return
	}

	org, err := db.CreateOrg(*name, *lineId)
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := uuid.ParseBytes(org.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("id: %s\n", id.String())
	fmt.Printf("name: %s\n", org.Name)
	fmt.Printf("line id: %s\n", org.LineId)
}

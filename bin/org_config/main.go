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

	flag.Parse()
	if len(*name) == 0 {
		fmt.Println("name is empty")
		return
	}

	godotenv.Load("../.env")
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

	bId, err := db.CreateOrg(*name)
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := uuid.FromBytes(bId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("id: %s\n", id.String())
}

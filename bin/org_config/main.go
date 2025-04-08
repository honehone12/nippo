package main

import (
	"flag"
	"fmt"
	"nippo/database/tidb"
	"nippo/input"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	name := flag.String("name", "", "org's name")
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

	db, err := tidb.Open(dbUri)
	if err != nil {
		fmt.Println(err)
		return
	}

	cName := input.Clean(*name)
	fmt.Printf("creating org with name '%s'\n", cName)

	bId, err := db.Once().CreateOrg(cName)
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

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
	email := flag.String("email", "", "org's email")
	flag.Parse()
	if len(*name) == 0 {
		fmt.Println("name is empty")
		return
	}
	if len(*email) == 0 {
		fmt.Println("email is empty")
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

	cEmail, err := input.ValidEmail(*email)
	if err != nil {
		fmt.Println("invalid email")
		return
	}

	bId, err := db.Once().CreateOrg(cName, cEmail)
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

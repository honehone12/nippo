package main

import (
	"flag"
	"fmt"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/input"
	"nippo/schema"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	orgIdStr := flag.String("org-id", "", "org's id string")
	name := flag.String("name", "", "user's name")
	email := flag.String("email", "", "user's email")
	flag.Parse()
	if len(*orgIdStr) == 0 {
		fmt.Println("org-id is empty")
		return
	}
	if len(*name) == 0 {
		fmt.Println("name is empty")
		return
	}
	if len(*email) == 0 {
		fmt.Println("email is empty")
		return
	}

	orgId, err := uuid.Parse(*orgIdStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	uUid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	lineId := schema.CleanText(fmt.Sprintf("U%x", uUid[:]))

	cName := input.Clean(*name)
	fmt.Printf("creating user with name '%s'\n", cName)

	cEmail, err := input.ValidEmail(*email)
	if err != nil {
		fmt.Println("invalid email")
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

	if err = db.Tx(func(tx database.Transaction) error {
		id, err := tx.CreateUser(orgId[:], lineId)
		if err != nil {
			return err
		}

		if err = tx.UpdateUserName(id, cName); err != nil {
			return err
		}

		err = tx.UpdateUserEmail(id, cEmail)
		return err
	}); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done")
}

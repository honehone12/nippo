package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := flag.String("id", "", "encoded id")
	flag.Parse()
	if len(*id) == 0 {
		fmt.Println("id is empty")
		return
	}

	dst := make([]byte, 16)
	n, err := base64.StdEncoding.Decode(dst, []byte(*id))
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != 16 {
		fmt.Println("size not 16")
		return
	}

	uUID, err := uuid.FromBytes(dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("id: %s\n", uUID.String())
}

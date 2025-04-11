package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "encoded url")
	flag.Parse()
	if len(*url) == 0 {
		fmt.Println("url is empty")
		return
	}

	s, err := base64.StdEncoding.DecodeString(*url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("url: %s\n", s)
}

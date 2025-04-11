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

	b := []byte(*url)
	len := base64.StdEncoding.DecodedLen(len(b))
	dst := make([]byte, len)
	_, err := base64.StdEncoding.Decode(dst, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("url: %s\n", string(dst))
}

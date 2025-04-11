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
	len := base64.StdEncoding.EncodedLen(len(b))
	dst := make([]byte, len)
	n, err := base64.StdEncoding.Decode(dst, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != len {
		fmt.Printf("size not %d\n", len)
		return
	}

	fmt.Printf("url: %s\n", string(dst))
}

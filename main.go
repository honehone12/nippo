package main

import (
	"nippo/log"

	"github.com/google/uuid"
)

func init() {
	log.FmtDefault()
}

func main() {
	uuid.NewString()

	log.L.Info(len(uuid.NewString()))
}

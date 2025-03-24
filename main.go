package main

import "nippo/log"

func init() {
	log.FmtDefault()
}

func main() {
	log.L.Info("hello")
}

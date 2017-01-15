package main

import (
	"log"

	"github.com/lsolniczek/mimus/setup"
)

func main() {
	if err := setup.Run().Start(); err != nil {
		log.Fatalln(err)
	}
}

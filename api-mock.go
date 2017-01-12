package main

import (
	"log"

	"github.com/lsolniczek/api-mock/setup"
)

func main() {
	if err := setup.Run().Start(); err != nil {
		log.Fatalln(err)
	}
}

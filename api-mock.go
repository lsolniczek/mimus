package main

import (
	"log"

	m "github.com/lsolniczek/api-mock/api-mock-main"
)

func main() {
	if err := m.Run().Start(); err != nil {
		log.Fatalln(err)
	}
}

package main

import m "github.com/lsolniczek/api-mock/api-mock-main"

func main() {
	inst := m.Run()
	inst.Start()
}

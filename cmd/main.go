package main

import (
	"iSpringTest/server"
)

func main() {
	serv := server.MakeServer("config.json")
	serv.Run()
}

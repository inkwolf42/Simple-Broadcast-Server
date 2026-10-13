package main

import (
	"os"

	"inkwolf.com/broadcast/cmd/client"
	"inkwolf.com/broadcast/cmd/server"
)

func main() {
	if len(os.Args) < 2 {
		println("No Command was entered.")
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "start":
		server.Run()
	case "connect":
		client.Run()
	default:
		println("Invalid Command.")
	}
}

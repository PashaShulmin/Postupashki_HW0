package main

import (
	"HW0/client"
	"HW0/server"
)

func main() {
	go server.StartServer()
	client.Connect()
}

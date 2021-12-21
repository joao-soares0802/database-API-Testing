package main

import (
	"awesomeProject/database"
	"awesomeProject/server"
)

func main() {
	go database.StartDB()
	server := server.NewServer()

	server.Run()
}

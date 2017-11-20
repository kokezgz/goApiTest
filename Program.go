package main

import (
	"log"

	"./Controllers"
)

func main() {
	log.Println("Server running")

	var controller Controllers.Controller
	controller.StartServer()
}

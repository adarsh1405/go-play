package main

import (
	"github.com/adarsh1405/go-play/server"
)

func main() {

	// Check Configurations of the Server
	server.CheckConfigs()

	// Setting up the Router
	server.Run()

}
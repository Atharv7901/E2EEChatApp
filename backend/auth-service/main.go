package main

import (
	"log"

	"github.com/Atharv7901/E2EEChatApp/config"
)

func main() {
	//initialize db connection
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	config.ConnectionTest(db)

	//initialize service, store and handlers

	//routes

	//start the servers
}

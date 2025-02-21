package main

import (
	"log"

	"github.com/Atharv7901/E2EEChatApp/config"
	authHandler "github.com/Atharv7901/E2EEChatApp/handler"
	authroutes "github.com/Atharv7901/E2EEChatApp/routes"
	authservice "github.com/Atharv7901/E2EEChatApp/service"
	authstore "github.com/Atharv7901/E2EEChatApp/store"
)

func main() {
	//initialize db connection
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	config.ConnectionTest(db)

	//initialize service, store and handlers
	authStore := authstore.NewAuthStore(db)
	authService := authservice.NewAuthService(authStore)
	authHandler := authHandler.NewAuthHandler(authService)

	//routes
	server := authroutes.NewAuthServer(":8080", db)
	if err := server.Run(authHandler); err != nil {
		log.Fatal(err)
	}
}

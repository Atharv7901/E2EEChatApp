package main

import (
	"log"

	"github.com/Atharv7901/E2EEChatApp/chat-service/config"
	chatHandler "github.com/Atharv7901/E2EEChatApp/chat-service/handler"
	chatroutes "github.com/Atharv7901/E2EEChatApp/chat-service/routes"
	chatservice "github.com/Atharv7901/E2EEChatApp/chat-service/service"
	chatstore "github.com/Atharv7901/E2EEChatApp/chat-service/store"
)

func main() {
	//initialize db connection
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	config.ConnectionTest(db)

	//initialize store, service and handler
	chatStore := chatstore.NewChatStore(db)
	chatService := chatservice.NewChatService(chatStore)
	chatHandler := chatHandler.NewChatHandler(chatService)

	//routes
	server := chatroutes.NewChatServer(":8081", db)
	if err := server.Run(chatHandler); err != nil {
		log.Fatal(err)
	}
}

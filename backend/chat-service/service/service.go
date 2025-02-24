package chatservice

import (
	"github.com/Atharv7901/E2EEChatApp/chat-service/db"
	chatstore "github.com/Atharv7901/E2EEChatApp/chat-service/store"
)

type ChatServiceInterface interface {
	GetAllUsers() ([]db.User, error)
	GetUserByID(id int) (*db.User, error)
}

type ChatService struct {
	chatstore chatstore.ChatStoreInterface
}

func NewChatService(chatStore chatstore.ChatStoreInterface) ChatServiceInterface {
	return &ChatService{chatstore: chatStore}
}

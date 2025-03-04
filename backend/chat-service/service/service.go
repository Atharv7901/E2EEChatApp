package chatservice

import (
	"github.com/Atharv7901/E2EEChatApp/chat-service/db"
	"github.com/Atharv7901/E2EEChatApp/chat-service/kafka"
	"github.com/Atharv7901/E2EEChatApp/chat-service/models"
	chatstore "github.com/Atharv7901/E2EEChatApp/chat-service/store"
)

type ChatServiceInterface interface {
	GetAllUsers() ([]db.User, error)
	GetUserByID(id int) (*db.User, error)
	ProduceMessage(msg interface{}, senderId int) error
	FetchMessages(broker string, groupID string, topic string, limit int) ([]models.ChatMessage, error)
}

type ChatService struct {
	chatstore  chatstore.ChatStoreInterface
	kafkaStore *kafka.KafkaProducer
}

func NewChatService(chatStore chatstore.ChatStoreInterface, kafkaStore *kafka.KafkaProducer) ChatServiceInterface {
	return &ChatService{chatstore: chatStore, kafkaStore: kafkaStore}
}

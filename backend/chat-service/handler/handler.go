package chatHandler

import (
	chatservice "github.com/Atharv7901/E2EEChatApp/chat-service/service"
	"github.com/gin-gonic/gin"
)

type ChatHandlerInterface interface {
	GetAllUsers(c *gin.Context)
}

type ChatHandler struct {
	chatservice chatservice.ChatServiceInterface
}

func NewChatHandler(chatService chatservice.ChatServiceInterface) ChatHandlerInterface {
	return &ChatHandler{chatservice: chatService}
}

package chatHandler

import (
	"net/http"

	"github.com/Atharv7901/E2EEChatApp/chat-service/models"
	"github.com/gin-gonic/gin"
)

func (h *ChatHandler) SendMessage(c *gin.Context) {
	var msg models.ChatMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.chatservice.ProduceMessage(msg, msg.SenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
}

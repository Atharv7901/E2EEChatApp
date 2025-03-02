package chatHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ChatHandler) FetchMessagesHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	messages, err := h.chatservice.FetchMessages("localhost:9092", "chat-group", "chat-messages", limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch messages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

package chatHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ChatHandler) GetAllUsers(c *gin.Context) {
	users, err := h.chatservice.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

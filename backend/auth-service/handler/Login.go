package authHandler

import (
	"net/http"

	authmodels "github.com/Atharv7901/E2EEChatApp/models"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Login(c *gin.Context) {
	var req authmodels.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

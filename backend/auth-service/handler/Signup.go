package authHandler

import (
	"net/http"

	authmodels "github.com/Atharv7901/E2EEChatApp/models"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req authmodels.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.authService.CreateUser(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created Successfully!", "user": user})
}

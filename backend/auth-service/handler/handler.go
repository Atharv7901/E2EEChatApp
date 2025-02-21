package authHandler

import (
	authservice "github.com/Atharv7901/E2EEChatApp/service"
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	SignUp(c *gin.Context)
}

type AuthHandler struct {
	authService authservice.AuthServiceInterface
}

func NewAuthHandler(authservice authservice.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{authService: authservice}
}

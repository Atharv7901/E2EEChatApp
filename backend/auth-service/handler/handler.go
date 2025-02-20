package authHandler

import (
	"net/http"

	authservice "github.com/Atharv7901/E2EEChatApp/service"
)

type AuthHandlerInterface interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	authService authservice.AuthServiceInterface
}

func NewAuthHandler(authservice authservice.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{authService: authservice}
}


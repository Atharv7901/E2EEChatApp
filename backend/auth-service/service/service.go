package authservice

import (
	authstore "github.com/Atharv7901/E2EEChatApp/store"
)

type AuthServiceInterface interface {
	CreateUser(firstName string, lastName string, email string, password string) (int, error)
}

type AuthService struct {
	authStore authstore.AuthStoreInterface
}

func NewAuthService(authStore authstore.AuthStoreInterface) AuthServiceInterface {
	return &AuthService{authStore: authStore}
}

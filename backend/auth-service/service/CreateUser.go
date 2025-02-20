package authservice

import (
	"errors"

	"github.com/Atharv7901/E2EEChatApp/db"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) CreateUser(firstName string, lastName string, email string, password string) (int, error) {
	existingUser, _ := s.authStore.GetUserByEmail(email)
	if existingUser != nil {
		return 0, errors.New("email already registered")
	}

	//hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &db.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
	}
	newUser, err := s.authStore.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return newUser, nil
}

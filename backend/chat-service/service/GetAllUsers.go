package chatservice

import "github.com/Atharv7901/E2EEChatApp/chat-service/db"

func (s *ChatService) GetAllUsers() ([]db.User, error) {
	return s.chatstore.GetAllUsers()
}

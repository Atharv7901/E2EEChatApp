package chatservice

import "github.com/Atharv7901/E2EEChatApp/chat-service/models"

func (s *ChatService) FetchMessages(broker string, groupID string, topic string, limit int) ([]models.ChatMessage, error) {
	messages, err := s.kafkaStore.FetchRecentMessages(broker, groupID, topic, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

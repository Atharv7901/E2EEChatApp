package chatservice

func (s *ChatService) ProduceMessage(message interface{}, senderId int) error {
	err := s.kafkaStore.ProduceMessage(message, senderId)
	if err != nil {
		return err
	}
	return nil
}

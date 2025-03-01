package models

type ChatMessage struct {
	SenderID   int    `json:"senderId"`
	ReceiverID int    `json:"receiverId"`
	Message    string `json:"message"`
}

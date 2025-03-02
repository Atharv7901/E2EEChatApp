package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Atharv7901/E2EEChatApp/chat-service/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (k *KafkaProducer) FetchRecentMessages(broker string, groupID string, topic string, limit int) ([]models.ChatMessage, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("Error creating kafka consumer: ", err)
		return nil, err
	}
	defer consumer.Close()

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Println("Error subscribing to Kafka topic:", err)
		return nil, err
	}

	messages := []models.ChatMessage{}

	timeout := time.After(5 * time.Second)
	for len(messages) < limit {
		select {
		case <-timeout:
			return messages, nil
		default:
			msg, err := consumer.ReadMessage(time.Second * 1)
			if err == nil {
				var chatMsg models.ChatMessage
				json.Unmarshal(msg.Value, &chatMsg)
				messages = append(messages, chatMsg)
			} else {
				fmt.Println("No new messages or error: ", err)
			}
		}
	}

	return messages, nil
}

package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Fatal("Failed to create Kafka producer: ", err)
	}
	return &KafkaProducer{Producer: p, Topic: topic}
}

func (k *KafkaProducer) ProduceMessage(message interface{}, userID int) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	partitionKey := fmt.Sprintf("%d", userID%3)

	err = k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.Topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(partitionKey),
		Value: messageJSON,
	}, nil)

	if err != nil {
		log.Println("Failed to produce message:", err)
		return err
	}

	return nil
}

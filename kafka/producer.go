package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func NewKafkaProducer(brokerAddress, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Logger:   log.New(os.Stdout, "kafka writer: ", 0),
	}
}

func ProduceMessage(producer *kafka.Writer, key, value []byte) error {
	return producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
}

// ProduceWithKey produces a message with a string key and value to the given topic.
func ProduceWithKey(brokerAddress, topic, key, message string) error {
	producer := NewKafkaProducer(brokerAddress, topic)
	defer producer.Close()
	return ProduceMessage(producer, []byte(key), []byte(message))
}

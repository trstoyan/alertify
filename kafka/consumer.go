package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(brokerAddress, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
		Logger:  log.New(os.Stdout, "kafka reader: ", 0),
	})
}

// ConsumeMessages reads messages in a loop and calls messageHandler with key and value.
func ConsumeMessages(consumer *kafka.Reader, messageHandler func(key, value string)) {
	for {
		msg, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while reading message: %v", err)
			continue
		}
		messageHandler(string(msg.Key), string(msg.Value))
	}
}

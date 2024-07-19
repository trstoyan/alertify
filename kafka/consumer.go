package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(handleMessage func(string), topic string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "10.25.8.39:9092",
		"group.id":          "notification-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			handleMessage(string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

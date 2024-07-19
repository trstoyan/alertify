package sms

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeSMS() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "10.25.8.39:9092",
		"group.id":          "sms-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"sms-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			SendSMS(string(msg.Value))
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

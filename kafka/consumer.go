package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consume(handleMessage func(string, string) string, topic string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "pkc-w7d6j.germanywestcentral.azure.confluent.cloud:9092",
		"sasl.username":     "UQLV7MXX2G4BMN4G",
		"sasl.password":     "tVNWl/C5NIUPTc+eHRk4U3cJbBW4lcvv89GkbZ1sBCkFOemlt1sWerbBfyQ2k7oQ",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"group.id":          "kafka-go-getting-started",
		"auto.offset.reset": "earliest"})
	// Check for errors when creating the consumer and connect to the Kafka cluster
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	type KafkaMessage struct {
		Key   string
		Value string
	}
	// Close the consumer when the function returns
	defer c.Close()

	err = c.SubscribeTopics([]string{topic}, nil)
	// Check for errors when subscribing to topics
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
	}

	// Create a channel for messages
	messageChan := make(chan KafkaMessage)

	// Number of worker goroutines
	numWorkers := 30

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for kafkaMsg := range messageChan {
				responseMessage := handleMessage(kafkaMsg.Key, kafkaMsg.Value)
				if responseMessage == "" {
					log.Printf("Failed to handle message: %s", responseMessage)
				} else {
					log.Printf("Message handled successfully: %s", responseMessage)
				}
			}
		}()
	}

	// Main loop to read messages and send them to the channel
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			messageChan <- KafkaMessage{Key: string(msg.Key), Value: string(msg.Value)} // Send message to worker goroutines only if there's a message
		} else {
			// Handle the error, e.g., log it or implement a retry mechanism
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	// close(messageChan)
}

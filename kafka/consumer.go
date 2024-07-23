package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
)

func Consume(handleMessage func(string, string) string, topic string) {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"sasl.username":     os.Getenv("KAFKA_SASL_USERNAME"),
		"sasl.password":     os.Getenv("KAFKA_SASL_PASSWORD"),
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
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
	numWorkers := 1

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for kafkaMsg := range messageChan {
				handleMessage(kafkaMsg.Key, kafkaMsg.Value)
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

package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
)

func ProduceWithKey(topic, key, message string) error {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %s", err)
		return err // Return error to allow for graceful handling
	}

	// Create a new producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     os.Getenv("KAFKA_SASL_USERNAME"),
		"sasl.password":     os.Getenv("KAFKA_SASL_PASSWORD"),
		"retries":           5,
		"retry.backoff.ms":  500,
	})
	if err != nil {
		log.Printf("Failed to create producer: %s", err)
		return err // Return error to allow for graceful handling
	}
	defer p.Close()

	// Produce the message
	if err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
	}, nil); err != nil {
		return err // Return error if message production fails
	}

	// Wait for all messages to be delivered
	p.Flush(5 * 1000) // Increase the timeout to ensure all messages are delivered

	return err
}

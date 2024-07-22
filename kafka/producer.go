package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ProduceWithKey(topic, key, message string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "pkc-w7d6j.germanywestcentral.azure.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     "UQLV7MXX2G4BMN4G",
		"sasl.password":     "tVNWl/C5NIUPTc+eHRk4U3cJbBW4lcvv89GkbZ1sBCkFOemlt1sWerbBfyQ2k7oQ",
		"retries":           5,   // Retry up to 5 times
		"retry.backoff.ms":  500, // Wait for 500ms before retrying
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
		return err // Return error instead of terminating the program
	}
	defer p.Close()

	deliveryChan := make(chan kafka.Event, 10000) // Use a delivery channel to handle reports

	// Produce message
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return err // Return error if message production fails
	}

	go func() {
		for e := range deliveryChan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				}
			}
		}
	}()

	// Wait for all messages to be delivered
	p.Flush(15 * 1000) // Adjust the timeout as needed

	return nil
}

package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(topic string, message string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.25.8.39:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer p.Close()

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	}

	p.Flush(5 * 1000)
}

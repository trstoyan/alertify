package sms

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceSMS(message string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.25.8.39:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer p.Close()

	// Start a go routine to asynchronously handle delivery reports
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("Successfully delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	topic := "sms-topic"
	// Produce message without waiting for delivery report
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	}

	// Note: In a real application, ensure proper shutdown logic to drain the Events channel and close the producer cleanly.
}

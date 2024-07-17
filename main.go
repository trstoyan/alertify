package main

import (
	"log"

	"alertify/kafka"
	"alertify/sms"
)

func main() {
	brokerAddress := "localhost:9092"
	smsTopic := "sms_topic"
	groupID := "sms_service"

	// Start SMS Consumer
	sms.StartSMSConsumer(brokerAddress, smsTopic, groupID)

	// Simulate producing an SMS message
	producer := kafka.NewKafkaProducer(brokerAddress, smsTopic)
	payload := sms.SMSPayload{
		PhoneNumber: "1234567890",
		Message:     "Hello, this is a test message.",
	}
	err := sms.ProduceSMSMessage(producer, payload)
	if err != nil {
		log.Fatalf("failed to produce SMS message: %v", err)
	}
}

package sms

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

type SMSPayload struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
}

func StartSMSConsumer(brokerAddress, topic, groupID string) {
	consumer := kafka.NewKafkaConsumer(brokerAddress, topic, groupID)
	go kafka.ConsumeMessages(consumer, handleSMSMessage)
}

func handleSMSMessage(key, value string) {
	var payload SMSPayload
	err := json.Unmarshal([]byte(value), &payload)
	if err != nil {
		log.Printf("error unmarshalling SMS message: %v", err)
		return
	}

	err = SendSMS(payload.PhoneNumber, payload.Message)
	if err != nil {
		log.Printf("error sending SMS: %v", err)
		return
	}

	log.Printf("SMS sent to %s: %s", payload.PhoneNumber, payload.Message)
}
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

func handleSMSMessage(msg kafka.Message) {
	var payload SMSPayload
	err := json.Unmarshal(msg.Value, &payload)
	if err != nil {
		log.Printf("error unmarshalling message: %v", err)
		return
	}

	err = SendSMS(payload.PhoneNumber, payload.Message)
	if err != nil {
		log.Printf("error sending SMS: %v", err)
		return
	}

	log.Printf("SMS sent to %s: %s", payload.PhoneNumber, payload.Message)
}

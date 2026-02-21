package sms

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

const smsTopic = "sms-topic"

func ProduceSMSMessage(brokerAddress string, payload SMSPayload) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return kafka.ProduceWithKey(brokerAddress, smsTopic, "", string(msg))
}

func ProduceSMSResponse(brokerAddress, key, message string) {
	const responseTopic = "sms-sent-topic"
	err := kafka.ProduceWithKey(brokerAddress, responseTopic, key, message)
	if err != nil {
		log.Printf("Failed to send SMS response: %s", err)
	} else {
		log.Printf("SMS response sent successfully with key: %s\n", key)
	}
}

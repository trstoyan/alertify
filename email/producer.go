package email

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

const emailTopic = "email-topic"

func ProduceEmailMessage(brokerAddress string, payload EmailPayload) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return kafka.ProduceWithKey(brokerAddress, emailTopic, "", string(msg))
}

func ProduceEmailResponse(brokerAddress, key, message string) {
	const responseTopic = "email-sent-topic"
	err := kafka.ProduceWithKey(brokerAddress, responseTopic, key, message)
	if err != nil {
		log.Printf("Failed to send email response: %s", err)
	} else {
		log.Printf("Email response sent successfully with key: %s\n", key)
	}
}

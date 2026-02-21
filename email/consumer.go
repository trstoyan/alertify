package email

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

// StartEmailConsumer starts a Kafka consumer that listens on the email topic
// and sends emails for each incoming message.
func StartEmailConsumer(brokerAddress, groupID string) {
	consumer := kafka.NewKafkaConsumer(brokerAddress, emailTopic, groupID)
	go kafka.ConsumeMessages(consumer, handleEmailMessage)
}

func handleEmailMessage(key, value string) {
	var payload EmailPayload
	err := json.Unmarshal([]byte(value), &payload)
	if err != nil {
		log.Printf("error unmarshalling email message: %v", err)
		return
	}

	err = SendEmail(payload)
	if err != nil {
		log.Printf("error sending email: %v", err)
		return
	}

	log.Printf("Email sent to %s: %s", payload.To, payload.Subject)
}

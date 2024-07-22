package email

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/trstoyan/alertify/kafka"
)

const topic string = "email-topic"

func ProduceEmail(message string) string {
	hasher := sha256.New()
	hasher.Write([]byte(message + time.Now().String()))
	msgKey := hex.EncodeToString(hasher.Sum(nil))

	err := kafka.ProduceWithKey(topic, msgKey, message)
	if err != nil {
		log.Fatalf("Failed to send Email: %s", err)
	} else {
		log.Printf("Email sent successfully with key: %s\n", msgKey)
	}
	return msgKey
}

const responseTopic string = "email-sent-topic"

func ProduceSMSResponse(message string) {
	err := kafka.ProduceWithKey(responseTopic, "", message)
	if err != nil {
		log.Fatalf("Failed to send Email response: %s", err)
	} else {
		log.Printf("Email response sent successfully with key: %s\n", message)
	}
}

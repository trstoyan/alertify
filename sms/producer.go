package sms

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/trstoyan/alertify/kafka"
)

const topic string = "sms-topic"

func ProduceSMS(message string) string {

	hasher := sha256.New()
	hasher.Write([]byte(message + time.Now().String()))
	msgKey := hex.EncodeToString(hasher.Sum(nil))

	err := kafka.ProduceWithKey(topic, msgKey, message)
	if err != nil {
		log.Fatalf("Failed to send SMS: %s", err)
	} else {
		log.Printf("SMS sent successfully with key: %s\n", msgKey)
	}
	return msgKey
}

const responseTopic string = "sms-sent-topic"

func ProduceSMSResponse(message string) {
	err := kafka.ProduceWithKey(responseTopic, "", message)
	if err != nil {
		log.Fatalf("Failed to send SMS response: %s", err)
	} else {
		log.Printf("SMS response sent successfully with key: %s\n", message)
	}
}

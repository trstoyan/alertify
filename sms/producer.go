package sms

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/trstoyan/alertify/kafka"
)

const topic string = "sms-topic"

func ProduceSMS(message string) (string, error) {

	fmt.Println("ProduceSMS called", message)

	hasher := sha256.New()
	hasher.Write([]byte(message + time.Now().String()))
	msgKey := hex.EncodeToString(hasher.Sum(nil))

	messageDict := map[string]string{
		"from":    "message_from",
		"to":      "message_to",
		"message": message,
	}

	// Convert messageDict to JSON string
	messageJSON, errjson := json.Marshal(messageDict)
	if errjson != nil {
		return "", fmt.Errorf("failed to marshal messageDict to JSON: %w", errjson)
	}

	// Use a Go routine to send the message

	err := kafka.ProduceWithKey(topic, msgKey, string(messageJSON))
	if err != nil {
		return "", fmt.Errorf("failed to send SMS: %w", err)
	}

	return msgKey, err
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

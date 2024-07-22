package email

import (
	"fmt"
	"log"

	"github.com/trstoyan/alertify/kafka"
)

func SendEmail(key, message string) string {

	// Replace with actual email gateway API call
	fmt.Println("Sending Email:", message, "with key:", key)
	// Simulate email send result (for example purposes)
	sendSuccess := false

	// Produce a response message to the Kafka response topic
	responseMessage := "Email sent successfully"
	if !sendSuccess {
		responseMessage = "failed_to_send"
	} else {
		responseMessage = "message_sent"
	}

	err := kafka.ProduceWithKey(responseTopic, key, responseMessage)
	if err != nil {
		log.Printf("Failed to produce email response: %s", err)
	} else {
		log.Printf("Produced email response for key: %s", key)
	}
	return responseMessage
}

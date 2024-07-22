package sms

import (
	"fmt"
	"log"

	"github.com/trstoyan/alertify/kafka"
)

func SendSMS(key, message string) string {

	// Replace with actual SMS gateway API call
	fmt.Println("Sending SMS:", message, "with key:", key)

	responseMessage := SendSMSTwilio("+19382531802", "+359884390195", message)

	err := kafka.ProduceWithKey(responseTopic, key, *responseMessage.Body)
	if err != nil {
		log.Printf("Failed to produce email response: %s", err)
	} else {
		log.Printf("Produced email response for key: %s", key)
	}

	return *responseMessage.Body
}

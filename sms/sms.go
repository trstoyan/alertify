package sms

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/trstoyan/alertify/kafka"
)

// Define a struct that matches the JSON structure
type SMS struct {
	From    string `json:"from"`
	Message string `json:"message"`
	To      string `json:"to"`
}

func SendSMS(key string, messageDict string) error {

	var sms SMS
	// Unmarshal the messageDict into the sms struct
	err := json.Unmarshal([]byte(messageDict), &sms)
	if err != nil {
		log.Printf("Failed to unmarshal messageDict: %s", err)
		return err // Return the error
	}

	fmt.Printf(sms.From, sms.To, sms.Message)
	// Assuming SendSMSTwilio returns an error
	response := "hello" //SendSMSTwilio(sms.From, sms.To, sms.Message)
	if response != "nil" {
		log.Printf("Failed to send SMS response: %s", response)
		return nil //response // Return the error
	} else {
		log.Printf("Sent SMS response for key: %s", key)
		// Produce a message to the response topic
		responseTopic := "your_response_topic_here" // Ensure this is correctly defined
		err := kafka.ProduceWithKey(responseTopic, key, "message_sent")
		if err != nil {
			log.Printf("Failed to produce SMS response: %s", err)
			return err // Return the error
		}
	}

	return nil
}

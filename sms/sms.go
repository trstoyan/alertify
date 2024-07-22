package sms

import (
	"fmt"
)

func SendSMS(key, message string) string {

	// Replace with actual SMS gateway API call
	fmt.Println("Sending SMS:", message, "with key:", key)

	responseMessage := "SMS sent successfully"

	return responseMessage
}

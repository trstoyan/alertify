package sms

import (
	"fmt"
	"time"
)

func SendSMS(message string) {

	time.Sleep(10 * time.Second)
	// Replace with actual SMS gateway API call
	fmt.Println("Sending SMS:", message)
}

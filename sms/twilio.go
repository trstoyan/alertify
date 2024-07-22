package sms

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMSTwilio(fromNumber string, toNumber string, message string) *openapi.ApiV2010Message {

	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Load environment variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Initialize the Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Set up message parameters
	params := &openapi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(fromNumber)
	params.SetBody(message)

	// Send the message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return resp
	} else {
		fmt.Println("SMS sent successfully!")
		fmt.Println(resp)
		return resp
	}

}

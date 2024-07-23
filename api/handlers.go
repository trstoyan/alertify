package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/trstoyan/alertify/sms"
)

func ValidateNotification(w http.ResponseWriter, r *http.Request) {
	var notification MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if notification.Channel == "" || notification.Message == "" {
		http.Error(w, "Invalid notification payload", http.StatusBadRequest)
		return
	}

	// Initialize a variable to hold the response message
	var response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	if notification.Channel == "sms" {
		message := fmt.Sprintf("From: %s\nTo: %s\nMessage: %s", notification.MessageFrom, notification.MessageTo, notification.Message)
		fmt.Println(message)
		msgKey, err := sms.ProduceSMS(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("SMS notification received: %+v\n", notification)
		response.Message = msgKey

	} else {
		http.Error(w, "Invalid notification channel", http.StatusBadRequest)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(response)
}

func HandleResponse(key, status string) {

	// Here you can handle the email response, for example, log it or store it in a database
	log.Printf("Received email response for key %s: %s", key, status)
}

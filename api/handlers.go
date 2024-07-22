package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trstoyan/alertify/email"
	"github.com/trstoyan/alertify/sms"
)

func ValidateNotification(w http.ResponseWriter, r *http.Request) {
	var notification Notification
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
		msgKey := sms.ProduceSMS(notification.Message)
		log.Printf("SMS notification received: %+v\n", notification)
		response.Message = msgKey

	} else if notification.Channel == "email" {
		msgKey := email.ProduceEmail(notification.Message)
		response.Message = msgKey

		log.Printf("Email notification received: %+v\n", notification)

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

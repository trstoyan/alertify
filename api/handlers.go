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
	if notification.Channel == "sms" {
		sms.ProduceSMS(notification.Message)
		log.Printf("SMS notification received: %+v\n", notification)
	}
	if notification.Channel == "email" {
		email.ProduceEmail(notification.Message)
		log.Printf("Email notification received: %+v\n", notification)
	}
	// Add routing to the appropriate channel service here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification received"))

}

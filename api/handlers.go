package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/trstoyan/Alertify/email"
	"github.com/trstoyan/Alertify/slack"
	"github.com/trstoyan/Alertify/sms"
)

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Channel == "" || req.Message == "" {
		http.Error(w, "channel and message are required", http.StatusBadRequest)
		return
	}

	brokerAddress := os.Getenv("KAFKA_BROKER")
	if brokerAddress == "" {
		brokerAddress = "localhost:9092"
	}

	var err error
	switch req.Channel {
	case "sms":
		err = sms.ProduceSMSMessage(brokerAddress, sms.SMSPayload{
			PhoneNumber: req.To,
			Message:     req.Message,
		})
	case "email":
		err = email.ProduceEmailMessage(brokerAddress, email.EmailPayload{
			To:      req.To,
			Subject: "Alertify Notification",
			Body:    req.Message,
		})
	case "slack":
		err = slack.ProduceSlackMessage(brokerAddress, slack.SlackPayload{
			Channel: req.To,
			Message: req.Message,
		})
	default:
		http.Error(w, "unsupported channel: use sms, email, or slack", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("failed to produce %s notification: %v", req.Channel, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NotificationResponse{Status: "error", Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(NotificationResponse{Status: "accepted", Message: "notification queued"})
}

package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// SlackPayload holds the data needed to send a Slack notification.
type SlackPayload struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type slackWebhookBody struct {
	Text    string `json:"text"`
	Channel string `json:"channel,omitempty"`
}

// SendSlack sends a message to a Slack channel via an incoming webhook.
// The webhook URL is read from the SLACK_WEBHOOK_URL environment variable.
func SendSlack(payload SlackPayload) error {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		// Mock implementation when Slack webhook is not configured
		log.Printf("Sending Slack message to %s: %s", payload.Channel, payload.Message)
		return nil
	}

	body := slackWebhookBody{
		Text:    payload.Message,
		Channel: payload.Channel,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal Slack payload: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to send Slack message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK response from Slack: %d", resp.StatusCode)
	}
	return nil
}

package slack

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

const slackTopic = "slack-topic"

func ProduceSlackMessage(brokerAddress string, payload SlackPayload) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return kafka.ProduceWithKey(brokerAddress, slackTopic, "", string(msg))
}

func ProduceSlackResponse(brokerAddress, key, message string) {
	const responseTopic = "slack-sent-topic"
	err := kafka.ProduceWithKey(brokerAddress, responseTopic, key, message)
	if err != nil {
		log.Printf("Failed to send Slack response: %s", err)
	} else {
		log.Printf("Slack response sent successfully with key: %s\n", key)
	}
}

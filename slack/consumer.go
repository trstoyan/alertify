package slack

import (
	"encoding/json"
	"log"

	"github.com/trstoyan/Alertify/kafka"
)

// StartSlackConsumer starts a Kafka consumer that listens on the Slack topic
// and sends Slack messages for each incoming payload.
func StartSlackConsumer(brokerAddress, groupID string) {
	consumer := kafka.NewKafkaConsumer(brokerAddress, slackTopic, groupID)
	go kafka.ConsumeMessages(consumer, handleSlackMessage)
}

func handleSlackMessage(key, value string) {
	var payload SlackPayload
	err := json.Unmarshal([]byte(value), &payload)
	if err != nil {
		log.Printf("error unmarshalling Slack message: %v", err)
		return
	}

	err = SendSlack(payload)
	if err != nil {
		log.Printf("error sending Slack message: %v", err)
		return
	}

	log.Printf("Slack message sent to %s: %s", payload.Channel, payload.Message)
}

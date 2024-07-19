package email

import (
	"github.com/trstoyan/alertify/kafka"
)

// ConsumerEmail sends an email message to the email-topic Kafka topic.
func ConsumeEmail(message string) {
	kafka.Produce(message, "email-topic")
	// Note: In a real application, ensure proper shutdown logic to drain the Events channel and close the producer cleanly.
}

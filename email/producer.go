package email

import (
	"github.com/trstoyan/alertify/kafka"
)

func ProduceEmail(message string) {
	kafka.Produce(message, "email-topic")
	// Note: In a real application, ensure proper shutdown logic to drain the Events channel and close the producer cleanly.
}

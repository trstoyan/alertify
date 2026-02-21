package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/trstoyan/Alertify/api"
	"github.com/trstoyan/Alertify/email"
	"github.com/trstoyan/Alertify/slack"
	"github.com/trstoyan/Alertify/sms"
)

func main() {
	brokerAddress := os.Getenv("KAFKA_BROKER")
	if brokerAddress == "" {
		brokerAddress = "localhost:9092"
	}

	// Start notification consumers
	sms.StartSMSConsumer(brokerAddress, "sms-topic", "sms-service")
	email.StartEmailConsumer(brokerAddress, "email-service")
	slack.StartSlackConsumer(brokerAddress, "slack-service")

	// Start HTTP API server in background
	go api.StartServer()

	// Block until a termination signal is received
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Received signal %s, shutting down", sig)
}

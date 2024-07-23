package main

import (
	"github.com/trstoyan/alertify/api"
	"github.com/trstoyan/alertify/kafka"
	"github.com/trstoyan/alertify/sms"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// main is the entry point of the program.
	// It starts the server, consumes messages from Kafka topics,
	// and handles graceful shutdown.

	go api.StartServer()

	go kafka.Consume(func(key, message string) string { sms.SendSMS(key, message); return "" }, "sms-topic")
	go kafka.Consume(func(key, message string) string { api.HandleResponse(key, message); return "" }, "email-sent-topic")

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")
}

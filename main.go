package main

import (
	"github.com/trstoyan/alertify/api"
	"github.com/trstoyan/alertify/email"
	"github.com/trstoyan/alertify/kafka"
	"github.com/trstoyan/alertify/sms"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go api.StartServer()

	go kafka.Consume("sms-topic", sms.SendSMS)
	go kafka.Consume("email-topic", email.SendEmail)

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")
}

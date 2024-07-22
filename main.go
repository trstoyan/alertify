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

	go kafka.Consume(func(key, message string) string { sms.SendSMS(key, message); return "" }, "sms-topic")
	go kafka.Consume(func(key, message string) string { email.SendEmail(key, message); return "" }, "email-topic")
	go kafka.Consume(func(key, message string) string { api.HandleResponse(key, message); return "" }, "email-sent-topic")

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")
}

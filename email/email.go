package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// EmailPayload holds the data needed to send an email.
type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// SendEmail sends an email using SMTP. Credentials and server are read from
// environment variables SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD, and
// EMAIL_FROM.
func SendEmail(payload EmailPayload) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("EMAIL_FROM")

	if host == "" || port == "" {
		// Mock implementation when SMTP is not configured
		log.Printf("Sending email to %s: [%s] %s", payload.To, payload.Subject, payload.Body)
		return nil
	}

	auth := smtp.PlainAuth("", user, password, host)
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, payload.To, payload.Subject, payload.Body)
	addr := fmt.Sprintf("%s:%s", host, port)
	return smtp.SendMail(addr, auth, from, []string{payload.To}, []byte(msg))
}

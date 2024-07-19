package email

import (
	"fmt"
	"time"
)

func SendEmail(message string) {
	// Replace with actual email gateway API call
	time.Sleep(10 * time.Second)
	fmt.Println("Sending Email:", message)
}

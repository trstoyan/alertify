package api

// NotificationRequest is the request body for the /notify endpoint.
type NotificationRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
	// To is the recipient: phone number for sms, email address for email, Slack channel for slack.
	To string `json:"to"`
}

// NotificationResponse is the response body for the /notify endpoint.
type NotificationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

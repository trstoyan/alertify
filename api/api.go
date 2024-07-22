package api

type Notification struct {
	ID      string `json:"id"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// RequestBody is the structure of the request body

type RequestBody struct {
	MessageTo   string `json:"message_to"`
	MessageFrom string `json:"message_from"`
	Channel     string `json:"channel"`
	Message     string `json:"message"`
}

// ResponseBody is the structure of the response body
type ResponseBody struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

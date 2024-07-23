package api

// RequestBody is the structure of the request body

type MessageRequest struct {
	Channel     string `json:"channel"`
	MessageTo   string `json:"message_to"`
	MessageFrom string `json:"message_from"`
	Message     string `json:"message"`
}

// ResponseBody is the structure of the response body
type ResponseBody struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

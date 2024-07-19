package api

type Notification struct {
	ID      string `json:"id"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

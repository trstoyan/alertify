package api

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/notify", ValidateNotification)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

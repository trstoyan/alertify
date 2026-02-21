package api

import (
	"log"
	"net/http"
	"os"
)

// StartServer starts the HTTP API server. The port is read from the PORT
// environment variable, defaulting to 8080.
func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/notify", notifyHandler)
	log.Printf("API server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

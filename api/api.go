package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/trstoyan/alertify/api/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/notify", handlers.NotificationHandler).Methods("POST")

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", router)
}

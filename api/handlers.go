package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/trstoyan/alertify/sms"
)

const (
	channel          = "channel"
	emailChannelName = "email"
	slackChannelName = "slack"
	smsChannelName   = "sms"
)

// NotificationHandler represents the interface the notification handler must conform to.
type NotificationHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// Notifier represents the interface methods for each notification channel's producer.
type Notifier interface {
	Produce(interface{}) error
}

// Handler is a type that implements the NotificationHandler interface.
type Handler struct {
	router        *mux.Router
	emailProducer Notifier
	smsProducer   Notifier
	slackProducer Notifier
}

// NewNotificationHandler creates and returns a new handler instance.
func NewNotificationHandler(smsSvc sms.SMSService) (*Handler, error) {
	h := &Handler{
		router:      mux.NewRouter(),
		smsProducer: sms.NewSMSProducer(),
	}

	h.router.HandleFunc("/notify", h.notificationHandler).Methods("POST")
	err := http.ListenAndServe(":8080", h.router)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *Handler) notificationHandler(w http.ResponseWriter, r *http.Request) {
	var (
		message interface{}
		err     error
	)

	if err = json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, errors.Wrap(err, "invalid JSON").Error(), http.StatusBadRequest) //bad request response
		return
	}

	channelName, ok := message.(map[string]interface{})[channel].(string)
	if !ok {
		http.Error(w, "failed to extract channel from request", http.StatusBadRequest) //bad request response
		return
	}

	// route request to the appropriate notification service
	switch channelName {
	case emailChannelName:
		h.emailProducer.Produce(message)
		fmt.Fprint(w, "sent email")

	case smsChannelName:
		h.smsProducer.Produce(message)
		fmt.Fprint(w, "sent text")

	case slackChannelName:
		h.slackProducer.Produce(message)
		fmt.Fprint(w, "sent slack notification")

	default:
		http.Error(w, fmt.Sprintf("Unknown channel: [%s]. Could not route request\n", channelName), http.StatusNotFound)
		return
	}
}

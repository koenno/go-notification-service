package creatorsrv

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
)

type NotificationCreator struct {
	in  chan api.Notification
	out chan dispatcher.Sendable
}

func NewNotificationCreator(in chan api.Notification, out chan dispatcher.Sendable) NotificationCreator {
	return NotificationCreator{
		in:  in,
		out: out,
	}
}

func (s NotificationCreator) notifications(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: got unsupported content type", contentType)
		return
	}
	var notif api.Notification
	err := json.NewDecoder(r.Body).Decode(&notif)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: unable to decode notification:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	s.in <- notif
}

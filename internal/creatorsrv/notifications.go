package creatorsrv

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
)

type NotificationServer struct {
	in  chan api.Notification
	out chan types.Sendable
}

func NewNotificationServer(in chan api.Notification, out chan types.Sendable) NotificationServer {
	return NotificationServer{
		in:  in,
		out: out,
	}
}

func (s NotificationServer) notifications(w http.ResponseWriter, r *http.Request) {
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
	validate := validator.New()
	err = validate.Struct(notif)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: required field is not set:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	s.in <- notif
}

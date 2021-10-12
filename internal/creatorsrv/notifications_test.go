package creatorsrv

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestNotifications(t *testing.T) {
	// given
	in := make(chan api.Notification)
	out := make(chan types.Sendable)
	notifCreator := NewNotificationServer(in, out)
	notif := getNotification()
	payload, _ := json.Marshal(notif)
	req := httptest.NewRequest(http.MethodPost, "/notifications", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	// when
	go func() {
		notifCreator.notifications(w, req)
	}()
	// then
	inNotif := <-in
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, notif, inNotif)
}

func TestShouldReturnBadRequestIfNoContentTypeHeader(t *testing.T) {
	// given
	in := make(chan api.Notification)
	out := make(chan types.Sendable)
	notifCreator := NewNotificationServer(in, out)
	notif := getNotification()
	payload, _ := json.Marshal(notif)
	req := httptest.NewRequest(http.MethodPost, "/notifications", bytes.NewReader(payload))
	w := httptest.NewRecorder()
	// when
	notifCreator.notifications(w, req)
	// then
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestShouldReturnBadRequestIfUnableToDecodeJson(t *testing.T) {
	// given
	in := make(chan api.Notification)
	out := make(chan types.Sendable)
	notifCreator := NewNotificationServer(in, out)
	payload := []byte("{ ")
	req := httptest.NewRequest(http.MethodPost, "/notifications", bytes.NewReader(payload))
	w := httptest.NewRecorder()
	// when
	notifCreator.notifications(w, req)
	// then
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func getNotification() api.Notification {
	dateLayout := "2006-01-02T15:04:05-0700"
	date := "2021-10-11T15:04:05Z"
	reservationDate, _ := time.Parse(dateLayout, date)
	return api.Notification{
		Reservation: api.Reservation{
			ID:       644365,
			Date:     reservationDate,
			Duration: time.Duration(5400000000000),
		},
		Room: api.Room{
			Name:        "Solar Winds",
			Number:      "2F.03",
			SeatsNumber: 10,
		},
		User: api.User{
			Name: "John Smith",
			Contact: map[string]string{
				"email": "john.smith@company.com",
				"sms":   "+0123456789",
			},
		},
	}
}

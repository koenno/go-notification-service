package creator

import (
	"testing"
	"time"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNotifications(t *testing.T) {
	// given
	in := make(chan api.Notification)
	out := make(chan interface{})
	creator := NewNotificationCreator(in, out)
	go func() {
		creator.Run()
	}()
	// when
	in <- getNotification()
	// then
	notif := <-out
	assert.IsType(t, dispatcher.EMail{}, notif)
	notif = <-out
	assert.IsType(t, dispatcher.SMS{}, notif)
}

func getNotification() api.Notification {
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

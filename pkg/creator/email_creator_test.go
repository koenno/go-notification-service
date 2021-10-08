package creator

import (
	"fmt"
	"testing"
	"time"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/stretchr/testify/assert"
)

const (
	date = "Friday October 08, 2021 at 11:00"
)

func TestShouldCreateEMail(t *testing.T) {
	// given
	dispatcherPipe := make(chan interface{})
	creator := NewEMailCreator(dispatcherPipe)
	expectedEMail := dispatcher.EMail{
		Recipients: dispatcher.Recipients{
			To: []string{"john.smith@company.com"},
		},
		Subject: "Your room reservation is confirmed",
		Message: fmt.Sprintf(`
	Hello John Smith
	The room "Solar Winds" with number 2F.03
	is reserved for 1h30m0s
	starting from %s.
	The room is equipped with 10 seats.
	To cancel your reservation click here.`,
			date),
	}
	// when
	go func() {
		creator.Create(getNotification())
	}()
	// then
	email := <-dispatcherPipe
	assert.Equal(t, expectedEMail, email)
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

package creator

import (
	"fmt"
	"testing"

	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
	"github.com/stretchr/testify/assert"
)

const (
	date = "Friday October 08, 2021 at 11:00"
)

func TestShouldCreateEMail(t *testing.T) {
	// given
	dispatcherPipe := make(chan types.Sendable)
	creator := NewEMailCreator(dispatcherPipe)
	expectedEMail := dispatcher.EMail{
		Recipients: dispatcher.Recipients{
			To: []string{"john.smith@company.com"},
		},
		Subject: "Your room reservation is confirmed",
		Message: fmt.Sprintf(`
	Hello John Smith
	The reservation 644365
	of the room "Solar Winds" with number 2F.03
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

func TestShouldNotCreateEMailIfNotificationIsNotEMailOne(t *testing.T) {
	// given
	dispatcherPipe := make(chan types.Sendable)
	creator := NewEMailCreator(dispatcherPipe)
	notif := getNotification()
	delete(notif.User.Contact, "email")
	// when
	creator.Create(notif)
	// then
	assert.Equal(t, 0, len(dispatcherPipe))
}

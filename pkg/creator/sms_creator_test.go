package creator

import (
	"fmt"
	"testing"

	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateSMS(t *testing.T) {
	// given
	dispatcherPipe := make(chan types.Sendable)
	creator := NewSMSCreator(dispatcherPipe)
	expectedSMS := dispatcher.SMS{
		TelephoneNumber: "+0123456789",
		Message: fmt.Sprintf(`
	Reservation 644365 is confirmed.
	Details:
	Room number: 2F.03
	Duration: 1h30m0s
	Start date: %s`,
			date),
	}
	// when
	go func() {
		creator.Create(getNotification())
	}()
	// then
	sms := <-dispatcherPipe
	assert.Equal(t, expectedSMS, sms)
}

func TestShouldNotCreateSMSIfNotificationIsNotSMSOne(t *testing.T) {
	// given
	dispatcherPipe := make(chan types.Sendable)
	creator := NewSMSCreator(dispatcherPipe)
	notif := getNotification()
	delete(notif.User.Contact, "sms")
	// when
	creator.Create(notif)
	// then
	assert.Equal(t, 0, len(dispatcherPipe))
}

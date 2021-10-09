package creator

import (
	"fmt"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
)

type SMSCreator struct {
	dispatcherPipe chan interface{}
}

func NewSMSCreator(dispatcher chan interface{}) SMSCreator {
	return SMSCreator{
		dispatcherPipe: dispatcher,
	}
}

func (c SMSCreator) Create(notif api.Notification) {
	if number, exist := notif.User.Contact["sms"]; exist {
		sms := dispatcher.SMS{
			TelephoneNumber: number,
			Message:         c.createMessage(notif),
		}
		c.dispatcherPipe <- sms
	}
}

func (c SMSCreator) createMessage(notif api.Notification) string {
	return fmt.Sprintf(`
	Reservation %d is confirmed.
	Details:
	Room number: %s
	Duration: %s
	Start date: %s`,
		notif.Reservation.ID,
		notif.Room.Number,
		notif.Reservation.Duration,
		notif.Reservation.Date.Format(dateLayout))
}

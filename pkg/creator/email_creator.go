package creator

import (
	"fmt"
	"log"

	"github.com/dustin/go-humanize/english"
	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/dispatcher"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
)

const (
	dateLayout = "Monday January 02, 2006 at 15:04"
)

type EMailCreator struct {
	dispatcherPipe chan types.Sendable
}

func NewEMailCreator(dispatcher chan types.Sendable) EMailCreator {
	return EMailCreator{
		dispatcherPipe: dispatcher,
	}
}

func (c EMailCreator) Create(notif api.Notification) {
	if address, exist := notif.User.Contact["email"]; exist {
		log.Printf("creating email notification for reservation %d", notif.Reservation.ID)
		email := dispatcher.EMail{
			Recipients: dispatcher.Recipients{
				To: []string{address},
			},
			Subject: c.createSubject(notif),
			Message: c.createMessage(notif),
		}
		c.dispatcherPipe <- email
	}
}

func (c EMailCreator) createSubject(notif api.Notification) string {
	return "Your room reservation is confirmed"
}

func (c EMailCreator) createMessage(notif api.Notification) string {
	seatsNumber := english.Plural(notif.Room.SeatsNumber, "seat", "")
	return fmt.Sprintf(`
	Hello %s
	The reservation %d
	of the room "%s" with number %s
	is reserved for %s
	starting from %s.
	The room is equipped with %s.
	To cancel your reservation click here.`,
		notif.User.Name,
		notif.Reservation.ID,
		notif.Room.Name, notif.Room.Number,
		notif.Reservation.Duration,
		notif.Reservation.Date.Format(dateLayout),
		seatsNumber)
}

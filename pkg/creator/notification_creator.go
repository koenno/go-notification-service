package creator

import (
	"log"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
	"github.com/koen-or-nant/go-notification-service/pkg/types"
)

type Creator interface {
	Create(api.Notification)
}

type NotificationCreator struct {
	notifs   chan api.Notification
	creators []Creator
}

type Sendable interface {
	Send()
}

func NewNotificationCreator(notifs chan api.Notification, sendPipe chan types.Sendable) NotificationCreator {
	return NotificationCreator{
		notifs: notifs,
		creators: []Creator{
			NewEMailCreator(sendPipe),
			NewSMSCreator(sendPipe),
		},
	}
}

func (c NotificationCreator) Run() {
	for notif := range c.notifs {
		log.Println("processing reservation", notif.Reservation.ID)
		for _, creator := range c.creators {
			creator.Create(notif)
		}
	}
}

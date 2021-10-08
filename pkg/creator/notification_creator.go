package creator

import (
	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

type Creator interface {
	Create(api.Notification)
}

type NotificationCreator struct {
	notifs   chan api.Notification
	creators []Creator
}

func NewNotificationCreator(notifs chan api.Notification, dispatcher chan interface{}) NotificationCreator {
	return NotificationCreator{
		notifs: notifs,
		creators: []Creator{
			NewEMailCreator(dispatcher),
			SMSCreator{},
		},
	}
}

func (c NotificationCreator) Run() {
	for notif := range c.notifs {
		for _, creator := range c.creators {
			creator.Create(notif)
		}
	}
}
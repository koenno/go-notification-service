package creator

import (
	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

type Creator interface {
	Create(api.Notification)
}

var creators []Creator

func init() {
	creators = []Creator{
		EMailCreator{},
		SMSCreator{},
	}
}

type NotificationCreator struct {
	notifs chan api.Notification
}

func NewNotificationCreator(notifs chan api.Notification) NotificationCreator {
	return NotificationCreator{
		notifs: notifs,
	}
}

func (c NotificationCreator) Run() {
	for notif := range c.notifs {
		for _, creator := range creators {
			creator.Create(notif)
		}
	}
}

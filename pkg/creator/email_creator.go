package creator

import (
	"log"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

type EMailCreator struct {
}

func (c EMailCreator) Create(notif api.Notification) {
	if address, exist := notif.User.Contact["email"]; exist {
		log.Print(address)
	}
}

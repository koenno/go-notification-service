package creator

import (
	"log"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

type SMSCreator struct {
}

func (c SMSCreator) Create(notif api.Notification) {
	if address, exist := notif.User.Contact["sms"]; exist {
		log.Print(address)
	}
}

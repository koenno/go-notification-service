package dispatcher

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/koen-or-nant/go-notification-service/internal/config"
)

type SMS struct {
	TelephoneNumber string `json:"telephoneNumber"`
	Message         string `json:"message"`
}

const (
	SMS_GATEWAY_ENV = "SMS_GATEWAY_ADDRESS"
)

func (s SMS) Send() {
	smsSendEndpoint, exist := config.GetConfig(SMS_GATEWAY_ENV)
	if !exist {
		log.Printf("ERROR: env %s is not set", SMS_GATEWAY_ENV)
		return
	}
	email, err := json.Marshal(s)
	if err != nil {
		log.Println("ERROR: unable to marshal sms:", err)
		return
	}
	resp, err := http.Post(smsSendEndpoint, "application/json", bytes.NewReader(email))
	if err != nil {
		log.Println("ERROR: unable to send an sms:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR: send an sms failed:", err)
		return
	}
}

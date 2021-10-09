package dispatcher

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/koen-or-nant/go-notification-service/internal/config"
)

type Recipients struct {
	To  []string
	CC  []string
	BCC []string
}

type EMail struct {
	Recipients Recipients
	Subject    string
	Message    string
}

const (
	EMAIL_SENDER_ENV = "EMAIL_SENDER_ADDRESS"
)

func (e EMail) Send() {
	emailSendEndpoint, exist := config.GetConfig(EMAIL_SENDER_ENV)
	if !exist {
		log.Printf("ERROR: env %s is not set", EMAIL_SENDER_ENV)
		return
	}
	email, err := json.Marshal(e)
	if err != nil {
		log.Println("ERROR: unable to marshal email:", err)
		return
	}
	resp, err := http.Post(emailSendEndpoint, "application/json", bytes.NewReader(email))
	if err != nil {
		log.Println("ERROR: unable to send an email:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR: send an email failed:", err)
		return
	}
}

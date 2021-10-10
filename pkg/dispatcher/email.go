package dispatcher

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/koen-or-nant/go-notification-service/internal/config"
)

type Recipients struct {
	To  []string `json:"to"`
	CC  []string `json:"cc"`
	BCC []string `json:"bcc"`
}

type EMail struct {
	Recipients Recipients `json:"recipients"`
	Subject    string     `json:"subject"`
	Message    string     `json:"message"`
}

const (
	EMAIL_GATEWAY_ENV = "EMAIL_GATEWAY_ADDRESS"
)

func (e EMail) Send() {
	emailSendEndpoint, exist := config.GetConfig(EMAIL_GATEWAY_ENV)
	if !exist {
		log.Printf("ERROR: env %s is not set", EMAIL_GATEWAY_ENV)
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

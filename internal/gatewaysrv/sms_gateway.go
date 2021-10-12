package gatewaysrv

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

func sms(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: got unsupported content type", contentType)
		return
	}
	var sms api.SMS
	err := json.NewDecoder(r.Body).Decode(&sms)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: unable to decode sms:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	printSMS(sms)
}

func printSMS(sms api.SMS) {
	log.Printf(`
	=============================== SMS ===============================
	Telephone number: %s
	-------------------------------------------------------------------
	%s
	===================================================================`,
		sms.TelephoneNumber,
		sms.Message)
}

package gatewaysrv

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/koen-or-nant/go-notification-service/pkg/api"
)

func email(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: got unsupported content type", contentType)
		return
	}
	var email api.EMail
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: unable to decode email:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	printEMail(email)
}

func printEMail(email api.EMail) {
	log.Printf(`
	============================== EMAIL ==============================
	Subject: %s
	-------------------------------------------------------------------
	TO: %s
	CC: %s
	BCC: %s
	-------------------------------------------------------------------
	%s
	===================================================================`,
		email.Subject,
		strings.Join(email.Recipients.To, ";"),
		strings.Join(email.Recipients.CC, ";"),
		strings.Join(email.Recipients.BCC, ";"),
		email.Message)
}

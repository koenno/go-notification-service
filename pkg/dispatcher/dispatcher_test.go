package dispatcher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldDispatchEMail(t *testing.T) {
	// given
	pipe := make(chan Sendable)
	sendPipe := make(chan EMail)
	dispatcher := NewDispatcher(pipe)
	expectedEMail := EMail{
		Recipients: Recipients{
			To: []string{"john.smith@company.com"},
		},
		Subject: "some subject",
		Message: "Hello Mr. Smith",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var email EMail
		err := json.NewDecoder(r.Body).Decode(&email)
		assert.Nil(t, err)
		sendPipe <- email
	}))
	defer server.Close()
	t.Setenv(EMAIL_GATEWAY_ENV, server.URL)
	go func() {
		dispatcher.Run()
	}()
	// when
	pipe <- EMail{
		Recipients: Recipients{
			To: []string{"john.smith@company.com"},
		},
		Subject: "some subject",
		Message: "Hello Mr. Smith",
	}
	// then
	email := <-sendPipe
	assert.Equal(t, expectedEMail, email)
}

func TestShouldDispatchSMS(t *testing.T) {
	// given
	pipe := make(chan Sendable)
	sendPipe := make(chan SMS)
	dispatcher := NewDispatcher(pipe)
	expectedSMS := SMS{
		TelephoneNumber: "+0123456789",
		Message:         "Hello Mr. Smith",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sms SMS
		err := json.NewDecoder(r.Body).Decode(&sms)
		assert.Nil(t, err)
		sendPipe <- sms
	}))
	defer server.Close()
	t.Setenv(SMS_GATEWAY_ENV, server.URL)
	go func() {
		dispatcher.Run()
	}()
	// when
	pipe <- SMS{
		TelephoneNumber: "+0123456789",
		Message:         "Hello Mr. Smith",
	}
	// then
	sms := <-sendPipe
	assert.Equal(t, expectedSMS, sms)
}

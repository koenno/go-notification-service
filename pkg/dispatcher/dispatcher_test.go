package dispatcher

import "testing"

func TestShouldDispatchEMail(t *testing.T) {
	// given
	email_sender_address := "http://sender.com/email"
	t.Setenv(EMAIL_SENDER_ENV, email_sender_address)
	pipe := make(chan Sendable)
	dispatcher := NewDispatcher(pipe)
	go func() {
		dispatcher.Run()
	}()
	// when
	pipe <- EMail{}
	// then

}

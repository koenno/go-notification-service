package dispatcher

import "github.com/koen-or-nant/go-notification-service/pkg/types"

type Sendable interface {
	Send()
}

type Dispatcher struct {
	inPipe chan types.Sendable
}

func NewDispatcher(pipe chan types.Sendable) Dispatcher {
	return Dispatcher{
		inPipe: pipe,
	}
}

func (d Dispatcher) Run() {
	for msg := range d.inPipe {
		msg.Send()
	}
}

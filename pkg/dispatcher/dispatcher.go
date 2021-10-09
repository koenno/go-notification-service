package dispatcher

type Sendable interface {
	Send()
}

type Dispatcher struct {
	inPipe chan Sendable
}

func NewDispatcher(pipe chan Sendable) Dispatcher {
	return Dispatcher{
		inPipe: pipe,
	}
}

func (d Dispatcher) Run() {
	for msg := range d.inPipe {
		msg.Send()
	}
}

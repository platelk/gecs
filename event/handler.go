package event

type ReceiverID uint32

type Receiver interface {
	ID() uint
	Received() []Eventer
}


type Handler interface {
	Type() string
	Send(e Eventer) error
	Register() Receiver
	Unregister(r Receiver)
}
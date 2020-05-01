package event

import "fmt"

type DefaultReceiver struct {
	id uint
	call func() []Eventer
}

func (d *DefaultReceiver) ID() uint {
	return d.id
}

func (d *DefaultReceiver) Received() []Eventer {
	return d.call()
}

type DefaultHandler struct {
	t          string
	eventQueue []Eventer
	readsQueue []uint32
	currentID  uint64
	maxListener uint
	listener   map[uint]uint64
}

func NewDefaultHandler(e Eventer) *DefaultHandler {
	return &DefaultHandler{
		t:        e.Type(),
		listener: make(map[uint]uint64),
	}
}

func (d *DefaultHandler) Type() string {
	return d.t
}

func (d *DefaultHandler) Send(e Eventer) error {
	if e.Type() != d.Type() {
		return fmt.Errorf("event of type %s is not of handler of type %s", e.Type(), d.Type())
	}
	d.eventQueue = append(d.eventQueue, e)
	d.readsQueue = append(d.readsQueue, uint32(len(d.listener)))
	return nil
}

func (d *DefaultHandler) Register() Receiver {
	listenerID := d.maxListener
	d.maxListener++
	f := func() []Eventer {
		events := d.read(listenerID)
		d.listener[listenerID] = d.currentID + uint64(len(d.eventQueue))
		d.clean()
		return events
	}
	d.listener[listenerID] = d.currentID + uint64(len(d.eventQueue))

	return &DefaultReceiver{
		id: listenerID,
		call: f,
	}
}

func (d *DefaultHandler) Unregister(r Receiver) {
	for i := range d.read(r.ID()) {
		d.readsQueue[i]--
	}
	delete(d.listener, r.ID())
	d.clean()
}

func (d *DefaultHandler) read(id uint) []Eventer {
	i, ok := d.listener[id]
	if !ok {
		return nil
	}
	return d.eventQueue[i-d.currentID:]
}

func (d *DefaultHandler) clean() {
	for len(d.readsQueue) > 0 && d.readsQueue[0] == 0 {
		d.readsQueue = d.readsQueue[1:]
		d.eventQueue = d.eventQueue[1:]
	}
}

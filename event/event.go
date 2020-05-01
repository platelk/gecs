package event

type Eventer interface {
	Data() interface{}
	Emitter() string
	ID() string
	Type() string
}

type Event struct {
	id string
	emitter string
	data interface{}
	t string
}

func (e *Event) Data() interface{} {
	return e.data
}

func (e *Event) Emitter() string {
	return e.emitter
}

func (e *Event) ID() string {
	return e.id
}

func (e *Event) Type() string {
	return e.t
}

func NewEvent(t, id, emitter string, data interface{}) *Event {
	return &Event{
		id: id,
		t: t,
		emitter: emitter,
		data: data,
	}
}

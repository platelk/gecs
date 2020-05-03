package dispatcher

import (
	"gecs/system"
)

type DispatchBuilder interface {
	With(system system.System) DispatchBuilder
	Build() Dispatcher
}

type Dispatcher interface {
	Dispatch(d *system.Data)
}

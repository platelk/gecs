package system

type Typer interface {
	Type() string
}

// Systems are centralized game engine subsystems that perform a specific function,
// such as rendering, physics, audio, etc. Every frame,
// they process each entity in the game world looking for components that are relevant to them,
// reading their contents, and performing actions.
// For example, a Rendering system could search for all entities that have Light, Mesh,
// or Emitter components and draw them to the screen.
type System interface {
	Typer
	Run(d *Data)
}

type systemImpl struct {
	t string
	f func(*Data)
}

func New(s string, f func(*Data)) System {
	return &systemImpl{
		t: s,
		f: f,
	}
}

func (s *systemImpl) Type() string {
	return s.t
}

func (s *systemImpl) Run(d *Data) {
	s.f(d)
}

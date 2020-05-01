# GECS

`gecs` is an implementation of the Entity Component System 
which define an other model to modelise game logic than standard OOP

# Example

```go

var PositionType = &Position{}

type Position struct {
	component.Base
	X, Y float32
}

func (p *Position) Type() string {
	return "position"
}

var DirectionType = &Direction{}

type Direction struct {
	component.Base
	X, Y float32
}

func (m *Direction) Type() string {
	return "direction"
}

var ComflabulationType = &Comflabulation{}

type Comflabulation struct {
	component.Base
	Thingy float32
	Dingy int
	Mingy bool
	Stringy string
}

func (m *Comflabulation) Type() string {
	return "comflabulation"
}

func newMoveSystem() system.System {
	return system.New("move", func(data *system.Data) {
		res := data.Query(query.Entities().With(DirectionType).With(PositionType))
		for _, row := range res.Rows() {
			m := row.Get(DirectionType).(*Direction)
			p := row.Get(PositionType).(*Position)
			p.X += m.X
			p.Y += m.Y
		}
	})
}

func newComflabSystem() system.System {
	return system.New("conflab", func(data *system.Data) {
		res := data.Query(query.Entities().With(ComflabulationType))
		for _, c := range res.Component(ComflabulationType) {
			comflab := c.(*Comflabulation)
			comflab.Thingy *= 1.000001
			comflab.Mingy = !comflab.Mingy
			comflab.Dingy++
		}
	})
}


func run() {
    entities := 25
    em := entity.NewSliceManager()
    sm := storage.NewDefaultManager()
	sm.Add(storage.NewMap(DirectionType))
	sm.Add(storage.NewMap(PositionType))
	sm.Add(storage.NewMap(ComflabulationType))
	data := system.NewData(em, sm)
	for i := 0; i < entities; i++ {
		builder := data.NewEntity().
			WithComponent(&Position{}).
			WithComponent(&Direction{})
		if i % 2 == 0 {
			builder.WithComponent(&Comflabulation{})
		}
		_, _ = builder.Build()
	}
	for i := 0; i < b.N; i++ {
		moveSys.Run(data)
		comfSys.Run(data)
	}
}
```
# GECS

`gecs` is an implementation of the Entity Component System 
which define an other model to modelise game logic than standard OOP

## Introduction

## Features

* Manage entity <=> components relation
* Handle multiple type of storage for component
* Systems can define relationship between each other to create execution plan
* Query entities based on their components 

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
	Thingy  float32
	Dingy   int
	Mingy   bool
	Stringy string
}

func (m *Comflabulation) Type() string {
	return "comflabulation"
}

func newMoveSystem() system.System {
	return system.NewWithPlan("move",
		system.Plan{}.WithRead(DirectionType).WithWrite(PositionType),
		func(data *system.Data) {
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
	return system.NewWithPlan("conflab",
		system.Plan{}.WithWrite(ComflabulationType),
		func(data *system.Data) {
			res := data.Query(query.Entities().With(ComflabulationType))
			for _, c := range res.Component(ComflabulationType) {
				comflab := c.(*Comflabulation)
				comflab.Thingy *= 1.000001
				comflab.Mingy = !comflab.Mingy
				comflab.Dingy++
			}
		})
}

func ExampleBasic() {
	em := entity.NewManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	data := system.NewData(em, sm)
	for i := 0; i < 25; i++ {
		builder := data.NewEntity().
			WithComponent(&Position{}).
			WithComponent(&Direction{})
		if i%2 == 0 {
			builder.WithComponent(&Comflabulation{})
		}
		_, _ = builder.Build()
	}
	d := dispatcher.NewSingleBuilder().
		With(newMoveSystem()).
		With(newComflabSystem()).
		Build()
	d.Dispatch(data)
}

```
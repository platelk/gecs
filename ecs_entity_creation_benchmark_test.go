package ecs

import (
	"fmt"
	"testing"
	"gecs/component"
	"gecs/entity"
	"gecs/query"
	"gecs/storage"
	"gecs/system"
)

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

var nbEntities = []int{
	10, 25, 50,
	100, 200, 400, 800,
	1600, 3200, 5000,
	10_000, 30_000,
	100_000, 500_000,
	1_000_000, 2_000_000, 5_000_000,
	10_000_000, 20_000_000,
}

func BenchmarkCreateEntities(b *testing.B) {
	entityManager := map[string]func() entity.Manager{
		"slice": func() entity.Manager { return entity.NewSliceManager() },
	}
	storageCreator := map[string]func(c component.Component) storage.Storage {
		"map": func(c component.Component) storage.Storage {
			return storage.NewMap(c)
		},
	}
	for name, emCreator := range entityManager {
		for sname, sCreator := range storageCreator {
			for _, entities := range nbEntities {
				it := entities
				b.Run(fmt.Sprintf("entity-creation/em=%s/sc=%s/size=%d", name, sname, it), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						n := it
						em := emCreator()
						sm := storage.NewDefaultManager()
						sm.Add(sCreator(DirectionType))
						sm.Add(sCreator(PositionType))
						sm.Add(sCreator(ComflabulationType))
						data := system.NewData(em, sm)
						for ; n > 0; n-- {
							_, _ = data.NewEntity().
								WithComponent(&Position{}).
								WithComponent(&Direction{}).
								WithComponent(&Comflabulation{}).
								Build()
						}
					}
				})
			}
		}
	}
}


func BenchmarkUpdateSystem(b *testing.B) {
	moveSys := newMoveSystem()
	comfSys := newComflabSystem()
	for _, n := range nbEntities {
		it := n
		b.Run(fmt.Sprintf("2-system-update/size=%d", it), func(b *testing.B) {
			em := entity.NewSliceManager()
			sm := storage.NewDefaultManager()
			sm.Add(storage.NewMap(DirectionType))
			sm.Add(storage.NewMap(PositionType))
			sm.Add(storage.NewMap(ComflabulationType))
			data := system.NewData(em, sm)
			for i := 0; i < it; i++ {
				builder := data.NewEntity().
					WithComponent(&Position{}).
					WithComponent(&Direction{})
				if it % 2 == 0 {
					builder.WithComponent(&Comflabulation{})
				}
				_, _ = builder.Build()

			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				moveSys.Run(data)
				comfSys.Run(data)
			}
		})
	}
}
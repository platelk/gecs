package query

import (
	"testing"
	"gecs/component"
	"gecs/entity"
	"gecs/storage"

	"github.com/stretchr/testify/require"
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
	Thingy  float32
	Dingy   int
	Mingy   bool
	Stringy string
}

func (m *Comflabulation) Type() string {
	return "comflabulation"
}

func TestEntities(t *testing.T) {
	em := entity.NewSliceManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	e := em.New()
	sm.GetStorage(DirectionType).Add(e, &Direction{})
	sm.GetStorage(PositionType).Add(e, &Position{})
	sm.GetStorage(ComflabulationType).Add(e, &Comflabulation{})
	res := Entities().With(DirectionType).With(PositionType).Exec(sm)
	require.Equal(t, len(res.Entities()), 1)
	require.Equal(t, res.Entities()[0].ID(), e.ID())
}

func TestEntities_select_based_on_component(t *testing.T) {
	em := entity.NewSliceManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	e1 := em.New()
	sm.GetStorage(DirectionType).Add(e1, &Direction{})
	sm.GetStorage(PositionType).Add(e1, &Position{})
	sm.GetStorage(ComflabulationType).Add(e1, &Comflabulation{})

	e2 := em.New()
	sm.GetStorage(PositionType).Add(e2, &Position{})
	res := Entities().With(DirectionType).With(PositionType).Exec(sm)
	require.Equal(t, len(res.Entities()), 1)
	require.Equal(t, res.Entities()[0].ID(), e1.ID())
}

func TestEntities_select_or_only_one_based_on_component(t *testing.T) {
	em := entity.NewSliceManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	e1 := em.New()
	sm.GetStorage(DirectionType).Add(e1, &Direction{})
	sm.GetStorage(PositionType).Add(e1, &Position{})
	sm.GetStorage(ComflabulationType).Add(e1, &Comflabulation{})

	e2 := em.New()
	sm.GetStorage(PositionType).Add(e2, &Position{})
	res := Entities().With(DirectionType).With(PositionType).OrWith(ComflabulationType).Exec(sm)
	require.Equal(t, len(res.Entities()), 1)
	require.Equal(t, res.Entities()[0].ID(), e1.ID())
}

func TestEntities_select_multiple_or_based_on_component(t *testing.T) {
	em := entity.NewSliceManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	e1 := em.New()
	sm.GetStorage(DirectionType).Add(e1, &Direction{})
	sm.GetStorage(PositionType).Add(e1, &Position{})

	e2 := em.New()
	sm.GetStorage(ComflabulationType).Add(e2, &Comflabulation{})
	res := Entities().With(DirectionType).With(PositionType).OrWith(ComflabulationType).Exec(sm)
	require.Equal(t, len(res.Entities()), 2)
}

func TestEntities_select_row_multiple_or_based_on_component(t *testing.T) {
	em := entity.NewSliceManager()
	sm := storage.NewMapManager()
	sm.AddStorage(storage.NewMap(DirectionType))
	sm.AddStorage(storage.NewMap(PositionType))
	sm.AddStorage(storage.NewMap(ComflabulationType))
	e1 := em.New()
	sm.GetStorage(DirectionType).Add(e1, &Direction{})

	e2 := em.New()
	sm.GetStorage(DirectionType).Add(e2, &Direction{})
	res := Entities().With(DirectionType).Exec(sm)
	require.Equal(t, 2, len(res.Rows()))
	for _, row := range res.Rows() {
		require.NotNil(t, row.Get(DirectionType))
	}
}

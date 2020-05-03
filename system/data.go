package system

import (
	"gecs/component"
	"gecs/entity"
	"gecs/event"
	"gecs/query"
	"gecs/storage"
)

type Data struct {
	entities   entity.Manager
	components storage.Manager
	events     map[string]event.Handler
}

func NewData(em entity.Manager, s storage.Manager) *Data {
	return &Data{
		entities:   em,
		components: s,
		events:     make(map[string]event.Handler),
	}
}

func (d *Data) AddComponents(s storage.Storage) {
	d.components.AddStorage(s)
}

func (d *Data) AddEvents(e event.Handler) {
	d.events[e.Type()] = e
}

func (d *Data) Entities() entity.Manager {
	return d.entities
}

func (d *Data) NewEntity() Builder {
	return NewBaseBuilder(d.entities, d.components)
}

func (d *Data) Components(t Typer) storage.Storage {
	return d.components.GetStorage(t)
}

func (d *Data) Events(t Typer) event.Handler {
	return d.events[t.Type()]
}

func (d *Data) Query(q query.Querier) *query.Result {
	return q.Exec(d.components)
}

type Builder interface {
	WithComponent(c component.Component) Builder
	Build() (entity.Entity, error)
}

type BaseBuilder struct {
	entityManager    entity.Manager
	componentManager storage.Manager
	components       []component.Component
}

func NewBaseBuilder(em entity.Manager, cm storage.Manager) *BaseBuilder {
	return &BaseBuilder{
		entityManager:    em,
		componentManager: cm,
	}
}

func (bb *BaseBuilder) WithComponent(c component.Component) Builder {
	bb.components = append(bb.components, c)

	return bb
}

func (bb *BaseBuilder) Build() (entity.Entity, error) {
	e := bb.entityManager.New()
	for _, c := range bb.components {
		if err := bb.componentManager.GetStorage(c).Add(e, c); err != nil {
			return entity.Entity{}, err
		}
	}

	return e, nil
}

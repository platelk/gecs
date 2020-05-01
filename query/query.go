package query

import (
	"gecs/component"
	"gecs/entity"
	"gecs/storage"
)

type Typer interface {
	Type() string
}

type Querier interface {
	Exec(sm storage.Manager) *Result
}

type Row struct {
	e          entity.Entity
	components map[string]component.Component
}

func (r *Row) Entity() entity.Entity {
	return r.e
}

func (r *Row) Get(t Typer) component.Component {
	return r.components[t.Type()]
}

type Result struct {
	entities   map[uint32]entity.Entity
	components map[uint32]map[string]component.Component
}

func (r *Result) Entities() []entity.Entity {
	entities := make([]entity.Entity, len(r.entities))
	var i int
	for _, e := range r.entities {
		entities[i] = e
		i++
	}
	return entities
}

func (r *Result) Component(t Typer) []component.Component {
	var components []component.Component
	for _, s := range r.components {
		components = append(components, s[t.Type()])
	}
	return components
}

func (r *Result) Merge(other *Result) {
	for id, e := range other.entities {
		r.entities[id] = e
	}
	for e, s := range other.components {
		for t, c := range s {
				if _, ok := r.components[e]; !ok {
					r.components[e] = make(map[string]component.Component)
				}
				r.components[e][t] = c
		}
	}
}

func (r *Result) Rows() []*Row {
	var rows []*Row
	for e, s := range r.components {
		rows = append(rows, &Row{
			e:          r.entities[e],
			components: s,
		})
	}
	return rows
}

type EntityQuery struct {
	and []Typer
	or  []*EntityQuery
}

func Entities() *EntityQuery {
	return &EntityQuery{}
}

func (eq *EntityQuery) With(c component.Component) *EntityQuery {
	eq.and = append(eq.and, c)
	return eq
}

func (eq *EntityQuery) OrWith(c component.Component) *EntityQuery {
	return &EntityQuery{
		and: []Typer{c},
		or:  []*EntityQuery{eq},
	}
}

func (eq *EntityQuery) Exec(sm storage.Manager) *Result {
	entities := make(map[uint32]entity.Entity)
	components := make(map[uint32]map[string]component.Component)
	if len(eq.and) > 0 {
		sm.Get(eq.and[0]).Iter(func(e entity.Entity, c component.Component) {
			cs := []component.Component{c}
			for _, t := range eq.and[1:] {
				s := sm.Get(t)
				if s != nil {
					c = s.Get(e)
					if c == nil {
						return
					}
					cs = append(cs, c)
				}
			}
			for _, c := range cs {
				if _, ok := components[e.ID()]; !ok {
					components[e.ID()] = make(map[string]component.Component)
				}
				components[e.ID()][c.Type()] = c
			}
			entities[e.ID()] = e
		})
	}
	res := &Result{
		entities:   entities,
		components: components,
	}
	for _, o := range eq.or {
		res.Merge(o.Exec(sm))
	}
	return res
}

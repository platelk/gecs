package storage

import (
	"fmt"
	"gecs/component"
	"gecs/entity"
)

// Map will store the component in a simple slice
type Map struct {
	t string
	maxID uint
	components map[entity.Entity]component.Component
}

func NewMap(c component.Component) *Map {
	return &Map{
		t: c.Type(),
		components: make(map[entity.Entity]component.Component),
	}
}

func (s *Map) Get(e entity.Entity) component.Component {
	return s.components[e]
}

func (s *Map) Add(e entity.Entity, c component.Component) error {
	if c.Type() != s.t {
		return fmt.Errorf("can't add component of type %s in storage of type %s", c.Type(), s.t)
	}
	c.SetID(s.maxID)
	s.maxID++
	s.components[e] = c
	return nil
}

func (s *Map) Del(e entity.Entity) error {
	delete(s.components, e)
	return nil
}

func (s *Map) Len() int {
	return len(s.components)
}

func (s *Map) Type() string {
	return s.t
}

func (s *Map) Iter(f func(e entity.Entity, c component.Component)) {
	for e, c := range s.components {
		f(e, c)
	}
}


package storage

import (
	"gecs/component"
	"gecs/entity"
)

type MapManager struct {
	components map[string]Storage
}

func NewMapManager() *MapManager {
	return &MapManager{
		components: map[string]Storage{},
	}
}

func (d *MapManager) Type() string {
	return "manager"
}

func (d *MapManager) Add(e entity.Entity, c component.Component) error {
	return d.GetStorage(c).Add(e, c)
}

func (d *MapManager) Del(e entity.Entity) error {
	for _, s := range d.components {
		if err := s.Del(e); err != nil {
			return err
		}
	}
	return nil
}

func (d *MapManager) Get(e entity.Entity) []component.Component {
	var cs []component.Component
	for _, s := range d.components {
		c := s.Get(e)
		if c != nil {
			cs = append(cs, c)
		}
	}
	return cs
}

func (d *MapManager) AddStorage(s Storage) Manager {
	d.components[s.Type()] = s
	return d
}

func (d *MapManager) DelStorage(s Storage) Manager {
	delete(d.components, s.Type())
	return d
}

func (d *MapManager) GetByType(t string) Storage {
	return d.components[t]
}

func (d *MapManager) GetStorage(t Typer) Storage {
	return d.GetByType(t.Type())
}

func (d *MapManager) Exist(t Typer) bool {
	_, ok := d.components[t.Type()]
	return ok
}

func (d *MapManager) Len() int {
	return len(d.components)
}

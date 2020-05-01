package storage

import (
	"gecs/component"
	"gecs/entity"
)

// Storage will define a storage definition for the component
type Storage interface {
	Type() string
	Add(e entity.Entity, c component.Component) error
	Del(e entity.Entity) error
	Get(e entity.Entity) component.Component
	Iter(func(e entity.Entity, c component.Component))
	Len() int
}
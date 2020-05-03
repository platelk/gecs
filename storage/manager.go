package storage

import (
	"gecs/component"
	"gecs/entity"
)

type Typer interface {
	Type() string
}

type Manager interface {
	Add(e entity.Entity, c component.Component) error
	Del(e entity.Entity) error
	Get(e entity.Entity) []component.Component
	AddStorage(s Storage) Manager
	DelStorage(s Storage) Manager
	GetStorage(t Typer) Storage
	Exist(t Typer) bool
	Len() int
}

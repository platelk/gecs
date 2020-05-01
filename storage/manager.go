package storage

type Typer interface {
	Type() string
}

type Manager interface {
	Add(s Storage) Manager
	Del(s Storage) Manager
	Get(t Typer) Storage
	Exist(t Typer) bool
	Len() int
}

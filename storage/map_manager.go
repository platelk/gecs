package storage

type DefaultManager struct {
	components map[string]Storage
}

func NewDefaultManager() *DefaultManager {
	return &DefaultManager{
		components: map[string]Storage{},
	}
}

func (d *DefaultManager) Add(s Storage) Manager {
	d.components[s.Type()] = s
	return d
}

func (d *DefaultManager) Del(s Storage) Manager {
	delete(d.components, s.Type())
	return d
}

func (d *DefaultManager) GetByType(t string) Storage {
	return d.components[t]
}

func (d *DefaultManager) Get(t Typer) Storage {
	return d.GetByType(t.Type())
}

func (d *DefaultManager) Exist(t Typer) bool {
	_, ok := d.components[t.Type()]
	return ok
}

func (d *DefaultManager) Len() int {
	return len(d.components)
}


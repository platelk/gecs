package entity

// Entity represents a single object in the game world. Has no functionality on its own.
// The world owns a collection of entities (either in a flat list or a hierarchy).
// Each entity has a unique identifier or name, for the sake of ease of use.
type Entity struct {
	id uint32
	version uint32
}

// ID will return the unique identifier attributed by the Manager
func (e Entity) ID() uint32 {
	return e.id
}

// Manager define the expected behavior of an entity manager
// its main role is to create new entity with unique identifier and be able to destroy and check them
type Manager interface {
	New() Entity
	IsValid(e Entity) bool
	Invalidate(e Entity)
	Len() int
}
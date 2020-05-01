package component

// Component is  plain-old-data structure that describes a certain trait an entity can have.
// Can be "attached" to entities to grant them certain abilities,
// e.g. a Light component contains parameters to make an entity glow,
// or a Collidable component can grant an entity collision detection properties.
// These components do not have any logic. They contain only data.
type Component interface {
	// ID is an unique identifier set by the Storage to be able to differentiate multiple component in the same storage
	ID() uint
	// SetID will set the unique identifier inside the storage
	SetID(id uint)
	// Type will return the underlying type of the component, this is mostly to avoid use of reflect package
	Type() string
}

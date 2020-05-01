package component

// Base define a basic component behavior and provide a simple way to provide repeating set of function
// just add it to your component definition
type Base struct {
	id uint
}

// SetID implements Component.SetID
func (c *Base) SetID(id uint) {
	c.id = id
}

// ID implements Component.ID
func (c *Base) ID() uint {
	return c.id
}

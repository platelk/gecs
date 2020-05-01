package component

// Tag define a simple to use empty component use to select specific entities
type Tag struct {
	Base
	value string
}

// NewTag will create an initialize Tag
func NewTag(value string) *Tag {
	return &Tag{
		Base:  Base{},
		value: value,
	}
}

// Type implements Component.Type
func (t *Tag) Type() string {
	return t.value
}

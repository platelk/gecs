package storage

import (
	"testing"
	"gecs/component"
)

func TestNewDefaultManager(t *testing.T) {
	runManagerTestSuite(
		t,
		func(t string) Storage {
			return NewMap(component.NewTag(t))
		},
		func(t string) component.Component {
			return component.NewTag(t)
		},
		NewDefaultManager(),
	)
}

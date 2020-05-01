package storage

import (
	"testing"
	"gecs/component"
	"gecs/entity"
)

func TestNewMap(t *testing.T) {
	runStorageTestSuite(t, entity.NewSliceManager(), func() component.Component {
		return component.NewTag("player")
	}, NewMap(component.NewTag("player")))
}

func BenchmarkNewMap(b *testing.B) {
	runStorageBenchmarkSuite(b, entity.NewSliceManager(), func(t string) component.Component {
		return component.NewTag(t)
	}, func(c component.Component) Storage {
		return NewMap(c)
	})
}
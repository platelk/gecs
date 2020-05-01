package storage

import (
	"fmt"
	"testing"
	"gecs/component"
	"gecs/entity"
)


var nbEntities = []int{
	10, 25, 50,
	100, 200, 400, 800,
	1600, 3200, 5000,
	10_000, 30_000,
	100_000, 500_000,
	1_000_000, 2_000_000, 5_000_000,
	10_000_000, 20_000_000,
}

func runStorageBenchmarkSuite(b *testing.B, em entity.Manager, compCreator func(string) component.Component, storageCreator func(component2 component.Component) Storage) {
	b.Run("multiple add", func(b *testing.B) {
		for _, nbEntity := range nbEntities {
			n := nbEntity
			b.Run(fmt.Sprintf("size=%d", nbEntity), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					c := compCreator("test1")
					s := storageCreator(c)
					for j := 0; j < n; j++ {
						_ = s.Add(em.New(), c)
					}
				}
			})
		}
	})
	b.Run("multiple add and delete", func(b *testing.B) {
		for _, nbEntity := range nbEntities {
			n := nbEntity
			b.Run(fmt.Sprintf("size=%d", nbEntity), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					c := compCreator("test1")
					s := storageCreator(c)
					for j := 0; j < n; j++ {
						e := em.New()
						_ = s.Add(e, c)
						_ = s.Del(e)
					}
				}
			})
		}
	})
}
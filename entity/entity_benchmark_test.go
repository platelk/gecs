package entity

import (
	"fmt"
	"testing"
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

func BenchmarkEntityManagerSuite(t *testing.B) {
	managers := map[string]func() Manager{
		"slice": func() Manager {
			return NewSliceManager()
		},
		"bitset": func() Manager {
			return NewBitSetManager()
		},
	}
	for name, managerCreator := range managers {
		for _, n := range nbEntities {
			entities := n
			t.Run(fmt.Sprintf("multiple-create/kind=%s/size=%d", name, entities), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					m := managerCreator()
					for n := 0; n < entities; n++ {
						m.New()
					}
				}
			})
			t.Run(fmt.Sprintf("1-create-then-1-invalidate/kind=%s/size=%d", name, entities), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					m := managerCreator()
					for n := 0; n < entities; n++ {
						e := m.New()
						m.Invalidate(e)
					}
				}
			})
			t.Run(fmt.Sprintf("multiple-create-then-invalidate-every-10/kind=%s/size=%d", name, entities), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					m := managerCreator()
					for n := 0; n < entities; n++ {
						e := m.New()
						if n%10 == 0 {
							m.Invalidate(e)
						}
					}

				}
			})
		}
	}
}

package entity

import (
	"testing"
)

func runEntityBenchmarkSuite(t *testing.B, managerCreator func(int) Manager) {
	t.Run("multiple create", func(b *testing.B) {
		m := managerCreator(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.New()
		}
	})
	t.Run("1 create then 1 invalidate", func(b *testing.B) {
		m := managerCreator(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			e := m.New()
			m.Invalidate(e)
		}
	})
	t.Run("multiple create then invalidate every 10", func(b *testing.B) {
		m := managerCreator(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			e := m.New()
			if b.N % 10 == 0 {
				m.Invalidate(e)
			}
		}
	})
}

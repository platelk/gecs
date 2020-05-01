package entity

import "testing"

func TestNewSliceManager(t *testing.T) {
	runEntityTestSuite(t, NewSliceManager())
}

func BenchmarkNewSliceManager(b *testing.B) {
	runEntityBenchmarkSuite(b, func(i int) Manager {
		return NewSliceManager()
	})
}
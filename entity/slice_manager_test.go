package entity

import "testing"

func TestNewSliceManager(t *testing.T) {
	runEntityTestSuite(t, NewSliceManager())
}

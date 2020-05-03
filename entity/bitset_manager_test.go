package entity

import (
	"testing"
)

func TestNewBitSetManager(t *testing.T) {
	runTestNewManager(t, NewBitSetManager())
}

package event

import (
	"fmt"
	"testing"
)

func TestNewDefaultHandler(t *testing.T) {
	var i int
	eg := func() Eventer {
		i++
		return NewEvent("player_move", fmt.Sprintf("%d", i), "test", nil)
	}
	runHandlerTestSuite(t, eg, NewDefaultHandler(eg()))
}

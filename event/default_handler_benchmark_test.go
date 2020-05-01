package event

import (
	"fmt"
	"testing"
)

func BenchmarkNewDefaultHandler(b *testing.B) {
	var i int
	eg := func() Eventer {
		i++
		return NewEvent("player_move", fmt.Sprintf("%d", i), "test", nil)
	}
	runHandlerBenchmarchSuite(b, eg, NewDefaultHandler(eg()))
}

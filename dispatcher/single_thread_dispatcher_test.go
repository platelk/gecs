package dispatcher

import (
	"testing"
)

func TestNewSingleBuilder(t *testing.T) {
	runDispatcherTestSuite(t, func() DispatchBuilder {
		return NewSingleBuilder()
	})
}

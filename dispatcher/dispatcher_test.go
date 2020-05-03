package dispatcher

import (
	"testing"
	"gecs/system"

	"github.com/stretchr/testify/require"
)

func runDispatcherTestSuite(t *testing.T, newBuilder func() DispatchBuilder) {
	t.Run("simple dispatch multiple layer without after", func(t *testing.T) {
		runSys := 0
		d := newBuilder().
			With(system.NewWithPlan("b",
				system.Plan{
					Read:  []string{"position"},
					Write: []string{"life"},
					After: nil,
				},
				func(data *system.Data) {
					require.Equal(t, 0, runSys)
					runSys++
				})).
			With(system.NewWithPlan("a", system.Plan{
				Read:  []string{"life"},
				Write: []string{"position"},
				After: nil,
			}, func(data *system.Data) {
				require.Equal(t, 1, runSys)
				runSys++
			})).
			With(system.NewWithPlan("c", system.Plan{
				Read:  []string{"life"},
				After: nil,
			}, func(data *system.Data) {
				require.Equal(t, 2, runSys)
				runSys++
			})).
			Build()
		d.Dispatch(system.NewData(nil, nil))
		require.Equal(t, 3, runSys)
	})
	t.Run("simple dispatch multiple layer with after", func(t *testing.T) {
		runSys := 0
		sysA := system.NewWithPlan("a", system.Plan{
			Read:  []string{"life"},
			Write: []string{"position"},
			After: nil,
		}, func(data *system.Data) {
			require.True(t, runSys < 2)
			runSys++
		})
		d := newBuilder().
			With(system.NewWithPlan("b", system.Plan{
				Read:  []string{"position"},
				Write: []string{"life"},
				After: []string{sysA.Type()},
			}, func(data *system.Data) {
				require.Equal(t, 2, runSys)
				runSys++
			})).
			With(sysA).
			With(system.NewWithPlan("c", system.Plan{
				Read:  []string{"life"},
				After: nil,
			}, func(data *system.Data) {
				require.True(t, runSys < 2)
				runSys++
			})).
			Build()
		d.Dispatch(system.NewData(nil, nil))
		require.Equal(t, 3, runSys)
	})
}

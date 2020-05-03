package dispatcher

import "gecs/system"

type SingleDispatcherBuilder struct {
	plan []*system.Plan
	sys  []system.System
}

func NewSingleBuilder() *SingleDispatcherBuilder {
	return &SingleDispatcherBuilder{}
}

func (s SingleDispatcherBuilder) With(sys system.System) DispatchBuilder {
	s.plan = append(s.plan, sys.Plan())
	s.sys = append(s.sys, sys)
	return s
}

func (s SingleDispatcherBuilder) Build() Dispatcher {
	return &SingleDispatcher{
		e: calculatePlan(s.plan, s.sys),
	}
}

type SingleDispatcher struct {
	e execTree
}

func (s *SingleDispatcher) Dispatch(d *system.Data) {
	for _, phase := range s.e.phases {
		for _, sys := range phase.parallel {
			sys.Run(d)
		}
	}
}

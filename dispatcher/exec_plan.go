package dispatcher

import "gecs/system"

type execTree struct {
	phases []execPhase
}

type execPhase struct {
	write    map[string]struct{}
	read     map[string]struct{}
	after    map[string]struct{}
	parallel map[string]system.System
}

func newExecPhase() execPhase {
	return execPhase{
		write:    make(map[string]struct{}),
		read:     make(map[string]struct{}),
		after:    make(map[string]struct{}),
		parallel: make(map[string]system.System),
	}
}

func calculatePlan(p []*system.Plan, s []system.System) execTree {
	et := execTree{
		phases: []execPhase{newExecPhase()},
	}

	for i := 0; i < len(p); i++ {
		plan, sys := p[i], s[i]
		insertIdx, before := insertInPhases(et, sys, plan)
		var phase execPhase
		switch {
		case before:
			phase = newExecPhase()
			et.phases = append([]execPhase{phase}, et.phases...)
		case insertIdx >= len(et.phases):
			phase = newExecPhase()
			et.phases = append(et.phases, phase)
		default:
			phase = et.phases[insertIdx]
		}
		addPlanInPhase(plan, sys, &phase)
	}

	return et
}

func insertInPhases(et execTree, sys system.System, p *system.Plan) (int, bool) {
	var before bool
	for i, phase := range et.phases {
		before = shouldBeBefore(sys, &phase)
		ok := canEnterInPhase(p, &phase)
		if ok {
			return i, before
		}
	}
	return len(et.phases), before
}

func addPlanInPhase(plan *system.Plan, sys system.System, phase *execPhase) {
	for _, c := range plan.Read {
		phase.read[c] = struct{}{}
	}
	for _, c := range plan.Write {
		phase.write[c] = struct{}{}
	}
	for _, s := range plan.After {
		phase.after[s] = struct{}{}
	}
	phase.parallel[sys.Type()] = sys
}

func shouldBeBefore(sys system.System, phase *execPhase) bool {
	for s := range phase.after {
		if sys.Type() == s {
			return true
		}
	}
	return false
}

func canEnterInPhase(plan *system.Plan, phase *execPhase) bool {
	for _, c := range plan.Write {
		if _, ok := phase.read[c]; ok {
			return false
		}
		if _, ok := phase.write[c]; ok {
			return false
		}
	}
	for _, c := range plan.Read {
		if _, ok := phase.write[c]; ok {
			return false
		}
	}
	for _, s := range plan.After {
		if _, ok := phase.parallel[s]; ok {
			return false
		}
	}

	return true
}

package system

import "gecs/component"

type Plan struct {
	Read  []string
	Write []string
	After []string
}

func (p Plan) WithRead(c component.Component) Plan {
	p.Read = append(p.Read, c.Type())
	return p
}

func (p Plan) WithWrite(c component.Component) Plan {
	p.Write = append(p.Write, c.Type())
	return p
}

func (p Plan) WithAfter(s System) Plan {
	p.Write = append(p.Write, s.Type())
	return p
}

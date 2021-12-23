package vm

import "github.com/numary/machine/core"

func (m *Machine) popValue() core.Value {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	return x
}

func (m *Machine) popAccount() core.Account {
	if a, ok := m.popValue().(core.Account); ok {
		return a
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popBool() bool {
	if b, ok := m.popValue().(core.Boolean); ok {
		return bool(b)
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popNumber() uint64 {
	if n, ok := m.popValue().(core.Number); ok {
		return uint64(n)
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popString() core.String {
	if s, ok := m.popValue().(core.String); ok {
		return s
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popMonetary() core.Monetary {
	if m, ok := m.popValue().(core.Monetary); ok {
		return m
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popAsset() core.Asset {
	if v, ok := m.popValue().(core.Asset); ok {
		return v
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popFunding() core.Funding {
	if v, ok := m.popValue().(core.Funding); ok {
		return v
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popAllotment() core.Allotment {
	if m, ok := m.popValue().(core.Allotment); ok {
		return m
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popPortion() core.Portion {
	if m, ok := m.popValue().(core.Portion); ok {
		return m
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) pushValue(v core.Value) {
	m.Stack = append(m.Stack, v)
}

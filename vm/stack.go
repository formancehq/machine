package vm

import (
	"fmt"

	"github.com/formancehq/machine/core"
)

func (m *Machine) popAccount() core.Account {
	x := m.popValue()
	if v, ok := x.(core.Account); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popNumber() core.Number {
	x := m.popValue()
	if v, ok := x.(core.Number); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popString() core.String {
	x := m.popValue()
	if v, ok := x.(core.String); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popMonetary() core.Monetary {
	x := m.popValue()
	if v, ok := x.(core.Monetary); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popAsset() core.Asset {
	x := m.popValue()
	if v, ok := x.(core.Asset); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popFunding() core.Funding {
	x := m.popValue()
	if v, ok := x.(core.Funding); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popAllotment() core.Allotment {
	x := m.popValue()
	if v, ok := x.(core.Allotment); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popPortion() core.Portion {
	x := m.popValue()
	if v, ok := x.(core.Portion); ok {
		return v
	}
	panic(fmt.Errorf("unexpected type '%T' on stack", x))
}

func (m *Machine) popValue() core.Value {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	return x
}

func (m *Machine) pushValue(v core.Value) {
	m.Stack = append(m.Stack, v)
}

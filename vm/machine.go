package vm

import (
	"encoding/binary"
	"fmt"

	ledger "github.com/numary/ledger/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

const (
	EXIT_OK = byte(iota + 1)
	EXIT_FAIL
)

func StdOutPrinter(c chan core.Value) {
	for v := range c {
		fmt.Println("OUT:", v)
	}
}

func NewMachine(p *program.Program) *Machine {
	printc := make(chan core.Value)

	m := Machine{
		Program:    p,
		print_chan: printc,
		Printer:    StdOutPrinter,
	}

	return &m
}

type Machine struct {
	P          uint
	Program    *program.Program
	Constants  []core.Value // Constants and Variables constitute the resources
	Variables  []core.Value
	Stack      []core.Value
	Printer    func(chan core.Value)
	Postings   []ledger.Posting
	print_chan chan core.Value
}

func (m *Machine) popValue() core.Value {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	return x
}

func (m *Machine) popAccount() core.Account {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	if a, ok := x.(core.Account); ok {
		return a
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popAsset() core.Asset {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	if a, ok := x.(core.Asset); ok {
		return a
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popNumber() uint64 {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	if n, ok := x.(core.Number); ok {
		return uint64(n)
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) popMonetary() core.Monetary {
	l := len(m.Stack)
	x := m.Stack[l-1]
	m.Stack = m.Stack[:l-1]
	if m, ok := x.(core.Monetary); ok {
		return m
	} else {
		panic("unexpected type on stack")
	}
}

func (m *Machine) pushNumber(x uint64) {
	num := core.Number(x)
	m.Stack = append(m.Stack, num)
}

func (m *Machine) getResource(addr program.Address) core.Value {
	a := uint16(addr)
	if a < (1 << 15) {
		return m.Constants[a]
	} else {
		a -= (1 << 15)
		return m.Variables[a]
	}
}

func (m *Machine) Tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	switch op {
	case program.OP_APUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+3]
		v := m.getResource(program.Address(binary.LittleEndian.Uint16(bytes)))
		m.Stack = append(m.Stack, v)
		m.P += 2
	case program.OP_IPUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+9]
		v := core.Number(binary.LittleEndian.Uint64(bytes))
		m.Stack = append(m.Stack, v)
		m.P += 8
	case program.OP_IADD:
		b := m.popNumber()
		a := m.popNumber()
		m.pushNumber(a + b)
	case program.OP_ISUB:
		b := m.popNumber()
		a := m.popNumber()
		m.pushNumber(a - b)
	case program.OP_PRINT:
		a := m.popValue()
		m.print_chan <- a
	case program.OP_FAIL:
		// fmt.Println("program failed")
		// fmt.Println("stack: ", m.Stack)
		return true, EXIT_FAIL
	case program.OP_SEND:
		dst := m.popAccount()
		src := m.popAccount()
		mon := m.popMonetary()
		p := ledger.Posting{
			Source:      string(src),
			Destination: string(dst),
			Asset:       string(mon.Asset),
			Amount:      int64(mon.Amount),
		}
		m.Postings = append(m.Postings, p)
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		// fmt.Println("end of program")
		// fmt.Println("stack: ", m.Stack)

		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) Execute(vars map[string]core.Value) byte {
	go m.Printer(m.print_chan)
	defer close(m.print_chan)

	m.Constants = m.Program.Constants
	m.Variables = []core.Value{}
	for _, name := range m.Program.Variables {
		if val, ok := vars[name]; ok {
			m.Variables = append(m.Variables, val)
		} else {
			return EXIT_FAIL
		}
	}

	for {
		finished, exit_code := m.Tick()
		if finished {
			return exit_code
		}
	}
}

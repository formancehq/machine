package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/numary/ledger/core"
	"github.com/numary/machine/vm/program"
)

const (
	EXIT_OK = byte(iota + 1)
	EXIT_FAIL
)

func StdOutPrinter(c chan uint64) {
	for v := range c {
		fmt.Println("OUT:", v)
	}
}

func NewMachine(p *program.Program) *Machine {
	printc := make(chan uint64)

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
	Constants  []string
	Stack      []byte
	Printer    func(chan uint64)
	Postings   []core.Posting
	print_chan chan uint64
}

func (m *Machine) popUint16() uint16 {
	l := len(m.Stack)
	bytes := m.Stack[l-2:]
	m.Stack = m.Stack[:l-2]
	return binary.LittleEndian.Uint16(bytes)
}

func (m *Machine) popUint64() uint64 {
	l := len(m.Stack)
	bytes := m.Stack[l-8:]
	m.Stack = m.Stack[:l-8]
	return binary.LittleEndian.Uint64(bytes)
}

func (m *Machine) pushUint64(v uint64) {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, v)
	m.Stack = append(m.Stack, bytes...)
}

func (m *Machine) Tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	switch op {
	case program.OP_PUSH2:
		m.Stack = append(m.Stack, m.Program.Instructions[m.P+1:m.P+3]...)
		m.P += 2
	case program.OP_PUSH8:
		m.Stack = append(m.Stack, m.Program.Instructions[m.P+1:m.P+9]...)
		m.P += 8
	case program.OP_IADD:
		b := m.popUint64()
		a := m.popUint64()
		m.pushUint64(a + b)
	case program.OP_ISUB:
		b := m.popUint64()
		a := m.popUint64()
		m.pushUint64(a - b)
	case program.OP_PRINT:
		a := m.popUint64()
		m.print_chan <- a
	case program.OP_FAIL:
		fmt.Println("program failed")
		fmt.Println("stack: ", m.Stack)
		return true, EXIT_FAIL
	case program.OP_SEND:
		dst := m.popUint16()
		src := m.popUint16()
		amount := m.popUint64()
		asset := m.popUint16()
		p := core.Posting{
			Source:      m.Constants[src],
			Destination: m.Constants[dst],
			Asset:       m.Constants[asset],
			Amount:      int64(amount),
		}
		m.Postings = append(m.Postings, p)
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		fmt.Println("end of program")
		fmt.Println("stack: ", m.Stack)

		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) Execute(vars map[string]string) byte {
	m.Constants = m.Program.Constants

	go m.Printer(m.print_chan)

	defer close(m.print_chan)

	for {
		finished, exit_code := m.Tick()
		if finished {
			return exit_code
		}
	}
}

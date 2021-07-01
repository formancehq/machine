package vm

import (
	"fmt"

	"github.com/numary/machine/vm/program"
)

func StdOutPrinter(c chan byte) {
	for v := range c {
		fmt.Println(v)
	}
}

func NewMachine(p program.Program) *Machine {
	out := make(chan byte)

	m := Machine{
		Program: p,
		outchan: out,
		Printer: StdOutPrinter,
	}

	return &m
}

type Machine struct {
	P       uint
	Program program.Program
	Stack   []byte
	Printer func(chan byte)
	outchan chan byte
}

func (m *Machine) Tick() bool {
	op := m.Program[m.P]

	switch op {
	case program.OP_IPUSH:
		m.P += 1
		m.Stack = append(m.Stack, m.Program[m.P])
	case program.OP_IADD:
		l := len(m.Stack)
		a := m.Stack[l-2]
		b := m.Stack[l-1]
		m.Stack = m.Stack[:l-2]
		m.Stack = append(m.Stack, a+b)
	case program.OP_ISUB:
		l := len(m.Stack)
		a := m.Stack[l-2]
		b := m.Stack[l-1]
		m.Stack = m.Stack[:l-2]
		m.Stack = append(m.Stack, a-b)
	case program.OP_PRINT:
		l := len(m.Stack)
		a := m.Stack[l-1]
		m.Stack = m.Stack[:l-1]
		m.outchan <- a
	}

	m.P += 1

	if int(m.P) >= len(m.Program) {
		fmt.Println("end of program")
		fmt.Println("stack: ", m.Stack)

		return false
	}

	return true
}

func (m *Machine) Execute() {
	go m.Printer(m.outchan)

	for {
		if !m.Tick() {
			break
		}
	}

	close(m.outchan)
}

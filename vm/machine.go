package vm

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"

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

func (m *Machine) getResource(addr core.Address) core.Value {
	a := uint16(addr)
	if a < (1 << 15) {
		return m.Constants[a]
	} else {
		a -= (1 << 15)
		return m.Variables[a]
	}
}

func (m *Machine) tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	switch op {
	case program.OP_APUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+3]
		v := m.getResource(core.Address(binary.LittleEndian.Uint16(bytes)))
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
		m.pushValue(core.Number(a + b))
	case program.OP_ISUB:
		b := m.popNumber()
		a := m.popNumber()
		m.pushValue(core.Number(a - b))
	case program.OP_PRINT:
		a := m.popValue()
		m.print_chan <- a
	case program.OP_FAIL:
		return true, EXIT_FAIL
	case program.OP_MAKEALLOC:
		ndest := m.popNumber()
		alloc := make([]core.AllocPart, ndest)
		for i := int(ndest - 1); i >= 0; i-- {
			acc := m.popAccount()
			b := int64(m.popNumber())
			a := int64(m.popNumber())
			alloc[i] = core.AllocPart{
				Ratio: big.NewRat(a, b),
				Dest:  acc,
			}
		}
		m.pushValue(core.Allocation(alloc))
	case program.OP_SEND:
		dest := m.popValue()
		src := m.popAccount()
		mon := m.popMonetary()
		if allocs, ok := dest.(core.Allocation); ok {
			posting_amounts := []int64{}
			// for every allocation, compute the floored amount and the total
			total := int64(0)
			for _, alloc := range allocs {
				var res big.Int
				res.Mul(alloc.Ratio.Num(), big.NewInt(int64(mon.Amount)))
				res.Div(&res, alloc.Ratio.Denom())
				posting_amounts = append(posting_amounts, res.Int64())
				total += res.Int64()
			}
			// generate postings
			for i := 0; i < len(allocs); i++ {
				// if the total is less than the target amount, add 1 unit to amounts until it isn't
				amt := posting_amounts[i]
				dest := allocs[i].Dest
				if total < int64(mon.Amount) {
					amt += 1
					total += 1
				}
				m.Postings = append(m.Postings, ledger.Posting{
					Source:      string(src),
					Destination: string(dest),
					Asset:       string(mon.Asset),
					Amount:      int64(amt),
				})
			}
		} else if dest, ok := dest.(core.Account); ok {
			m.Postings = append(m.Postings, ledger.Posting{
				Source:      string(src),
				Destination: string(dest),
				Asset:       string(mon.Asset),
				Amount:      int64(mon.Amount),
			})
		}
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) execute(vars []core.Value) byte {
	go m.Printer(m.print_chan)
	defer close(m.print_chan)

	m.Constants = m.Program.Constants
	m.Variables = vars

	for {
		finished, exit_code := m.tick()
		if finished {
			return exit_code
		}
	}
}

func (m *Machine) Execute(vars map[string]core.Value) (byte, error) {
	v, err := m.Program.ParseVariables(vars)
	if err != nil {
		return 0, err
	}
	return m.execute(v), nil
}

func (m *Machine) ExecuteFromJSON(vars map[string]json.RawMessage) (byte, error) {
	v, err := m.Program.ParseVariablesJSON(vars)
	if err != nil {
		return 0, err
	}
	return m.execute(v), nil
}

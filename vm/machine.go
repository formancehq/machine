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
	Balances   map[string]map[string]uint64
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
	case program.OP_SOURCE:
		mon := m.popMonetary()
		n := m.popNumber()
		sources := []core.Account{}
		for i := uint64(0); i < n; i++ {
			sources = append(sources, m.popAccount())
		}
		asset := mon.Asset
		target := mon.Amount

		var n_actual_src uint64
		for _, src := range sources {
			src_funds := m.Balances[string(src)][asset]
			var amt_to_withdraw uint64
			if src_funds > target {
				amt_to_withdraw = target
			} else {
				amt_to_withdraw = src_funds
			}
			m.Balances[string(src)][asset] -= amt_to_withdraw
			target -= amt_to_withdraw
			m.pushValue(src)
			m.pushValue(core.Monetary{
				Asset:  asset,
				Amount: amt_to_withdraw,
			})
			n_actual_src++
		}
		m.pushValue(core.Number(n_actual_src))
	case program.OP_ALLOC:
		allotment := m.popAllotment()
		n := m.popNumber()
		source_accounts := make([]core.Account, n)
		source_amounts := make([]uint64, n)
		total_src := uint64(0)
		var asset *string
		// extract accounts from stack while checking the assets correspond
		for i := uint64(0); i < n; i++ {
			source_accounts[i] = m.popAccount()
			mon := m.popMonetary()
			source_amounts[i] = mon.Amount
			// check that the assets correspond
			if asset == nil {
				asset = &mon.Asset
			} else if *asset != mon.Asset {
				return true, EXIT_FAIL
			}
			total_src += mon.Amount
		}
		parts := []uint64{}
		total_allocated := uint64(0)
		// for every part in the allotment, calculate the floored value
		for _, part := range allotment {
			var res big.Int
			res.Mul(part.Num(), new(big.Int).SetUint64(total_src))
			res.Div(&res, part.Denom())
			parts = append(parts, res.Uint64())
			total_allocated += res.Uint64()
		}
		// for every part in the floored values, fetch them from the sources
		for _, part := range parts {
			// if the total allocated is less than the target amount, add 1 unit until it isn't
			if total_allocated < uint64(total_src) {
				part += 1
				total_allocated += 1
			}
			n := 0 // number of sources needed to fill this part
			for i, acc := range source_accounts {
				amt := source_amounts[i] // amount to withdraw from the account
				if source_amounts[i] > part {
					amt = part
				}
				part -= amt
				m.pushValue(core.Monetary{Asset: *asset, Amount: amt})
				m.pushValue(acc)
				n += 1
			}
			m.pushValue(core.Number(n))
		}
	case program.OP_SEND:
		dest := m.popAccount()
		n := m.popNumber()
		for i := uint64(0); i < n; i++ {
			src := m.popAccount()
			mon := m.popMonetary()
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

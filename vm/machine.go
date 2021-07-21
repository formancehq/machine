package vm

import (
	"encoding/binary"
	"encoding/json"
	"errors"
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
		Constants:  p.Constants,
		print_chan: printc,
		Printer:    StdOutPrinter,
	}

	return &m
}

type Machine struct {
	P          uint
	Program    *program.Program
	Constants  []core.Value // Constants and Variables
	Variables  []core.Value // constitute the resources
	Stack      []core.Value
	Postings   []ledger.Posting             // accumulates postings throughout execution
	Balances   map[string]map[string]uint64 // keeps tracks of balances througout execution
	Printer    func(chan core.Value)
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

func (m *Machine) getBalance(account core.Account, asset core.Asset) (uint64, error) {
	if account == "world" {
		return 0, errors.New("tried to use balance of world")
	}
	if acc_balance, ok := m.Balances[string(account)]; ok {
		if balance, ok := acc_balance[string(asset)]; ok {
			return balance, nil
		} else {
			return 0, fmt.Errorf("missing %v balance of %v", asset, account)
		}
	} else {
		return 0, fmt.Errorf("missing balance of %v", account)
	}
}

func (m *Machine) withdraw(account core.Account, asset core.Asset, amount uint64) bool {
	if account == "world" {
		return true
	}
	withdraw_ok := false
	if acc_balance, ok := m.Balances[string(account)]; ok {
		if balance, ok := acc_balance[string(asset)]; ok {
			if balance >= amount {
				acc_balance[string(asset)] -= amount
				withdraw_ok = true
			}
		}
	}
	return withdraw_ok
}

func (m *Machine) credit(account core.Account, asset core.Asset, amount uint64) {
	if acc_balance, ok := m.Balances[string(account)]; ok {
		if _, ok := acc_balance[string(asset)]; ok {
			acc_balance[string(asset)] += amount
		} else {
			acc_balance[string(asset)] = amount
		}
	} else {
		m.Balances[string(account)] = map[string]uint64{
			string(asset): amount,
		}
	}
}

func (m *Machine) resolveMonetary(account core.Account, mon core.Monetary) (uint64, error) {
	var res uint64
	if mon.Amount.All {
		a, err := m.getBalance(account, mon.Asset)
		if err != nil {
			return 0, err
		}
		res = a
	} else {
		res = mon.Amount.Specific
	}
	return res, nil
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
		n := m.popNumber()
		sources := make([]core.Account, n)
		for i := uint64(0); i < n; i++ {
			sources[i] = m.popAccount()
		}
		mon := m.popMonetary()
		asset := mon.Asset
		target := mon.Amount

		type part struct {
			mon core.Monetary
			acc core.Account
		}
		result := []part{}

		var n_actual_src uint64
		for _, src := range sources {
			src_funds := m.Balances[string(src)][string(asset)]
			var amt_to_withdraw uint64
			if target.All {
				if src == "world" {
					return true, EXIT_FAIL
				}
				amt_to_withdraw = src_funds
			} else {
				if target.Specific == 0 {
					break
				}
				if src_funds > target.Specific || src == "world" {
					amt_to_withdraw = target.Specific
				} else {
					amt_to_withdraw = src_funds
				}
				target.Specific -= amt_to_withdraw
			}
			result = append(result, part{
				mon: core.Monetary{
					Asset:  asset,
					Amount: core.NewAmountSpecific(amt_to_withdraw),
				},
				acc: src,
			})
			n_actual_src++
		}
		if !target.All && target.Specific != 0 {
			return true, EXIT_FAIL
		}
		for i := len(result) - 1; i >= 0; i-- {
			m.pushValue(result[i].mon)
			m.pushValue(result[i].acc)
		}
		m.pushValue(core.Number(n_actual_src))
	case program.OP_MAKE_ALLOTMENT:
		n := m.popNumber()
		portions := make([]core.Portion, n)
		for i := uint64(0); i < n; i++ {
			p := m.popPortion()
			portions[i] = p
		}
		allotment, err := core.NewAllotment(portions)
		if err != nil {
			return true, EXIT_FAIL
		}
		m.pushValue(*allotment)

	case program.OP_ALLOC:
		allotment := m.popAllotment()
		nparts := len(allotment)
		nsources := m.popNumber()
		source_accounts := make([]core.Account, nsources)
		source_amounts := make([]uint64, nsources)
		total_src := uint64(0)
		var asset *core.Asset
		// extract accounts from stack while checking the assets correspond
		for i := uint64(0); i < nsources; i++ {
			source_accounts[i] = m.popAccount()
			mon := m.popMonetary()
			amt, err := m.resolveMonetary(source_accounts[i], mon)
			if err != nil {
				return true, EXIT_FAIL
			}
			source_amounts[i] = amt
			// check that the assets correspond
			if asset == nil {
				asset = &mon.Asset
			} else if *asset != mon.Asset {
				return true, EXIT_FAIL
			}
			total_src += amt
		}
		parts := make([]uint64, nparts)
		total_allocated := uint64(0)
		// for every part in the allotment, calculate the floored value
		for i, allot := range allotment {
			var res big.Int
			res.Mul(new(big.Int).SetUint64(total_src), allot.Num())
			res.Div(&res, allot.Denom())
			parts[i] = res.Uint64()
			total_allocated += res.Uint64()
		}

		type subpart struct {
			mon core.Monetary
			acc core.Account
		}
		result := make([][]subpart, nparts)

		// for every part in the floored values, fetch them from the sources
		first_non_empty_idx := 0
		for ipart, part := range parts {
			// if the total allocated is less than the target amount, add 1 unit until it isn't
			if total_allocated < uint64(total_src) {
				part += 1
				total_allocated += 1
			}

			result[ipart] = []subpart{}
			n := 0 // number of sources needed to fill this part
			// start at the first non empty source
			for i := first_non_empty_idx; i < len(source_accounts); i++ {
				if part == 0 { // if we finished filling this part
					break
				}
				amt := source_amounts[i] // amount to withdraw from the account
				if amt > part {
					// if we have more than enough to fill, don't give too much to this part
					amt = part
				} else { // if we had to empty the source
					first_non_empty_idx++
				}
				part -= amt
				source_amounts[i] -= amt
				result[ipart] = append(result[ipart], subpart{
					mon: core.Monetary{Asset: *asset, Amount: core.NewAmountSpecific(amt)},
					acc: source_accounts[i],
				})
				n += 1
			}
		}
		for i := len(result) - 1; i >= 0; i-- {
			part := result[i]
			for j := len(part) - 1; j >= 0; j-- {
				subpart := part[j]
				m.pushValue(subpart.mon)
				m.pushValue(subpart.acc)
			}
			m.pushValue(core.Number(len(part)))
		}
	case program.OP_SEND:
		dest := m.popAccount()
		n := m.popNumber()
		for i := uint64(0); i < n; i++ {
			src := m.popAccount()
			mon := m.popMonetary()
			amt, err := m.resolveMonetary(src, mon)
			if err != nil {
				return true, EXIT_FAIL
			}
			if amt == 0 {
				continue
			}
			// verify and withdraw funds
			if ok := m.withdraw(src, mon.Asset, amt); !ok {
				return true, EXIT_FAIL
			}
			m.credit(dest, mon.Asset, amt)
			m.Postings = append(m.Postings, ledger.Posting{
				Source:      string(src),
				Destination: string(dest),
				Asset:       string(mon.Asset),
				Amount:      int64(amt),
			})
		}
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) Execute() (byte, error) {
	go m.Printer(m.print_chan)
	defer close(m.print_chan)

	if m.Variables == nil {
		return 0, errors.New("variables haven't been initialized")
	} else if m.Balances == nil {
		return 0, errors.New("balances haven't been initialized")
	}

	for {
		finished, exit_code := m.tick()
		if finished {
			return exit_code, nil
		}
	}
}

func (m *Machine) GetNeededBalances() (map[string]map[string]struct{}, error) {
	needed := map[string]map[string]struct{}{}
	for addr, needed_assets := range m.Program.NeededBalances {
		account := m.getResource(addr)
		if account, ok := account.(core.Account); ok {
			if string(account) == "world" {
				continue
			}
			needed[string(account)] = map[string]struct{}{}
			for addr := range needed_assets {
				mon := m.getResource(addr)
				if mon, ok := mon.(core.Monetary); ok {
					needed[string(account)][string(mon.Asset)] = struct{}{}
				} else {
					return nil, errors.New("incorrect program")
				}
			}
		} else {
			return nil, errors.New("incorrect program")
		}
	}
	return needed, nil
}

func (m *Machine) SetBalances(balances map[string]map[string]uint64) error {
	// for every account that we need balances of, check if it's there
	for addr, needed_assets := range m.Program.NeededBalances {
		account := m.getResource(addr)
		if account, ok := account.(core.Account); ok {
			if string(account) == "world" {
				continue
			}
			if b, ok := balances[string(account)]; ok {
				// for every asset that we need balances of on that account
				for addr := range needed_assets {
					mon := m.getResource(addr)
					if mon, ok := mon.(core.Monetary); ok {
						if _, ok := b[string(mon.Asset)]; !ok {
							return fmt.Errorf("missing %v balance of %v", mon.Asset, account)
						}
					} else {
						return errors.New("incorrect program")
					}
				}
			} else {
				return fmt.Errorf("missing balances of %v", account)
			}
		} else {
			return errors.New("incorrect program")
		}
	}
	m.Balances = balances
	return nil
}

func (m *Machine) SetVars(vars map[string]core.Value) error {
	v, err := m.Program.ParseVariables(vars)
	if err != nil {
		return err
	}
	m.Variables = v
	return nil
}

func (m *Machine) SetVarsFromJSON(vars map[string]json.RawMessage) error {
	v, err := m.Program.ParseVariablesJSON(vars)
	if err != nil {
		return err
	}
	m.Variables = v
	return nil
}

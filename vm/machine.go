/*
Provides `Machine`, which executes programs and outputs postings.
1: Create New Machine
2: Set Variables (with `core.Value`s or JSON)
3: Resolve Resources (answer requests on channel)
4: Resolve Balances (answer requests on channel)
6: Execute
*/
package vm

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/logrusorgru/aurora"
	ledger "github.com/numary/ledger/pkg/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

const (
	EXIT_OK = byte(iota + 1)
	EXIT_FAIL
	EXIT_FAIL_INVALID
	EXIT_FAIL_INSUFFICIENT_FUNDS
)

func StdOutPrinter(c chan core.Value) {
	for v := range c {
		fmt.Println("OUT:", v)
	}
}

func NewMachine(p program.Program) *Machine {
	printc := make(chan core.Value)

	m := Machine{
		Program:             p,
		UnresolvedResources: p.Resources,
		Resources:           make([]core.Value, 0),
		print_chan:          printc,
		Printer:             StdOutPrinter,
		Postings:            make([]ledger.Posting, 0),
		TxMeta:              map[string]core.Value{},
	}

	return &m
}

type Machine struct {
	P                   uint
	Program             program.Program
	Vars                map[string]core.Value
	UnresolvedResources []program.Resource
	Resources           []core.Value // Constants and Variables
	resolve_called      bool
	Balances            map[core.Account]map[core.Asset]ledger.MonetaryInt // keeps tracks of balances througout execution
	set_balance_called  bool
	Stack               []core.Value
	Postings            []ledger.Posting      // accumulates postings throughout execution
	TxMeta              map[string]core.Value // accumulates transaction meta throughout execution
	Printer             func(chan core.Value)
	print_chan          chan core.Value
	Debug               bool
}

func (m *Machine) GetTxMetaJson() ledger.Metadata {
	meta := make(ledger.Metadata)
	for k, v := range m.TxMeta {
		val_json, _ := json.Marshal(v)
		v, _ := json.Marshal(core.ValueJSON{
			Type:  v.GetType().String(),
			Value: val_json,
		})
		meta[k] = v
	}
	return meta
}

func (m *Machine) getResource(addr core.Address) (*core.Value, bool) {
	a := int(addr)
	if a >= len(m.Resources) {
		return nil, false
	}
	return &m.Resources[a], true
}

func (m *Machine) withdrawAll(account core.Account, asset core.Asset, overdraft ledger.MonetaryInt) (*core.Funding, error) {
	if acc_balance, ok := m.Balances[account]; ok {
		if balance, ok := acc_balance[asset]; ok {
			amount_taken := *ledger.NewMonetaryInt(0)
			if balance.Add(&overdraft).Gt(ledger.NewMonetaryInt(0)) {
				amount_taken = *balance.Add(&overdraft)
				acc_balance[asset] = *overdraft.Neg()
			}

			return &core.Funding{
				Asset: asset,
				Parts: []core.FundingPart{{
					Account: account,
					Amount:  amount_taken,
				}},
			}, nil
		}
	}
	return nil, fmt.Errorf("missing %v balance from %v", asset, account)
}

func (m *Machine) withdrawAlways(account core.Account, mon core.Monetary) (*core.Funding, error) {
	if acc_balance, ok := m.Balances[account]; ok {
		if balance, ok := acc_balance[mon.Asset]; ok {
			acc_balance[mon.Asset] = *balance.Sub(&mon.Amount)
			return &core.Funding{
				Asset: mon.Asset,
				Parts: []core.FundingPart{{
					Account: account,
					Amount:  mon.Amount,
				}},
			}, nil
		}
	}
	return nil, fmt.Errorf("missing %v balance from %v", mon.Asset, account)
}

func (m *Machine) credit(account core.Account, funding core.Funding) {
	if account == "world" {
		return
	}
	if acc_balance, ok := m.Balances[account]; ok {
		if _, ok := acc_balance[funding.Asset]; ok {
			for _, part := range funding.Parts {
				balance := acc_balance[funding.Asset]
				acc_balance[funding.Asset] = *balance.Add(&part.Amount)
			}
		}
	}
}

func (m *Machine) repay(funding core.Funding) {
	for _, part := range funding.Parts {
		if part.Account == "world" {
			continue
		}
		balance := m.Balances[part.Account][funding.Asset]
		m.Balances[part.Account][funding.Asset] = *balance.Add(&part.Amount)
	}
}

func (m *Machine) tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	if m.Debug {
		fmt.Println("STATE ---------------------------------------------------------------------")
		fmt.Printf("    %v\n", aurora.Blue(m.Stack))
		fmt.Printf("    %v\n", aurora.Cyan(m.Balances))
		fmt.Printf("    %v\n", program.OpcodeName(op))
	}

	switch op {
	case program.OP_APUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+3]
		v, ok := m.getResource(core.Address(binary.LittleEndian.Uint16(bytes)))
		if !ok {
			return true, EXIT_FAIL
		}
		m.Stack = append(m.Stack, *v)
		m.P += 2
	case program.OP_IPUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+9]
		v := core.Number(*big.NewInt(int64(binary.LittleEndian.Uint64(bytes))))
		m.Stack = append(m.Stack, v)
		m.P += 8
	case program.OP_BUMP:
		n := big.Int(m.popNumber())
		idx := len(m.Stack) - int(n.Uint64()) - 1
		v := m.Stack[idx]
		m.Stack = append(m.Stack[:idx], m.Stack[idx+1:]...)
		m.Stack = append(m.Stack, v)
	case program.OP_DELETE:
		n := m.popValue()
		if n.GetType() == core.TYPE_FUNDING {
			return true, EXIT_FAIL_INVALID
		}
	case program.OP_IADD:
		b := m.popNumber()
		a := m.popNumber()

		m.pushValue(*a.Add(&b))
	case program.OP_ISUB:
		b := m.popNumber()
		a := m.popNumber()
		m.pushValue(*a.Sub(&b))
	case program.OP_PRINT:
		a := m.popValue()
		m.print_chan <- a
	case program.OP_FAIL:
		return true, EXIT_FAIL
	case program.OP_ASSET:
		v := m.popValue()
		switch v := v.(type) {
		case core.Asset:
			m.pushValue(v)
		case core.Monetary:
			m.pushValue(v.Asset)
		case core.Funding:
			m.pushValue(v.Asset)
		default:
			return true, EXIT_FAIL_INVALID
		}

	case program.OP_MONETARY_NEW:
		amount := m.popNumber()
		asset := m.popAsset()
		m.pushValue(core.Monetary{
			Asset:  asset,
			Amount: ledger.MonetaryInt(amount),
		})

	case program.OP_MONETARY_ADD:
		b := m.popMonetary()
		a := m.popMonetary()
		if a.Asset != b.Asset {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(core.Monetary{
			Asset:  a.Asset,
			Amount: *a.Amount.Add(&b.Amount),
		})

	case program.OP_MAKE_ALLOTMENT:
		n := m.popNumber()
		portions := make([]core.Portion, n.Uint64())
		for i := uint64(0); i < n.Uint64(); i++ {
			p := m.popPortion()
			portions[i] = p
		}
		allotment, err := core.NewAllotment(portions)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*allotment)
	case program.OP_TAKE_ALL:
		overdraft := m.popMonetary()
		account := m.popAccount()
		funding, err := m.withdrawAll(account, overdraft.Asset, overdraft.Amount)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*funding)

	case program.OP_TAKE_ALWAYS:
		mon := m.popMonetary()
		account := m.popAccount()
		funding, err := m.withdrawAlways(account, mon)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*funding)

	case program.OP_TAKE:
		mon := m.popMonetary()
		funding := m.popFunding()
		if funding.Asset != mon.Asset {
			return true, EXIT_FAIL_INVALID
		}
		result, remainder, err := funding.Take(mon.Amount)
		if err != nil {
			return true, EXIT_FAIL_INSUFFICIENT_FUNDS
		}
		m.pushValue(remainder)
		m.pushValue(result)

	case program.OP_TAKE_MAX:
		mon := m.popMonetary()
		funding := m.popFunding()
		if funding.Asset != mon.Asset {
			return true, EXIT_FAIL_INVALID
		}
		missing := *ledger.NewMonetaryInt(0)
		total := funding.Total()
		if mon.Amount.Gt(&total) {
			missing = *mon.Amount.Sub(&total)
		}
		result, remainder := funding.TakeMax(mon.Amount)
		m.pushValue(core.Monetary{
			Asset:  mon.Asset,
			Amount: missing,
		})
		m.pushValue(remainder)
		m.pushValue(result)

	case program.OP_FUNDING_ASSEMBLE:
		num := m.popNumber()
		n := int(num.Uint64())
		if n == 0 {
			return true, EXIT_FAIL_INVALID
		}
		first := m.popFunding()
		result := core.Funding{
			Asset: first.Asset,
		}
		fundings_rev := make([]core.Funding, n)
		fundings_rev[0] = first
		for i := 1; i < n; i++ {
			f := m.popFunding()
			if f.Asset != result.Asset {
				return true, EXIT_FAIL_INVALID
			}
			fundings_rev[i] = f
		}
		for i := 0; i < n; i++ {
			res, err := result.Concat(fundings_rev[n-1-i])
			if err != nil {
				return true, EXIT_FAIL_INVALID
			}
			result = res
		}
		m.pushValue(result)

	case program.OP_FUNDING_SUM:
		funding := m.popFunding()
		sum := funding.Total()
		m.pushValue(funding)
		m.pushValue(core.Monetary{
			Asset:  funding.Asset,
			Amount: sum,
		})

	case program.OP_FUNDING_REVERSE:
		funding := m.popFunding()
		result := funding.Reverse()
		m.pushValue(result)

	case program.OP_ALLOC:
		allotment := m.popAllotment()
		monetary := m.popMonetary()
		total := monetary.Amount
		parts := allotment.Allocate(total)
		for i := len(parts) - 1; i >= 0; i-- {
			m.pushValue(core.Monetary{
				Asset:  monetary.Asset,
				Amount: parts[i],
			})
		}

	case program.OP_REPAY:
		m.repay(m.popFunding())

	case program.OP_SEND:
		dest := m.popAccount()
		funding := m.popFunding()
		m.credit(dest, funding)
		for _, part := range funding.Parts {
			src := part.Account
			amt := part.Amount
			if amt.Eq(ledger.NewMonetaryInt(0)) {
				continue
			}
			m.Postings = append(m.Postings, ledger.Posting{
				Source:      string(src),
				Destination: string(dest),
				Asset:       string(funding.Asset),
				Amount:      &amt,
			})
		}
	case program.OP_TX_META:
		k := m.popString()
		v := m.popValue()
		m.TxMeta[string(k)] = v

	default:
		return true, EXIT_FAIL_INVALID
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

	if len(m.Resources) != len(m.UnresolvedResources) {
		return 0, errors.New("resources haven't been initialized")
	} else if m.Balances == nil {
		return 0, errors.New("balances haven't been initialized")
	}

	for {
		finished, exit_code := m.tick()
		if finished {
			if exit_code == EXIT_OK && len(m.Stack) != 0 {
				return EXIT_FAIL_INVALID, nil
			} else {
				return exit_code, nil
			}
		}
	}
}

type BalanceRequest struct {
	Account  string
	Asset    string
	Response chan ledger.MonetaryInt
	Error    error
}

func (m *Machine) ResolveBalances() (chan BalanceRequest, error) {
	if len(m.Resources) != len(m.UnresolvedResources) {
		return nil, errors.New("tried to resolve balances before resources")
	}
	if m.set_balance_called {
		return nil, errors.New("tried to call ResolveBalances twice")
	}
	m.set_balance_called = true
	ch := make(chan BalanceRequest)
	go func() {
		defer close(ch)
		m.Balances = make(map[core.Account]map[core.Asset]ledger.MonetaryInt)
		// for every account that we need balances of, check if it's there
		for addr, needed_assets := range m.Program.NeededBalances {
			account, ok := m.getResource(addr)
			if !ok {
				ch <- BalanceRequest{
					Error: errors.New("invalid program (resolve balances: invalid address of account)"),
				}
				return
			}
			if account, ok := (*account).(core.Account); ok {
				m.Balances[account] = make(map[core.Asset]ledger.MonetaryInt)
				// for every asset, send request
				for addr := range needed_assets {
					mon, ok := m.getResource(addr)
					if !ok {
						ch <- BalanceRequest{
							Error: errors.New("invalid program (resolve balances: invalid address of monetary)"),
						}
						return
					}
					if ha, ok := (*mon).(core.HasAsset); ok {
						asset := ha.GetAsset()
						if string(account) == "world" {
							m.Balances[account][asset] = *ledger.NewMonetaryInt(0)
							continue
						}
						resp_ch := make(chan ledger.MonetaryInt)
						ch <- BalanceRequest{
							Account:  string(account),
							Asset:    string(asset),
							Response: resp_ch,
						}
						resp, ok := <-resp_ch
						close(resp_ch)
						if !ok {
							ch <- BalanceRequest{
								Error: errors.New("error on response channel"),
							}
							return
						}
						m.Balances[account][asset] = resp
					} else {
						ch <- BalanceRequest{
							Error: errors.New("invalid program (resolve balances: not an asset)"),
						}
						return
					}
				}
			} else {
				ch <- BalanceRequest{
					Error: errors.New("incorrect program (resolve balances: not an account)"),
				}
				return
			}
		}
	}()
	return ch, nil
}

type MetadataRequest struct {
	Account  string
	Key      string
	Response chan core.Value
	Error    error
}

func (m *Machine) ResolveResources() (chan MetadataRequest, error) {
	if m.resolve_called {
		return nil, errors.New("tried to call ResolveResources twice")
	}
	m.resolve_called = true
	ch := make(chan MetadataRequest)
	go func() {
		defer close(ch)
		for len(m.Resources) != len(m.UnresolvedResources) {
			idx := len(m.Resources)
			res := m.UnresolvedResources[idx]
			var val core.Value
			switch res := res.(type) {
			case program.Constant:
				val = res.Inner
			case program.Parameter:
				var ok bool
				val, ok = m.Vars[res.Name]
				if !ok {
					ch <- MetadataRequest{
						Error: fmt.Errorf("missing variable: %v", res.Name),
					}
					return
				}
			case program.Metadata:
				source_account, ok := m.getResource(res.SourceAccount)
				if !ok {
					ch <- MetadataRequest{
						Error: errors.New("tried to request metadata of an account which has not yet been solved"),
					}
					return
				}
				if (*source_account).GetType() != core.TYPE_ACCOUNT {
					ch <- MetadataRequest{
						Error: fmt.Errorf("tried to request metadata on wrong entity: %v instead of ACCOUNT", (*source_account).GetType()),
					}
					return
				}
				account := (*source_account).(core.Account)
				resp := make(chan core.Value)
				ch <- MetadataRequest{
					Account:  string(account),
					Key:      res.Key,
					Response: resp,
				}
				val = <-resp
				close(resp)
				if val == nil {
					ch <- MetadataRequest{
						Error: errors.New("tried to set nil as resource"),
					}
					return
				}
				if val.GetType() != res.Typ {
					ch <- MetadataRequest{
						Error: fmt.Errorf("wrong type: expected %v, got %v", res.Typ, val.GetType()),
					}
					return
				}
			}
			m.Resources = append(m.Resources, val)
		}
	}()
	return ch, nil
}

func (m *Machine) SetVars(vars map[string]core.Value) error {
	v, err := m.Program.ParseVariables(vars)
	if err != nil {
		return err
	}
	m.Vars = v
	return nil
}

func (m *Machine) SetVarsFromJSON(vars map[string]json.RawMessage) error {
	v, err := m.Program.ParseVariablesJSON(vars)
	if err != nil {
		return err
	}
	m.Vars = v
	return nil
}

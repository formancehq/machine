package compiler

import (
	"errors"
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

type FallbackAccount core.Address

// Returns the resource addresses of all the accounts
func (p *parseVisitor) VisitValueAwareSource(c parser.IValueAwareSourceContext, push_asset func(), mon_addr *core.Address) (map[core.Address]struct{}, *CompileError) {
	needed_accounts := map[core.Address]struct{}{}
	is_all := mon_addr == nil
	switch c := c.(type) {
	case *parser.SrcContext:
		accounts, _, unbounded, err := p.VisitSource(c.Source(), push_asset, is_all)
		if err != nil {
			return nil, err
		}
		for k, v := range accounts {
			needed_accounts[k] = v
		}
		if !is_all {
			p.PushAddress(*mon_addr)
			err := p.TakeFromSource(unbounded, push_asset)
			if err != nil {
				return nil, LogicError(c, err)
			}
		}
	case *parser.SrcAllotmentContext:
		if is_all {
			return nil, LogicError(c, errors.New("cannot take all balance of an allotment source"))
		}
		p.PushAddress(*mon_addr)
		p.VisitAllotment(c.SourceAllotment(), c.SourceAllotment().GetPortions())
		p.AppendInstruction(program.OP_ALLOC)

		sources := c.SourceAllotment().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			accounts, _, fallback, cerr := p.VisitSource(sources[i], push_asset, is_all)
			if cerr != nil {
				return nil, cerr
			}
			for k, v := range accounts {
				needed_accounts[k] = v
			}
			err := p.Bump(int64(i + 1))
			if err != nil {
				return nil, LogicError(c, err)
			}
			err = p.TakeFromSource(fallback, push_asset)
			if err != nil {
				return nil, LogicError(c, err)
			}
		}
		err := p.PushInteger(*core.NewNumber(int64(n)))
		if err != nil {
			return nil, LogicError(c, err)
		}
		p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
	}
	return needed_accounts, nil
}

func (p *parseVisitor) TakeFromSource(fallback *FallbackAccount, push_asset func()) error {
	if fallback == nil {
		p.AppendInstruction(program.OP_TAKE)
		err := p.Bump(1)
		if err != nil {
			return err
		}
		p.AppendInstruction(program.OP_REPAY)
	} else {
		p.AppendInstruction(program.OP_TAKE_MAX)
		err := p.Bump(1)
		if err != nil {
			return err
		}
		p.AppendInstruction(program.OP_REPAY)
		p.PushAddress(core.Address(*fallback))
		err = p.Bump(2)
		if err != nil {
			return err
		}
		p.AppendInstruction(program.OP_TAKE_ALWAYS)
		err = p.PushInteger(*core.NewNumber(2))
		if err != nil {
			return err
		}
		p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
	}
	return nil
}

// Returns the resource addresses of all the accounts,
// the addresses of accounts already emptied,
// and possibly a fallback account if the source has an unbounded overdraft allowance or contains @world
func (p *parseVisitor) VisitSource(c parser.ISourceContext, push_asset func(), is_all bool) (map[core.Address]struct{}, map[core.Address]struct{}, *FallbackAccount, *CompileError) {
	needed_accounts := map[core.Address]struct{}{}
	emptied_accounts := map[core.Address]struct{}{}
	var fallback *FallbackAccount
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, acc_addr, err := p.VisitExpr(c.SourceAccount().GetAccount(), true)
		if err != nil {
			return nil, nil, nil, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, nil, nil, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
		}
		if p.isWorld(*acc_addr) {
			f := FallbackAccount(*acc_addr)
			fallback = &f
		}

		overdraft := c.SourceAccount().GetOverdraft()
		if overdraft == nil {
			// no overdraft: use zero monetary
			push_asset()
			err := p.PushInteger(*core.NewNumber(0))
			if err != nil {
				return nil, nil, nil, LogicError(c, err)
			}
			p.AppendInstruction(program.OP_MONETARY_NEW)
			p.AppendInstruction(program.OP_TAKE_ALL)
		} else {
			if p.isWorld(*acc_addr) {
				return nil, nil, nil, LogicError(c, errors.New("@world is already set to an unbounded overdraft"))
			}
			switch c := overdraft.(type) {
			case *parser.SrcAccountOverdraftSpecificContext:
				ty, _, err := p.VisitExpr(c.GetSpecific(), true)
				if err != nil {
					return nil, nil, nil, err
				}
				if ty != core.TYPE_MONETARY {
					return nil, nil, nil, LogicError(c, errors.New("wrong type: expected monetary"))
				}
				p.AppendInstruction(program.OP_TAKE_ALL)
			case *parser.SrcAccountOverdraftUnboundedContext:
				push_asset()
				err := p.PushInteger(*core.NewNumber(0))
				if err != nil {
					return nil, nil, nil, LogicError(c, err)
				}
				p.AppendInstruction(program.OP_MONETARY_NEW)
				p.AppendInstruction(program.OP_TAKE_ALL)
				f := FallbackAccount(*acc_addr)
				fallback = &f
			}
		}
		needed_accounts[*acc_addr] = struct{}{}
		emptied_accounts[*acc_addr] = struct{}{}

		if fallback != nil && is_all {
			return nil, nil, nil, LogicError(c, errors.New("cannot take all balance of an unbounded source"))
		}

	case *parser.SrcMaxedContext:
		accounts, _, subsource_fallback, cerr := p.VisitSource(c.SourceMaxed().GetSrc(), push_asset, false)
		if cerr != nil {
			return nil, nil, nil, cerr
		}
		ty, _, cerr := p.VisitExpr(c.SourceMaxed().GetMax(), true)
		if cerr != nil {
			return nil, nil, nil, cerr
		}
		if ty != core.TYPE_MONETARY {
			return nil, nil, nil, LogicError(c, errors.New("wrong type: expected monetary as max"))
		}
		for k, v := range accounts {
			needed_accounts[k] = v
		}
		p.AppendInstruction(program.OP_TAKE_MAX)
		err := p.Bump(1)
		if err != nil {
			return nil, nil, nil, LogicError(c, err)
		}
		p.AppendInstruction(program.OP_REPAY)
		if subsource_fallback != nil {
			p.PushAddress(core.Address(*subsource_fallback))
			err := p.Bump(2)
			if err != nil {
				return nil, nil, nil, LogicError(c, err)
			}
			p.AppendInstruction(program.OP_TAKE_ALL)
			err = p.PushInteger(*core.NewNumber(2))
			if err != nil {
				return nil, nil, nil, LogicError(c, err)
			}
			p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
		} else {
			err := p.Bump(1)
			if err != nil {
				return nil, nil, nil, LogicError(c, err)
			}
			p.AppendInstruction(program.OP_DELETE)
		}
	case *parser.SrcInOrderContext:
		sources := c.SourceInOrder().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			accounts, emptied, subsource_fallback, err := p.VisitSource(sources[i], push_asset, is_all)
			if err != nil {
				return nil, nil, nil, err
			}
			fallback = subsource_fallback
			if subsource_fallback != nil && i != n-1 {
				return nil, nil, nil, LogicError(c, errors.New("an unbounded subsource can only be in last position"))
			}
			for k, v := range accounts {
				needed_accounts[k] = v
			}
			for k, v := range emptied {
				if _, ok := emptied_accounts[k]; ok {
					return nil, nil, nil, LogicError(sources[i], fmt.Errorf("%v is already empty at this stage", p.resources[k]))
				}
				emptied_accounts[k] = v
			}
		}
		err := p.PushInteger(*core.NewNumber(int64(n)))
		if err != nil {
			return nil, nil, nil, LogicError(c, err)
		}
		p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
	}
	return needed_accounts, emptied_accounts, fallback, nil
}

package compiler

import (
	"errors"
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

// Returns the resource addresses of all the accounts
func (p *parseVisitor) VisitValueAwareSource(c parser.IValueAwareSourceContext, push_asset func(), mon_addr *core.Address) (map[core.Address]struct{}, *CompileError) {
	needed_accounts := map[core.Address]struct{}{}
	is_all := mon_addr == nil
	switch c := c.(type) {
	case *parser.SrcContext:
		accounts, _, _, err := p.VisitSource(c.Source(), push_asset, is_all)
		if err != nil {
			return nil, err
		}
		for k, v := range accounts {
			needed_accounts[k] = v
		}
		if !is_all {
			p.PushAddress(*mon_addr)
			p.instructions = append(p.instructions, program.OP_TAKE)
			p.PushInteger(core.Number(1))
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.instructions = append(p.instructions, program.OP_REPAY)
		}
	case *parser.SrcAllotmentContext:
		if is_all {
			return nil, LogicError(c, errors.New("cannot take all balance of an allotment source"))
		}
		p.PushAddress(*mon_addr)
		p.VisitAllotment(c.SourceAllotment(), c.SourceAllotment().GetPortions())
		p.instructions = append(p.instructions, program.OP_ALLOC_MON)

		sources := c.SourceAllotment().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			accounts, _, _, err := p.VisitSource(sources[i], push_asset, is_all)
			if err != nil {
				return nil, err
			}
			for k, v := range accounts {
				needed_accounts[k] = v
			}
			p.PushInteger(core.Number(i + 1))
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.instructions = append(p.instructions, program.OP_TAKE)
			p.PushInteger(core.Number(1))
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.instructions = append(p.instructions, program.OP_REPAY)
		}
		p.PushInteger(core.Number(n))
		p.instructions = append(p.instructions, program.OP_ASSEMBLE)
	}
	return needed_accounts, nil
}

// Returns the resource addresses of all the accounts,
// the addresses of accounts already emptied,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitSource(c parser.ISourceContext, push_asset func(), is_all bool) (map[core.Address]struct{}, map[core.Address]struct{}, bool, *CompileError) {
	needed_accounts := map[core.Address]struct{}{}
	emptied_accounts := map[core.Address]struct{}{}
	bottomless := false
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, acc_addr, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return nil, nil, false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, nil, false, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
		}
		bottomless = bottomless || p.isWorld(*acc_addr)
		if p.isWorld(*acc_addr) && is_all {
			return nil, nil, false, LogicError(c, errors.New("cannot take all balance of world"))
		} else {
			push_asset()
			p.instructions = append(p.instructions, program.OP_TAKE_ALL)
			needed_accounts[*acc_addr] = struct{}{}
			emptied_accounts[*acc_addr] = struct{}{}
		}
	case *parser.SrcMaxedContext:
		accounts, _, _, err := p.VisitSource(c.SourceMaxed().GetSrc(), push_asset, false)
		if err != nil {
			return nil, nil, false, err
		}
		ty, _, err := p.VisitExpr(c.SourceMaxed().GetMax(), true)
		if err != nil {
			return nil, nil, false, err
		}
		if ty != core.TYPE_MONETARY {
			return nil, nil, false, LogicError(c, errors.New("wrong type: expected monetary as max"))
		}
		for k, v := range accounts {
			needed_accounts[k] = v
		}
		p.instructions = append(p.instructions, program.OP_TAKE_MAX)
		p.PushInteger(core.Number(1))
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.instructions = append(p.instructions, program.OP_REPAY)
	case *parser.SrcInOrderContext:
		sources := c.SourceInOrder().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			accounts, emptied, subsource_bottomless, err := p.VisitSource(sources[i], push_asset, is_all)
			if err != nil {
				return nil, nil, false, err
			}
			bottomless = bottomless || subsource_bottomless
			if subsource_bottomless && i != n-1 {
				return nil, nil, false, LogicError(c, errors.New("world can only be in last position of source"))
			}
			for k, v := range accounts {
				needed_accounts[k] = v
			}
			for k, v := range emptied {
				if _, ok := emptied_accounts[k]; ok {
					return nil, nil, false, LogicError(sources[i], fmt.Errorf("%v is already empty at this stage", p.resources[k]))
				}
				emptied_accounts[k] = v
			}
		}
		p.PushInteger(core.Number(n))
		p.instructions = append(p.instructions, program.OP_ASSEMBLE)
	}
	return needed_accounts, emptied_accounts, bottomless, nil
}

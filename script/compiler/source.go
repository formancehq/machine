package compiler

import (
	"errors"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

// Returns the resource addresses of all the accounts,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitValueAwareSource(c parser.IValueAwareSourceContext, push_asset func(), mon_addr *core.Address) ([]core.Address, *CompileError) {
	needed_accounts := []core.Address{}
	is_all := mon_addr == nil
	switch c := c.(type) {
	case *parser.SrcContext:
		subsource_needed_accounts, _, err := p.VisitSource(c.Source(), push_asset, is_all)
		if err != nil {
			return nil, err
		}
		needed_accounts = append(needed_accounts, subsource_needed_accounts...)
		if !is_all {
			p.PushAddress(*mon_addr)
			p.instructions = append(p.instructions, program.OP_TAKE)
		}
	case *parser.SrcAllotmentContext:
		if is_all {
			return nil, LogicError(c, errors.New("cannot take all balance of an allotment source"))
		}
		sources := c.SourceAllotment().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			subsource_needed_accounts, _, err := p.VisitSource(sources[i], push_asset, is_all)
			if err != nil {
				return nil, err
			}
			needed_accounts = append(needed_accounts, subsource_needed_accounts...)
		}
		p.VisitAllotment(c.SourceAllotment(), c.SourceAllotment().GetPortions())
		p.PushAddress(*mon_addr)
		p.instructions = append(p.instructions, program.OP_TAKE_SPLIT)
	}
	return needed_accounts, nil
}

// Returns the resource addresses of all the accounts,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitSource(c parser.ISourceContext, push_asset func(), is_all bool) ([]core.Address, bool, *CompileError) {
	needed_accounts := []core.Address{}
	bottomless := false
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, acc_addr, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, false, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
		}
		needed_accounts = append(needed_accounts, *acc_addr)
		bottomless = bottomless || p.isWorld(*acc_addr)
		if p.isWorld(*acc_addr) && is_all {
			return nil, false, LogicError(c, errors.New("cannot take all balance of world"))
		} else {
			push_asset()
			p.instructions = append(p.instructions, program.OP_TAKE_ALL)
		}
	case *parser.SrcMaxedContext:
		accounts, _, err := p.VisitSource(c.SourceMaxed().GetSrc(), push_asset, false)
		if err != nil {
			return nil, false, err
		}
		ty, _, err := p.VisitExpr(c.SourceMaxed().GetMax(), true)
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_MONETARY {
			return nil, false, LogicError(c, errors.New("wrong type: expected monetary as max"))
		}
		needed_accounts = append(needed_accounts, accounts...)
		p.instructions = append(p.instructions, program.OP_TAKE_MAX)
	case *parser.SrcInOrderContext:
		sources := c.SourceInOrder().GetSources()
		n := len(sources)
		for i := 0; i < n; i++ {
			accounts, subsource_bottomless, err := p.VisitSource(sources[i], push_asset, is_all)
			if err != nil {
				return nil, false, err
			}
			needed_accounts = append(needed_accounts, accounts...)
			bottomless = bottomless || subsource_bottomless
			if subsource_bottomless && i != n-1 {
				return nil, false, LogicError(c, errors.New("world can only be in last position of source"))
			}
		}
		p.PushInteger(core.Number(n))
		p.instructions = append(p.instructions, program.OP_ASSEMBLE)
	}
	return needed_accounts, bottomless, nil
}

package compiler

import (
	"errors"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

// Returns the resource addresses of all the accounts,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitSource(c parser.ISourceContext, push_asset func(), mon_addr *core.Address) ([]core.Address, *CompileError) {
	needed_accounts := []core.Address{}
	is_all := mon_addr == nil
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, acc_addr, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return nil, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
		}
		needed_accounts = append(needed_accounts, *acc_addr)
		if is_all {
			if p.isWorld(*acc_addr) {
				return nil, LogicError(c, errors.New("cannot take all balance of world"))
			}
			push_asset()
			p.instructions = append(p.instructions, program.OP_TAKE_ACC_ALL)
		} else {
			p.PushAddress(*mon_addr)
			p.instructions = append(p.instructions, program.OP_TAKE_ACC)
		}
	case *parser.SrcBlockContext:
		sources := c.SourceBlock().GetSources()
		n := len(sources)
		for i := n - 1; i >= 0; i-- {
			ty, acc_addr, err := p.VisitExpr(sources[i], true)
			if err != nil {
				return nil, err
			}
			if ty != core.TYPE_ACCOUNT {
				return nil, LogicError(c, errors.New("wrong type: expected only accounts in sources"))
			}
			needed_accounts = append(needed_accounts, *acc_addr)
			if is_all {
				if p.isWorld(*acc_addr) {
					return nil, LogicError(c, errors.New("cannot take all balance of world"))
				}
				push_asset()
				p.instructions = append(p.instructions, program.OP_TAKE_ACC_ALL)
			} else {
				if p.isWorld(*acc_addr) && i != n-1 {
					return nil, LogicError(c, errors.New("world can only be in last position of source"))
				} else if p.isWorld(*acc_addr) {
					p.PushAddress(*mon_addr)
					p.instructions = append(p.instructions, program.OP_TAKE_ACC)
				} else {
					push_asset()
					p.instructions = append(p.instructions, program.OP_TAKE_ACC_ALL)
				}
			}
		}
		p.PushInteger(core.Number(n))
		p.instructions = append(p.instructions, program.OP_ASSEMBLE)
		if !is_all {
			p.PushAddress(*mon_addr)
			p.instructions = append(p.instructions, program.OP_TAKE)
		}
	case *parser.SrcAllotmentContext:
		if is_all {
			return nil, LogicError(c, errors.New("cannot take all balance of allocation"))
		}
		sources := c.SourceAllotment().GetSrcs()
		n := len(sources)
		for i := n - 1; i >= 0; i-- {
			ty, acc_addr, err := p.VisitExpr(sources[i], true)
			if err != nil {
				return nil, err
			}
			if ty != core.TYPE_ACCOUNT {
				return nil, LogicError(c, errors.New("wrong type: expected only accounts in sources"))
			}
			needed_accounts = append(needed_accounts, *acc_addr)
			if p.isWorld(*acc_addr) {
				p.PushAddress(*mon_addr)
				p.instructions = append(p.instructions, program.OP_TAKE_ACC)
			} else {
				push_asset()
				p.instructions = append(p.instructions, program.OP_TAKE_ACC_ALL)
			}
		}
		err := p.VisitAllotment(c, c.SourceAllotment().GetPortions())
		if err != nil {
			return nil, err
		}
		p.PushAddress(*mon_addr)
		p.instructions = append(p.instructions, program.OP_TAKE_SPLIT)

	}
	return needed_accounts, nil
}

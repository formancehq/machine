package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

// Returns the resource addresses of all the accounts,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitSource(c parser.ISourceContext) ([]core.Address, bool, *CompileError) {
	source_accounts := []core.Address{}
	has_world := false
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, addr, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, false, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
		}
		has_world = has_world || p.isWorld(*addr)
		source_accounts = append(source_accounts, *addr)
	case *parser.SrcInOrderContext:
		sources := c.SourceInOrder().GetSrcs()
		n := len(sources)
		for i := n - 1; i >= 0; i-- {
			subsource_accounts, subsource_has_world, err := p.VisitSource(sources[i])
			if err != nil {
				return nil, false, err
			}
			has_world = has_world || subsource_has_world
			source_accounts = append(source_accounts, subsource_accounts...)
		}
		p.PushInteger(core.Number(n))
		p.instructions = append(p.instructions, program.OP_MAKE_SOURCE_IN_ORDER)
	case *parser.SrcAllotmentContext:
		sources := c.SourceAllotment().GetSrcs()
		n := len(sources)
		for i := 0; i < n; i++ {
			subsource_accounts, subsource_has_world, err := p.VisitSource(sources[i])
			if err != nil {
				return nil, false, err
			}
			has_world = has_world || subsource_has_world
			source_accounts = append(source_accounts, subsource_accounts...)
		}
		err := p.VisitAllotment(c.SourceAllotment())
		if err != nil {
			return nil, false, err
		}
		p.instructions = append(p.instructions, program.OP_MAKE_SOURCE_ALLOTMENT)
	case *parser.SrcLimitedContext:
		subsource_accounts, subsource_has_world, err := p.VisitSource(c.SourceLimited().GetSrc())
		if err != nil {
			return nil, false, err
		}
		has_world = has_world || subsource_has_world
		source_accounts = append(source_accounts, subsource_accounts...)
		ty, _, err := p.VisitExpr(c.SourceLimited().GetMax(), true)
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_MONETARY {
			return nil, false, LogicError(c, errors.New("wrong type: expected monetary"))
		}
		p.instructions = append(p.instructions, program.OP_MAKE_SOURCE_MAXED)
	}
	return source_accounts, has_world, nil
}

func (p *parseVisitor) VisitAllotment(c parser.ISourceAllotmentContext) *CompileError {
	total := big.NewRat(0, 1)
	has_remaining := false
	has_variables := false
	portions := c.GetPortions()
	for i := len(portions) - 1; i >= 0; i-- {
		c := portions[i]
		switch c := c.(type) {
		case *parser.AllotmentPortionConstContext:
			portion, err := core.ParsePortionSpecific(c.GetText())
			if err != nil {
				return LogicError(c, err)
			}
			rat := big.Rat(*portion.Specific)
			total.Add(&rat, total)
			addr, err := p.AllocateResource(program.Constant{Inner: core.Portion(*portion)})
			if err != nil {
				return LogicError(c, err)
			}
			p.PushAddress(addr)
		case *parser.AllotmentPortionVarContext:
			ty, _, err := p.VisitVariable(c.GetPor(), true)
			if err != nil {
				return err
			}
			if ty != core.TYPE_PORTION {
				return LogicError(c,
					fmt.Errorf("wrong type: expected type portion for variable: %v", ty),
				)
			}
			has_variables = true
		case *parser.AllotmentPortionRemainingContext:
			if has_remaining {
				return LogicError(c,
					errors.New("two uses of `remaining` in the same allocation"),
				)
			}
			addr, err := p.AllocateResource(program.Constant{Inner: core.NewPortionRemaining()})
			if err != nil {
				return LogicError(c, err)
			}
			p.PushAddress(addr)
			has_remaining = true
		}
	}
	if has_variables && !has_remaining {
		return LogicError(c,
			errors.New("allocation has variable portions but no 'remaining'"),
		)
	}
	if (has_remaining && total.Cmp(big.NewRat(1, 1)) != -1) || (!has_remaining && total.Cmp(big.NewRat(1, 1)) != 0) {
		return LogicError(c,
			errors.New("sum of known portions is already equal or is greater than 100%"),
		)
	}
	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)
	return nil
}

// if !has_remaining {
// 		return LogicError(c,
// 			errors.New("allocation has variable portions but no 'remaining'"),
// 		)
// 	}
// 	p.PushValue(core.Number(len(portions)))

// 	if total.Cmp(big.NewRat(1, 1)) != -1 {
// 		return LogicError(c,
// 			errors.New("sum of known portions is already equal or is greater than 100%"),
// 		)
// 	}
// 	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)

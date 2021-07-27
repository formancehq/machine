package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) VisitAllocation(c parser.IAllocationContext) *CompileError {
	switch c := c.(type) {
	case *parser.AllocAccountContext:
		ty, _, err := p.VisitExpr(c.Expression())
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return LogicError(c,
				errors.New("expected account or allocation as destination"),
			)
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	case *parser.AllocConstContext:
		err := p.VisitAllocBlockConst(c.AllocBlockConst())
		return err
	case *parser.AllocDynContext:
		err := p.VisitAllocBlockDyn(c.AllocBlockDyn())
		return err
	}
	return nil
}

func (p *parseVisitor) VisitAllocBlockConst(c parser.IAllocBlockConstContext) *CompileError {
	// allocate
	portions := []core.Portion{}
	for _, c := range c.GetPortions() {
		switch c := c.(type) {
		case *parser.AllocPartConstConstContext:
			portion, err := core.ParsePortionSpecific(c.GetText())
			if err != nil {
				return LogicError(c, err)
			}
			portions = append(portions, *portion)
		case *parser.AllocPartConstRemainingContext:
			portions = append(portions, core.NewPortionRemaining())
			// presence of too many 'remaining' checked on allotment creation
		}
	}
	allotment, err := core.NewAllotment(portions)
	if err != nil {
		c.GetStart()
		return LogicError(c, err)
	}
	p.PushValue(*allotment)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocBlockDyn(c parser.IAllocBlockDynContext) *CompileError {
	total := big.NewRat(0, 1)
	// allocate
	portions := c.GetPortions()
	has_remaining := false
	for i := len(portions) - 1; i >= 0; i-- {
		c := portions[i]
		switch c := c.(type) {
		case *parser.AllocPartDynConstContext:
			portion, err := core.ParsePortionSpecific(c.GetText())
			if err != nil {
				return LogicError(c, err)
			}
			rat := big.Rat(*portion.Specific)
			total.Add(&rat, total)
			p.PushValue(core.Portion(*portion))
		case *parser.AllocPartDynVarContext:
			ty, _, err := p.VisitVariable(c.GetPor())
			if err != nil {
				return err
			}
			if ty != core.TYPE_PORTION {
				return LogicError(c,
					fmt.Errorf("wrong type: expected type portion for variable: %v", ty),
				)
			}
		case *parser.AllocPartDynRemainingContext:
			if has_remaining {
				return LogicError(c,
					errors.New("two uses of `remaining` in the same allocation"),
				)
			}
			p.PushValue(core.NewPortionRemaining())
			has_remaining = true
		}
	}
	if !has_remaining {
		return LogicError(c,
			errors.New("allocation has variable portions but no 'remaining'"),
		)
	}
	p.PushValue(core.Number(len(portions)))

	if total.Cmp(big.NewRat(1, 1)) != -1 {
		return LogicError(c,
			errors.New("sum of known portions is already equal or is greater than 100%"),
		)
	}
	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	// distribute to destination accounts

	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IExpressionContext) *CompileError {
	for _, dest := range dests {
		ty, _, err := p.VisitExpr(dest)
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return LogicError(dest, errors.New("expected account as destination of allocation line"))
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	}
	return nil
}

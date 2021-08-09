package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) VisitDestination(c parser.IDestinationContext) *CompileError {
	switch c := c.(type) {
	case *parser.DestAccountContext:
		ty, _, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return LogicError(c,
				errors.New("expected account or allocation as destination"),
			)
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	case *parser.DestAllotmentContext:
		err := p.VisitDestinationAllotment(c.DestinationAllotment())
		return err
	}
	return nil
}

func (p *parseVisitor) VisitDestinationAllotment(c parser.IDestinationAllotmentContext) *CompileError {
	total := big.NewRat(0, 1)
	// allocate
	portions := c.GetPortions()
	has_variable := false
	has_remaining := false
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
			p.PushAddress(*addr)
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
			has_variable = true
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
			p.PushAddress(*addr)
			has_remaining = true
		}
	}
	if has_variable && !has_remaining {
		return LogicError(c,
			errors.New("allotment has variable portions but no 'remaining'"),
		)
	}
	p.PushInteger(core.Number(len(portions)))

	if has_remaining && total.Cmp(big.NewRat(1, 1)) != -1 {
		return LogicError(c,
			errors.New("allotment has variable portions but sum of known portions is already equal or is greater than 100%"),
		)
	}
	if !has_variable && !has_remaining && total.Cmp(big.NewRat(1, 1)) != 0 {
		return LogicError(c,
			errors.New("allotment has no variable or remaining portions but sum is not 100%"),
		)
	}
	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	// distribute to destination accounts

	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IExpressionContext) *CompileError {
	for _, dest := range dests {
		ty, _, err := p.VisitExpr(dest, true)
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

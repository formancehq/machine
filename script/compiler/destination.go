package compiler

import (
	"errors"

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
	err := p.VisitAllotment(c, c.GetPortions())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_ALLOC)
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

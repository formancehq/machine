package compiler

import (
	"errors"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) VisitDestination(c parser.IDestinationContext) *CompileError {
	err := p.VisitDestinationRecursive(c)
	if err != nil {
		return err
	}
	p.AppendInstruction(program.OP_REPAY)
	return nil
}

func (p *parseVisitor) VisitDestinationRecursive(c parser.IDestinationContext) *CompileError {
	switch c := c.(type) {
	case *parser.DestAccountContext:
		p.AppendInstruction(program.OP_FUNDING_SUM)
		p.AppendInstruction(program.OP_TAKE)
		ty, _, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return LogicError(c,
				errors.New("wrong type: expected account as destination"),
			)
		}
		p.AppendInstruction(program.OP_SEND)
		return nil
	case *parser.DestInOrderContext:
		dests := c.DestinationInOrder().GetDests()
		amounts := c.DestinationInOrder().GetAmounts()
		n := len(dests)

		// initialize the `kept` accumulator
		p.AppendInstruction(program.OP_FUNDING_SUM)
		p.AppendInstruction(program.OP_ASSET)
		p.PushInteger(*core.NewNumber(0))
		p.AppendInstruction(program.OP_MONETARY_NEW)

		p.Bump(1)

		for i := 0; i < n; i++ {

			ty, _, err := p.VisitExpr(amounts[i], true)
			if err != nil {
				return err
			}
			if ty != core.TYPE_MONETARY {
				return LogicError(c, errors.New("wrong type: expected monetary as max"))
			}
			p.AppendInstruction(program.OP_TAKE_MAX)
			p.Bump(2)
			p.AppendInstruction(program.OP_DELETE)
			err = p.VisitKeptOrDestination(dests[i])
			if err != nil {
				return err
			}
			p.AppendInstruction(program.OP_FUNDING_SUM)
			p.Bump(3)
			p.AppendInstruction(program.OP_MONETARY_ADD)
			p.Bump(1)
			p.Bump(2)
			p.PushInteger(*core.NewNumber(2))
			p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
		}
		p.AppendInstruction(program.OP_FUNDING_REVERSE)
		p.Bump(1)
		p.AppendInstruction(program.OP_TAKE)
		p.AppendInstruction(program.OP_FUNDING_REVERSE)
		p.Bump(1)
		p.AppendInstruction(program.OP_FUNDING_REVERSE)
		err := p.VisitKeptOrDestination(c.DestinationInOrder().GetRemainingDest())
		if err != nil {
			return err
		}
		p.Bump(1)
		p.PushInteger(*core.NewNumber(2))
		p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
		return nil
	case *parser.DestAllotmentContext:
		err := p.VisitDestinationAllotment(c.DestinationAllotment())
		return err
	default:
		return InternalError(c)
	}
}

func (p *parseVisitor) VisitKeptOrDestination(c parser.IKeptOrDestinationContext) *CompileError {
	switch c := c.(type) {
	case *parser.IsKeptContext:
		return nil
	case *parser.IsDestinationContext:
		err := p.VisitDestinationRecursive(c.Destination())
		return err
	default:
		return InternalError(c)
	}
}

func (p *parseVisitor) VisitDestinationAllotment(c parser.IDestinationAllotmentContext) *CompileError {
	p.AppendInstruction(program.OP_FUNDING_SUM)
	err := p.VisitAllotment(c, c.GetPortions())
	if err != nil {
		return err
	}
	p.AppendInstruction(program.OP_ALLOC)
	err = p.VisitAllocDestination(c.GetDests())
	if err != nil {
		return err
	}
	return nil
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IKeptOrDestinationContext) *CompileError {
	p.Bump(int64(len(dests)))
	for _, dest := range dests {
		p.Bump(1)
		p.AppendInstruction(program.OP_TAKE)
		err := p.VisitKeptOrDestination(dest)
		if err != nil {
			return err
		}
		p.Bump(1)
		p.PushInteger(*core.NewNumber(2))
		p.AppendInstruction(program.OP_FUNDING_ASSEMBLE)
	}
	return nil
}

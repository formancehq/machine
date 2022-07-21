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
	p.instructions = append(p.instructions, program.OP_REPAY)
	return nil
}

func (p *parseVisitor) VisitDestinationRecursive(c parser.IDestinationContext) *CompileError {
	switch c := c.(type) {
	case *parser.DestAccountContext:
		p.instructions = append(p.instructions, program.OP_FUNDING_SUM)
		p.instructions = append(p.instructions, program.OP_TAKE)
		ty, _, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return LogicError(c,
				errors.New("wrong type: expected account as destination"),
			)
		}
		p.instructions = append(p.instructions, program.OP_SEND)
		return nil
	case *parser.DestInOrderContext:
		dests := c.DestinationInOrder().GetDests()
		amounts := c.DestinationInOrder().GetAmounts()
		n := len(dests)

		// initialize the `kept` accumulator
		p.instructions = append(p.instructions, program.OP_FUNDING_SUM)
		p.instructions = append(p.instructions, program.OP_ASSET)
		p.PushInteger(0)
		p.instructions = append(p.instructions, program.OP_MONETARY_NEW)

		p.PushInteger(1)
		p.instructions = append(p.instructions, program.OP_BUMP)

		for i := 0; i < n; i++ {

			ty, _, err := p.VisitExpr(amounts[i], true)
			if err != nil {
				return err
			}
			if ty != core.TYPE_MONETARY {
				return LogicError(c, errors.New("wrong type: expected monetary as max"))
			}
			p.instructions = append(p.instructions, program.OP_TAKE_MAX)
			err = p.VisitKeptOrDestination(dests[i])
			if err != nil {
				return err
			}
			p.instructions = append(p.instructions, program.OP_FUNDING_SUM)
			p.PushInteger(3)
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.instructions = append(p.instructions, program.OP_MONETARY_ADD)
			p.PushInteger(1)
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.PushInteger(2)
			p.instructions = append(p.instructions, program.OP_BUMP)
			p.PushInteger(2)
			p.instructions = append(p.instructions, program.OP_FUNDING_ASSEMBLE)
		}
		p.instructions = append(p.instructions, program.OP_FUNDING_REVERSE)
		p.PushInteger(1)
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.instructions = append(p.instructions, program.OP_TAKE)
		p.instructions = append(p.instructions, program.OP_FUNDING_REVERSE)
		p.PushInteger(1)
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.instructions = append(p.instructions, program.OP_FUNDING_REVERSE)
		err := p.VisitKeptOrDestination(c.DestinationInOrder().GetRemainingDest())
		if err != nil {
			return err
		}
		p.PushInteger(1)
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.PushInteger(2)
		p.instructions = append(p.instructions, program.OP_FUNDING_ASSEMBLE)
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
	p.instructions = append(p.instructions, program.OP_FUNDING_SUM)
	err := p.VisitAllotment(c, c.GetPortions())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_ALLOC)
	err = p.VisitAllocDestination(c.GetDests())
	if err != nil {
		return err
	}
	return nil
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IKeptOrDestinationContext) *CompileError {
	p.PushInteger(core.Number(len(dests)))
	p.instructions = append(p.instructions, program.OP_BUMP)
	for _, dest := range dests {
		p.PushInteger(core.Number(1))
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.instructions = append(p.instructions, program.OP_TAKE)
		err := p.VisitKeptOrDestination(dest)
		if err != nil {
			return err
		}
		p.PushInteger(core.Number(1))
		p.instructions = append(p.instructions, program.OP_BUMP)
		p.PushInteger(2)
		p.instructions = append(p.instructions, program.OP_FUNDING_ASSEMBLE)
	}
	return nil
}

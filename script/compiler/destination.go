package compiler

import (
	"errors"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) VisitDestination(c parser.IDestinationContext) *CompileError {
	bottomless, err := p.VisitDestinationRec(c)
	if err != nil {
		return err
	}
	if !bottomless {
		return LogicError(c, errors.New("this destination must not be capped"))
	}
	return nil
}

func (p *parseVisitor) VisitDestinationRec(c parser.IDestinationContext) (bool, *CompileError) {
	switch c := c.(type) {
	case *parser.DestAccountContext:
		ty, _, err := p.VisitExpr(c.Expression(), true)
		if err != nil {
			return false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return false, LogicError(c,
				errors.New("expected account or allocation as destination"),
			)
		}
		p.instructions = append(p.instructions, program.OP_SEND)
		return true, nil
	case *parser.DestMaxedContext:
		ty, _, err := p.VisitExpr(c.DestinationMaxed().GetMax(), true)
		if err != nil {
			return false, err
		}
		if ty != core.TYPE_MONETARY {
			return false, LogicError(c, errors.New("wrong type: expected monetary as max"))
		}
		p.instructions = append(p.instructions, program.OP_TAKE_MAX)
		_, err = p.VisitDestinationRec(c.DestinationMaxed().GetDest())
		if err != nil {
			return false, err
		}
		return false, nil
	case *parser.DestInOrderContext:
		dests := c.DestinationInOrder().GetDests()
		n := len(dests)
		for i := 0; i < n; i++ {
			bottomless, err := p.VisitDestinationRec(dests[i])
			if err != nil {
				return false, err
			}
			if bottomless && i != n-1 {
				return false, LogicError(dests[i+1], errors.New("the value has already been entirely distributed to previous destinations at this stage"))
			}
		}
		return true, nil
	case *parser.DestAllotmentContext:
		err := p.VisitDestinationAllotment(c.DestinationAllotment())
		return true, err
	default:
		return false, InternalError(c)
	}
}

func (p *parseVisitor) VisitDestinationAllotment(c parser.IDestinationAllotmentContext) *CompileError {
	err := p.VisitAllotment(c, c.GetPortions())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_ALLOC)
	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IDestinationContext) *CompileError {
	for _, dest := range dests {
		bottomless, err := p.VisitDestinationRec(dest)
		if err != nil {
			return err
		}
		if !bottomless {
			return LogicError(dest, errors.New("this destination must not be capped"))
		}
	}
	return nil
}

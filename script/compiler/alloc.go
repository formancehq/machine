package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) VisitAllocation(c parser.IAllocationContext) error {
	switch c := c.(type) {
	case *parser.AllocAccountContext:
		ty, _, err := p.VisitExpr(c.Expression())
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return errors.New("expected account or allocation as destination")
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

func (p *parseVisitor) VisitAllocBlockConst(c parser.IAllocBlockConstContext) error {
	// allocate
	portions := []*big.Rat{}
	for _, c := range c.GetPortions() {
		switch c := c.(type) {
		case *parser.AllocPartConstConstContext:
			portion, ok := core.ParsePortion(c.PortionConst().GetPor().GetText())
			if !ok {
				return errors.New("invalid format for allocation portion")
			}
			rat := big.Rat(*portion)
			portions = append(portions, &rat)
		case *parser.AllocPartConstRemainingContext:
			portions = append(portions, nil)
		}
	}
	allotment, err := core.NewAllotment(portions)
	if err != nil {
		return err
	}
	p.PushValue(*allotment)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocBlockDyn(c parser.IAllocBlockDynContext) error {
	total := big.NewRat(0, 1)
	// allocate
	portions := c.GetPortions()
	has_remaining := false
	for i := len(portions) - 1; i >= 0; i-- {
		c := portions[i]
		switch c := c.(type) {
		case *parser.AllocPartDynConstContext:
			portion, ok := core.ParsePortion(c.PortionConst().GetPor().GetText())
			if !ok {
				return errors.New("invalid format for allocation portion")
			}
			rat := big.Rat(*portion)
			total.Add(&rat, total)
			p.PushValue(core.Portion(*portion))
		case *parser.AllocPartDynVarContext:
			ty, _, err := p.VisitVariable(c.PortionVar().GetPor())
			if err != nil {
				return err
			}
			if ty != core.TYPE_PORTION {
				return fmt.Errorf("tried to use wrong variable type for portion of allocation: %v", ty)
			}
		case *parser.AllocPartDynRemainingContext:
			if has_remaining {
				return errors.New("two uses of `remaining` in the same allocation")
			}
			p.PushValue(core.Number(0)) // use Number(0) as indicator of 'remaining' portion in the stack
			has_remaining = true
		}
	}
	p.PushValue(core.Number(len(portions)))
	if total.Cmp(big.NewRat(1, 1)) != -1 {
		return errors.New("sum of portions did not equal 100%")
	}
	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	// distribute to destination accounts
	return p.VisitAllocDestination(c.GetDests())
}

func (p *parseVisitor) VisitAllocDestination(dests []parser.IExpressionContext) error {
	for _, dest := range dests {
		ty, _, err := p.VisitExpr(dest)
		if err != nil {
			return err
		}
		if ty != core.TYPE_ACCOUNT {
			return errors.New("expected account as destination of allocation line")
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	}
	return nil
}

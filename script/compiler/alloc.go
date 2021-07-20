package compiler

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

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
	total := big.NewRat(0, 1)
	// allocate
	allotment := []big.Rat{}
	portions := c.GetPortions()
	for _, v := range portions {
		frac, err := p.VisitPortion(v.GetPor())
		if err != nil {
			return err
		}
		total.Add(frac, total)
		allotment = append(allotment, *frac)
	}
	if total.Cmp(big.NewRat(1, 1)) != 0 {
		return errors.New("sum of fractions did not equal 100%")
	}
	p.PushValue(core.Allotment(allotment))
	p.instructions = append(p.instructions, program.OP_ALLOC)
	// distribute to destination accounts
	for _, dest := range c.GetDests() {
		ty, _, err := p.VisitExpr(dest)
		if err != nil {
			return nil
		}
		if ty != core.TYPE_ACCOUNT {
			return errors.New("expected account as destination of allocation line")
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	}
	return nil
}

func (p *parseVisitor) VisitAllocBlockDyn(c parser.IAllocBlockDynContext) error {
	total := big.NewRat(0, 1)
	// allocate
	portions := c.GetPortions()
	has_remaining := false
	for _, v := range portions {
		switch c := v.(type) {
		case *parser.AllocPartDynConstContext:
			portion, err := p.VisitPortion(c.PortionConst().GetPor())
			if err != nil {
				return err
			}
			total.Add(portion, total)
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
			p.PushValue(core.Number(0)) // use Number(0) as indicator of remaining in the stack
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
	for _, dest := range c.GetDests() {
		ty, _, err := p.VisitExpr(dest)
		if err != nil {
			return nil
		}
		if ty != core.TYPE_ACCOUNT {
			return errors.New("expected account as destination of allocation line")
		}
		p.instructions = append(p.instructions, program.OP_SEND)
	}
	return nil
}

func (p *parseVisitor) VisitPortion(c parser.IPortionContext) (*big.Rat, error) {
	switch c := c.(type) {
	case *parser.RatioContext:
		v := strings.Split(c.GetR().GetText(), "/")
		ns := strings.Trim(v[0], " \t")
		ds := strings.Trim(v[1], " \t")
		n, err := strconv.ParseInt(ns, 10, 64)
		if err != nil {
			return nil, err
		}
		if n <= 0 {
			return nil, errors.New("numerator must be greater than zero")
		}
		d, err := strconv.ParseInt(ds, 10, 64)
		if err != nil {
			return nil, err
		}
		if d <= 0 {
			return nil, errors.New("denominator must be greater than zero")
		}
		return big.NewRat(int64(n), int64(d)), nil
	case *parser.PercentageContext:
		pint := c.GetPint().GetText()
		maybe_pfrac := c.GetPfrac()
		var pfrac string
		if maybe_pfrac != nil {
			pfrac = maybe_pfrac.GetText()
		}
		res, ok := new(big.Rat).SetString(pint + "." + pfrac)
		res.Mul(res, big.NewRat(1, 100))
		fmt.Println(res)
		if !ok {
			return nil, errors.New("percentage was not in a valid format")
		}
		if res.Cmp(big.NewRat(0, 1)) != 1 || res.Cmp(big.NewRat(1, 1)) != -1 {
			return nil, errors.New("percentage must be greater than zero and less than 100")
		}
		return res, nil
	default:
		panic("internal compiler error")
	}
}

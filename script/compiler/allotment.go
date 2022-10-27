package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/formancehq/machine/core"
	"github.com/formancehq/machine/script/parser"
	"github.com/formancehq/machine/vm/program"
)

func (p *parseVisitor) VisitAllotment(c antlr.ParserRuleContext, portions []parser.IAllotmentPortionContext) *CompileError {
	total := big.NewRat(0, 1)
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
	if total.Cmp(big.NewRat(1, 1)) == 1 {
		return LogicError(c,
			errors.New("the sum of known portions is greater than 100%"),
		)
	}
	if total.Cmp(big.NewRat(1, 1)) == -1 && !has_remaining {
		return LogicError(c,
			errors.New("the sum of portions might be less than 100%"),
		)
	}
	if total.Cmp(big.NewRat(1, 1)) == 0 && has_variable {
		return LogicError(c,
			errors.New("the sum of portions might be greater than 100%"),
		)
	}
	if total.Cmp(big.NewRat(1, 1)) == 0 && has_remaining {
		return LogicError(c,
			errors.New("known portions are already equal to 100%"),
		)
	}
	err := p.PushInteger(*core.NewNumber(int64(len(portions))))
	if err != nil {
		return LogicError(c, err)
	}
	p.AppendInstruction(program.OP_MAKE_ALLOTMENT)
	return nil
}

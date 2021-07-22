package compiler

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

type parseVisitor struct {
	elistener       *ErrorListener
	instructions    []byte
	constants       []core.Value // must not exceed 32768 elements
	variables       map[string]program.VarInfo
	needed_balances map[core.Address]map[core.Address]struct{} // for each account, set of monetaries (for their assets)
}

func (p *parseVisitor) LogicError(c antlr.ParserRuleContext, err error) error {
	p.elistener.Errors = append(p.elistener.Errors, CompileError{
		line:   c.GetStart().GetLine(),
		column: c.GetStart().GetColumn(),
		len:    len(c.GetText()),
		msg:    err.Error(),
	})
	return err
}

// Allocates constants if it hasn't already been,
// and returns its resource address.
func (p *parseVisitor) AllocateConstant(v core.Value) (core.Address, error) {
	for i := 0; i < len(p.constants); i++ {
		if p.constants[i] == v {
			return core.Address(i), nil
		}
	}
	if len(p.constants) >= 32768 {
		return 0, errors.New("number of unique constants exceeded 32768")
	}
	p.constants = append(p.constants, v)
	return core.Address(len(p.constants) - 1), nil
}

func (p *parseVisitor) PushValue(val core.Value) (*core.Address, error) {
	switch val := val.(type) {
	case core.Account, core.Asset, core.Monetary, core.Allotment, core.Portion:
		p.instructions = append(p.instructions, program.OP_APUSH)
		addr, err := p.AllocateConstant(val)
		if err != nil {
			return nil, err
		}
		bytes := addr.ToBytes()
		p.instructions = append(p.instructions, bytes...)
		return &addr, nil
	case core.Number:
		p.instructions = append(p.instructions, program.OP_IPUSH)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(val))
		p.instructions = append(p.instructions, bytes...)
		return nil, nil
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) isMonetaryAll(addr core.Address) bool {
	if addr.IsConstant() {
		idx := addr.ToIdx()
		if idx < len(p.constants) {
			if mon, ok := p.constants[idx].(core.Monetary); ok {
				return mon.Amount.All
			}
		}
	}
	return false
}

func (p *parseVisitor) isWorld(addr core.Address) bool {
	if addr.IsConstant() {
		idx := addr.ToIdx()
		if idx < len(p.constants) {
			if acc, ok := p.constants[idx].(core.Account); ok {
				if string(acc) == "world" {
					return true
				}
			}
		}
	}
	return false
}

func (p *parseVisitor) VisitVariable(c parser.IVariableContext) (core.Type, *core.Address, error) {
	name := c.GetText()[1:] // strip '$' prefix
	if info, ok := p.variables[name]; ok {
		p.instructions = append(p.instructions, program.OP_APUSH)
		bytes := info.Addr.ToBytes()
		p.instructions = append(p.instructions, bytes...)
		return info.Ty, &info.Addr, nil
	} else {
		return 0, nil, p.LogicError(c, errors.New("variable not declared"))
	}
}

func (p *parseVisitor) VisitExpr(c parser.IExpressionContext) (core.Type, *core.Address, error) {
	switch c := c.(type) {
	case *parser.ExprAddSubContext:
		ty, _, err := p.VisitExpr(c.GetLhs())
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, p.LogicError(c, errors.New("tried to do arithmetic with wrong type"))
		}
		ty, _, err = p.VisitExpr(c.GetRhs())
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, p.LogicError(c, errors.New("tried to do arithmetic with wrong type"))
		}
		switch c.GetOp().GetTokenType() {
		case parser.NumScriptLexerOP_ADD:
			p.instructions = append(p.instructions, program.OP_IADD)
		case parser.NumScriptLexerOP_SUB:
			p.instructions = append(p.instructions, program.OP_ISUB)
		}
		return core.TYPE_NUMBER, nil, nil
	case *parser.ExprLiteralContext:
		ty, addr, err := p.VisitLit(c.GetLit())
		if err != nil {
			return 0, nil, err
		}
		return ty, addr, nil
	case *parser.ExprVariableContext:
		ty, addr, err := p.VisitVariable(c.GetVar_())
		return ty, addr, err
	default:
		panic("internal compiler error")
	}
}

// pushes a value from a literal onto the stack
func (p *parseVisitor) VisitLit(c parser.ILiteralContext) (core.Type, *core.Address, error) {
	switch c := c.(type) {
	case *parser.LitAccountContext:
		account := core.Account(c.GetText()[1:])
		addr, err := p.PushValue(account)
		if err != nil {
			return 0, nil, p.LogicError(c, err)
		}
		return core.TYPE_ACCOUNT, addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		addr, err := p.PushValue(asset)
		if err != nil {
			return 0, nil, p.LogicError(c, err)
		}
		return core.TYPE_ASSET, addr, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return 0, nil, p.LogicError(c, err)
		}
		number := core.Number(n)
		p.PushValue(number) // does not fail with numbers
		return core.TYPE_NUMBER, nil, nil
	case *parser.LitMonetaryContext:
		asset := c.Monetary().GetAsset().GetText()
		var amount core.Amount
		switch c := c.Monetary().GetAmt().(type) {
		case *parser.AmountAllContext:
			amount = core.NewAmountAll()
		case *parser.AmountSpecificContext:
			a, err := strconv.ParseUint(c.GetText(), 10, 64)
			if err != nil {
				return 0, nil, p.LogicError(c, err)
			}
			amount = core.NewAmountSpecific(a)
		}
		monetary := core.Monetary{
			Asset:  core.Asset(asset),
			Amount: amount,
		}
		addr, err := p.PushValue(monetary)
		if err != nil {
			return 0, nil, p.LogicError(c, err)
		}
		return core.TYPE_MONETARY, addr, nil
	default:
		panic("internal compiler error")
	}
}

// Returns the resource addresses of all the accounts,
// and true if the source is bottomless (contains @world)
func (p *parseVisitor) VisitSource(c parser.ISourceContext) ([]core.Address, bool, error) {
	needed_accounts := []core.Address{}
	bottomless := false
	switch c := c.(type) {
	case *parser.SrcAccountContext:
		ty, addr, err := p.VisitExpr(c.Expression())
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, false, p.LogicError(c, errors.New("expected account or allocation as destination"))
		}
		bottomless = bottomless || p.isWorld(*addr)
		needed_accounts = append(needed_accounts, *addr)
		p.PushValue(core.Number(1))
	case *parser.SrcBlockContext:
		sources := c.SourceBlock().GetSources()
		n := len(sources)
		for i := n - 1; i >= 0; i-- {
			ty, addr, err := p.VisitExpr(sources[i])
			if err != nil {
				return nil, false, err
			}
			if ty != core.TYPE_ACCOUNT {
				return nil, false, p.LogicError(c, errors.New("expected only accounts in sources"))
			}
			bottomless = bottomless || p.isWorld(*addr)
			needed_accounts = append(needed_accounts, *addr)
		}
		p.PushValue(core.Number(len(c.SourceBlock().GetSources())))
		p.instructions = append(p.instructions, program.OP_SOURCE)
	}
	return needed_accounts, bottomless, nil
}

func (p *parseVisitor) VisitSend(c *parser.SendContext) error {
	ty, mon_addr, err := p.VisitExpr(c.GetMon())
	if err != nil {
		return err
	}
	if ty != core.TYPE_MONETARY {
		return p.LogicError(c, errors.New("wrong type for monetary value"))
	}

	needed_accounts, bottomless, err := p.VisitSource(c.GetSrc())
	if err != nil {
		return err
	}
	if p.isMonetaryAll(*mon_addr) && bottomless {
		return p.LogicError(c, errors.New("cannot take all balance of world"))
	}

	// add source accounts to the needed balances
	for _, acc := range needed_accounts {
		if b, ok := p.needed_balances[acc]; ok {
			b[*mon_addr] = struct{}{}
		} else {
			p.needed_balances[acc] = map[core.Address]struct{}{
				*mon_addr: {},
			}
		}
	}

	err = p.VisitAllocation(c.GetDest())
	if err != nil {
		return err
	}
	return nil
}

func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) error {
	_, _, err := p.VisitExpr(ctx.GetExpr())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

func (p *parseVisitor) VisitVars(c *parser.VarListDeclContext) error {
	if len(c.GetV()) > 32768 {
		return p.LogicError(c, fmt.Errorf("number of variables exceeded %v", 32768))
	}
	for _, v := range c.GetV() {
		name := v.GetName().GetText()[1:]
		if _, ok := p.variables[name]; ok {
			return p.LogicError(c, errors.New("duplicate variable"))
		}
		var ty core.Type
		switch v.GetTy().GetText() {
		case "account":
			ty = core.TYPE_ACCOUNT
		case "asset":
			ty = core.TYPE_ASSET
		case "number":
			ty = core.TYPE_NUMBER
		case "monetary":
			ty = core.TYPE_MONETARY
		case "portion":
			ty = core.TYPE_PORTION
		default:
			return errors.New("internal compiler error")
		}
		addr := core.NewVarAddress(uint16(len(p.variables)))
		p.variables[name] = program.VarInfo{
			Ty:   ty,
			Addr: addr,
		}
	}
	return nil
}

func (p *parseVisitor) VisitScript(c parser.IScriptContext) error {
	switch c := c.(type) {
	case *parser.ScriptContext:
		vars := c.GetVars()
		if vars != nil {
			switch c := vars.(type) {
			case *parser.VarListDeclContext:
				err := p.VisitVars(c)
				if err != nil {
					return err
				}
			default:
				panic("internal compiler error")
			}
		}
		for _, stmt := range c.GetStmts() {
			switch c := stmt.(type) {
			case *parser.PrintContext:
				err := p.VisitPrint(c)
				if err != nil {
					return err
				}
			case *parser.FailContext:
				p.instructions = append(p.instructions, program.OP_FAIL)
			case *parser.SendContext:
				err := p.VisitSend(c)
				if err != nil {
					return err
				}
			default:
				panic("Invalid context")
			}
		}
	default:
		panic("Invalid context")
	}
	return nil
}

type arg struct {
	name string
	ty   core.Type
}

func Compile(input string) (*program.Program, error) {
	elistener := &ErrorListener{}

	is := antlr.NewInputStream(input)
	lexer := parser.NewNumScriptLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(elistener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewNumScriptParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(elistener)

	p.BuildParseTrees = true

	tree := p.Script()

	if len(elistener.Errors) != 0 {
		err := CompileErrorList{
			errors: elistener.Errors,
			source: input,
		}
		return nil, &err
	}

	visitor := parseVisitor{
		elistener:       elistener,
		instructions:    make([]byte, 0),
		constants:       make([]core.Value, 0),
		variables:       make(map[string]program.VarInfo),
		needed_balances: make(map[core.Address]map[core.Address]struct{}),
	}

	err := visitor.VisitScript(tree)

	if err != nil {
		err := CompileErrorList{
			errors: elistener.Errors,
			source: input,
		}
		return nil, &err
	}

	return &program.Program{
		Instructions:   visitor.instructions,
		Constants:      visitor.constants,
		Variables:      visitor.variables,
		NeededBalances: visitor.needed_balances,
	}, nil
}

package compiler

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/numary/machine/core"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

type parseVisitor struct {
	elistener       *ErrorListener
	instructions    []byte
	resources       []program.Resource                         // must not exceed 65536 elements
	var_idx         map[string]core.Address                    // maps name to resource index
	needed_balances map[core.Address]map[core.Address]struct{} // for each account, set of assets needed
}

// Allocates constants if it hasn't already been,
// and returns its resource address.
func (p *parseVisitor) findConstant(constant program.Constant) (*core.Address, bool) {
	for i := 0; i < len(p.resources); i++ {
		if c, ok := p.resources[i].(program.Constant); ok {
			if c == constant {
				addr := core.Address(i)
				return &addr, true
			}
		}
	}
	return nil, false
}

func (p *parseVisitor) AllocateResource(res program.Resource) (*core.Address, error) {
	if c, ok := res.(program.Constant); ok {
		idx, ok := p.findConstant(c)
		if ok {
			return idx, nil
		}
	}
	if len(p.resources) >= 65536 {
		return nil, errors.New("number of unique constants exceeded 65536")
	}
	p.resources = append(p.resources, res)
	addr := core.NewAddress(uint16(len(p.resources) - 1))
	return &addr, nil
}

func (p *parseVisitor) PushAddress(addr *core.Address) {
	p.instructions = append(p.instructions, program.OP_APUSH)
	bytes := addr.ToBytes()
	p.instructions = append(p.instructions, bytes...)
}

func (p *parseVisitor) PushInteger(val core.Number) {
	p.instructions = append(p.instructions, program.OP_IPUSH)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(val))
	p.instructions = append(p.instructions, bytes...)
}

func (p *parseVisitor) isMonetaryAll(addr core.Address) bool {
	idx := int(addr)
	if idx < len(p.resources) {
		if c, ok := p.resources[idx].(program.Constant); ok {
			if mon, ok := c.Inner.(core.Monetary); ok {
				return mon.Amount.All
			}
		}
	}
	return false
}

func (p *parseVisitor) isWorld(addr core.Address) bool {
	idx := int(addr)
	if idx < len(p.resources) {
		if c, ok := p.resources[idx].(program.Constant); ok {
			if acc, ok := c.Inner.(core.Account); ok {
				if string(acc) == "world" {
					return true
				}
			}
		}
	}
	return false
}

func (p *parseVisitor) VisitVariable(c parser.IVariableContext, push bool) (core.Type, *core.Address, *CompileError) {
	name := c.GetText()[1:] // strip '$' prefix
	if idx, ok := p.var_idx[name]; ok {
		res := p.resources[idx]
		if push {
			p.instructions = append(p.instructions, program.OP_APUSH)
			bytes := idx.ToBytes()
			p.instructions = append(p.instructions, bytes...)
		}
		return res.GetType(), &idx, nil
	} else {
		return 0, nil, LogicError(c, errors.New("variable not declared"))
	}
}

func (p *parseVisitor) VisitExpr(c parser.IExpressionContext, push bool) (core.Type, *core.Address, *CompileError) {
	switch c := c.(type) {
	case *parser.ExprAddSubContext:
		ty, _, err := p.VisitExpr(c.GetLhs(), push)
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, LogicError(c, errors.New("tried to do arithmetic with wrong type"))
		}
		ty, _, err = p.VisitExpr(c.GetRhs(), push)
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, LogicError(c, errors.New("tried to do arithmetic with wrong type"))
		}
		if push {
			switch c.GetOp().GetTokenType() {
			case parser.NumScriptLexerOP_ADD:
				p.instructions = append(p.instructions, program.OP_IADD)
			case parser.NumScriptLexerOP_SUB:
				p.instructions = append(p.instructions, program.OP_ISUB)
			}
		}
		return core.TYPE_NUMBER, nil, nil
	case *parser.ExprLiteralContext:
		ty, addr, err := p.VisitLit(c.GetLit(), push)
		if err != nil {
			return 0, nil, err
		}
		return ty, addr, nil
	case *parser.ExprVariableContext:
		ty, addr, err := p.VisitVariable(c.GetVar_(), push)
		return ty, addr, err
	default:
		return 0, nil, InternalError(c)
	}
}

// pushes a value from a literal onto the stack
func (p *parseVisitor) VisitLit(c parser.ILiteralContext, push bool) (core.Type, *core.Address, *CompileError) {
	switch c := c.(type) {
	case *parser.LitAccountContext:
		account := core.Account(c.GetText()[1:])
		addr, err := p.AllocateResource(program.Constant{Inner: account})
		if err != nil {
			return 0, nil, LogicError(c, err)
		}
		if push {
			p.PushAddress(addr)
		}
		return core.TYPE_ACCOUNT, addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		addr, err := p.AllocateResource(program.Constant{Inner: asset})
		if err != nil {
			return 0, nil, LogicError(c, err)
		}
		return core.TYPE_ASSET, addr, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return 0, nil, LogicError(c, err)
		}
		number := core.Number(n)
		if push {
			p.PushInteger(number)
		}
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
				return 0, nil, LogicError(c, err)
			}
			amount = core.NewAmountSpecific(a)
		}
		monetary := core.Monetary{
			Asset:  core.Asset(asset),
			Amount: amount,
		}
		addr, err := p.AllocateResource(program.Constant{Inner: monetary})
		if err != nil {
			return 0, nil, LogicError(c, err)
		}
		if push {
			p.PushAddress(addr)
		}
		return core.TYPE_MONETARY, addr, nil
	default:
		return 0, nil, InternalError(c)
	}
}

// // Returns the resource addresses of all the accounts,
// // and true if the source is bottomless (contains @world)
// func (p *parseVisitor) VisitSource(c parser.ISourceContext) ([]core.Address, bool, *CompileError) {
// 	needed_accounts := []core.Address{}
// 	bottomless := false
// 	switch c := c.(type) {
// 	case *parser.SrcAccountContext:
// 		ty, addr, err := p.VisitExpr(c.Expression(), true)
// 		if err != nil {
// 			return nil, false, err
// 		}
// 		if ty != core.TYPE_ACCOUNT {
// 			return nil, false, LogicError(c, errors.New("wrong type: expected account or allocation as destination"))
// 		}
// 		bottomless = bottomless || p.isWorld(*addr)
// 		needed_accounts = append(needed_accounts, *addr)
// 		p.PushInteger(core.Number(1))
// 	case *parser.SrcBlockContext:
// 		sources := c.SourceBlock().GetSources()
// 		n := len(sources)
// 		for i := n - 1; i >= 0; i-- {
// 			ty, addr, err := p.VisitExpr(sources[i], true)
// 			if err != nil {
// 				return nil, false, err
// 			}
// 			if ty != core.TYPE_ACCOUNT {
// 				return nil, false, LogicError(c, errors.New("wrong type: expected only accounts in sources"))
// 			}
// 			bottomless = bottomless || p.isWorld(*addr)
// 			needed_accounts = append(needed_accounts, *addr)
// 		}
// 		p.PushInteger(core.Number(len(c.SourceBlock().GetSources())))
// 		p.instructions = append(p.instructions, program.OP_SOURCE)
// 	}
// 	return needed_accounts, bottomless, nil
// }

// send statement
func (p *parseVisitor) VisitSend(c *parser.SendContext) *CompileError {
	ty, mon_addr, err := p.VisitExpr(c.GetMon(), true)
	if err != nil {
		return err
	}
	if ty != core.TYPE_MONETARY {
		return LogicError(c, errors.New("wrong type for monetary value"))
	}
	needed_accounts, bottomless, err := p.VisitSource(c.GetSrc())
	if err != nil {
		return err
	}
	if p.isMonetaryAll(*mon_addr) && bottomless {
		return LogicError(c, errors.New("cannot take all balance of world"))
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
	err = p.VisitDestination(c.GetDest())
	if err != nil {
		return err
	}
	return nil
}

// print statement
func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) *CompileError {
	_, _, err := p.VisitExpr(ctx.GetExpr(), true)
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

// vars declaration block
func (p *parseVisitor) VisitVars(c *parser.VarListDeclContext) *CompileError {
	if len(c.GetV()) > 32768 {
		return LogicError(c, fmt.Errorf("number of variables exceeded %v", 32768))
	}
	for _, v := range c.GetV() {
		name := v.GetName().GetText()[1:]
		if _, ok := p.var_idx[name]; ok {
			return LogicError(c, errors.New("duplicate variable"))
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
			return InternalError(c)
		}

		var addr core.Address
		c_orig := v.GetOrig()
		if c_orig != nil {
			src_ty, src, cerr := p.VisitExpr(c_orig.GetAcc(), false)
			if cerr != nil {
				return cerr
			}
			if src_ty != core.TYPE_ACCOUNT {
				return LogicError(c_orig, errors.New("wrong type: expected account"))
			}
			key := strings.Trim(c_orig.GetKey().GetText(), `"`)
			a, err := p.AllocateResource(program.Metadata{SourceAccount: *src, Key: key, Typ: ty})
			if err != nil {
				return LogicError(c_orig, err)
			}
			addr = *a
		} else {
			a, err := p.AllocateResource(program.Parameter{Typ: ty, Name: name})
			if err != nil {
				return LogicError(c_orig, err)
			}
			addr = *a
		}
		p.var_idx[name] = addr
	}
	return nil
}

func (p *parseVisitor) VisitScript(c parser.IScriptContext) *CompileError {
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
				return InternalError(c)
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
				return InternalError(c)
			}
		}
	default:
		return InternalError(c)
	}
	return nil
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
		resources:       make([]program.Resource, 0),
		var_idx:         make(map[string]core.Address),
		needed_balances: make(map[core.Address]map[core.Address]struct{}),
	}

	err := visitor.VisitScript(tree)

	if err != nil {
		err := CompileErrorList{
			errors: []CompileError{*err},
			source: input,
		}
		return nil, &err
	}

	return &program.Program{
		Instructions:   visitor.instructions,
		Resources:      visitor.resources,
		NeededBalances: visitor.needed_balances,
	}, nil
}

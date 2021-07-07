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
	elistener    *ErrorListener
	instructions []byte
	constants    []core.Value /* must not exceed 32768 elements */
	variables    map[string]program.VarInfo
}

// Allocates constants if it hasn't already been,
// and returns its resource address.
func (p *parseVisitor) AllocateValue(v core.Value) (program.Address, error) {
	for i := 0; i < len(p.constants); i++ {
		if p.constants[i] == v {
			return program.Address(i), nil
		}
	}
	if len(p.constants) >= 32768 {
		return 0, errors.New("number of unique constants exceeded 32768")
	}
	p.constants = append(p.constants, v)
	return program.Address(len(p.constants) - 1), nil
}

func (p *parseVisitor) PushValue(val core.Value) error {
	switch val := val.(type) {
	case core.Account, core.Asset, core.Monetary:
		p.instructions = append(p.instructions, program.OP_APUSH)
		addr, err := p.AllocateValue(val)
		if err != nil {
			return err
		}
		bytes := addr.ToBytes()
		p.instructions = append(p.instructions, bytes...)
	case core.Number:
		p.instructions = append(p.instructions, program.OP_IPUSH)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(val))
		p.instructions = append(p.instructions, bytes...)
	default:
		panic("internal compiler error")
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
				p.VisitVars(c)
			default:
				panic("internal compiler error")
			}
		}
		for _, stmt := range c.GetStmts() {
			switch c := stmt.(type) {
			case *parser.PrintContext:
				err := p.VisitPrint(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
					return err
				}
			case *parser.FailContext:
				p.VisitFail(c)
			case *parser.SendContext:
				err := p.VisitSend(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
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

func (p *parseVisitor) VisitVars(c *parser.VarListDeclContext) error {
	if len(c.GetV()) > 32768 {
		return fmt.Errorf("number of variables exceeded %v", 32768)
	}
	for _, v := range c.GetV() {
		name := v.GetName().GetText()[1:]
		if _, ok := p.variables[name]; ok {
			return fmt.Errorf("duplicate variable: %v", name)
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
		}
		addr := program.NewVarAddress(uint16(len(p.variables)))
		p.variables[name] = program.VarInfo{
			Ty:   ty,
			Addr: addr,
		}
	}
	return nil
}

func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) error {
	_, err := p.VisitExpr(ctx.GetExpr())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

func (p *parseVisitor) VisitExpr(ctx parser.IExpressionContext) (core.Type, error) {
	switch ctx := ctx.(type) {
	case *parser.ExprAddSubContext:
		ty, err := p.VisitExpr(ctx.GetLhs())
		if err != nil {
			return 0, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, errors.New("tried to do arithmetic with wrong type")
		}
		ty, err = p.VisitExpr(ctx.GetRhs())
		if err != nil {
			return 0, nil
		}
		if ty != core.TYPE_NUMBER {
			return 0, errors.New("tried to do arithmetic with wrong type")
		}
		switch ctx.GetOp().GetTokenType() {
		case parser.NumScriptLexerOP_ADD:
			p.instructions = append(p.instructions, program.OP_IADD)
		case parser.NumScriptLexerOP_SUB:
			p.instructions = append(p.instructions, program.OP_ISUB)
		}
		return core.TYPE_NUMBER, nil
	case *parser.ExprLiteralContext:
		val, err := p.VisitLit(ctx.GetLit())
		if err != nil {
			return 0, err
		}
		err = p.PushValue(val)
		if err != nil {
			return 0, err
		}
		return val.GetType(), nil
	case *parser.ExprVariableContext:
		name := ctx.GetVariable().GetText()[1:]
		if info, ok := p.variables[name]; ok {
			p.instructions = append(p.instructions, program.OP_APUSH)
			bytes := info.Addr.ToBytes()
			p.instructions = append(p.instructions, bytes...)
			return info.Ty, nil
		} else {
			return 0, fmt.Errorf("variable not declared : %v", name)
		}
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) VisitLit(c parser.ILiteralContext) (core.Value, error) {
	switch c := c.(type) {
	case *parser.LitAccountContext:
		addr := core.Account(c.GetText()[1:])
		return addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		return asset, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		number := core.Number(n)
		return number, nil
	case *parser.LitMonetaryContext:
		asset := c.Monetary().GetAsset().GetText()
		amount, err := strconv.ParseUint(c.Monetary().GetAmount().GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		monetary := core.Monetary{
			Asset:  asset,
			Amount: amount,
		}
		return monetary, nil
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) VisitFail(ctx *parser.FailContext) {
	p.instructions = append(p.instructions, program.OP_FAIL)
}

func (p *parseVisitor) VisitSend(ctx *parser.SendContext) error {
	err := p.VisitArgs(ctx.GetArgs(), []arg{
		{
			name: "value",
			ty:   core.TYPE_MONETARY,
		},
		{
			name: "source",
			ty:   core.TYPE_ACCOUNT,
		},
		{
			name: "destination",
			ty:   core.TYPE_ACCOUNT,
		},
	})
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_SEND)
	return nil
}

type arg struct {
	name string
	ty   core.Type
}

// pushes them on the stack in the order given
func (p *parseVisitor) VisitArgs(cs []parser.IArgumentContext, args []arg) error {
	// put all parsed arguments in a map
	arg_expr := make(map[string]parser.IExpressionContext)
	for _, c := range cs {
		name := c.GetName().GetText()
		expr := c.GetVal()
		if _, ok := arg_expr[name]; ok {
			return fmt.Errorf("duplicate argument: %s", name)
		}
		arg_expr[name] = expr
	}
	// push arguments in right order
	for _, arg := range args {
		expr, ok := arg_expr[arg.name]
		if !ok {
			return fmt.Errorf("missing argument: %s", arg.name)
		}
		ty, err := p.VisitExpr(expr)
		if err != nil {
			return err
		}
		if ty != arg.ty {
			return fmt.Errorf("wrong argument type")
		}
		delete(arg_expr, arg.name)
	}
	// check for excess arguments
	for name := range arg_expr {
		return fmt.Errorf("unrecognized argument: %s", name)
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
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	visitor := parseVisitor{
		elistener:    elistener,
		instructions: make([]byte, 0),
		constants:    make([]core.Value, 0),
		variables:    make(map[string]program.VarInfo, 0),
	}

	_ = visitor.VisitScript(tree)

	if len(elistener.Errors) != 0 {
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	return &program.Program{
		Instructions: visitor.instructions,
		Constants:    visitor.constants,
		Variables:    visitor.variables,
	}, nil
}

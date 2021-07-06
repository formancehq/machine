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
	constants    []string /* must not exceed 32768 elements */
}

// Allocates constants if it hasn't already been,
// and returns its resource address.
func (p *parseVisitor) AllocateString(str string) (uint16, error) {
	for i := 0; i < len(p.constants); i++ {
		if p.constants[i] == str {
			return uint16(i), nil
		}
	}
	if len(p.constants) >= 32768 {
		return 0, errors.New("number of unique resource literals exceeded 32768")
	}
	p.constants = append(p.constants, str)
	return uint16(len(p.constants) - 1), nil
}

func (p *parseVisitor) PushValue(val core.Value) error {
	switch val := val.(type) {
	case *core.Address:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(string(*val))
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)
	case *core.Asset:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(string(*val))
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)
	case *core.Number:
		p.instructions = append(p.instructions, program.OP_PUSH8)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(*val))
		p.instructions = append(p.instructions, bytes...)
	case *core.Monetary:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(val.Asset)
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)

		p.instructions = append(p.instructions, program.OP_PUSH8)
		bytes = make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, val.Number)
		p.instructions = append(p.instructions, bytes...)
	default:
		panic("unreachable")
	}
	return nil
}

func (p *parseVisitor) VisitScript(c parser.IScriptContext) error {
	switch c := c.(type) {
	case *parser.ScriptContext:
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

func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) error {
	_, err := p.VisitExpr(ctx.GetExpr())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

func (p *parseVisitor) VisitExpr(ctx parser.IExpressionContext) (core.ValueType, error) {
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
		p.PushValue(val)
		return val.GetType(), nil
	default:
		panic("unreachable")
	}
}

func (p *parseVisitor) VisitLit(c parser.ILiteralContext) (core.Value, error) {
	switch c := c.(type) {
	case *parser.LitAddressContext:
		addr := core.Address(c.GetText())
		return &addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		return &asset, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		number := core.Number(n)
		return &number, nil
	case *parser.LitMonetaryContext:
		asset := c.Monetary().GetAsset().GetText()
		number, err := strconv.ParseUint(c.Monetary().GetNumber().GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		monetary := core.Monetary{
			Asset:  asset,
			Number: number,
		}
		return &monetary, nil
	default:
		panic("unreachable")
	}
}

func (p *parseVisitor) VisitFail(ctx *parser.FailContext) {
	p.instructions = append(p.instructions, program.OP_FAIL)
}

func (p *parseVisitor) VisitSend(ctx *parser.SendContext) error {
	args, err := p.VisitArgs(ctx.GetArgs(), map[string]core.ValueType{
		"source":      core.TYPE_ADDRESS,
		"destination": core.TYPE_ADDRESS,
		"monetary":    core.TYPE_MONETARY,
	})
	if err != nil {
		return err
	}
	mon := args["monetary"]
	src := args["source"]
	dst := args["destination"]
	p.PushValue(mon)
	p.PushValue(src)
	p.PushValue(dst)
	p.instructions = append(p.instructions, program.OP_SEND)
	return nil
}

func (p *parseVisitor) VisitArgs(cs []parser.IArgumentContext, args map[string]core.ValueType) (map[string]core.Value, error) {
	res := make(map[string]core.Value)
	for _, c := range cs {
		name := c.GetName().GetText()
		val, err := p.VisitLit(c.GetLit())
		if err != nil {
			return nil, err
		}
		val_ty := val.GetType()
		if _, ok := res[name]; ok {
			return nil, fmt.Errorf("duplicate argument: %s", name)
		}
		if ty, ok := args[name]; ok && ty == val_ty {
			delete(args, name)
		} else {
			return nil, fmt.Errorf("argument is not valid")
		}
		res[name] = val
	}
	for name := range args {
		return nil, fmt.Errorf("missing argument: %s", name)
	}
	return res, nil
}

type CompileError struct {
	line, column int
	msg          string
}

type CompileErrorList []CompileError

func (c *CompileErrorList) Error() string {
	s := ""
	for _, e := range *c {
		s = fmt.Sprintf("%v\nerror:%v:%v   %v", s, e.line, e.column, e.msg)
	}
	return s
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors CompileErrorList
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, CompileError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *ErrorListener) LogicError(token antlr.Token, err error) {
	l.Errors = append(l.Errors, CompileError{
		line:   token.GetLine(),
		column: token.GetColumn(),
		msg:    err.Error(),
	})
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
		constants:    make([]string, 0),
	}

	_ = visitor.VisitScript(tree)

	if len(elistener.Errors) != 0 {
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	return &program.Program{
		Instructions: visitor.instructions,
		Constants:    visitor.constants,
		Variables:    make([]string, 0),
	}, nil
}

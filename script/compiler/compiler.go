package compiler

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/numary/machine/script/parser"
	"github.com/numary/machine/vm/program"
)

type parseListener struct {
	*parser.BaseNumScriptListener
	program program.Program
}

func (p *parseListener) ExitAddSub(c *parser.AddSubContext) {
	switch c.GetOp().GetTokenType() {
	case parser.NumScriptParserOP_ADD:
		p.program = append(p.program, program.OP_IADD)
	case parser.NumScriptLexerOP_SUB:
		p.program = append(p.program, program.OP_ISUB)
	}
}

func (p *parseListener) ExitNumber(c *parser.NumberContext) {
	a, _ := strconv.Atoi(c.GetText())
	p.program = append(p.program, program.OP_IPUSH, byte(a))
}

func (p *parseListener) ExitScript(c *parser.ScriptContext) {
	p.program = append(p.program, program.OP_PRINT)
}

func Compile(input string) program.Program {
	listener := parseListener{
		program: program.Program{},
	}

	is := antlr.NewInputStream(input)
	lexer := parser.NewNumScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewNumScriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Script())

	return listener.program
}

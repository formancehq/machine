package compiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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

package compiler

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/logrusorgru/aurora"
)

type CompileError struct {
	line, column int
	len          int
	msg          string
}

type CompileErrorList struct {
	errors []CompileError
	source string
}

func (c *CompileErrorList) Error() string {
	source := strings.ReplaceAll(c.source, "\t", " ")
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")

	txt_bar := aurora.Blue("|")

	s := ""
	for _, e := range c.errors {
		s += fmt.Sprintf("%v error:%v:%v\n", aurora.Red("-->"), e.line, e.column)
		s += fmt.Sprintf("%v\n", txt_bar)
		s += fmt.Sprintf("%v %v\n", txt_bar, lines[e.line-1])
		s += fmt.Sprintf("%v %v%v\n", txt_bar, strings.Repeat(" ", e.column), aurora.Red(strings.Repeat("^", e.len)+" "+e.msg))
	}
	return s
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []CompileError
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	length := 1
	if token, ok := offendingSymbol.(antlr.Token); ok {
		length = len(token.GetText())
	}
	l.Errors = append(l.Errors, CompileError{
		line:   line,
		column: column,
		len:    length,
		msg:    msg,
	})
}

func (l *ErrorListener) LogicError(c antlr.ParserRuleContext, err error) error {
	l.Errors = append(l.Errors, CompileError{
		line:   c.GetStart().GetLine(),
		column: c.GetStart().GetColumn(),
		len:    len(c.GetText()),
		msg:    err.Error(),
	})
	return err
}

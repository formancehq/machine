package compiler

import (
	"fmt"
	"math"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/logrusorgru/aurora"
)

type CompileError struct {
	startl, startc int
	endl, endc     int
	msg            string
}

type CompileErrorList struct {
	errors []CompileError
	source string
}

func (c *CompileErrorList) Error() string {
	source := strings.ReplaceAll(c.source, "\t", " ")
	lines := strings.SplitAfter(strings.ReplaceAll(source, "\r\n", "\n"), "\n")
	lines[len(lines)-1] += "\n"

	txt_bar_good := aurora.Blue("|")

	s := ""
	for _, e := range c.errors {
		ln_pad := int(math.Log10(float64(e.endl))) + 1 // line number padding
		// error indicator
		s += fmt.Sprintf("%v error:%v:%v\n", aurora.Red("-->"), e.startl, e.startc)
		// initial empty line
		s += fmt.Sprintf("%v %v\n", strings.Repeat(" ", ln_pad), txt_bar_good)
		// offending lines
		for l := e.startl; l <= e.endl; l++ { // "print fail"
			line := lines[l-1]
			before := ""
			after := ""
			start := 0
			if l == e.startl {
				before = line[:e.startc]
				line = line[e.startc:]
				start = e.startc
			}
			if l == e.endl {
				idx := e.endc - start + 1
				if idx >= len(line) { // because newline was erased
					idx = len(line) - 1
				}
				after = line[idx:]
				line = line[:idx]
			}
			s += aurora.Red(fmt.Sprintf("%0*d | ", ln_pad, l)).String()
			s += fmt.Sprintf("%v%v%v", aurora.BrightBlack(before), line, aurora.BrightBlack(after))
		}
		// message
		start := strings.IndexFunc(lines[e.endl-1], func(r rune) bool {
			return r != ' '
		})
		span := e.endc - start + 1
		if e.startl == e.endl {
			start = e.startc
			span = e.endc - e.startc
		}
		if span == 0 {
			span = 1
		}
		s += fmt.Sprintf("%v %v %v%v %v\n", strings.Repeat(" ", ln_pad), txt_bar_good, strings.Repeat(" ", start), aurora.Red(strings.Repeat("^", span)), e.msg)
	}
	return s
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []CompileError
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, startl, startc int, msg string, e antlr.RecognitionException) {
	length := 1
	if token, ok := offendingSymbol.(antlr.Token); ok {
		length = len(token.GetText())
	}
	endl := startl
	endc := startc + length - 1 // -1 so that end caracter is inside the offending token
	l.Errors = append(l.Errors, CompileError{
		startl: startl,
		startc: startc,
		endl:   endl,
		endc:   endc,
		msg:    msg,
	})
}

func LogicError(c antlr.ParserRuleContext, err error) *CompileError {
	endc := c.GetStop().GetColumn() + len(c.GetStop().GetText())
	// fmt.Println(c.GetStart().GetLine(),
	// 	c.GetStart().GetColumn(),
	// 	c.GetStop().GetLine(),
	// 	endc)
	return &CompileError{
		startl: c.GetStart().GetLine(),
		startc: c.GetStart().GetColumn(),
		endl:   c.GetStop().GetLine(),
		endc:   endc,
		msg:    err.Error(),
	}
}

const INTERNAL_ERROR_MSG = "internal compiler error, please report to the issue tracker"

func InternalError(c antlr.ParserRuleContext) *CompileError {
	return &CompileError{
		startl: c.GetStart().GetLine(),
		startc: c.GetStart().GetColumn(),
		endl:   c.GetStop().GetLine(),
		endc:   c.GetStop().GetColumn(),
		msg:    INTERNAL_ERROR_MSG,
	}
}

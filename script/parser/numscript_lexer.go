// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 66, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 6, 2, 25,
	10, 2, 13, 2, 14, 2, 26, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 6, 5, 40, 10, 5, 13, 5, 14, 5, 41, 3, 6, 6, 6, 45, 10,
	6, 13, 6, 14, 6, 46, 3, 7, 6, 7, 50, 10, 7, 13, 7, 14, 7, 51, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 61, 10, 11, 13, 11, 14, 11,
	62, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 3, 2, 7, 4, 2, 12, 12, 15, 15, 3, 2, 99, 124, 3,
	2, 67, 92, 3, 2, 50, 59, 4, 2, 11, 12, 34, 34, 2, 70, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 3, 24, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7, 33, 3, 2,
	2, 2, 9, 39, 3, 2, 2, 2, 11, 44, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 53,
	3, 2, 2, 2, 17, 55, 3, 2, 2, 2, 19, 57, 3, 2, 2, 2, 21, 60, 3, 2, 2, 2,
	23, 25, 9, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 24, 3,
	2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 4, 3, 2, 2, 2, 28, 29, 7, 101, 2, 2, 29,
	30, 7, 99, 2, 2, 30, 31, 7, 110, 2, 2, 31, 32, 7, 101, 2, 2, 32, 6, 3,
	2, 2, 2, 33, 34, 7, 104, 2, 2, 34, 35, 7, 99, 2, 2, 35, 36, 7, 107, 2,
	2, 36, 37, 7, 110, 2, 2, 37, 8, 3, 2, 2, 2, 38, 40, 9, 3, 2, 2, 39, 38,
	3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2,
	42, 10, 3, 2, 2, 2, 43, 45, 9, 4, 2, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3,
	2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 12, 3, 2, 2, 2, 48,
	50, 9, 5, 2, 2, 49, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 49, 3, 2, 2,
	2, 51, 52, 3, 2, 2, 2, 52, 14, 3, 2, 2, 2, 53, 54, 7, 45, 2, 2, 54, 16,
	3, 2, 2, 2, 55, 56, 7, 47, 2, 2, 56, 18, 3, 2, 2, 2, 57, 58, 7, 61, 2,
	2, 58, 20, 3, 2, 2, 2, 59, 61, 9, 6, 2, 2, 60, 59, 3, 2, 2, 2, 61, 62,
	3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2,
	64, 65, 8, 11, 2, 2, 65, 22, 3, 2, 2, 2, 8, 2, 26, 41, 46, 51, 62, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'calc'", "'fail'", "", "", "", "'+'", "'-'", "';'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "CALC", "FAIL", "IDENTIFIER", "ASSET", "NUMBER", "OP_ADD",
	"OP_SUB", "SEP", "WHITESPACE",
}

var lexerRuleNames = []string{
	"NEWLINE", "CALC", "FAIL", "IDENTIFIER", "ASSET", "NUMBER", "OP_ADD", "OP_SUB",
	"SEP", "WHITESPACE",
}

type NumScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewNumScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *NumScriptLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNumScriptLexer(input antlr.CharStream) *NumScriptLexer {
	l := new(NumScriptLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NumScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumScriptLexer tokens.
const (
	NumScriptLexerNEWLINE    = 1
	NumScriptLexerCALC       = 2
	NumScriptLexerFAIL       = 3
	NumScriptLexerIDENTIFIER = 4
	NumScriptLexerASSET      = 5
	NumScriptLexerNUMBER     = 6
	NumScriptLexerOP_ADD     = 7
	NumScriptLexerOP_SUB     = 8
	NumScriptLexerSEP        = 9
	NumScriptLexerWHITESPACE = 10
)

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 98, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 3, 2, 3, 2, 3, 3, 6, 3, 41, 10, 3, 13, 3, 14, 3, 42, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 76, 10, 14, 13, 14, 14, 14,
	77, 3, 15, 6, 15, 81, 10, 15, 13, 15, 14, 15, 82, 3, 16, 6, 16, 86, 10,
	16, 13, 16, 14, 16, 87, 3, 17, 3, 17, 3, 18, 6, 18, 93, 10, 18, 13, 18,
	14, 18, 94, 3, 18, 3, 18, 2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 3, 2, 7, 4, 2, 12, 12, 15, 15, 5, 2, 60, 60, 97, 97, 99,
	124, 3, 2, 50, 59, 4, 2, 49, 59, 67, 92, 4, 2, 11, 12, 34, 34, 2, 102,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 40, 3, 2, 2, 2, 7, 44,
	3, 2, 2, 2, 9, 50, 3, 2, 2, 2, 11, 55, 3, 2, 2, 2, 13, 60, 3, 2, 2, 2,
	15, 62, 3, 2, 2, 2, 17, 64, 3, 2, 2, 2, 19, 66, 3, 2, 2, 2, 21, 68, 3,
	2, 2, 2, 23, 70, 3, 2, 2, 2, 25, 72, 3, 2, 2, 2, 27, 75, 3, 2, 2, 2, 29,
	80, 3, 2, 2, 2, 31, 85, 3, 2, 2, 2, 33, 89, 3, 2, 2, 2, 35, 92, 3, 2, 2,
	2, 37, 38, 7, 46, 2, 2, 38, 4, 3, 2, 2, 2, 39, 41, 9, 2, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 6, 3, 2, 2, 2, 44, 45, 7, 114, 2, 2, 45, 46, 7, 116, 2, 2, 46, 47,
	7, 107, 2, 2, 47, 48, 7, 112, 2, 2, 48, 49, 7, 118, 2, 2, 49, 8, 3, 2,
	2, 2, 50, 51, 7, 104, 2, 2, 51, 52, 7, 99, 2, 2, 52, 53, 7, 107, 2, 2,
	53, 54, 7, 110, 2, 2, 54, 10, 3, 2, 2, 2, 55, 56, 7, 117, 2, 2, 56, 57,
	7, 103, 2, 2, 57, 58, 7, 112, 2, 2, 58, 59, 7, 102, 2, 2, 59, 12, 3, 2,
	2, 2, 60, 61, 7, 45, 2, 2, 61, 14, 3, 2, 2, 2, 62, 63, 7, 47, 2, 2, 63,
	16, 3, 2, 2, 2, 64, 65, 7, 42, 2, 2, 65, 18, 3, 2, 2, 2, 66, 67, 7, 43,
	2, 2, 67, 20, 3, 2, 2, 2, 68, 69, 7, 93, 2, 2, 69, 22, 3, 2, 2, 2, 70,
	71, 7, 95, 2, 2, 71, 24, 3, 2, 2, 2, 72, 73, 7, 63, 2, 2, 73, 26, 3, 2,
	2, 2, 74, 76, 9, 3, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75,
	3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 28, 3, 2, 2, 2, 79, 81, 9, 4, 2, 2,
	80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3,
	2, 2, 2, 83, 30, 3, 2, 2, 2, 84, 86, 9, 5, 2, 2, 85, 84, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 32, 3, 2, 2,
	2, 89, 90, 7, 61, 2, 2, 90, 34, 3, 2, 2, 2, 91, 93, 9, 6, 2, 2, 92, 91,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 96, 3, 2, 2, 2, 96, 97, 8, 18, 2, 2, 97, 36, 3, 2, 2, 2, 8, 2, 42,
	77, 82, 87, 94, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "'print'", "'fail'", "'send'", "'+'", "'-'", "'('", "')'",
	"'['", "']'", "'='", "", "", "", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "EQ", "IDENTIFIER", "NUMBER", "ASSET", "SEP",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "EQ", "IDENTIFIER", "NUMBER", "ASSET", "SEP",
	"WHITESPACE",
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
	NumScriptLexerT__0       = 1
	NumScriptLexerNEWLINE    = 2
	NumScriptLexerPRINT      = 3
	NumScriptLexerFAIL       = 4
	NumScriptLexerSEND       = 5
	NumScriptLexerOP_ADD     = 6
	NumScriptLexerOP_SUB     = 7
	NumScriptLexerLPAREN     = 8
	NumScriptLexerRPAREN     = 9
	NumScriptLexerLBRACK     = 10
	NumScriptLexerRBRACK     = 11
	NumScriptLexerEQ         = 12
	NumScriptLexerIDENTIFIER = 13
	NumScriptLexerNUMBER     = 14
	NumScriptLexerASSET      = 15
	NumScriptLexerSEP        = 16
	NumScriptLexerWHITESPACE = 17
)

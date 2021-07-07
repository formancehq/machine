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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 167,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 3, 6, 3, 59, 10, 3, 13, 3, 14, 3, 60, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	6, 21, 133, 10, 21, 13, 21, 14, 21, 134, 3, 22, 6, 22, 138, 10, 22, 13,
	22, 14, 22, 139, 3, 23, 3, 23, 6, 23, 144, 10, 23, 13, 23, 14, 23, 145,
	3, 24, 3, 24, 6, 24, 150, 10, 24, 13, 24, 14, 24, 151, 3, 25, 6, 25, 155,
	10, 25, 13, 25, 14, 25, 156, 3, 26, 3, 26, 3, 27, 6, 27, 162, 10, 27, 13,
	27, 14, 27, 163, 3, 27, 3, 27, 2, 2, 28, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 3, 2, 7, 4, 2, 12, 12, 15, 15, 3, 2, 50, 59, 5, 2,
	50, 60, 97, 97, 99, 124, 4, 2, 49, 59, 67, 92, 4, 2, 11, 11, 34, 34, 2,
	173, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 3, 55, 3, 2,
	2, 2, 5, 58, 3, 2, 2, 2, 7, 62, 3, 2, 2, 2, 9, 67, 3, 2, 2, 2, 11, 73,
	3, 2, 2, 2, 13, 78, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2,
	19, 87, 3, 2, 2, 2, 21, 89, 3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 93, 3,
	2, 2, 2, 27, 95, 3, 2, 2, 2, 29, 97, 3, 2, 2, 2, 31, 99, 3, 2, 2, 2, 33,
	101, 3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37, 115, 3, 2, 2, 2, 39, 122, 3,
	2, 2, 2, 41, 132, 3, 2, 2, 2, 43, 137, 3, 2, 2, 2, 45, 141, 3, 2, 2, 2,
	47, 147, 3, 2, 2, 2, 49, 154, 3, 2, 2, 2, 51, 158, 3, 2, 2, 2, 53, 161,
	3, 2, 2, 2, 55, 56, 7, 46, 2, 2, 56, 4, 3, 2, 2, 2, 57, 59, 9, 2, 2, 2,
	58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3,
	2, 2, 2, 61, 6, 3, 2, 2, 2, 62, 63, 7, 120, 2, 2, 63, 64, 7, 99, 2, 2,
	64, 65, 7, 116, 2, 2, 65, 66, 7, 117, 2, 2, 66, 8, 3, 2, 2, 2, 67, 68,
	7, 114, 2, 2, 68, 69, 7, 116, 2, 2, 69, 70, 7, 107, 2, 2, 70, 71, 7, 112,
	2, 2, 71, 72, 7, 118, 2, 2, 72, 10, 3, 2, 2, 2, 73, 74, 7, 104, 2, 2, 74,
	75, 7, 99, 2, 2, 75, 76, 7, 107, 2, 2, 76, 77, 7, 110, 2, 2, 77, 12, 3,
	2, 2, 2, 78, 79, 7, 117, 2, 2, 79, 80, 7, 103, 2, 2, 80, 81, 7, 112, 2,
	2, 81, 82, 7, 102, 2, 2, 82, 14, 3, 2, 2, 2, 83, 84, 7, 45, 2, 2, 84, 16,
	3, 2, 2, 2, 85, 86, 7, 47, 2, 2, 86, 18, 3, 2, 2, 2, 87, 88, 7, 42, 2,
	2, 88, 20, 3, 2, 2, 2, 89, 90, 7, 43, 2, 2, 90, 22, 3, 2, 2, 2, 91, 92,
	7, 93, 2, 2, 92, 24, 3, 2, 2, 2, 93, 94, 7, 95, 2, 2, 94, 26, 3, 2, 2,
	2, 95, 96, 7, 125, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98, 7, 127, 2, 2, 98,
	30, 3, 2, 2, 2, 99, 100, 7, 63, 2, 2, 100, 32, 3, 2, 2, 2, 101, 102, 7,
	99, 2, 2, 102, 103, 7, 101, 2, 2, 103, 104, 7, 101, 2, 2, 104, 105, 7,
	113, 2, 2, 105, 106, 7, 119, 2, 2, 106, 107, 7, 112, 2, 2, 107, 108, 7,
	118, 2, 2, 108, 34, 3, 2, 2, 2, 109, 110, 7, 99, 2, 2, 110, 111, 7, 117,
	2, 2, 111, 112, 7, 117, 2, 2, 112, 113, 7, 103, 2, 2, 113, 114, 7, 118,
	2, 2, 114, 36, 3, 2, 2, 2, 115, 116, 7, 112, 2, 2, 116, 117, 7, 119, 2,
	2, 117, 118, 7, 111, 2, 2, 118, 119, 7, 100, 2, 2, 119, 120, 7, 103, 2,
	2, 120, 121, 7, 116, 2, 2, 121, 38, 3, 2, 2, 2, 122, 123, 7, 111, 2, 2,
	123, 124, 7, 113, 2, 2, 124, 125, 7, 112, 2, 2, 125, 126, 7, 103, 2, 2,
	126, 127, 7, 118, 2, 2, 127, 128, 7, 99, 2, 2, 128, 129, 7, 116, 2, 2,
	129, 130, 7, 123, 2, 2, 130, 40, 3, 2, 2, 2, 131, 133, 9, 3, 2, 2, 132,
	131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135,
	3, 2, 2, 2, 135, 42, 3, 2, 2, 2, 136, 138, 9, 4, 2, 2, 137, 136, 3, 2,
	2, 2, 138, 139, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2,
	140, 44, 3, 2, 2, 2, 141, 143, 7, 38, 2, 2, 142, 144, 9, 4, 2, 2, 143,
	142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146,
	3, 2, 2, 2, 146, 46, 3, 2, 2, 2, 147, 149, 7, 66, 2, 2, 148, 150, 9, 4,
	2, 2, 149, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2,
	151, 152, 3, 2, 2, 2, 152, 48, 3, 2, 2, 2, 153, 155, 9, 5, 2, 2, 154, 153,
	3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2,
	2, 2, 157, 50, 3, 2, 2, 2, 158, 159, 7, 61, 2, 2, 159, 52, 3, 2, 2, 2,
	160, 162, 9, 6, 2, 2, 161, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163,
	161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166,
	8, 27, 2, 2, 166, 54, 3, 2, 2, 2, 10, 2, 60, 134, 139, 145, 151, 156, 163,
	3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "'vars'", "'print'", "'fail'", "'send'", "'+'", "'-'", "'('",
	"')'", "'['", "']'", "'{'", "'}'", "'='", "'account'", "'asset'", "'number'",
	"'monetary'", "", "", "", "", "", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT",
	"TY_ASSET", "TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "SEP", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT",
	"TY_ASSET", "TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "SEP", "WHITESPACE",
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
	NumScriptLexerT__0          = 1
	NumScriptLexerNEWLINE       = 2
	NumScriptLexerVARS          = 3
	NumScriptLexerPRINT         = 4
	NumScriptLexerFAIL          = 5
	NumScriptLexerSEND          = 6
	NumScriptLexerOP_ADD        = 7
	NumScriptLexerOP_SUB        = 8
	NumScriptLexerLPAREN        = 9
	NumScriptLexerRPAREN        = 10
	NumScriptLexerLBRACK        = 11
	NumScriptLexerRBRACK        = 12
	NumScriptLexerLBRACE        = 13
	NumScriptLexerRBRACE        = 14
	NumScriptLexerEQ            = 15
	NumScriptLexerTY_ACCOUNT    = 16
	NumScriptLexerTY_ASSET      = 17
	NumScriptLexerTY_NUMBER     = 18
	NumScriptLexerTY_MONETARY   = 19
	NumScriptLexerNUMBER        = 20
	NumScriptLexerIDENTIFIER    = 21
	NumScriptLexerVARIABLE_NAME = 22
	NumScriptLexerACCOUNT       = 23
	NumScriptLexerASSET         = 24
	NumScriptLexerSEP           = 25
	NumScriptLexerWHITESPACE    = 26
)

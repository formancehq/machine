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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 170,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 6, 4, 59, 10, 4, 13, 4, 14, 4, 60, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 6, 22, 135, 10, 22, 13, 22, 14, 22, 136, 3, 23, 6, 23, 140,
	10, 23, 13, 23, 14, 23, 141, 3, 24, 3, 24, 6, 24, 146, 10, 24, 13, 24,
	14, 24, 147, 3, 24, 6, 24, 151, 10, 24, 13, 24, 14, 24, 152, 3, 25, 3,
	25, 6, 25, 157, 10, 25, 13, 25, 14, 25, 158, 3, 25, 6, 25, 162, 10, 25,
	13, 25, 14, 25, 163, 3, 26, 6, 26, 167, 10, 26, 13, 26, 14, 26, 168, 2,
	2, 27, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 3, 2, 9, 4, 2, 12,
	12, 15, 15, 4, 2, 11, 11, 34, 34, 3, 2, 50, 59, 5, 2, 50, 60, 97, 97, 99,
	124, 4, 2, 97, 97, 99, 124, 5, 2, 50, 59, 97, 97, 99, 124, 4, 2, 49, 59,
	67, 92, 2, 177, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 53, 3, 2, 2, 2, 5,
	55, 3, 2, 2, 2, 7, 58, 3, 2, 2, 2, 9, 64, 3, 2, 2, 2, 11, 69, 3, 2, 2,
	2, 13, 75, 3, 2, 2, 2, 15, 80, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2, 19, 87,
	3, 2, 2, 2, 21, 89, 3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 93, 3, 2, 2, 2,
	27, 95, 3, 2, 2, 2, 29, 97, 3, 2, 2, 2, 31, 99, 3, 2, 2, 2, 33, 101, 3,
	2, 2, 2, 35, 103, 3, 2, 2, 2, 37, 111, 3, 2, 2, 2, 39, 117, 3, 2, 2, 2,
	41, 124, 3, 2, 2, 2, 43, 134, 3, 2, 2, 2, 45, 139, 3, 2, 2, 2, 47, 143,
	3, 2, 2, 2, 49, 154, 3, 2, 2, 2, 51, 166, 3, 2, 2, 2, 53, 54, 7, 46, 2,
	2, 54, 4, 3, 2, 2, 2, 55, 56, 9, 2, 2, 2, 56, 6, 3, 2, 2, 2, 57, 59, 9,
	3, 2, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60,
	61, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 8, 4, 2, 2, 63, 8, 3, 2, 2,
	2, 64, 65, 7, 120, 2, 2, 65, 66, 7, 99, 2, 2, 66, 67, 7, 116, 2, 2, 67,
	68, 7, 117, 2, 2, 68, 10, 3, 2, 2, 2, 69, 70, 7, 114, 2, 2, 70, 71, 7,
	116, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73, 7, 112, 2, 2, 73, 74, 7, 118,
	2, 2, 74, 12, 3, 2, 2, 2, 75, 76, 7, 104, 2, 2, 76, 77, 7, 99, 2, 2, 77,
	78, 7, 107, 2, 2, 78, 79, 7, 110, 2, 2, 79, 14, 3, 2, 2, 2, 80, 81, 7,
	117, 2, 2, 81, 82, 7, 103, 2, 2, 82, 83, 7, 112, 2, 2, 83, 84, 7, 102,
	2, 2, 84, 16, 3, 2, 2, 2, 85, 86, 7, 45, 2, 2, 86, 18, 3, 2, 2, 2, 87,
	88, 7, 47, 2, 2, 88, 20, 3, 2, 2, 2, 89, 90, 7, 42, 2, 2, 90, 22, 3, 2,
	2, 2, 91, 92, 7, 43, 2, 2, 92, 24, 3, 2, 2, 2, 93, 94, 7, 93, 2, 2, 94,
	26, 3, 2, 2, 2, 95, 96, 7, 95, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98, 7, 125,
	2, 2, 98, 30, 3, 2, 2, 2, 99, 100, 7, 127, 2, 2, 100, 32, 3, 2, 2, 2, 101,
	102, 7, 63, 2, 2, 102, 34, 3, 2, 2, 2, 103, 104, 7, 99, 2, 2, 104, 105,
	7, 101, 2, 2, 105, 106, 7, 101, 2, 2, 106, 107, 7, 113, 2, 2, 107, 108,
	7, 119, 2, 2, 108, 109, 7, 112, 2, 2, 109, 110, 7, 118, 2, 2, 110, 36,
	3, 2, 2, 2, 111, 112, 7, 99, 2, 2, 112, 113, 7, 117, 2, 2, 113, 114, 7,
	117, 2, 2, 114, 115, 7, 103, 2, 2, 115, 116, 7, 118, 2, 2, 116, 38, 3,
	2, 2, 2, 117, 118, 7, 112, 2, 2, 118, 119, 7, 119, 2, 2, 119, 120, 7, 111,
	2, 2, 120, 121, 7, 100, 2, 2, 121, 122, 7, 103, 2, 2, 122, 123, 7, 116,
	2, 2, 123, 40, 3, 2, 2, 2, 124, 125, 7, 111, 2, 2, 125, 126, 7, 113, 2,
	2, 126, 127, 7, 112, 2, 2, 127, 128, 7, 103, 2, 2, 128, 129, 7, 118, 2,
	2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 116, 2, 2, 131, 132, 7, 123, 2,
	2, 132, 42, 3, 2, 2, 2, 133, 135, 9, 4, 2, 2, 134, 133, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 44, 3,
	2, 2, 2, 138, 140, 9, 5, 2, 2, 139, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2,
	2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 46, 3, 2, 2, 2, 143,
	145, 7, 38, 2, 2, 144, 146, 9, 6, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147,
	3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 150, 3, 2,
	2, 2, 149, 151, 9, 7, 2, 2, 150, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 48, 3, 2, 2, 2, 154, 156,
	7, 66, 2, 2, 155, 157, 9, 6, 2, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 3, 2, 2, 2,
	160, 162, 9, 5, 2, 2, 161, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163,
	161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 50, 3, 2, 2, 2, 165, 167, 9,
	8, 2, 2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 166, 3, 2, 2,
	2, 168, 169, 3, 2, 2, 2, 169, 52, 3, 2, 2, 2, 11, 2, 60, 136, 141, 147,
	152, 158, 163, 168, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "", "'vars'", "'print'", "'fail'", "'send'", "'+'", "'-'",
	"'('", "')'", "'['", "']'", "'{'", "'}'", "'='", "'account'", "'asset'",
	"'number'", "'monetary'",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ",
	"TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER",
	"VARIABLE_NAME", "ACCOUNT", "ASSET",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ",
	"TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER",
	"VARIABLE_NAME", "ACCOUNT", "ASSET",
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
	NumScriptLexerWHITESPACE    = 3
	NumScriptLexerVARS          = 4
	NumScriptLexerPRINT         = 5
	NumScriptLexerFAIL          = 6
	NumScriptLexerSEND          = 7
	NumScriptLexerOP_ADD        = 8
	NumScriptLexerOP_SUB        = 9
	NumScriptLexerLPAREN        = 10
	NumScriptLexerRPAREN        = 11
	NumScriptLexerLBRACK        = 12
	NumScriptLexerRBRACK        = 13
	NumScriptLexerLBRACE        = 14
	NumScriptLexerRBRACE        = 15
	NumScriptLexerEQ            = 16
	NumScriptLexerTY_ACCOUNT    = 17
	NumScriptLexerTY_ASSET      = 18
	NumScriptLexerTY_NUMBER     = 19
	NumScriptLexerTY_MONETARY   = 20
	NumScriptLexerNUMBER        = 21
	NumScriptLexerIDENTIFIER    = 22
	NumScriptLexerVARIABLE_NAME = 23
	NumScriptLexerACCOUNT       = 24
	NumScriptLexerASSET         = 25
)

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 32, 222,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 6, 4, 70, 10, 4, 13, 4, 14, 4, 71, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 6, 25, 174,
	10, 25, 13, 25, 14, 25, 175, 3, 26, 6, 26, 179, 10, 26, 13, 26, 14, 26,
	180, 3, 26, 3, 26, 6, 26, 185, 10, 26, 13, 26, 14, 26, 186, 3, 27, 3, 27,
	3, 28, 6, 28, 192, 10, 28, 13, 28, 14, 28, 193, 3, 29, 3, 29, 6, 29, 198,
	10, 29, 13, 29, 14, 29, 199, 3, 29, 6, 29, 203, 10, 29, 13, 29, 14, 29,
	204, 3, 30, 3, 30, 6, 30, 209, 10, 30, 13, 30, 14, 30, 210, 3, 30, 6, 30,
	214, 10, 30, 13, 30, 14, 30, 215, 3, 31, 6, 31, 219, 10, 31, 13, 31, 14,
	31, 220, 2, 2, 32, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28,
	55, 29, 57, 30, 59, 31, 61, 32, 3, 2, 9, 4, 2, 12, 12, 15, 15, 4, 2, 11,
	11, 34, 34, 3, 2, 50, 59, 5, 2, 50, 60, 97, 97, 99, 124, 4, 2, 97, 97,
	99, 124, 5, 2, 50, 59, 97, 97, 99, 124, 4, 2, 49, 59, 67, 92, 2, 231, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49,
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2,
	57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 3, 63, 3, 2, 2, 2,
	5, 66, 3, 2, 2, 2, 7, 69, 3, 2, 2, 2, 9, 75, 3, 2, 2, 2, 11, 80, 3, 2,
	2, 2, 13, 86, 3, 2, 2, 2, 15, 91, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19, 103,
	3, 2, 2, 2, 21, 115, 3, 2, 2, 2, 23, 124, 3, 2, 2, 2, 25, 126, 3, 2, 2,
	2, 27, 128, 3, 2, 2, 2, 29, 130, 3, 2, 2, 2, 31, 132, 3, 2, 2, 2, 33, 134,
	3, 2, 2, 2, 35, 136, 3, 2, 2, 2, 37, 138, 3, 2, 2, 2, 39, 140, 3, 2, 2,
	2, 41, 142, 3, 2, 2, 2, 43, 150, 3, 2, 2, 2, 45, 156, 3, 2, 2, 2, 47, 163,
	3, 2, 2, 2, 49, 173, 3, 2, 2, 2, 51, 178, 3, 2, 2, 2, 53, 188, 3, 2, 2,
	2, 55, 191, 3, 2, 2, 2, 57, 195, 3, 2, 2, 2, 59, 206, 3, 2, 2, 2, 61, 218,
	3, 2, 2, 2, 63, 64, 7, 118, 2, 2, 64, 65, 7, 113, 2, 2, 65, 4, 3, 2, 2,
	2, 66, 67, 9, 2, 2, 2, 67, 6, 3, 2, 2, 2, 68, 70, 9, 3, 2, 2, 69, 68, 3,
	2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72,
	73, 3, 2, 2, 2, 73, 74, 8, 4, 2, 2, 74, 8, 3, 2, 2, 2, 75, 76, 7, 120,
	2, 2, 76, 77, 7, 99, 2, 2, 77, 78, 7, 116, 2, 2, 78, 79, 7, 117, 2, 2,
	79, 10, 3, 2, 2, 2, 80, 81, 7, 114, 2, 2, 81, 82, 7, 116, 2, 2, 82, 83,
	7, 107, 2, 2, 83, 84, 7, 112, 2, 2, 84, 85, 7, 118, 2, 2, 85, 12, 3, 2,
	2, 2, 86, 87, 7, 104, 2, 2, 87, 88, 7, 99, 2, 2, 88, 89, 7, 107, 2, 2,
	89, 90, 7, 110, 2, 2, 90, 14, 3, 2, 2, 2, 91, 92, 7, 117, 2, 2, 92, 93,
	7, 103, 2, 2, 93, 94, 7, 112, 2, 2, 94, 95, 7, 102, 2, 2, 95, 16, 3, 2,
	2, 2, 96, 97, 7, 117, 2, 2, 97, 98, 7, 113, 2, 2, 98, 99, 7, 119, 2, 2,
	99, 100, 7, 116, 2, 2, 100, 101, 7, 101, 2, 2, 101, 102, 7, 103, 2, 2,
	102, 18, 3, 2, 2, 2, 103, 104, 7, 102, 2, 2, 104, 105, 7, 103, 2, 2, 105,
	106, 7, 117, 2, 2, 106, 107, 7, 118, 2, 2, 107, 108, 7, 107, 2, 2, 108,
	109, 7, 112, 2, 2, 109, 110, 7, 99, 2, 2, 110, 111, 7, 118, 2, 2, 111,
	112, 7, 107, 2, 2, 112, 113, 7, 113, 2, 2, 113, 114, 7, 112, 2, 2, 114,
	20, 3, 2, 2, 2, 115, 116, 7, 99, 2, 2, 116, 117, 7, 110, 2, 2, 117, 118,
	7, 110, 2, 2, 118, 119, 7, 113, 2, 2, 119, 120, 7, 101, 2, 2, 120, 121,
	7, 99, 2, 2, 121, 122, 7, 118, 2, 2, 122, 123, 7, 103, 2, 2, 123, 22, 3,
	2, 2, 2, 124, 125, 7, 45, 2, 2, 125, 24, 3, 2, 2, 2, 126, 127, 7, 47, 2,
	2, 127, 26, 3, 2, 2, 2, 128, 129, 7, 42, 2, 2, 129, 28, 3, 2, 2, 2, 130,
	131, 7, 43, 2, 2, 131, 30, 3, 2, 2, 2, 132, 133, 7, 93, 2, 2, 133, 32,
	3, 2, 2, 2, 134, 135, 7, 95, 2, 2, 135, 34, 3, 2, 2, 2, 136, 137, 7, 125,
	2, 2, 137, 36, 3, 2, 2, 2, 138, 139, 7, 127, 2, 2, 139, 38, 3, 2, 2, 2,
	140, 141, 7, 63, 2, 2, 141, 40, 3, 2, 2, 2, 142, 143, 7, 99, 2, 2, 143,
	144, 7, 101, 2, 2, 144, 145, 7, 101, 2, 2, 145, 146, 7, 113, 2, 2, 146,
	147, 7, 119, 2, 2, 147, 148, 7, 112, 2, 2, 148, 149, 7, 118, 2, 2, 149,
	42, 3, 2, 2, 2, 150, 151, 7, 99, 2, 2, 151, 152, 7, 117, 2, 2, 152, 153,
	7, 117, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 118, 2, 2, 155, 44,
	3, 2, 2, 2, 156, 157, 7, 112, 2, 2, 157, 158, 7, 119, 2, 2, 158, 159, 7,
	111, 2, 2, 159, 160, 7, 100, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162, 7,
	116, 2, 2, 162, 46, 3, 2, 2, 2, 163, 164, 7, 111, 2, 2, 164, 165, 7, 113,
	2, 2, 165, 166, 7, 112, 2, 2, 166, 167, 7, 103, 2, 2, 167, 168, 7, 118,
	2, 2, 168, 169, 7, 99, 2, 2, 169, 170, 7, 116, 2, 2, 170, 171, 7, 123,
	2, 2, 171, 48, 3, 2, 2, 2, 172, 174, 9, 4, 2, 2, 173, 172, 3, 2, 2, 2,
	174, 175, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	50, 3, 2, 2, 2, 177, 179, 9, 4, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3,
	2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2,
	2, 182, 184, 7, 49, 2, 2, 183, 185, 9, 4, 2, 2, 184, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 52, 3,
	2, 2, 2, 188, 189, 7, 39, 2, 2, 189, 54, 3, 2, 2, 2, 190, 192, 9, 5, 2,
	2, 191, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 56, 3, 2, 2, 2, 195, 197, 7, 38, 2, 2, 196, 198,
	9, 6, 2, 2, 197, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 197, 3, 2,
	2, 2, 199, 200, 3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201, 203, 9, 7, 2, 2,
	202, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 58, 3, 2, 2, 2, 206, 208, 7, 66, 2, 2, 207, 209,
	9, 6, 2, 2, 208, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 208, 3, 2,
	2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2, 2, 2, 212, 214, 9, 5, 2, 2,
	213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 60, 3, 2, 2, 2, 217, 219, 9, 8, 2, 2, 218, 217, 3,
	2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2,
	2, 221, 62, 3, 2, 2, 2, 13, 2, 71, 175, 180, 186, 193, 199, 204, 210, 215,
	220, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'to'", "", "", "'vars'", "'print'", "'fail'", "'send'", "'source'",
	"'destination'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'='", "'account'", "'asset'", "'number'", "'monetary'",
	"", "", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "SOURCE",
	"DESTINATION", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN", "LBRACK",
	"RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER",
	"TY_MONETARY", "NUMBER", "RATIO", "PERCENT", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "SOURCE",
	"DESTINATION", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN", "LBRACK",
	"RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER",
	"TY_MONETARY", "NUMBER", "RATIO", "PERCENT", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
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
	NumScriptLexerSOURCE        = 8
	NumScriptLexerDESTINATION   = 9
	NumScriptLexerALLOCATE      = 10
	NumScriptLexerOP_ADD        = 11
	NumScriptLexerOP_SUB        = 12
	NumScriptLexerLPAREN        = 13
	NumScriptLexerRPAREN        = 14
	NumScriptLexerLBRACK        = 15
	NumScriptLexerRBRACK        = 16
	NumScriptLexerLBRACE        = 17
	NumScriptLexerRBRACE        = 18
	NumScriptLexerEQ            = 19
	NumScriptLexerTY_ACCOUNT    = 20
	NumScriptLexerTY_ASSET      = 21
	NumScriptLexerTY_NUMBER     = 22
	NumScriptLexerTY_MONETARY   = 23
	NumScriptLexerNUMBER        = 24
	NumScriptLexerRATIO         = 25
	NumScriptLexerPERCENT       = 26
	NumScriptLexerIDENTIFIER    = 27
	NumScriptLexerVARIABLE_NAME = 28
	NumScriptLexerACCOUNT       = 29
	NumScriptLexerASSET         = 30
)

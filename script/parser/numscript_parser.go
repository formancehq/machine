// Code generated from NumScript.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // NumScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 266,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 62, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6,
	69, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77, 11, 6,
	3, 7, 3, 7, 3, 7, 5, 7, 82, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8,
	89, 10, 8, 13, 8, 14, 8, 90, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 107, 10, 10, 13,
	10, 14, 10, 108, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 117,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 124, 10, 12, 13, 12,
	14, 12, 125, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 5, 14, 138, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 6, 15, 147, 10, 15, 13, 15, 14, 15, 148, 3, 15, 3, 15, 3, 16, 3,
	16, 5, 16, 155, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 162,
	10, 17, 13, 17, 14, 17, 163, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 182, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 202, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 208, 10, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 5, 21, 223, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 6, 22,
	230, 10, 22, 13, 22, 14, 22, 231, 6, 22, 234, 10, 22, 13, 22, 14, 22, 235,
	3, 22, 3, 22, 3, 22, 3, 23, 7, 23, 242, 10, 23, 12, 23, 14, 23, 245, 11,
	23, 3, 23, 5, 23, 248, 10, 23, 3, 23, 3, 23, 3, 23, 7, 23, 253, 10, 23,
	12, 23, 14, 23, 256, 11, 23, 3, 23, 7, 23, 259, 10, 23, 12, 23, 14, 23,
	262, 11, 23, 3, 23, 3, 23, 3, 23, 2, 3, 10, 24, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 4, 3, 2,
	22, 23, 3, 2, 31, 37, 2, 275, 2, 46, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6,
	61, 3, 2, 2, 2, 8, 63, 3, 2, 2, 2, 10, 68, 3, 2, 2, 2, 12, 81, 3, 2, 2,
	2, 14, 83, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 99, 3, 2, 2, 2, 20, 116,
	3, 2, 2, 2, 22, 118, 3, 2, 2, 2, 24, 129, 3, 2, 2, 2, 26, 137, 3, 2, 2,
	2, 28, 139, 3, 2, 2, 2, 30, 154, 3, 2, 2, 2, 32, 156, 3, 2, 2, 2, 34, 207,
	3, 2, 2, 2, 36, 209, 3, 2, 2, 2, 38, 211, 3, 2, 2, 2, 40, 218, 3, 2, 2,
	2, 42, 224, 3, 2, 2, 2, 44, 243, 3, 2, 2, 2, 46, 47, 7, 26, 2, 2, 47, 48,
	7, 46, 2, 2, 48, 49, 7, 42, 2, 2, 49, 50, 7, 27, 2, 2, 50, 3, 3, 2, 2,
	2, 51, 52, 7, 26, 2, 2, 52, 53, 7, 46, 2, 2, 53, 54, 7, 3, 2, 2, 54, 55,
	7, 27, 2, 2, 55, 5, 3, 2, 2, 2, 56, 62, 7, 45, 2, 2, 57, 62, 7, 46, 2,
	2, 58, 62, 7, 42, 2, 2, 59, 62, 7, 38, 2, 2, 60, 62, 5, 2, 2, 2, 61, 56,
	3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	61, 60, 3, 2, 2, 2, 62, 7, 3, 2, 2, 2, 63, 64, 7, 44, 2, 2, 64, 9, 3, 2,
	2, 2, 65, 66, 8, 6, 1, 2, 66, 69, 5, 6, 4, 2, 67, 69, 5, 8, 5, 2, 68, 65,
	3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 75, 3, 2, 2, 2, 70, 71, 12, 5, 2, 2,
	71, 72, 9, 2, 2, 2, 72, 74, 5, 10, 6, 6, 73, 70, 3, 2, 2, 2, 74, 77, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 78, 82, 7, 39, 2, 2, 79, 82, 5, 8, 5, 2, 80, 82, 7, 40,
	2, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 13,
	3, 2, 2, 2, 83, 84, 7, 28, 2, 2, 84, 88, 7, 5, 2, 2, 85, 86, 5, 20, 11,
	2, 86, 87, 7, 5, 2, 2, 87, 89, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 93, 7, 29, 2, 2, 93, 15, 3, 2, 2, 2, 94, 95, 7, 18, 2, 2, 95, 96, 5,
	10, 6, 2, 96, 97, 7, 20, 2, 2, 97, 98, 5, 20, 11, 2, 98, 17, 3, 2, 2, 2,
	99, 100, 7, 28, 2, 2, 100, 106, 7, 5, 2, 2, 101, 102, 5, 12, 7, 2, 102,
	103, 7, 20, 2, 2, 103, 104, 5, 20, 11, 2, 104, 105, 7, 5, 2, 2, 105, 107,
	3, 2, 2, 2, 106, 101, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 106, 3, 2,
	2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 7, 29, 2, 2,
	111, 19, 3, 2, 2, 2, 112, 117, 5, 10, 6, 2, 113, 117, 5, 16, 9, 2, 114,
	117, 5, 14, 8, 2, 115, 117, 5, 18, 10, 2, 116, 112, 3, 2, 2, 2, 116, 113,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 115, 3, 2, 2, 2, 117, 21, 3, 2,
	2, 2, 118, 119, 7, 28, 2, 2, 119, 123, 7, 5, 2, 2, 120, 121, 5, 26, 14,
	2, 121, 122, 7, 5, 2, 2, 122, 124, 3, 2, 2, 2, 123, 120, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 128, 7, 29, 2, 2, 128, 23, 3, 2, 2, 2, 129, 130, 7, 18,
	2, 2, 130, 131, 5, 10, 6, 2, 131, 132, 7, 17, 2, 2, 132, 133, 5, 26, 14,
	2, 133, 25, 3, 2, 2, 2, 134, 138, 5, 10, 6, 2, 135, 138, 5, 24, 13, 2,
	136, 138, 5, 22, 12, 2, 137, 134, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137,
	136, 3, 2, 2, 2, 138, 27, 3, 2, 2, 2, 139, 140, 7, 28, 2, 2, 140, 146,
	7, 5, 2, 2, 141, 142, 5, 12, 7, 2, 142, 143, 7, 17, 2, 2, 143, 144, 5,
	26, 14, 2, 144, 145, 7, 5, 2, 2, 145, 147, 3, 2, 2, 2, 146, 141, 3, 2,
	2, 2, 147, 148, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 151, 7, 29, 2, 2, 151, 29, 3, 2, 2, 2, 152,
	155, 5, 26, 14, 2, 153, 155, 5, 28, 15, 2, 154, 152, 3, 2, 2, 2, 154, 153,
	3, 2, 2, 2, 155, 31, 3, 2, 2, 2, 156, 157, 7, 14, 2, 2, 157, 158, 5, 8,
	5, 2, 158, 161, 7, 28, 2, 2, 159, 160, 7, 5, 2, 2, 160, 162, 5, 34, 18,
	2, 161, 159, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163,
	164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 5, 2, 2, 166, 167,
	7, 29, 2, 2, 167, 33, 3, 2, 2, 2, 168, 169, 7, 12, 2, 2, 169, 208, 5, 10,
	6, 2, 170, 171, 7, 11, 2, 2, 171, 172, 7, 24, 2, 2, 172, 173, 7, 38, 2,
	2, 173, 174, 7, 4, 2, 2, 174, 175, 5, 10, 6, 2, 175, 176, 7, 25, 2, 2,
	176, 208, 3, 2, 2, 2, 177, 208, 7, 13, 2, 2, 178, 181, 7, 15, 2, 2, 179,
	182, 5, 10, 6, 2, 180, 182, 5, 4, 3, 2, 181, 179, 3, 2, 2, 2, 181, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 7, 24, 2, 2, 184, 201, 7, 5,
	2, 2, 185, 186, 7, 16, 2, 2, 186, 187, 7, 30, 2, 2, 187, 188, 5, 30, 16,
	2, 188, 189, 7, 5, 2, 2, 189, 190, 7, 19, 2, 2, 190, 191, 7, 30, 2, 2,
	191, 192, 5, 20, 11, 2, 192, 202, 3, 2, 2, 2, 193, 194, 7, 19, 2, 2, 194,
	195, 7, 30, 2, 2, 195, 196, 5, 20, 11, 2, 196, 197, 7, 5, 2, 2, 197, 198,
	7, 16, 2, 2, 198, 199, 7, 30, 2, 2, 199, 200, 5, 30, 16, 2, 200, 202, 3,
	2, 2, 2, 201, 185, 3, 2, 2, 2, 201, 193, 3, 2, 2, 2, 202, 203, 3, 2, 2,
	2, 203, 204, 7, 5, 2, 2, 204, 205, 7, 25, 2, 2, 205, 208, 3, 2, 2, 2, 206,
	208, 5, 32, 17, 2, 207, 168, 3, 2, 2, 2, 207, 170, 3, 2, 2, 2, 207, 177,
	3, 2, 2, 2, 207, 178, 3, 2, 2, 2, 207, 206, 3, 2, 2, 2, 208, 35, 3, 2,
	2, 2, 209, 210, 9, 3, 2, 2, 210, 37, 3, 2, 2, 2, 211, 212, 7, 10, 2, 2,
	212, 213, 7, 24, 2, 2, 213, 214, 5, 10, 6, 2, 214, 215, 7, 4, 2, 2, 215,
	216, 7, 38, 2, 2, 216, 217, 7, 25, 2, 2, 217, 39, 3, 2, 2, 2, 218, 219,
	5, 36, 19, 2, 219, 222, 5, 8, 5, 2, 220, 221, 7, 30, 2, 2, 221, 223, 5,
	38, 20, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 41, 3, 2, 2,
	2, 224, 225, 7, 9, 2, 2, 225, 226, 7, 28, 2, 2, 226, 233, 7, 5, 2, 2, 227,
	229, 5, 40, 21, 2, 228, 230, 7, 5, 2, 2, 229, 228, 3, 2, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 234, 3, 2,
	2, 2, 233, 227, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2,
	235, 236, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 7, 29, 2, 2, 238,
	239, 7, 5, 2, 2, 239, 43, 3, 2, 2, 2, 240, 242, 7, 5, 2, 2, 241, 240, 3,
	2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2,
	2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 248, 5, 42, 22, 2,
	247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	254, 5, 34, 18, 2, 250, 251, 7, 5, 2, 2, 251, 253, 5, 34, 18, 2, 252, 250,
	3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2,
	2, 2, 255, 260, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 259, 7, 5, 2, 2,
	258, 257, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260,
	261, 3, 2, 2, 2, 261, 263, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 264,
	7, 2, 2, 3, 264, 45, 3, 2, 2, 2, 24, 61, 68, 75, 81, 90, 108, 116, 125,
	137, 148, 154, 163, 181, 201, 207, 222, 231, 235, 243, 247, 254, 260,
}
var literalNames = []string{
	"", "'*'", "','", "", "", "", "", "'vars'", "'meta'", "'set_tx_meta'",
	"'print'", "'fail'", "'if'", "'send'", "'source'", "'from'", "'max'", "'destination'",
	"'to'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'", "'{'",
	"'}'", "'='", "'account'", "'asset'", "'number'", "'monetary'", "'portion'",
	"'boolean'", "'string'", "", "", "'remaining'", "", "", "'%'",
}
var symbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "META", "SET_TX_META", "PRINT", "FAIL", "IF", "SEND", "SOURCE",
	"FROM", "MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "TY_PORTION", "TY_BOOLEAN", "TY_STRING", "STRING",
	"PORTION", "PORTION_REMAINING", "BOOLEAN", "NUMBER", "PERCENT", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"monetary", "monetaryAll", "literal", "variable", "expression", "allotmentPortion",
	"destinationInOrder", "destinationMaxed", "destinationAllotment", "destination",
	"sourceInOrder", "sourceMaxed", "source", "sourceAllotment", "valueAwareSource",
	"ifStatement", "statement", "type_", "origin", "varDecl", "varListDecl",
	"script",
}

type NumScriptParser struct {
	*antlr.BaseParser
}

// NewNumScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *NumScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNumScriptParser(input antlr.TokenStream) *NumScriptParser {
	this := new(NumScriptParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NumScript.g4"

	return this
}

// NumScriptParser tokens.
const (
	NumScriptParserEOF               = antlr.TokenEOF
	NumScriptParserT__0              = 1
	NumScriptParserT__1              = 2
	NumScriptParserNEWLINE           = 3
	NumScriptParserWHITESPACE        = 4
	NumScriptParserMULTILINE_COMMENT = 5
	NumScriptParserLINE_COMMENT      = 6
	NumScriptParserVARS              = 7
	NumScriptParserMETA              = 8
	NumScriptParserSET_TX_META       = 9
	NumScriptParserPRINT             = 10
	NumScriptParserFAIL              = 11
	NumScriptParserIF                = 12
	NumScriptParserSEND              = 13
	NumScriptParserSOURCE            = 14
	NumScriptParserFROM              = 15
	NumScriptParserMAX               = 16
	NumScriptParserDESTINATION       = 17
	NumScriptParserTO                = 18
	NumScriptParserALLOCATE          = 19
	NumScriptParserOP_ADD            = 20
	NumScriptParserOP_SUB            = 21
	NumScriptParserLPAREN            = 22
	NumScriptParserRPAREN            = 23
	NumScriptParserLBRACK            = 24
	NumScriptParserRBRACK            = 25
	NumScriptParserLBRACE            = 26
	NumScriptParserRBRACE            = 27
	NumScriptParserEQ                = 28
	NumScriptParserTY_ACCOUNT        = 29
	NumScriptParserTY_ASSET          = 30
	NumScriptParserTY_NUMBER         = 31
	NumScriptParserTY_MONETARY       = 32
	NumScriptParserTY_PORTION        = 33
	NumScriptParserTY_BOOLEAN        = 34
	NumScriptParserTY_STRING         = 35
	NumScriptParserSTRING            = 36
	NumScriptParserPORTION           = 37
	NumScriptParserPORTION_REMAINING = 38
	NumScriptParserBOOLEAN           = 39
	NumScriptParserNUMBER            = 40
	NumScriptParserPERCENT           = 41
	NumScriptParserVARIABLE_NAME     = 42
	NumScriptParserACCOUNT           = 43
	NumScriptParserASSET             = 44
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_monetary             = 0
	NumScriptParserRULE_monetaryAll          = 1
	NumScriptParserRULE_literal              = 2
	NumScriptParserRULE_variable             = 3
	NumScriptParserRULE_expression           = 4
	NumScriptParserRULE_allotmentPortion     = 5
	NumScriptParserRULE_destinationInOrder   = 6
	NumScriptParserRULE_destinationMaxed     = 7
	NumScriptParserRULE_destinationAllotment = 8
	NumScriptParserRULE_destination          = 9
	NumScriptParserRULE_sourceInOrder        = 10
	NumScriptParserRULE_sourceMaxed          = 11
	NumScriptParserRULE_source               = 12
	NumScriptParserRULE_sourceAllotment      = 13
	NumScriptParserRULE_valueAwareSource     = 14
	NumScriptParserRULE_ifStatement          = 15
	NumScriptParserRULE_statement            = 16
	NumScriptParserRULE_type_                = 17
	NumScriptParserRULE_origin               = 18
	NumScriptParserRULE_varDecl              = 19
	NumScriptParserRULE_varListDecl          = 20
	NumScriptParserRULE_script               = 21
)

// IMonetaryContext is an interface to support dynamic dispatch.
type IMonetaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsset returns the asset token.
	GetAsset() antlr.Token

	// GetAmt returns the amt token.
	GetAmt() antlr.Token

	// SetAsset sets the asset token.
	SetAsset(antlr.Token)

	// SetAmt sets the amt token.
	SetAmt(antlr.Token)

	// IsMonetaryContext differentiates from other interfaces.
	IsMonetaryContext()
}

type MonetaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	asset  antlr.Token
	amt    antlr.Token
}

func NewEmptyMonetaryContext() *MonetaryContext {
	var p = new(MonetaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_monetary
	return p
}

func (*MonetaryContext) IsMonetaryContext() {}

func NewMonetaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonetaryContext {
	var p = new(MonetaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_monetary

	return p
}

func (s *MonetaryContext) GetParser() antlr.Parser { return s.parser }

func (s *MonetaryContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryContext) GetAmt() antlr.Token { return s.amt }

func (s *MonetaryContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryContext) SetAmt(v antlr.Token) { s.amt = v }

func (s *MonetaryContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACK, 0)
}

func (s *MonetaryContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACK, 0)
}

func (s *MonetaryContext) ASSET() antlr.TerminalNode {
	return s.GetToken(NumScriptParserASSET, 0)
}

func (s *MonetaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserNUMBER, 0)
}

func (s *MonetaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonetaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterMonetary(s)
	}
}

func (s *MonetaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitMonetary(s)
	}
}

func (p *NumScriptParser) Monetary() (localctx IMonetaryContext) {
	this := p
	_ = this

	localctx = NewMonetaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NumScriptParserRULE_monetary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(45)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(46)

		var _m = p.Match(NumScriptParserNUMBER)

		localctx.(*MonetaryContext).amt = _m
	}
	{
		p.SetState(47)
		p.Match(NumScriptParserRBRACK)
	}

	return localctx
}

// IMonetaryAllContext is an interface to support dynamic dispatch.
type IMonetaryAllContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsset returns the asset token.
	GetAsset() antlr.Token

	// SetAsset sets the asset token.
	SetAsset(antlr.Token)

	// IsMonetaryAllContext differentiates from other interfaces.
	IsMonetaryAllContext()
}

type MonetaryAllContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	asset  antlr.Token
}

func NewEmptyMonetaryAllContext() *MonetaryAllContext {
	var p = new(MonetaryAllContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_monetaryAll
	return p
}

func (*MonetaryAllContext) IsMonetaryAllContext() {}

func NewMonetaryAllContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonetaryAllContext {
	var p = new(MonetaryAllContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_monetaryAll

	return p
}

func (s *MonetaryAllContext) GetParser() antlr.Parser { return s.parser }

func (s *MonetaryAllContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryAllContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryAllContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACK, 0)
}

func (s *MonetaryAllContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACK, 0)
}

func (s *MonetaryAllContext) ASSET() antlr.TerminalNode {
	return s.GetToken(NumScriptParserASSET, 0)
}

func (s *MonetaryAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAllContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonetaryAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterMonetaryAll(s)
	}
}

func (s *MonetaryAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitMonetaryAll(s)
	}
}

func (p *NumScriptParser) MonetaryAll() (localctx IMonetaryAllContext) {
	this := p
	_ = this

	localctx = NewMonetaryAllContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NumScriptParserRULE_monetaryAll)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(50)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryAllContext).asset = _m
	}
	{
		p.SetState(51)
		p.Match(NumScriptParserT__0)
	}
	{
		p.SetState(52)
		p.Match(NumScriptParserRBRACK)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LitStringContext struct {
	*LiteralContext
}

func NewLitStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitStringContext {
	var p = new(LitStringContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSTRING, 0)
}

func (s *LitStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterLitString(s)
	}
}

func (s *LitStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitLitString(s)
	}
}

type LitAccountContext struct {
	*LiteralContext
}

func NewLitAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitAccountContext {
	var p = new(LitAccountContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitAccountContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(NumScriptParserACCOUNT, 0)
}

func (s *LitAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterLitAccount(s)
	}
}

func (s *LitAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitLitAccount(s)
	}
}

type LitAssetContext struct {
	*LiteralContext
}

func NewLitAssetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitAssetContext {
	var p = new(LitAssetContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitAssetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitAssetContext) ASSET() antlr.TerminalNode {
	return s.GetToken(NumScriptParserASSET, 0)
}

func (s *LitAssetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterLitAsset(s)
	}
}

func (s *LitAssetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitLitAsset(s)
	}
}

type LitMonetaryContext struct {
	*LiteralContext
}

func NewLitMonetaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitMonetaryContext {
	var p = new(LitMonetaryContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitMonetaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitMonetaryContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *LitMonetaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterLitMonetary(s)
	}
}

func (s *LitMonetaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitLitMonetary(s)
	}
}

type LitNumberContext struct {
	*LiteralContext
}

func NewLitNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitNumberContext {
	var p = new(LitNumberContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserNUMBER, 0)
}

func (s *LitNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterLitNumber(s)
	}
}

func (s *LitNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitLitNumber(s)
	}
}

func (p *NumScriptParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NumScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserSTRING:
		localctx = NewLitStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Match(NumScriptParserSTRING)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.Monetary()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(NumScriptParserVARIABLE_NAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *NumScriptParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NumScriptParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(NumScriptParserVARIABLE_NAME)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprAddSubContext struct {
	*ExpressionContext
	lhs IExpressionContext
	op  antlr.Token
	rhs IExpressionContext
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ExprAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprAddSubContext) GetLhs() IExpressionContext { return s.lhs }

func (s *ExprAddSubContext) GetRhs() IExpressionContext { return s.rhs }

func (s *ExprAddSubContext) SetLhs(v IExpressionContext) { s.lhs = v }

func (s *ExprAddSubContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprAddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprAddSubContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(NumScriptParserOP_ADD, 0)
}

func (s *ExprAddSubContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(NumScriptParserOP_SUB, 0)
}

func (s *ExprAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterExprAddSub(s)
	}
}

func (s *ExprAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitExprAddSub(s)
	}
}

type ExprLiteralContext struct {
	*ExpressionContext
	lit ILiteralContext
}

func NewExprLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLiteralContext {
	var p = new(ExprLiteralContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprLiteralContext) GetLit() ILiteralContext { return s.lit }

func (s *ExprLiteralContext) SetLit(v ILiteralContext) { s.lit = v }

func (s *ExprLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLiteralContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterExprLiteral(s)
	}
}

func (s *ExprLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitExprLiteral(s)
	}
}

type ExprVariableContext struct {
	*ExpressionContext
	var_ IVariableContext
}

func NewExprVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprVariableContext {
	var p = new(ExprVariableContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprVariableContext) GetVar_() IVariableContext { return s.var_ }

func (s *ExprVariableContext) SetVar_(v IVariableContext) { s.var_ = v }

func (s *ExprVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprVariableContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExprVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterExprVariable(s)
	}
}

func (s *ExprVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitExprVariable(s)
	}
}

func (p *NumScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *NumScriptParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, NumScriptParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserSTRING, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(64)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)

			var _x = p.Variable()

			localctx.(*ExprVariableContext).var_ = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, NumScriptParserRULE_expression)
			p.SetState(68)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(69)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprAddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == NumScriptParserOP_ADD || _la == NumScriptParserOP_SUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprAddSubContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(70)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IAllotmentPortionContext is an interface to support dynamic dispatch.
type IAllotmentPortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllotmentPortionContext differentiates from other interfaces.
	IsAllotmentPortionContext()
}

type AllotmentPortionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllotmentPortionContext() *AllotmentPortionContext {
	var p = new(AllotmentPortionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allotmentPortion
	return p
}

func (*AllotmentPortionContext) IsAllotmentPortionContext() {}

func NewAllotmentPortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllotmentPortionContext {
	var p = new(AllotmentPortionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allotmentPortion

	return p
}

func (s *AllotmentPortionContext) GetParser() antlr.Parser { return s.parser }

func (s *AllotmentPortionContext) CopyFrom(ctx *AllotmentPortionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AllotmentPortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllotmentPortionRemainingContext struct {
	*AllotmentPortionContext
}

func NewAllotmentPortionRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionRemainingContext {
	var p = new(AllotmentPortionRemainingContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionRemainingContext) PORTION_REMAINING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION_REMAINING, 0)
}

func (s *AllotmentPortionRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllotmentPortionRemaining(s)
	}
}

func (s *AllotmentPortionRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllotmentPortionRemaining(s)
	}
}

type AllotmentPortionConstContext struct {
	*AllotmentPortionContext
}

func NewAllotmentPortionConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionConstContext {
	var p = new(AllotmentPortionConstContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionConstContext) PORTION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION, 0)
}

func (s *AllotmentPortionConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllotmentPortionConst(s)
	}
}

func (s *AllotmentPortionConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllotmentPortionConst(s)
	}
}

type AllotmentPortionVarContext struct {
	*AllotmentPortionContext
	por IVariableContext
}

func NewAllotmentPortionVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionVarContext {
	var p = new(AllotmentPortionVarContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionVarContext) GetPor() IVariableContext { return s.por }

func (s *AllotmentPortionVarContext) SetPor(v IVariableContext) { s.por = v }

func (s *AllotmentPortionVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionVarContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AllotmentPortionVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllotmentPortionVar(s)
	}
}

func (s *AllotmentPortionVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllotmentPortionVar(s)
	}
}

func (p *NumScriptParser) AllotmentPortion() (localctx IAllotmentPortionContext) {
	this := p
	_ = this

	localctx = NewAllotmentPortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NumScriptParserRULE_allotmentPortion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPORTION:
		localctx = NewAllotmentPortionConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(NumScriptParserPORTION)
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewAllotmentPortionVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)

			var _x = p.Variable()

			localctx.(*AllotmentPortionVarContext).por = _x
		}

	case NumScriptParserPORTION_REMAINING:
		localctx = NewAllotmentPortionRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(NumScriptParserPORTION_REMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDestinationInOrderContext is an interface to support dynamic dispatch.
type IDestinationInOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_destination returns the _destination rule contexts.
	Get_destination() IDestinationContext

	// Set_destination sets the _destination rule contexts.
	Set_destination(IDestinationContext)

	// GetDests returns the dests rule context list.
	GetDests() []IDestinationContext

	// SetDests sets the dests rule context list.
	SetDests([]IDestinationContext)

	// IsDestinationInOrderContext differentiates from other interfaces.
	IsDestinationInOrderContext()
}

type DestinationInOrderContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_destination IDestinationContext
	dests        []IDestinationContext
}

func NewEmptyDestinationInOrderContext() *DestinationInOrderContext {
	var p = new(DestinationInOrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_destinationInOrder
	return p
}

func (*DestinationInOrderContext) IsDestinationInOrderContext() {}

func NewDestinationInOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationInOrderContext {
	var p = new(DestinationInOrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_destinationInOrder

	return p
}

func (s *DestinationInOrderContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationInOrderContext) Get_destination() IDestinationContext { return s._destination }

func (s *DestinationInOrderContext) Set_destination(v IDestinationContext) { s._destination = v }

func (s *DestinationInOrderContext) GetDests() []IDestinationContext { return s.dests }

func (s *DestinationInOrderContext) SetDests(v []IDestinationContext) { s.dests = v }

func (s *DestinationInOrderContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *DestinationInOrderContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *DestinationInOrderContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *DestinationInOrderContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *DestinationInOrderContext) AllDestination() []IDestinationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDestinationContext)(nil)).Elem())
	var tst = make([]IDestinationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDestinationContext)
		}
	}

	return tst
}

func (s *DestinationInOrderContext) Destination(i int) IDestinationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
}

func (s *DestinationInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationInOrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestinationInOrder(s)
	}
}

func (s *DestinationInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestinationInOrder(s)
	}
}

func (p *NumScriptParser) DestinationInOrder() (localctx IDestinationInOrderContext) {
	this := p
	_ = this

	localctx = NewDestinationInOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumScriptParserRULE_destinationInOrder)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(82)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(NumScriptParserMAX-16))|(1<<(NumScriptParserLBRACK-16))|(1<<(NumScriptParserLBRACE-16))|(1<<(NumScriptParserSTRING-16))|(1<<(NumScriptParserNUMBER-16))|(1<<(NumScriptParserVARIABLE_NAME-16))|(1<<(NumScriptParserACCOUNT-16))|(1<<(NumScriptParserASSET-16)))) != 0) {
		{
			p.SetState(83)

			var _x = p.Destination()

			localctx.(*DestinationInOrderContext)._destination = _x
		}
		localctx.(*DestinationInOrderContext).dests = append(localctx.(*DestinationInOrderContext).dests, localctx.(*DestinationInOrderContext)._destination)
		{
			p.SetState(84)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IDestinationMaxedContext is an interface to support dynamic dispatch.
type IDestinationMaxedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMax returns the max rule contexts.
	GetMax() IExpressionContext

	// GetDest returns the dest rule contexts.
	GetDest() IDestinationContext

	// SetMax sets the max rule contexts.
	SetMax(IExpressionContext)

	// SetDest sets the dest rule contexts.
	SetDest(IDestinationContext)

	// IsDestinationMaxedContext differentiates from other interfaces.
	IsDestinationMaxedContext()
}

type DestinationMaxedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	max    IExpressionContext
	dest   IDestinationContext
}

func NewEmptyDestinationMaxedContext() *DestinationMaxedContext {
	var p = new(DestinationMaxedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_destinationMaxed
	return p
}

func (*DestinationMaxedContext) IsDestinationMaxedContext() {}

func NewDestinationMaxedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationMaxedContext {
	var p = new(DestinationMaxedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_destinationMaxed

	return p
}

func (s *DestinationMaxedContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationMaxedContext) GetMax() IExpressionContext { return s.max }

func (s *DestinationMaxedContext) GetDest() IDestinationContext { return s.dest }

func (s *DestinationMaxedContext) SetMax(v IExpressionContext) { s.max = v }

func (s *DestinationMaxedContext) SetDest(v IDestinationContext) { s.dest = v }

func (s *DestinationMaxedContext) MAX() antlr.TerminalNode {
	return s.GetToken(NumScriptParserMAX, 0)
}

func (s *DestinationMaxedContext) TO() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTO, 0)
}

func (s *DestinationMaxedContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestinationMaxedContext) Destination() IDestinationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
}

func (s *DestinationMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationMaxedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestinationMaxed(s)
	}
}

func (s *DestinationMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestinationMaxed(s)
	}
}

func (p *NumScriptParser) DestinationMaxed() (localctx IDestinationMaxedContext) {
	this := p
	_ = this

	localctx = NewDestinationMaxedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumScriptParserRULE_destinationMaxed)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(NumScriptParserMAX)
	}
	{
		p.SetState(93)

		var _x = p.expression(0)

		localctx.(*DestinationMaxedContext).max = _x
	}
	{
		p.SetState(94)
		p.Match(NumScriptParserTO)
	}
	{
		p.SetState(95)

		var _x = p.Destination()

		localctx.(*DestinationMaxedContext).dest = _x
	}

	return localctx
}

// IDestinationAllotmentContext is an interface to support dynamic dispatch.
type IDestinationAllotmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allotmentPortion returns the _allotmentPortion rule contexts.
	Get_allotmentPortion() IAllotmentPortionContext

	// Get_destination returns the _destination rule contexts.
	Get_destination() IDestinationContext

	// Set_allotmentPortion sets the _allotmentPortion rule contexts.
	Set_allotmentPortion(IAllotmentPortionContext)

	// Set_destination sets the _destination rule contexts.
	Set_destination(IDestinationContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllotmentPortionContext

	// GetDests returns the dests rule context list.
	GetDests() []IDestinationContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetDests sets the dests rule context list.
	SetDests([]IDestinationContext)

	// IsDestinationAllotmentContext differentiates from other interfaces.
	IsDestinationAllotmentContext()
}

type DestinationAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_destination      IDestinationContext
	dests             []IDestinationContext
}

func NewEmptyDestinationAllotmentContext() *DestinationAllotmentContext {
	var p = new(DestinationAllotmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_destinationAllotment
	return p
}

func (*DestinationAllotmentContext) IsDestinationAllotmentContext() {}

func NewDestinationAllotmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationAllotmentContext {
	var p = new(DestinationAllotmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_destinationAllotment

	return p
}

func (s *DestinationAllotmentContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationAllotmentContext) Get_allotmentPortion() IAllotmentPortionContext {
	return s._allotmentPortion
}

func (s *DestinationAllotmentContext) Get_destination() IDestinationContext { return s._destination }

func (s *DestinationAllotmentContext) Set_allotmentPortion(v IAllotmentPortionContext) {
	s._allotmentPortion = v
}

func (s *DestinationAllotmentContext) Set_destination(v IDestinationContext) { s._destination = v }

func (s *DestinationAllotmentContext) GetPortions() []IAllotmentPortionContext { return s.portions }

func (s *DestinationAllotmentContext) GetDests() []IDestinationContext { return s.dests }

func (s *DestinationAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *DestinationAllotmentContext) SetDests(v []IDestinationContext) { s.dests = v }

func (s *DestinationAllotmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *DestinationAllotmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *DestinationAllotmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *DestinationAllotmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *DestinationAllotmentContext) AllTO() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserTO)
}

func (s *DestinationAllotmentContext) TO(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserTO, i)
}

func (s *DestinationAllotmentContext) AllAllotmentPortion() []IAllotmentPortionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem())
	var tst = make([]IAllotmentPortionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllotmentPortionContext)
		}
	}

	return tst
}

func (s *DestinationAllotmentContext) AllotmentPortion(i int) IAllotmentPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllotmentPortionContext)
}

func (s *DestinationAllotmentContext) AllDestination() []IDestinationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDestinationContext)(nil)).Elem())
	var tst = make([]IDestinationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDestinationContext)
		}
	}

	return tst
}

func (s *DestinationAllotmentContext) Destination(i int) IDestinationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
}

func (s *DestinationAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationAllotmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestinationAllotment(s)
	}
}

func (s *DestinationAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestinationAllotment(s)
	}
}

func (p *NumScriptParser) DestinationAllotment() (localctx IDestinationAllotmentContext) {
	this := p
	_ = this

	localctx = NewDestinationAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumScriptParserRULE_destinationAllotment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(98)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NumScriptParserPORTION-37))|(1<<(NumScriptParserPORTION_REMAINING-37))|(1<<(NumScriptParserVARIABLE_NAME-37)))) != 0) {
		{
			p.SetState(99)

			var _x = p.AllotmentPortion()

			localctx.(*DestinationAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*DestinationAllotmentContext).portions = append(localctx.(*DestinationAllotmentContext).portions, localctx.(*DestinationAllotmentContext)._allotmentPortion)
		{
			p.SetState(100)
			p.Match(NumScriptParserTO)
		}
		{
			p.SetState(101)

			var _x = p.Destination()

			localctx.(*DestinationAllotmentContext)._destination = _x
		}
		localctx.(*DestinationAllotmentContext).dests = append(localctx.(*DestinationAllotmentContext).dests, localctx.(*DestinationAllotmentContext)._destination)
		{
			p.SetState(102)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(108)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IDestinationContext is an interface to support dynamic dispatch.
type IDestinationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestinationContext differentiates from other interfaces.
	IsDestinationContext()
}

type DestinationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationContext() *DestinationContext {
	var p = new(DestinationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_destination
	return p
}

func (*DestinationContext) IsDestinationContext() {}

func NewDestinationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationContext {
	var p = new(DestinationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_destination

	return p
}

func (s *DestinationContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationContext) CopyFrom(ctx *DestinationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DestinationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DestAccountContext struct {
	*DestinationContext
}

func NewDestAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestAccountContext {
	var p = new(DestAccountContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestAccountContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestAccount(s)
	}
}

func (s *DestAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestAccount(s)
	}
}

type DestAllotmentContext struct {
	*DestinationContext
}

func NewDestAllotmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestAllotmentContext {
	var p = new(DestAllotmentContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestAllotmentContext) DestinationAllotment() IDestinationAllotmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationAllotmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationAllotmentContext)
}

func (s *DestAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestAllotment(s)
	}
}

func (s *DestAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestAllotment(s)
	}
}

type DestMaxedContext struct {
	*DestinationContext
}

func NewDestMaxedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestMaxedContext {
	var p = new(DestMaxedContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestMaxedContext) DestinationMaxed() IDestinationMaxedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationMaxedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationMaxedContext)
}

func (s *DestMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestMaxed(s)
	}
}

func (s *DestMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestMaxed(s)
	}
}

type DestInOrderContext struct {
	*DestinationContext
}

func NewDestInOrderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestInOrderContext {
	var p = new(DestInOrderContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestInOrderContext) DestinationInOrder() IDestinationInOrderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationInOrderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationInOrderContext)
}

func (s *DestInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterDestInOrder(s)
	}
}

func (s *DestInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitDestInOrder(s)
	}
}

func (p *NumScriptParser) Destination() (localctx IDestinationContext) {
	this := p
	_ = this

	localctx = NewDestinationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NumScriptParserRULE_destination)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDestAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.expression(0)
		}

	case 2:
		localctx = NewDestMaxedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.DestinationMaxed()
		}

	case 3:
		localctx = NewDestInOrderContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.DestinationInOrder()
		}

	case 4:
		localctx = NewDestAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.DestinationAllotment()
		}

	}

	return localctx
}

// ISourceInOrderContext is an interface to support dynamic dispatch.
type ISourceInOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_source returns the _source rule contexts.
	Get_source() ISourceContext

	// Set_source sets the _source rule contexts.
	Set_source(ISourceContext)

	// GetSources returns the sources rule context list.
	GetSources() []ISourceContext

	// SetSources sets the sources rule context list.
	SetSources([]ISourceContext)

	// IsSourceInOrderContext differentiates from other interfaces.
	IsSourceInOrderContext()
}

type SourceInOrderContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_source ISourceContext
	sources []ISourceContext
}

func NewEmptySourceInOrderContext() *SourceInOrderContext {
	var p = new(SourceInOrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_sourceInOrder
	return p
}

func (*SourceInOrderContext) IsSourceInOrderContext() {}

func NewSourceInOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceInOrderContext {
	var p = new(SourceInOrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_sourceInOrder

	return p
}

func (s *SourceInOrderContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceInOrderContext) Get_source() ISourceContext { return s._source }

func (s *SourceInOrderContext) Set_source(v ISourceContext) { s._source = v }

func (s *SourceInOrderContext) GetSources() []ISourceContext { return s.sources }

func (s *SourceInOrderContext) SetSources(v []ISourceContext) { s.sources = v }

func (s *SourceInOrderContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *SourceInOrderContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *SourceInOrderContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *SourceInOrderContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *SourceInOrderContext) AllSource() []ISourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceContext)(nil)).Elem())
	var tst = make([]ISourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceContext)
		}
	}

	return tst
}

func (s *SourceInOrderContext) Source(i int) ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceInOrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSourceInOrder(s)
	}
}

func (s *SourceInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSourceInOrder(s)
	}
}

func (p *NumScriptParser) SourceInOrder() (localctx ISourceInOrderContext) {
	this := p
	_ = this

	localctx = NewSourceInOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NumScriptParserRULE_sourceInOrder)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(117)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(NumScriptParserMAX-16))|(1<<(NumScriptParserLBRACK-16))|(1<<(NumScriptParserLBRACE-16))|(1<<(NumScriptParserSTRING-16))|(1<<(NumScriptParserNUMBER-16))|(1<<(NumScriptParserVARIABLE_NAME-16))|(1<<(NumScriptParserACCOUNT-16))|(1<<(NumScriptParserASSET-16)))) != 0) {
		{
			p.SetState(118)

			var _x = p.Source()

			localctx.(*SourceInOrderContext)._source = _x
		}
		localctx.(*SourceInOrderContext).sources = append(localctx.(*SourceInOrderContext).sources, localctx.(*SourceInOrderContext)._source)
		{
			p.SetState(119)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// ISourceMaxedContext is an interface to support dynamic dispatch.
type ISourceMaxedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMax returns the max rule contexts.
	GetMax() IExpressionContext

	// GetSrc returns the src rule contexts.
	GetSrc() ISourceContext

	// SetMax sets the max rule contexts.
	SetMax(IExpressionContext)

	// SetSrc sets the src rule contexts.
	SetSrc(ISourceContext)

	// IsSourceMaxedContext differentiates from other interfaces.
	IsSourceMaxedContext()
}

type SourceMaxedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	max    IExpressionContext
	src    ISourceContext
}

func NewEmptySourceMaxedContext() *SourceMaxedContext {
	var p = new(SourceMaxedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_sourceMaxed
	return p
}

func (*SourceMaxedContext) IsSourceMaxedContext() {}

func NewSourceMaxedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceMaxedContext {
	var p = new(SourceMaxedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_sourceMaxed

	return p
}

func (s *SourceMaxedContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceMaxedContext) GetMax() IExpressionContext { return s.max }

func (s *SourceMaxedContext) GetSrc() ISourceContext { return s.src }

func (s *SourceMaxedContext) SetMax(v IExpressionContext) { s.max = v }

func (s *SourceMaxedContext) SetSrc(v ISourceContext) { s.src = v }

func (s *SourceMaxedContext) MAX() antlr.TerminalNode {
	return s.GetToken(NumScriptParserMAX, 0)
}

func (s *SourceMaxedContext) FROM() antlr.TerminalNode {
	return s.GetToken(NumScriptParserFROM, 0)
}

func (s *SourceMaxedContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SourceMaxedContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceMaxedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSourceMaxed(s)
	}
}

func (s *SourceMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSourceMaxed(s)
	}
}

func (p *NumScriptParser) SourceMaxed() (localctx ISourceMaxedContext) {
	this := p
	_ = this

	localctx = NewSourceMaxedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NumScriptParserRULE_sourceMaxed)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(NumScriptParserMAX)
	}
	{
		p.SetState(128)

		var _x = p.expression(0)

		localctx.(*SourceMaxedContext).max = _x
	}
	{
		p.SetState(129)
		p.Match(NumScriptParserFROM)
	}
	{
		p.SetState(130)

		var _x = p.Source()

		localctx.(*SourceMaxedContext).src = _x
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) CopyFrom(ctx *SourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SrcAccountContext struct {
	*SourceContext
}

func NewSrcAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcAccountContext {
	var p = new(SrcAccountContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcAccountContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcAccount(s)
	}
}

func (s *SrcAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcAccount(s)
	}
}

type SrcMaxedContext struct {
	*SourceContext
}

func NewSrcMaxedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcMaxedContext {
	var p = new(SrcMaxedContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcMaxedContext) SourceMaxed() ISourceMaxedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceMaxedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceMaxedContext)
}

func (s *SrcMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcMaxed(s)
	}
}

func (s *SrcMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcMaxed(s)
	}
}

type SrcInOrderContext struct {
	*SourceContext
}

func NewSrcInOrderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcInOrderContext {
	var p = new(SrcInOrderContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcInOrderContext) SourceInOrder() ISourceInOrderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceInOrderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceInOrderContext)
}

func (s *SrcInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcInOrder(s)
	}
}

func (s *SrcInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcInOrder(s)
	}
}

func (p *NumScriptParser) Source() (localctx ISourceContext) {
	this := p
	_ = this

	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NumScriptParserRULE_source)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserSTRING, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.expression(0)
		}

	case NumScriptParserMAX:
		localctx = NewSrcMaxedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.SourceMaxed()
		}

	case NumScriptParserLBRACE:
		localctx = NewSrcInOrderContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.SourceInOrder()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISourceAllotmentContext is an interface to support dynamic dispatch.
type ISourceAllotmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allotmentPortion returns the _allotmentPortion rule contexts.
	Get_allotmentPortion() IAllotmentPortionContext

	// Get_source returns the _source rule contexts.
	Get_source() ISourceContext

	// Set_allotmentPortion sets the _allotmentPortion rule contexts.
	Set_allotmentPortion(IAllotmentPortionContext)

	// Set_source sets the _source rule contexts.
	Set_source(ISourceContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllotmentPortionContext

	// GetSources returns the sources rule context list.
	GetSources() []ISourceContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetSources sets the sources rule context list.
	SetSources([]ISourceContext)

	// IsSourceAllotmentContext differentiates from other interfaces.
	IsSourceAllotmentContext()
}

type SourceAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_source           ISourceContext
	sources           []ISourceContext
}

func NewEmptySourceAllotmentContext() *SourceAllotmentContext {
	var p = new(SourceAllotmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_sourceAllotment
	return p
}

func (*SourceAllotmentContext) IsSourceAllotmentContext() {}

func NewSourceAllotmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceAllotmentContext {
	var p = new(SourceAllotmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_sourceAllotment

	return p
}

func (s *SourceAllotmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceAllotmentContext) Get_allotmentPortion() IAllotmentPortionContext {
	return s._allotmentPortion
}

func (s *SourceAllotmentContext) Get_source() ISourceContext { return s._source }

func (s *SourceAllotmentContext) Set_allotmentPortion(v IAllotmentPortionContext) {
	s._allotmentPortion = v
}

func (s *SourceAllotmentContext) Set_source(v ISourceContext) { s._source = v }

func (s *SourceAllotmentContext) GetPortions() []IAllotmentPortionContext { return s.portions }

func (s *SourceAllotmentContext) GetSources() []ISourceContext { return s.sources }

func (s *SourceAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *SourceAllotmentContext) SetSources(v []ISourceContext) { s.sources = v }

func (s *SourceAllotmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *SourceAllotmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *SourceAllotmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *SourceAllotmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *SourceAllotmentContext) AllFROM() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserFROM)
}

func (s *SourceAllotmentContext) FROM(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserFROM, i)
}

func (s *SourceAllotmentContext) AllAllotmentPortion() []IAllotmentPortionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem())
	var tst = make([]IAllotmentPortionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllotmentPortionContext)
		}
	}

	return tst
}

func (s *SourceAllotmentContext) AllotmentPortion(i int) IAllotmentPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllotmentPortionContext)
}

func (s *SourceAllotmentContext) AllSource() []ISourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceContext)(nil)).Elem())
	var tst = make([]ISourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceContext)
		}
	}

	return tst
}

func (s *SourceAllotmentContext) Source(i int) ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceAllotmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSourceAllotment(s)
	}
}

func (s *SourceAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSourceAllotment(s)
	}
}

func (p *NumScriptParser) SourceAllotment() (localctx ISourceAllotmentContext) {
	this := p
	_ = this

	localctx = NewSourceAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NumScriptParserRULE_sourceAllotment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(138)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NumScriptParserPORTION-37))|(1<<(NumScriptParserPORTION_REMAINING-37))|(1<<(NumScriptParserVARIABLE_NAME-37)))) != 0) {
		{
			p.SetState(139)

			var _x = p.AllotmentPortion()

			localctx.(*SourceAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*SourceAllotmentContext).portions = append(localctx.(*SourceAllotmentContext).portions, localctx.(*SourceAllotmentContext)._allotmentPortion)
		{
			p.SetState(140)
			p.Match(NumScriptParserFROM)
		}
		{
			p.SetState(141)

			var _x = p.Source()

			localctx.(*SourceAllotmentContext)._source = _x
		}
		localctx.(*SourceAllotmentContext).sources = append(localctx.(*SourceAllotmentContext).sources, localctx.(*SourceAllotmentContext)._source)
		{
			p.SetState(142)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IValueAwareSourceContext is an interface to support dynamic dispatch.
type IValueAwareSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueAwareSourceContext differentiates from other interfaces.
	IsValueAwareSourceContext()
}

type ValueAwareSourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAwareSourceContext() *ValueAwareSourceContext {
	var p = new(ValueAwareSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_valueAwareSource
	return p
}

func (*ValueAwareSourceContext) IsValueAwareSourceContext() {}

func NewValueAwareSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAwareSourceContext {
	var p = new(ValueAwareSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_valueAwareSource

	return p
}

func (s *ValueAwareSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAwareSourceContext) CopyFrom(ctx *ValueAwareSourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueAwareSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAwareSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SrcContext struct {
	*ValueAwareSourceContext
}

func NewSrcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcContext {
	var p = new(SrcContext)

	p.ValueAwareSourceContext = NewEmptyValueAwareSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueAwareSourceContext))

	return p
}

func (s *SrcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SrcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrc(s)
	}
}

func (s *SrcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrc(s)
	}
}

type SrcAllotmentContext struct {
	*ValueAwareSourceContext
}

func NewSrcAllotmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcAllotmentContext {
	var p = new(SrcAllotmentContext)

	p.ValueAwareSourceContext = NewEmptyValueAwareSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueAwareSourceContext))

	return p
}

func (s *SrcAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcAllotmentContext) SourceAllotment() ISourceAllotmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceAllotmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceAllotmentContext)
}

func (s *SrcAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcAllotment(s)
	}
}

func (s *SrcAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcAllotment(s)
	}
}

func (p *NumScriptParser) ValueAwareSource() (localctx IValueAwareSourceContext) {
	this := p
	_ = this

	localctx = NewValueAwareSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NumScriptParserRULE_valueAwareSource)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSrcContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Source()
		}

	case 2:
		localctx = NewSrcAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.SourceAllotment()
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IVariableContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IVariableContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStmts returns the stmts rule context list.
	GetStmts() []IStatementContext

	// SetStmts sets the stmts rule context list.
	SetStmts([]IStatementContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	condition  IVariableContext
	_statement IStatementContext
	stmts      []IStatementContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetCondition() IVariableContext { return s.condition }

func (s *IfStatementContext) Get_statement() IStatementContext { return s._statement }

func (s *IfStatementContext) SetCondition(v IVariableContext) { s.condition = v }

func (s *IfStatementContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *IfStatementContext) GetStmts() []IStatementContext { return s.stmts }

func (s *IfStatementContext) SetStmts(v []IStatementContext) { s.stmts = v }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(NumScriptParserIF, 0)
}

func (s *IfStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *IfStatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *IfStatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *IfStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *IfStatementContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *NumScriptParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NumScriptParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(NumScriptParserIF)
	}
	{
		p.SetState(155)

		var _x = p.Variable()

		localctx.(*IfStatementContext).condition = _x
	}
	{
		p.SetState(156)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(157)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(158)

				var _x = p.Statement()

				localctx.(*IfStatementContext)._statement = _x
			}
			localctx.(*IfStatementContext).stmts = append(localctx.(*IfStatementContext).stmts, localctx.(*IfStatementContext)._statement)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	{
		p.SetState(163)
		p.Match(NumScriptParserNEWLINE)
	}
	{
		p.SetState(164)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	*StatementContext
	expr IExpressionContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintContext) GetExpr() IExpressionContext { return s.expr }

func (s *PrintContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPRINT, 0)
}

func (s *PrintContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitPrint(s)
	}
}

type IfStmtContext struct {
	*StatementContext
	stmt IIfStatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStmtContext) GetStmt() IIfStatementContext { return s.stmt }

func (s *IfStmtContext) SetStmt(v IIfStatementContext) { s.stmt = v }

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

type SetTxMetaContext struct {
	*StatementContext
	key   antlr.Token
	value IExpressionContext
}

func NewSetTxMetaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetTxMetaContext {
	var p = new(SetTxMetaContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SetTxMetaContext) GetKey() antlr.Token { return s.key }

func (s *SetTxMetaContext) SetKey(v antlr.Token) { s.key = v }

func (s *SetTxMetaContext) GetValue() IExpressionContext { return s.value }

func (s *SetTxMetaContext) SetValue(v IExpressionContext) { s.value = v }

func (s *SetTxMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTxMetaContext) SET_TX_META() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSET_TX_META, 0)
}

func (s *SetTxMetaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLPAREN, 0)
}

func (s *SetTxMetaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRPAREN, 0)
}

func (s *SetTxMetaContext) STRING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSTRING, 0)
}

func (s *SetTxMetaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SetTxMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSetTxMeta(s)
	}
}

func (s *SetTxMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSetTxMeta(s)
	}
}

type FailContext struct {
	*StatementContext
}

func NewFailContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FailContext {
	var p = new(FailContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailContext) FAIL() antlr.TerminalNode {
	return s.GetToken(NumScriptParserFAIL, 0)
}

func (s *FailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterFail(s)
	}
}

func (s *FailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitFail(s)
	}
}

type SendContext struct {
	*StatementContext
	mon    IExpressionContext
	monAll IMonetaryAllContext
	src    IValueAwareSourceContext
	dest   IDestinationContext
}

func NewSendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendContext {
	var p = new(SendContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SendContext) GetMon() IExpressionContext { return s.mon }

func (s *SendContext) GetMonAll() IMonetaryAllContext { return s.monAll }

func (s *SendContext) GetSrc() IValueAwareSourceContext { return s.src }

func (s *SendContext) GetDest() IDestinationContext { return s.dest }

func (s *SendContext) SetMon(v IExpressionContext) { s.mon = v }

func (s *SendContext) SetMonAll(v IMonetaryAllContext) { s.monAll = v }

func (s *SendContext) SetSrc(v IValueAwareSourceContext) { s.src = v }

func (s *SendContext) SetDest(v IDestinationContext) { s.dest = v }

func (s *SendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendContext) SEND() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSEND, 0)
}

func (s *SendContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLPAREN, 0)
}

func (s *SendContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *SendContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *SendContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRPAREN, 0)
}

func (s *SendContext) SOURCE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSOURCE, 0)
}

func (s *SendContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserEQ)
}

func (s *SendContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserEQ, i)
}

func (s *SendContext) DESTINATION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserDESTINATION, 0)
}

func (s *SendContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SendContext) MonetaryAll() IMonetaryAllContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryAllContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryAllContext)
}

func (s *SendContext) ValueAwareSource() IValueAwareSourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueAwareSourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueAwareSourceContext)
}

func (s *SendContext) Destination() IDestinationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
}

func (s *SendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSend(s)
	}
}

func (s *SendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSend(s)
	}
}

func (p *NumScriptParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NumScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(205)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(167)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserSET_TX_META:
		localctx = NewSetTxMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Match(NumScriptParserSET_TX_META)
		}
		{
			p.SetState(169)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(170)

			var _m = p.Match(NumScriptParserSTRING)

			localctx.(*SetTxMetaContext).key = _m
		}
		{
			p.SetState(171)
			p.Match(NumScriptParserT__1)
		}
		{
			p.SetState(172)

			var _x = p.expression(0)

			localctx.(*SetTxMetaContext).value = _x
		}
		{
			p.SetState(173)
			p.Match(NumScriptParserRPAREN)
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(NumScriptParserSEND)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(177)

				var _x = p.expression(0)

				localctx.(*SendContext).mon = _x
			}

		case 2:
			{
				p.SetState(178)

				var _x = p.MonetaryAll()

				localctx.(*SendContext).monAll = _x
			}

		}
		{
			p.SetState(181)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(182)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(183)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(184)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(185)

				var _x = p.ValueAwareSource()

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(186)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(187)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(188)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(189)

				var _x = p.Destination()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(191)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(192)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(193)

				var _x = p.Destination()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(194)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(195)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(196)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(197)

				var _x = p.ValueAwareSource()

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(201)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(202)
			p.Match(NumScriptParserRPAREN)
		}

	case NumScriptParserIF:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(204)

			var _x = p.IfStatement()

			localctx.(*IfStmtContext).stmt = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TY_ACCOUNT() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_ACCOUNT, 0)
}

func (s *Type_Context) TY_ASSET() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_ASSET, 0)
}

func (s *Type_Context) TY_NUMBER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_NUMBER, 0)
}

func (s *Type_Context) TY_STRING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_STRING, 0)
}

func (s *Type_Context) TY_MONETARY() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_MONETARY, 0)
}

func (s *Type_Context) TY_PORTION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_PORTION, 0)
}

func (s *Type_Context) TY_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_BOOLEAN, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *NumScriptParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NumScriptParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(NumScriptParserTY_ACCOUNT-29))|(1<<(NumScriptParserTY_ASSET-29))|(1<<(NumScriptParserTY_NUMBER-29))|(1<<(NumScriptParserTY_MONETARY-29))|(1<<(NumScriptParserTY_PORTION-29))|(1<<(NumScriptParserTY_BOOLEAN-29))|(1<<(NumScriptParserTY_STRING-29)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOriginContext is an interface to support dynamic dispatch.
type IOriginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetAcc returns the acc rule contexts.
	GetAcc() IExpressionContext

	// SetAcc sets the acc rule contexts.
	SetAcc(IExpressionContext)

	// IsOriginContext differentiates from other interfaces.
	IsOriginContext()
}

type OriginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	acc    IExpressionContext
	key    antlr.Token
}

func NewEmptyOriginContext() *OriginContext {
	var p = new(OriginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_origin
	return p
}

func (*OriginContext) IsOriginContext() {}

func NewOriginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OriginContext {
	var p = new(OriginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_origin

	return p
}

func (s *OriginContext) GetParser() antlr.Parser { return s.parser }

func (s *OriginContext) GetKey() antlr.Token { return s.key }

func (s *OriginContext) SetKey(v antlr.Token) { s.key = v }

func (s *OriginContext) GetAcc() IExpressionContext { return s.acc }

func (s *OriginContext) SetAcc(v IExpressionContext) { s.acc = v }

func (s *OriginContext) META() antlr.TerminalNode {
	return s.GetToken(NumScriptParserMETA, 0)
}

func (s *OriginContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLPAREN, 0)
}

func (s *OriginContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRPAREN, 0)
}

func (s *OriginContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OriginContext) STRING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSTRING, 0)
}

func (s *OriginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OriginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OriginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterOrigin(s)
	}
}

func (s *OriginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitOrigin(s)
	}
}

func (p *NumScriptParser) Origin() (localctx IOriginContext) {
	this := p
	_ = this

	localctx = NewOriginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NumScriptParserRULE_origin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(NumScriptParserMETA)
	}
	{
		p.SetState(210)
		p.Match(NumScriptParserLPAREN)
	}
	{
		p.SetState(211)

		var _x = p.expression(0)

		localctx.(*OriginContext).acc = _x
	}
	{
		p.SetState(212)
		p.Match(NumScriptParserT__1)
	}
	{
		p.SetState(213)

		var _m = p.Match(NumScriptParserSTRING)

		localctx.(*OriginContext).key = _m
	}
	{
		p.SetState(214)
		p.Match(NumScriptParserRPAREN)
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() IType_Context

	// GetName returns the name rule contexts.
	GetName() IVariableContext

	// GetOrig returns the orig rule contexts.
	GetOrig() IOriginContext

	// SetTy sets the ty rule contexts.
	SetTy(IType_Context)

	// SetName sets the name rule contexts.
	SetName(IVariableContext)

	// SetOrig sets the orig rule contexts.
	SetOrig(IOriginContext)

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     IType_Context
	name   IVariableContext
	orig   IOriginContext
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) GetTy() IType_Context { return s.ty }

func (s *VarDeclContext) GetName() IVariableContext { return s.name }

func (s *VarDeclContext) GetOrig() IOriginContext { return s.orig }

func (s *VarDeclContext) SetTy(v IType_Context) { s.ty = v }

func (s *VarDeclContext) SetName(v IVariableContext) { s.name = v }

func (s *VarDeclContext) SetOrig(v IOriginContext) { s.orig = v }

func (s *VarDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDeclContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarDeclContext) EQ() antlr.TerminalNode {
	return s.GetToken(NumScriptParserEQ, 0)
}

func (s *VarDeclContext) Origin() IOriginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOriginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOriginContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *NumScriptParser) VarDecl() (localctx IVarDeclContext) {
	this := p
	_ = this

	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NumScriptParserRULE_varDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(217)

		var _x = p.Variable()

		localctx.(*VarDeclContext).name = _x
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserEQ {
		{
			p.SetState(218)
			p.Match(NumScriptParserEQ)
		}
		{
			p.SetState(219)

			var _x = p.Origin()

			localctx.(*VarDeclContext).orig = _x
		}

	}

	return localctx
}

// IVarListDeclContext is an interface to support dynamic dispatch.
type IVarListDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_varDecl returns the _varDecl rule contexts.
	Get_varDecl() IVarDeclContext

	// Set_varDecl sets the _varDecl rule contexts.
	Set_varDecl(IVarDeclContext)

	// GetV returns the v rule context list.
	GetV() []IVarDeclContext

	// SetV sets the v rule context list.
	SetV([]IVarDeclContext)

	// IsVarListDeclContext differentiates from other interfaces.
	IsVarListDeclContext()
}

type VarListDeclContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_varDecl IVarDeclContext
	v        []IVarDeclContext
}

func NewEmptyVarListDeclContext() *VarListDeclContext {
	var p = new(VarListDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_varListDecl
	return p
}

func (*VarListDeclContext) IsVarListDeclContext() {}

func NewVarListDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarListDeclContext {
	var p = new(VarListDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_varListDecl

	return p
}

func (s *VarListDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarListDeclContext) Get_varDecl() IVarDeclContext { return s._varDecl }

func (s *VarListDeclContext) Set_varDecl(v IVarDeclContext) { s._varDecl = v }

func (s *VarListDeclContext) GetV() []IVarDeclContext { return s.v }

func (s *VarListDeclContext) SetV(v []IVarDeclContext) { s.v = v }

func (s *VarListDeclContext) VARS() antlr.TerminalNode {
	return s.GetToken(NumScriptParserVARS, 0)
}

func (s *VarListDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *VarListDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *VarListDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *VarListDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *VarListDeclContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *VarListDeclContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarListDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarListDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarListDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterVarListDecl(s)
	}
}

func (s *VarListDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitVarListDecl(s)
	}
}

func (p *NumScriptParser) VarListDecl() (localctx IVarListDeclContext) {
	this := p
	_ = this

	localctx = NewVarListDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NumScriptParserRULE_varListDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(223)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(224)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(NumScriptParserTY_ACCOUNT-29))|(1<<(NumScriptParserTY_ASSET-29))|(1<<(NumScriptParserTY_NUMBER-29))|(1<<(NumScriptParserTY_MONETARY-29))|(1<<(NumScriptParserTY_PORTION-29))|(1<<(NumScriptParserTY_BOOLEAN-29))|(1<<(NumScriptParserTY_STRING-29)))) != 0) {
		{
			p.SetState(225)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(226)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)
		p.Match(NumScriptParserRBRACE)
	}
	{
		p.SetState(236)
		p.Match(NumScriptParserNEWLINE)
	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVars returns the vars rule contexts.
	GetVars() IVarListDeclContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// SetVars sets the vars rule contexts.
	SetVars(IVarListDeclContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStmts returns the stmts rule context list.
	GetStmts() []IStatementContext

	// SetStmts sets the stmts rule context list.
	SetStmts([]IStatementContext)

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	vars       IVarListDeclContext
	_statement IStatementContext
	stmts      []IStatementContext
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) GetVars() IVarListDeclContext { return s.vars }

func (s *ScriptContext) Get_statement() IStatementContext { return s._statement }

func (s *ScriptContext) SetVars(v IVarListDeclContext) { s.vars = v }

func (s *ScriptContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ScriptContext) GetStmts() []IStatementContext { return s.stmts }

func (s *ScriptContext) SetStmts(v []IStatementContext) { s.stmts = v }

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(NumScriptParserEOF, 0)
}

func (s *ScriptContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ScriptContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ScriptContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *ScriptContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *ScriptContext) VarListDecl() IVarListDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarListDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarListDeclContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *NumScriptParser) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NumScriptParserRULE_script)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(238)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(244)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(247)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(248)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(249)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(255)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)
		p.Match(NumScriptParserEOF)
	}

	return localctx
}

func (p *NumScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NumScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

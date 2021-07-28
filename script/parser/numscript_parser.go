// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 220,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 5, 2, 41, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 59, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12, 6, 14, 6,
	67, 11, 6, 3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 6, 8, 80, 10, 8, 13, 8, 14, 8, 81, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 5, 9, 89, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6,
	10, 98, 10, 10, 13, 10, 14, 10, 99, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	5, 11, 107, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 114, 10,
	12, 13, 12, 14, 12, 115, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 122, 10, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 5, 14, 147, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 152, 10,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 166, 10, 17, 3, 18, 3, 18, 3, 18, 6, 18, 171, 10,
	18, 13, 18, 14, 18, 172, 3, 18, 3, 18, 6, 18, 177, 10, 18, 13, 18, 14,
	18, 178, 6, 18, 181, 10, 18, 13, 18, 14, 18, 182, 3, 18, 3, 18, 6, 18,
	187, 10, 18, 13, 18, 14, 18, 188, 3, 19, 7, 19, 192, 10, 19, 12, 19, 14,
	19, 195, 11, 19, 3, 19, 5, 19, 198, 10, 19, 3, 19, 3, 19, 6, 19, 202, 10,
	19, 13, 19, 14, 19, 203, 3, 19, 7, 19, 207, 10, 19, 12, 19, 14, 19, 210,
	11, 19, 3, 19, 7, 19, 213, 10, 19, 12, 19, 14, 19, 216, 11, 19, 3, 19,
	3, 19, 3, 19, 2, 3, 10, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 2, 4, 3, 2, 17, 18, 3, 2, 27, 31, 2, 229, 2, 40,
	3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10,
	58, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 88, 3, 2, 2,
	2, 18, 90, 3, 2, 2, 2, 20, 106, 3, 2, 2, 2, 22, 108, 3, 2, 2, 2, 24, 121,
	3, 2, 2, 2, 26, 151, 3, 2, 2, 2, 28, 153, 3, 2, 2, 2, 30, 155, 3, 2, 2,
	2, 32, 161, 3, 2, 2, 2, 34, 167, 3, 2, 2, 2, 36, 193, 3, 2, 2, 2, 38, 41,
	7, 35, 2, 2, 39, 41, 7, 25, 2, 2, 40, 38, 3, 2, 2, 2, 40, 39, 3, 2, 2,
	2, 41, 3, 3, 2, 2, 2, 42, 43, 7, 21, 2, 2, 43, 44, 7, 39, 2, 2, 44, 45,
	5, 2, 2, 2, 45, 46, 7, 22, 2, 2, 46, 5, 3, 2, 2, 2, 47, 52, 7, 38, 2, 2,
	48, 52, 7, 39, 2, 2, 49, 52, 7, 35, 2, 2, 50, 52, 5, 4, 3, 2, 51, 47, 3,
	2, 2, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	7, 3, 2, 2, 2, 53, 54, 7, 37, 2, 2, 54, 9, 3, 2, 2, 2, 55, 56, 8, 6, 1,
	2, 56, 59, 5, 6, 4, 2, 57, 59, 5, 8, 5, 2, 58, 55, 3, 2, 2, 2, 58, 57,
	3, 2, 2, 2, 59, 65, 3, 2, 2, 2, 60, 61, 12, 5, 2, 2, 61, 62, 9, 2, 2, 2,
	62, 64, 5, 10, 6, 6, 63, 60, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3,
	2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68,
	71, 7, 33, 2, 2, 69, 71, 7, 34, 2, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3, 2,
	2, 2, 71, 13, 3, 2, 2, 2, 72, 73, 7, 23, 2, 2, 73, 79, 7, 6, 2, 2, 74,
	75, 5, 12, 7, 2, 75, 76, 7, 3, 2, 2, 76, 77, 5, 10, 6, 2, 77, 78, 7, 6,
	2, 2, 78, 80, 3, 2, 2, 2, 79, 74, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 7, 24, 2, 2,
	84, 15, 3, 2, 2, 2, 85, 89, 7, 33, 2, 2, 86, 89, 5, 8, 5, 2, 87, 89, 7,
	34, 2, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89,
	17, 3, 2, 2, 2, 90, 91, 7, 23, 2, 2, 91, 97, 7, 6, 2, 2, 92, 93, 5, 16,
	9, 2, 93, 94, 7, 3, 2, 2, 94, 95, 5, 10, 6, 2, 95, 96, 7, 6, 2, 2, 96,
	98, 3, 2, 2, 2, 97, 92, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2,
	2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 7, 24, 2, 2, 102,
	19, 3, 2, 2, 2, 103, 107, 5, 14, 8, 2, 104, 107, 5, 18, 10, 2, 105, 107,
	5, 10, 6, 2, 106, 103, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2,
	2, 2, 107, 21, 3, 2, 2, 2, 108, 109, 7, 23, 2, 2, 109, 113, 7, 6, 2, 2,
	110, 111, 5, 10, 6, 2, 111, 112, 7, 6, 2, 2, 112, 114, 3, 2, 2, 2, 113,
	110, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116,
	3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 7, 24, 2, 2, 118, 23, 3, 2,
	2, 2, 119, 122, 5, 22, 12, 2, 120, 122, 5, 10, 6, 2, 121, 119, 3, 2, 2,
	2, 121, 120, 3, 2, 2, 2, 122, 25, 3, 2, 2, 2, 123, 124, 7, 11, 2, 2, 124,
	152, 5, 10, 6, 2, 125, 152, 7, 12, 2, 2, 126, 127, 7, 13, 2, 2, 127, 128,
	5, 10, 6, 2, 128, 129, 7, 19, 2, 2, 129, 146, 7, 6, 2, 2, 130, 131, 7,
	14, 2, 2, 131, 132, 7, 26, 2, 2, 132, 133, 5, 24, 13, 2, 133, 134, 7, 6,
	2, 2, 134, 135, 7, 15, 2, 2, 135, 136, 7, 26, 2, 2, 136, 137, 5, 20, 11,
	2, 137, 147, 3, 2, 2, 2, 138, 139, 7, 15, 2, 2, 139, 140, 7, 26, 2, 2,
	140, 141, 5, 20, 11, 2, 141, 142, 7, 6, 2, 2, 142, 143, 7, 14, 2, 2, 143,
	144, 7, 26, 2, 2, 144, 145, 5, 24, 13, 2, 145, 147, 3, 2, 2, 2, 146, 130,
	3, 2, 2, 2, 146, 138, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 7, 6,
	2, 2, 149, 150, 7, 20, 2, 2, 150, 152, 3, 2, 2, 2, 151, 123, 3, 2, 2, 2,
	151, 125, 3, 2, 2, 2, 151, 126, 3, 2, 2, 2, 152, 27, 3, 2, 2, 2, 153, 154,
	9, 3, 2, 2, 154, 29, 3, 2, 2, 2, 155, 156, 7, 4, 2, 2, 156, 157, 5, 10,
	6, 2, 157, 158, 7, 5, 2, 2, 158, 159, 7, 32, 2, 2, 159, 160, 7, 20, 2,
	2, 160, 31, 3, 2, 2, 2, 161, 162, 5, 28, 15, 2, 162, 165, 5, 8, 5, 2, 163,
	164, 7, 26, 2, 2, 164, 166, 5, 30, 16, 2, 165, 163, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 33, 3, 2, 2, 2, 167, 168, 7, 10, 2, 2, 168, 170, 7, 23,
	2, 2, 169, 171, 7, 6, 2, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 180, 3, 2, 2, 2, 174,
	176, 5, 32, 17, 2, 175, 177, 7, 6, 2, 2, 176, 175, 3, 2, 2, 2, 177, 178,
	3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2,
	2, 2, 180, 174, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2,
	182, 183, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 7, 24, 2, 2, 185,
	187, 7, 6, 2, 2, 186, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 186,
	3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 35, 3, 2, 2, 2, 190, 192, 7, 6,
	2, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2,
	193, 194, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 196,
	198, 5, 34, 18, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199,
	3, 2, 2, 2, 199, 208, 5, 26, 14, 2, 200, 202, 7, 6, 2, 2, 201, 200, 3,
	2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2,
	2, 204, 205, 3, 2, 2, 2, 205, 207, 5, 26, 14, 2, 206, 201, 3, 2, 2, 2,
	207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209,
	214, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 213, 7, 6, 2, 2, 212, 211,
	3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2,
	2, 2, 215, 217, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 218, 7, 2, 2, 3,
	218, 37, 3, 2, 2, 2, 25, 40, 51, 58, 65, 70, 81, 88, 99, 106, 115, 121,
	146, 151, 165, 172, 178, 182, 188, 193, 197, 203, 208, 214,
}
var literalNames = []string{
	"", "'to'", "'metadata('", "','", "", "", "", "", "'vars'", "'print'",
	"'fail'", "'send'", "'source'", "'destination'", "'allocate'", "'+'", "'-'",
	"'('", "')'", "'['", "']'", "'{'", "'}'", "'*'", "'='", "'account'", "'asset'",
	"'number'", "'monetary'", "'portion'", "", "", "'remaining'", "", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "PRINT", "FAIL", "SEND", "SOURCE", "DESTINATION", "ALLOCATE", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "ALL",
	"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
	"STRING", "PORTION", "PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"amount", "monetary", "literal", "variable", "expression", "allocPartConst",
	"allocBlockConst", "allocPartDyn", "allocBlockDyn", "allocation", "sourceBlock",
	"source", "statement", "type_", "origin", "varDecl", "varListDecl", "script",
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
	NumScriptParserT__2              = 3
	NumScriptParserNEWLINE           = 4
	NumScriptParserWHITESPACE        = 5
	NumScriptParserMULTILINE_COMMENT = 6
	NumScriptParserLINE_COMMENT      = 7
	NumScriptParserVARS              = 8
	NumScriptParserPRINT             = 9
	NumScriptParserFAIL              = 10
	NumScriptParserSEND              = 11
	NumScriptParserSOURCE            = 12
	NumScriptParserDESTINATION       = 13
	NumScriptParserALLOCATE          = 14
	NumScriptParserOP_ADD            = 15
	NumScriptParserOP_SUB            = 16
	NumScriptParserLPAREN            = 17
	NumScriptParserRPAREN            = 18
	NumScriptParserLBRACK            = 19
	NumScriptParserRBRACK            = 20
	NumScriptParserLBRACE            = 21
	NumScriptParserRBRACE            = 22
	NumScriptParserALL               = 23
	NumScriptParserEQ                = 24
	NumScriptParserTY_ACCOUNT        = 25
	NumScriptParserTY_ASSET          = 26
	NumScriptParserTY_NUMBER         = 27
	NumScriptParserTY_MONETARY       = 28
	NumScriptParserTY_PORTION        = 29
	NumScriptParserSTRING            = 30
	NumScriptParserPORTION           = 31
	NumScriptParserPORTION_REMAINING = 32
	NumScriptParserNUMBER            = 33
	NumScriptParserPERCENT           = 34
	NumScriptParserVARIABLE_NAME     = 35
	NumScriptParserACCOUNT           = 36
	NumScriptParserASSET             = 37
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_amount          = 0
	NumScriptParserRULE_monetary        = 1
	NumScriptParserRULE_literal         = 2
	NumScriptParserRULE_variable        = 3
	NumScriptParserRULE_expression      = 4
	NumScriptParserRULE_allocPartConst  = 5
	NumScriptParserRULE_allocBlockConst = 6
	NumScriptParserRULE_allocPartDyn    = 7
	NumScriptParserRULE_allocBlockDyn   = 8
	NumScriptParserRULE_allocation      = 9
	NumScriptParserRULE_sourceBlock     = 10
	NumScriptParserRULE_source          = 11
	NumScriptParserRULE_statement       = 12
	NumScriptParserRULE_type_           = 13
	NumScriptParserRULE_origin          = 14
	NumScriptParserRULE_varDecl         = 15
	NumScriptParserRULE_varListDecl     = 16
	NumScriptParserRULE_script          = 17
)

// IAmountContext is an interface to support dynamic dispatch.
type IAmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmountContext differentiates from other interfaces.
	IsAmountContext()
}

type AmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmountContext() *AmountContext {
	var p = new(AmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_amount
	return p
}

func (*AmountContext) IsAmountContext() {}

func NewAmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmountContext {
	var p = new(AmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_amount

	return p
}

func (s *AmountContext) GetParser() antlr.Parser { return s.parser }

func (s *AmountContext) CopyFrom(ctx *AmountContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AmountAllContext struct {
	*AmountContext
}

func NewAmountAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AmountAllContext {
	var p = new(AmountAllContext)

	p.AmountContext = NewEmptyAmountContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AmountContext))

	return p
}

func (s *AmountAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountAllContext) ALL() antlr.TerminalNode {
	return s.GetToken(NumScriptParserALL, 0)
}

func (s *AmountAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAmountAll(s)
	}
}

func (s *AmountAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAmountAll(s)
	}
}

type AmountSpecificContext struct {
	*AmountContext
	num antlr.Token
}

func NewAmountSpecificContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AmountSpecificContext {
	var p = new(AmountSpecificContext)

	p.AmountContext = NewEmptyAmountContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AmountContext))

	return p
}

func (s *AmountSpecificContext) GetNum() antlr.Token { return s.num }

func (s *AmountSpecificContext) SetNum(v antlr.Token) { s.num = v }

func (s *AmountSpecificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountSpecificContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserNUMBER, 0)
}

func (s *AmountSpecificContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAmountSpecific(s)
	}
}

func (s *AmountSpecificContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAmountSpecific(s)
	}
}

func (p *NumScriptParser) Amount() (localctx IAmountContext) {
	localctx = NewAmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NumScriptParserRULE_amount)

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

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserNUMBER:
		localctx = NewAmountSpecificContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*AmountSpecificContext).num = _m
		}

	case NumScriptParserALL:
		localctx = NewAmountAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(NumScriptParserALL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMonetaryContext is an interface to support dynamic dispatch.
type IMonetaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsset returns the asset token.
	GetAsset() antlr.Token

	// SetAsset sets the asset token.
	SetAsset(antlr.Token)

	// GetAmt returns the amt rule contexts.
	GetAmt() IAmountContext

	// SetAmt sets the amt rule contexts.
	SetAmt(IAmountContext)

	// IsMonetaryContext differentiates from other interfaces.
	IsMonetaryContext()
}

type MonetaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	asset  antlr.Token
	amt    IAmountContext
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

func (s *MonetaryContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryContext) GetAmt() IAmountContext { return s.amt }

func (s *MonetaryContext) SetAmt(v IAmountContext) { s.amt = v }

func (s *MonetaryContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACK, 0)
}

func (s *MonetaryContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACK, 0)
}

func (s *MonetaryContext) ASSET() antlr.TerminalNode {
	return s.GetToken(NumScriptParserASSET, 0)
}

func (s *MonetaryContext) Amount() IAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmountContext)
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
	localctx = NewMonetaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NumScriptParserRULE_monetary)

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
		p.SetState(40)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(41)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(42)

		var _x = p.Amount()

		localctx.(*MonetaryContext).amt = _x
	}
	{
		p.SetState(43)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
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
		p.SetState(51)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(54)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)

			var _x = p.Variable()

			localctx.(*ExprVariableContext).var_ = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, NumScriptParserRULE_expression)
			p.SetState(58)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(59)

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
				p.SetState(60)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAllocPartConstContext is an interface to support dynamic dispatch.
type IAllocPartConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllocPartConstContext differentiates from other interfaces.
	IsAllocPartConstContext()
}

type AllocPartConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllocPartConstContext() *AllocPartConstContext {
	var p = new(AllocPartConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocPartConst
	return p
}

func (*AllocPartConstContext) IsAllocPartConstContext() {}

func NewAllocPartConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocPartConstContext {
	var p = new(AllocPartConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocPartConst

	return p
}

func (s *AllocPartConstContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocPartConstContext) CopyFrom(ctx *AllocPartConstContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AllocPartConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllocPartConstRemainingContext struct {
	*AllocPartConstContext
}

func NewAllocPartConstRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartConstRemainingContext {
	var p = new(AllocPartConstRemainingContext)

	p.AllocPartConstContext = NewEmptyAllocPartConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartConstContext))

	return p
}

func (s *AllocPartConstRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartConstRemainingContext) PORTION_REMAINING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION_REMAINING, 0)
}

func (s *AllocPartConstRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocPartConstRemaining(s)
	}
}

func (s *AllocPartConstRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocPartConstRemaining(s)
	}
}

type AllocPartConstConstContext struct {
	*AllocPartConstContext
}

func NewAllocPartConstConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartConstConstContext {
	var p = new(AllocPartConstConstContext)

	p.AllocPartConstContext = NewEmptyAllocPartConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartConstContext))

	return p
}

func (s *AllocPartConstConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartConstConstContext) PORTION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION, 0)
}

func (s *AllocPartConstConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocPartConstConst(s)
	}
}

func (s *AllocPartConstConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocPartConstConst(s)
	}
}

func (p *NumScriptParser) AllocPartConst() (localctx IAllocPartConstContext) {
	localctx = NewAllocPartConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NumScriptParserRULE_allocPartConst)

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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPORTION:
		localctx = NewAllocPartConstConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(NumScriptParserPORTION)
		}

	case NumScriptParserPORTION_REMAINING:
		localctx = NewAllocPartConstRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(NumScriptParserPORTION_REMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAllocBlockConstContext is an interface to support dynamic dispatch.
type IAllocBlockConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allocPartConst returns the _allocPartConst rule contexts.
	Get_allocPartConst() IAllocPartConstContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_allocPartConst sets the _allocPartConst rule contexts.
	Set_allocPartConst(IAllocPartConstContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllocPartConstContext

	// GetDests returns the dests rule context list.
	GetDests() []IExpressionContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllocPartConstContext)

	// SetDests sets the dests rule context list.
	SetDests([]IExpressionContext)

	// IsAllocBlockConstContext differentiates from other interfaces.
	IsAllocBlockConstContext()
}

type AllocBlockConstContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_allocPartConst IAllocPartConstContext
	portions        []IAllocPartConstContext
	_expression     IExpressionContext
	dests           []IExpressionContext
}

func NewEmptyAllocBlockConstContext() *AllocBlockConstContext {
	var p = new(AllocBlockConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocBlockConst
	return p
}

func (*AllocBlockConstContext) IsAllocBlockConstContext() {}

func NewAllocBlockConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocBlockConstContext {
	var p = new(AllocBlockConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocBlockConst

	return p
}

func (s *AllocBlockConstContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocBlockConstContext) Get_allocPartConst() IAllocPartConstContext {
	return s._allocPartConst
}

func (s *AllocBlockConstContext) Get_expression() IExpressionContext { return s._expression }

func (s *AllocBlockConstContext) Set_allocPartConst(v IAllocPartConstContext) { s._allocPartConst = v }

func (s *AllocBlockConstContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AllocBlockConstContext) GetPortions() []IAllocPartConstContext { return s.portions }

func (s *AllocBlockConstContext) GetDests() []IExpressionContext { return s.dests }

func (s *AllocBlockConstContext) SetPortions(v []IAllocPartConstContext) { s.portions = v }

func (s *AllocBlockConstContext) SetDests(v []IExpressionContext) { s.dests = v }

func (s *AllocBlockConstContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *AllocBlockConstContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *AllocBlockConstContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *AllocBlockConstContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *AllocBlockConstContext) AllAllocPartConst() []IAllocPartConstContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllocPartConstContext)(nil)).Elem())
	var tst = make([]IAllocPartConstContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllocPartConstContext)
		}
	}

	return tst
}

func (s *AllocBlockConstContext) AllocPartConst(i int) IAllocPartConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocPartConstContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllocPartConstContext)
}

func (s *AllocBlockConstContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AllocBlockConstContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllocBlockConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocBlockConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocBlockConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocBlockConst(s)
	}
}

func (s *AllocBlockConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocBlockConst(s)
	}
}

func (p *NumScriptParser) AllocBlockConst() (localctx IAllocBlockConstContext) {
	localctx = NewAllocBlockConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumScriptParserRULE_allocBlockConst)
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
		p.SetState(70)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(71)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserPORTION || _la == NumScriptParserPORTION_REMAINING {
		{
			p.SetState(72)

			var _x = p.AllocPartConst()

			localctx.(*AllocBlockConstContext)._allocPartConst = _x
		}
		localctx.(*AllocBlockConstContext).portions = append(localctx.(*AllocBlockConstContext).portions, localctx.(*AllocBlockConstContext)._allocPartConst)
		{
			p.SetState(73)
			p.Match(NumScriptParserT__0)
		}
		{
			p.SetState(74)

			var _x = p.expression(0)

			localctx.(*AllocBlockConstContext)._expression = _x
		}
		localctx.(*AllocBlockConstContext).dests = append(localctx.(*AllocBlockConstContext).dests, localctx.(*AllocBlockConstContext)._expression)
		{
			p.SetState(75)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IAllocPartDynContext is an interface to support dynamic dispatch.
type IAllocPartDynContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllocPartDynContext differentiates from other interfaces.
	IsAllocPartDynContext()
}

type AllocPartDynContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllocPartDynContext() *AllocPartDynContext {
	var p = new(AllocPartDynContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocPartDyn
	return p
}

func (*AllocPartDynContext) IsAllocPartDynContext() {}

func NewAllocPartDynContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocPartDynContext {
	var p = new(AllocPartDynContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocPartDyn

	return p
}

func (s *AllocPartDynContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocPartDynContext) CopyFrom(ctx *AllocPartDynContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AllocPartDynContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartDynContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllocPartDynVarContext struct {
	*AllocPartDynContext
	por IVariableContext
}

func NewAllocPartDynVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartDynVarContext {
	var p = new(AllocPartDynVarContext)

	p.AllocPartDynContext = NewEmptyAllocPartDynContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartDynContext))

	return p
}

func (s *AllocPartDynVarContext) GetPor() IVariableContext { return s.por }

func (s *AllocPartDynVarContext) SetPor(v IVariableContext) { s.por = v }

func (s *AllocPartDynVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartDynVarContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AllocPartDynVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocPartDynVar(s)
	}
}

func (s *AllocPartDynVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocPartDynVar(s)
	}
}

type AllocPartDynConstContext struct {
	*AllocPartDynContext
}

func NewAllocPartDynConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartDynConstContext {
	var p = new(AllocPartDynConstContext)

	p.AllocPartDynContext = NewEmptyAllocPartDynContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartDynContext))

	return p
}

func (s *AllocPartDynConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartDynConstContext) PORTION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION, 0)
}

func (s *AllocPartDynConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocPartDynConst(s)
	}
}

func (s *AllocPartDynConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocPartDynConst(s)
	}
}

type AllocPartDynRemainingContext struct {
	*AllocPartDynContext
}

func NewAllocPartDynRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartDynRemainingContext {
	var p = new(AllocPartDynRemainingContext)

	p.AllocPartDynContext = NewEmptyAllocPartDynContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartDynContext))

	return p
}

func (s *AllocPartDynRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartDynRemainingContext) PORTION_REMAINING() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPORTION_REMAINING, 0)
}

func (s *AllocPartDynRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocPartDynRemaining(s)
	}
}

func (s *AllocPartDynRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocPartDynRemaining(s)
	}
}

func (p *NumScriptParser) AllocPartDyn() (localctx IAllocPartDynContext) {
	localctx = NewAllocPartDynContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumScriptParserRULE_allocPartDyn)

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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPORTION:
		localctx = NewAllocPartDynConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(NumScriptParserPORTION)
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewAllocPartDynVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)

			var _x = p.Variable()

			localctx.(*AllocPartDynVarContext).por = _x
		}

	case NumScriptParserPORTION_REMAINING:
		localctx = NewAllocPartDynRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Match(NumScriptParserPORTION_REMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAllocBlockDynContext is an interface to support dynamic dispatch.
type IAllocBlockDynContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allocPartDyn returns the _allocPartDyn rule contexts.
	Get_allocPartDyn() IAllocPartDynContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_allocPartDyn sets the _allocPartDyn rule contexts.
	Set_allocPartDyn(IAllocPartDynContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllocPartDynContext

	// GetDests returns the dests rule context list.
	GetDests() []IExpressionContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllocPartDynContext)

	// SetDests sets the dests rule context list.
	SetDests([]IExpressionContext)

	// IsAllocBlockDynContext differentiates from other interfaces.
	IsAllocBlockDynContext()
}

type AllocBlockDynContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	_allocPartDyn IAllocPartDynContext
	portions      []IAllocPartDynContext
	_expression   IExpressionContext
	dests         []IExpressionContext
}

func NewEmptyAllocBlockDynContext() *AllocBlockDynContext {
	var p = new(AllocBlockDynContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocBlockDyn
	return p
}

func (*AllocBlockDynContext) IsAllocBlockDynContext() {}

func NewAllocBlockDynContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocBlockDynContext {
	var p = new(AllocBlockDynContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocBlockDyn

	return p
}

func (s *AllocBlockDynContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocBlockDynContext) Get_allocPartDyn() IAllocPartDynContext { return s._allocPartDyn }

func (s *AllocBlockDynContext) Get_expression() IExpressionContext { return s._expression }

func (s *AllocBlockDynContext) Set_allocPartDyn(v IAllocPartDynContext) { s._allocPartDyn = v }

func (s *AllocBlockDynContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AllocBlockDynContext) GetPortions() []IAllocPartDynContext { return s.portions }

func (s *AllocBlockDynContext) GetDests() []IExpressionContext { return s.dests }

func (s *AllocBlockDynContext) SetPortions(v []IAllocPartDynContext) { s.portions = v }

func (s *AllocBlockDynContext) SetDests(v []IExpressionContext) { s.dests = v }

func (s *AllocBlockDynContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *AllocBlockDynContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *AllocBlockDynContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *AllocBlockDynContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *AllocBlockDynContext) AllAllocPartDyn() []IAllocPartDynContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllocPartDynContext)(nil)).Elem())
	var tst = make([]IAllocPartDynContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllocPartDynContext)
		}
	}

	return tst
}

func (s *AllocBlockDynContext) AllocPartDyn(i int) IAllocPartDynContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocPartDynContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllocPartDynContext)
}

func (s *AllocBlockDynContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AllocBlockDynContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllocBlockDynContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocBlockDynContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocBlockDynContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocBlockDyn(s)
	}
}

func (s *AllocBlockDynContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocBlockDyn(s)
	}
}

func (p *NumScriptParser) AllocBlockDyn() (localctx IAllocBlockDynContext) {
	localctx = NewAllocBlockDynContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumScriptParserRULE_allocBlockDyn)
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
		p.SetState(88)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(89)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(NumScriptParserPORTION-31))|(1<<(NumScriptParserPORTION_REMAINING-31))|(1<<(NumScriptParserVARIABLE_NAME-31)))) != 0) {
		{
			p.SetState(90)

			var _x = p.AllocPartDyn()

			localctx.(*AllocBlockDynContext)._allocPartDyn = _x
		}
		localctx.(*AllocBlockDynContext).portions = append(localctx.(*AllocBlockDynContext).portions, localctx.(*AllocBlockDynContext)._allocPartDyn)
		{
			p.SetState(91)
			p.Match(NumScriptParserT__0)
		}
		{
			p.SetState(92)

			var _x = p.expression(0)

			localctx.(*AllocBlockDynContext)._expression = _x
		}
		localctx.(*AllocBlockDynContext).dests = append(localctx.(*AllocBlockDynContext).dests, localctx.(*AllocBlockDynContext)._expression)
		{
			p.SetState(93)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// IAllocationContext is an interface to support dynamic dispatch.
type IAllocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllocationContext differentiates from other interfaces.
	IsAllocationContext()
}

type AllocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllocationContext() *AllocationContext {
	var p = new(AllocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocation
	return p
}

func (*AllocationContext) IsAllocationContext() {}

func NewAllocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocationContext {
	var p = new(AllocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocation

	return p
}

func (s *AllocationContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocationContext) CopyFrom(ctx *AllocationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AllocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllocConstContext struct {
	*AllocationContext
}

func NewAllocConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocConstContext {
	var p = new(AllocConstContext)

	p.AllocationContext = NewEmptyAllocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocationContext))

	return p
}

func (s *AllocConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocConstContext) AllocBlockConst() IAllocBlockConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocBlockConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllocBlockConstContext)
}

func (s *AllocConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocConst(s)
	}
}

func (s *AllocConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocConst(s)
	}
}

type AllocAccountContext struct {
	*AllocationContext
}

func NewAllocAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocAccountContext {
	var p = new(AllocAccountContext)

	p.AllocationContext = NewEmptyAllocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocationContext))

	return p
}

func (s *AllocAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocAccountContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllocAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocAccount(s)
	}
}

func (s *AllocAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocAccount(s)
	}
}

type AllocDynContext struct {
	*AllocationContext
}

func NewAllocDynContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocDynContext {
	var p = new(AllocDynContext)

	p.AllocationContext = NewEmptyAllocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocationContext))

	return p
}

func (s *AllocDynContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocDynContext) AllocBlockDyn() IAllocBlockDynContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocBlockDynContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllocBlockDynContext)
}

func (s *AllocDynContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocDyn(s)
	}
}

func (s *AllocDynContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocDyn(s)
	}
}

func (p *NumScriptParser) Allocation() (localctx IAllocationContext) {
	localctx = NewAllocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NumScriptParserRULE_allocation)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAllocConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.AllocBlockConst()
		}

	case 2:
		localctx = NewAllocDynContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.AllocBlockDyn()
		}

	case 3:
		localctx = NewAllocAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.expression(0)
		}

	}

	return localctx
}

// ISourceBlockContext is an interface to support dynamic dispatch.
type ISourceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetSources returns the sources rule context list.
	GetSources() []IExpressionContext

	// SetSources sets the sources rule context list.
	SetSources([]IExpressionContext)

	// IsSourceBlockContext differentiates from other interfaces.
	IsSourceBlockContext()
}

type SourceBlockContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_expression IExpressionContext
	sources     []IExpressionContext
}

func NewEmptySourceBlockContext() *SourceBlockContext {
	var p = new(SourceBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_sourceBlock
	return p
}

func (*SourceBlockContext) IsSourceBlockContext() {}

func NewSourceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceBlockContext {
	var p = new(SourceBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_sourceBlock

	return p
}

func (s *SourceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceBlockContext) Get_expression() IExpressionContext { return s._expression }

func (s *SourceBlockContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *SourceBlockContext) GetSources() []IExpressionContext { return s.sources }

func (s *SourceBlockContext) SetSources(v []IExpressionContext) { s.sources = v }

func (s *SourceBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *SourceBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *SourceBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *SourceBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *SourceBlockContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SourceBlockContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SourceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSourceBlock(s)
	}
}

func (s *SourceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSourceBlock(s)
	}
}

func (p *NumScriptParser) SourceBlock() (localctx ISourceBlockContext) {
	localctx = NewSourceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NumScriptParserRULE_sourceBlock)
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
		p.SetState(106)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(107)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(NumScriptParserLBRACK-19))|(1<<(NumScriptParserNUMBER-19))|(1<<(NumScriptParserVARIABLE_NAME-19))|(1<<(NumScriptParserACCOUNT-19))|(1<<(NumScriptParserASSET-19)))) != 0) {
		{
			p.SetState(108)

			var _x = p.expression(0)

			localctx.(*SourceBlockContext)._expression = _x
		}
		localctx.(*SourceBlockContext).sources = append(localctx.(*SourceBlockContext).sources, localctx.(*SourceBlockContext)._expression)
		{
			p.SetState(109)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(NumScriptParserRBRACE)
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

type SrcBlockContext struct {
	*SourceContext
}

func NewSrcBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcBlockContext {
	var p = new(SrcBlockContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcBlockContext) SourceBlock() ISourceBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceBlockContext)
}

func (s *SrcBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcBlock(s)
	}
}

func (s *SrcBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcBlock(s)
	}
}

func (p *NumScriptParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NumScriptParserRULE_source)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACE:
		localctx = NewSrcBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.SourceBlock()
		}

	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	mon  IExpressionContext
	src  ISourceContext
	dest IAllocationContext
}

func NewSendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendContext {
	var p = new(SendContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SendContext) GetMon() IExpressionContext { return s.mon }

func (s *SendContext) GetSrc() ISourceContext { return s.src }

func (s *SendContext) GetDest() IAllocationContext { return s.dest }

func (s *SendContext) SetMon(v IExpressionContext) { s.mon = v }

func (s *SendContext) SetSrc(v ISourceContext) { s.src = v }

func (s *SendContext) SetDest(v IAllocationContext) { s.dest = v }

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

func (s *SendContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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

func (s *SendContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SendContext) Allocation() IAllocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllocationContext)
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
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NumScriptParserRULE_statement)

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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(122)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(125)

			var _x = p.expression(0)

			localctx.(*SendContext).mon = _x
		}
		{
			p.SetState(126)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(127)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(128)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(129)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(130)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(131)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(132)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(133)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(134)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(136)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(137)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(138)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(139)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(140)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(141)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(142)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(146)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(147)
			p.Match(NumScriptParserRPAREN)
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

func (s *Type_Context) TY_MONETARY() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_MONETARY, 0)
}

func (s *Type_Context) TY_PORTION() antlr.TerminalNode {
	return s.GetToken(NumScriptParserTY_PORTION, 0)
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
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NumScriptParserRULE_type_)
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
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY)|(1<<NumScriptParserTY_PORTION))) != 0) {
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
	localctx = NewOriginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NumScriptParserRULE_origin)

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
		p.SetState(153)
		p.Match(NumScriptParserT__1)
	}
	{
		p.SetState(154)

		var _x = p.expression(0)

		localctx.(*OriginContext).acc = _x
	}
	{
		p.SetState(155)
		p.Match(NumScriptParserT__2)
	}
	{
		p.SetState(156)

		var _m = p.Match(NumScriptParserSTRING)

		localctx.(*OriginContext).key = _m
	}
	{
		p.SetState(157)
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
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NumScriptParserRULE_varDecl)
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
		p.SetState(159)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(160)

		var _x = p.Variable()

		localctx.(*VarDeclContext).name = _x
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserEQ {
		{
			p.SetState(161)
			p.Match(NumScriptParserEQ)
		}
		{
			p.SetState(162)

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

func (s *VarListDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *VarListDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *VarListDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
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
	localctx = NewVarListDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NumScriptParserRULE_varListDecl)
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
		p.SetState(165)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(166)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(167)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY)|(1<<NumScriptParserTY_PORTION))) != 0) {
		{
			p.SetState(172)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(173)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(NumScriptParserRBRACE)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(183)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NumScriptParserRULE_script)
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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(188)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(194)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(197)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
				{
					p.SetState(198)
					p.Match(NumScriptParserNEWLINE)
				}

				p.SetState(201)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(203)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(209)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(215)
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
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

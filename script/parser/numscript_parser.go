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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 226,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 5, 2, 45, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10,
	4, 3, 4, 5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 65, 10, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 72, 10, 7, 3, 7, 3, 7, 3, 7, 7, 7, 77, 10,
	7, 12, 7, 14, 7, 80, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 95, 10, 11, 13, 11, 14,
	11, 96, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 104, 10, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 6, 13, 113, 10, 13, 13, 13, 14,
	13, 114, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 122, 10, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 6, 15, 129, 10, 15, 13, 15, 14, 15, 130, 3,
	15, 3, 15, 3, 16, 3, 16, 5, 16, 137, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 162,
	10, 17, 3, 17, 3, 17, 3, 17, 5, 17, 167, 10, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 6, 20, 177, 10, 20, 13, 20, 14, 20, 178,
	3, 20, 3, 20, 6, 20, 183, 10, 20, 13, 20, 14, 20, 184, 6, 20, 187, 10,
	20, 13, 20, 14, 20, 188, 3, 20, 3, 20, 6, 20, 193, 10, 20, 13, 20, 14,
	20, 194, 3, 21, 7, 21, 198, 10, 21, 12, 21, 14, 21, 201, 11, 21, 3, 21,
	5, 21, 204, 10, 21, 3, 21, 3, 21, 6, 21, 208, 10, 21, 13, 21, 14, 21, 209,
	3, 21, 7, 21, 213, 10, 21, 12, 21, 14, 21, 216, 11, 21, 3, 21, 7, 21, 219,
	10, 21, 12, 21, 14, 21, 222, 11, 21, 3, 21, 3, 21, 3, 21, 2, 3, 12, 22,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 2, 4, 3, 2, 15, 16, 3, 2, 25, 29, 2, 233, 2, 44, 3, 2, 2, 2, 4, 46,
	3, 2, 2, 2, 6, 58, 3, 2, 2, 2, 8, 64, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12,
	71, 3, 2, 2, 2, 14, 81, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 85, 3, 2, 2,
	2, 20, 87, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 105, 3, 2, 2, 2, 26, 121,
	3, 2, 2, 2, 28, 123, 3, 2, 2, 2, 30, 136, 3, 2, 2, 2, 32, 166, 3, 2, 2,
	2, 34, 168, 3, 2, 2, 2, 36, 170, 3, 2, 2, 2, 38, 173, 3, 2, 2, 2, 40, 199,
	3, 2, 2, 2, 42, 45, 7, 31, 2, 2, 43, 45, 7, 23, 2, 2, 44, 42, 3, 2, 2,
	2, 44, 43, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 47, 7, 19, 2, 2, 47, 48,
	7, 36, 2, 2, 48, 49, 5, 2, 2, 2, 49, 50, 7, 20, 2, 2, 50, 5, 3, 2, 2, 2,
	51, 59, 7, 30, 2, 2, 52, 55, 7, 31, 2, 2, 53, 54, 7, 3, 2, 2, 54, 56, 7,
	31, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	59, 7, 32, 2, 2, 58, 51, 3, 2, 2, 2, 58, 52, 3, 2, 2, 2, 59, 7, 3, 2, 2,
	2, 60, 65, 7, 35, 2, 2, 61, 65, 7, 36, 2, 2, 62, 65, 7, 31, 2, 2, 63, 65,
	5, 4, 3, 2, 64, 60, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 63, 3, 2, 2, 2, 65, 9, 3, 2, 2, 2, 66, 67, 7, 34, 2, 2, 67, 11, 3,
	2, 2, 2, 68, 69, 8, 7, 1, 2, 69, 72, 5, 8, 5, 2, 70, 72, 5, 10, 6, 2, 71,
	68, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 78, 3, 2, 2, 2, 73, 74, 12, 5,
	2, 2, 74, 75, 9, 2, 2, 2, 75, 77, 5, 12, 7, 6, 76, 73, 3, 2, 2, 2, 77,
	80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 13, 3, 2, 2,
	2, 80, 78, 3, 2, 2, 2, 81, 82, 5, 6, 4, 2, 82, 15, 3, 2, 2, 2, 83, 84,
	5, 10, 6, 2, 84, 17, 3, 2, 2, 2, 85, 86, 7, 4, 2, 2, 86, 19, 3, 2, 2, 2,
	87, 88, 7, 21, 2, 2, 88, 94, 7, 6, 2, 2, 89, 90, 5, 14, 8, 2, 90, 91, 7,
	5, 2, 2, 91, 92, 5, 12, 7, 2, 92, 93, 7, 6, 2, 2, 93, 95, 3, 2, 2, 2, 94,
	89, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 98, 3, 2, 2, 2, 98, 99, 7, 22, 2, 2, 99, 21, 3, 2, 2, 2, 100, 104,
	5, 14, 8, 2, 101, 104, 5, 16, 9, 2, 102, 104, 5, 18, 10, 2, 103, 100, 3,
	2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 23, 3, 2, 2,
	2, 105, 106, 7, 21, 2, 2, 106, 112, 7, 6, 2, 2, 107, 108, 5, 22, 12, 2,
	108, 109, 7, 5, 2, 2, 109, 110, 5, 12, 7, 2, 110, 111, 7, 6, 2, 2, 111,
	113, 3, 2, 2, 2, 112, 107, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 112,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 7, 22,
	2, 2, 117, 25, 3, 2, 2, 2, 118, 122, 5, 20, 11, 2, 119, 122, 5, 24, 13,
	2, 120, 122, 5, 12, 7, 2, 121, 118, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121,
	120, 3, 2, 2, 2, 122, 27, 3, 2, 2, 2, 123, 124, 7, 21, 2, 2, 124, 128,
	7, 6, 2, 2, 125, 126, 5, 12, 7, 2, 126, 127, 7, 6, 2, 2, 127, 129, 3, 2,
	2, 2, 128, 125, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2,
	130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 7, 22, 2, 2, 133,
	29, 3, 2, 2, 2, 134, 137, 5, 28, 15, 2, 135, 137, 5, 12, 7, 2, 136, 134,
	3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 31, 3, 2, 2, 2, 138, 139, 7, 9,
	2, 2, 139, 167, 5, 12, 7, 2, 140, 167, 7, 10, 2, 2, 141, 142, 7, 11, 2,
	2, 142, 143, 5, 12, 7, 2, 143, 144, 7, 17, 2, 2, 144, 161, 7, 6, 2, 2,
	145, 146, 7, 12, 2, 2, 146, 147, 7, 24, 2, 2, 147, 148, 5, 30, 16, 2, 148,
	149, 7, 6, 2, 2, 149, 150, 7, 13, 2, 2, 150, 151, 7, 24, 2, 2, 151, 152,
	5, 26, 14, 2, 152, 162, 3, 2, 2, 2, 153, 154, 7, 13, 2, 2, 154, 155, 7,
	24, 2, 2, 155, 156, 5, 26, 14, 2, 156, 157, 7, 6, 2, 2, 157, 158, 7, 12,
	2, 2, 158, 159, 7, 24, 2, 2, 159, 160, 5, 30, 16, 2, 160, 162, 3, 2, 2,
	2, 161, 145, 3, 2, 2, 2, 161, 153, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163,
	164, 7, 6, 2, 2, 164, 165, 7, 18, 2, 2, 165, 167, 3, 2, 2, 2, 166, 138,
	3, 2, 2, 2, 166, 140, 3, 2, 2, 2, 166, 141, 3, 2, 2, 2, 167, 33, 3, 2,
	2, 2, 168, 169, 9, 3, 2, 2, 169, 35, 3, 2, 2, 2, 170, 171, 5, 34, 18, 2,
	171, 172, 5, 10, 6, 2, 172, 37, 3, 2, 2, 2, 173, 174, 7, 8, 2, 2, 174,
	176, 7, 21, 2, 2, 175, 177, 7, 6, 2, 2, 176, 175, 3, 2, 2, 2, 177, 178,
	3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 186, 3, 2,
	2, 2, 180, 182, 5, 36, 19, 2, 181, 183, 7, 6, 2, 2, 182, 181, 3, 2, 2,
	2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	187, 3, 2, 2, 2, 186, 180, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 186,
	3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 192, 7, 22,
	2, 2, 191, 193, 7, 6, 2, 2, 192, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2,
	194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 39, 3, 2, 2, 2, 196, 198,
	7, 6, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2,
	2, 2, 199, 200, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2,
	202, 204, 5, 38, 20, 2, 203, 202, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 214, 5, 32, 17, 2, 206, 208, 7, 6, 2, 2, 207, 206,
	3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2,
	2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 5, 32, 17, 2, 212, 207, 3, 2, 2,
	2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	220, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 219, 7, 6, 2, 2, 218, 217,
	3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2,
	2, 2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 224, 7, 2, 2, 3,
	224, 41, 3, 2, 2, 2, 25, 44, 55, 58, 64, 71, 78, 96, 103, 114, 121, 130,
	136, 161, 166, 178, 184, 188, 194, 199, 203, 209, 214, 220,
}
var literalNames = []string{
	"", "'.'", "'remaining'", "'to'", "", "", "'vars'", "'print'", "'fail'",
	"'send'", "'source'", "'destination'", "'allocate'", "'+'", "'-'", "'('",
	"')'", "'['", "']'", "'{'", "'}'", "'*'", "'='", "'account'", "'asset'",
	"'number'", "'monetary'", "'portion'", "", "", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND",
	"SOURCE", "DESTINATION", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN",
	"LBRACK", "RBRACK", "LBRACE", "RBRACE", "ALL", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "TY_PORTION", "RATIO", "NUMBER", "PERCENT",
	"IDENTIFIER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"amount", "monetary", "portion", "literal", "variable", "expression", "portionConst",
	"portionVar", "portionRemaining", "allocBlockConst", "allocPartDyn", "allocBlockDyn",
	"allocation", "sourceBlock", "source", "statement", "type_", "varDecl",
	"varListDecl", "script",
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
	NumScriptParserEOF           = antlr.TokenEOF
	NumScriptParserT__0          = 1
	NumScriptParserT__1          = 2
	NumScriptParserT__2          = 3
	NumScriptParserNEWLINE       = 4
	NumScriptParserWHITESPACE    = 5
	NumScriptParserVARS          = 6
	NumScriptParserPRINT         = 7
	NumScriptParserFAIL          = 8
	NumScriptParserSEND          = 9
	NumScriptParserSOURCE        = 10
	NumScriptParserDESTINATION   = 11
	NumScriptParserALLOCATE      = 12
	NumScriptParserOP_ADD        = 13
	NumScriptParserOP_SUB        = 14
	NumScriptParserLPAREN        = 15
	NumScriptParserRPAREN        = 16
	NumScriptParserLBRACK        = 17
	NumScriptParserRBRACK        = 18
	NumScriptParserLBRACE        = 19
	NumScriptParserRBRACE        = 20
	NumScriptParserALL           = 21
	NumScriptParserEQ            = 22
	NumScriptParserTY_ACCOUNT    = 23
	NumScriptParserTY_ASSET      = 24
	NumScriptParserTY_NUMBER     = 25
	NumScriptParserTY_MONETARY   = 26
	NumScriptParserTY_PORTION    = 27
	NumScriptParserRATIO         = 28
	NumScriptParserNUMBER        = 29
	NumScriptParserPERCENT       = 30
	NumScriptParserIDENTIFIER    = 31
	NumScriptParserVARIABLE_NAME = 32
	NumScriptParserACCOUNT       = 33
	NumScriptParserASSET         = 34
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_amount           = 0
	NumScriptParserRULE_monetary         = 1
	NumScriptParserRULE_portion          = 2
	NumScriptParserRULE_literal          = 3
	NumScriptParserRULE_variable         = 4
	NumScriptParserRULE_expression       = 5
	NumScriptParserRULE_portionConst     = 6
	NumScriptParserRULE_portionVar       = 7
	NumScriptParserRULE_portionRemaining = 8
	NumScriptParserRULE_allocBlockConst  = 9
	NumScriptParserRULE_allocPartDyn     = 10
	NumScriptParserRULE_allocBlockDyn    = 11
	NumScriptParserRULE_allocation       = 12
	NumScriptParserRULE_sourceBlock      = 13
	NumScriptParserRULE_source           = 14
	NumScriptParserRULE_statement        = 15
	NumScriptParserRULE_type_            = 16
	NumScriptParserRULE_varDecl          = 17
	NumScriptParserRULE_varListDecl      = 18
	NumScriptParserRULE_script           = 19
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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserNUMBER:
		localctx = NewAmountSpecificContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*AmountSpecificContext).num = _m
		}

	case NumScriptParserALL:
		localctx = NewAmountAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
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

		var _x = p.Amount()

		localctx.(*MonetaryContext).amt = _x
	}
	{
		p.SetState(47)
		p.Match(NumScriptParserRBRACK)
	}

	return localctx
}

// IPortionContext is an interface to support dynamic dispatch.
type IPortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortionContext differentiates from other interfaces.
	IsPortionContext()
}

type PortionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortionContext() *PortionContext {
	var p = new(PortionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_portion
	return p
}

func (*PortionContext) IsPortionContext() {}

func NewPortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionContext {
	var p = new(PortionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_portion

	return p
}

func (s *PortionContext) GetParser() antlr.Parser { return s.parser }

func (s *PortionContext) CopyFrom(ctx *PortionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PercentageContext struct {
	*PortionContext
	pint  antlr.Token
	pfrac antlr.Token
}

func NewPercentageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentageContext {
	var p = new(PercentageContext)

	p.PortionContext = NewEmptyPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PortionContext))

	return p
}

func (s *PercentageContext) GetPint() antlr.Token { return s.pint }

func (s *PercentageContext) GetPfrac() antlr.Token { return s.pfrac }

func (s *PercentageContext) SetPint(v antlr.Token) { s.pint = v }

func (s *PercentageContext) SetPfrac(v antlr.Token) { s.pfrac = v }

func (s *PercentageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentageContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPERCENT, 0)
}

func (s *PercentageContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNUMBER)
}

func (s *PercentageContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNUMBER, i)
}

func (s *PercentageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterPercentage(s)
	}
}

func (s *PercentageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitPercentage(s)
	}
}

type RatioContext struct {
	*PortionContext
	r antlr.Token
}

func NewRatioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RatioContext {
	var p = new(RatioContext)

	p.PortionContext = NewEmptyPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PortionContext))

	return p
}

func (s *RatioContext) GetR() antlr.Token { return s.r }

func (s *RatioContext) SetR(v antlr.Token) { s.r = v }

func (s *RatioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RatioContext) RATIO() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRATIO, 0)
}

func (s *RatioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterRatio(s)
	}
}

func (s *RatioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitRatio(s)
	}
}

func (p *NumScriptParser) Portion() (localctx IPortionContext) {
	localctx = NewPortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NumScriptParserRULE_portion)
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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserRATIO:
		localctx = NewRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _m = p.Match(NumScriptParserRATIO)

			localctx.(*RatioContext).r = _m
		}

	case NumScriptParserNUMBER:
		localctx = NewPercentageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*PercentageContext).pint = _m
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NumScriptParserT__0 {
			{
				p.SetState(51)
				p.Match(NumScriptParserT__0)
			}
			{
				p.SetState(52)

				var _m = p.Match(NumScriptParserNUMBER)

				localctx.(*PercentageContext).pfrac = _m
			}

		}
		{
			p.SetState(55)
			p.Match(NumScriptParserPERCENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 6, NumScriptParserRULE_literal)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
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
	p.EnterRule(localctx, 8, NumScriptParserRULE_variable)

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
		p.SetState(64)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, NumScriptParserRULE_expression, _p)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(67)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)

			var _x = p.Variable()

			localctx.(*ExprVariableContext).var_ = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, NumScriptParserRULE_expression)
			p.SetState(71)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(72)

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
				p.SetState(73)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IPortionConstContext is an interface to support dynamic dispatch.
type IPortionConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPor returns the por rule contexts.
	GetPor() IPortionContext

	// SetPor sets the por rule contexts.
	SetPor(IPortionContext)

	// IsPortionConstContext differentiates from other interfaces.
	IsPortionConstContext()
}

type PortionConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	por    IPortionContext
}

func NewEmptyPortionConstContext() *PortionConstContext {
	var p = new(PortionConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_portionConst
	return p
}

func (*PortionConstContext) IsPortionConstContext() {}

func NewPortionConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionConstContext {
	var p = new(PortionConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_portionConst

	return p
}

func (s *PortionConstContext) GetParser() antlr.Parser { return s.parser }

func (s *PortionConstContext) GetPor() IPortionContext { return s.por }

func (s *PortionConstContext) SetPor(v IPortionContext) { s.por = v }

func (s *PortionConstContext) Portion() IPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionContext)
}

func (s *PortionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortionConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterPortionConst(s)
	}
}

func (s *PortionConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitPortionConst(s)
	}
}

func (p *NumScriptParser) PortionConst() (localctx IPortionConstContext) {
	localctx = NewPortionConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumScriptParserRULE_portionConst)

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
		p.SetState(79)

		var _x = p.Portion()

		localctx.(*PortionConstContext).por = _x
	}

	return localctx
}

// IPortionVarContext is an interface to support dynamic dispatch.
type IPortionVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPor returns the por rule contexts.
	GetPor() IVariableContext

	// SetPor sets the por rule contexts.
	SetPor(IVariableContext)

	// IsPortionVarContext differentiates from other interfaces.
	IsPortionVarContext()
}

type PortionVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	por    IVariableContext
}

func NewEmptyPortionVarContext() *PortionVarContext {
	var p = new(PortionVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_portionVar
	return p
}

func (*PortionVarContext) IsPortionVarContext() {}

func NewPortionVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionVarContext {
	var p = new(PortionVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_portionVar

	return p
}

func (s *PortionVarContext) GetParser() antlr.Parser { return s.parser }

func (s *PortionVarContext) GetPor() IVariableContext { return s.por }

func (s *PortionVarContext) SetPor(v IVariableContext) { s.por = v }

func (s *PortionVarContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PortionVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortionVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterPortionVar(s)
	}
}

func (s *PortionVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitPortionVar(s)
	}
}

func (p *NumScriptParser) PortionVar() (localctx IPortionVarContext) {
	localctx = NewPortionVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumScriptParserRULE_portionVar)

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

		var _x = p.Variable()

		localctx.(*PortionVarContext).por = _x
	}

	return localctx
}

// IPortionRemainingContext is an interface to support dynamic dispatch.
type IPortionRemainingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortionRemainingContext differentiates from other interfaces.
	IsPortionRemainingContext()
}

type PortionRemainingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortionRemainingContext() *PortionRemainingContext {
	var p = new(PortionRemainingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_portionRemaining
	return p
}

func (*PortionRemainingContext) IsPortionRemainingContext() {}

func NewPortionRemainingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionRemainingContext {
	var p = new(PortionRemainingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_portionRemaining

	return p
}

func (s *PortionRemainingContext) GetParser() antlr.Parser { return s.parser }
func (s *PortionRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionRemainingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortionRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterPortionRemaining(s)
	}
}

func (s *PortionRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitPortionRemaining(s)
	}
}

func (p *NumScriptParser) PortionRemaining() (localctx IPortionRemainingContext) {
	localctx = NewPortionRemainingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumScriptParserRULE_portionRemaining)

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
		p.SetState(83)
		p.Match(NumScriptParserT__1)
	}

	return localctx
}

// IAllocBlockConstContext is an interface to support dynamic dispatch.
type IAllocBlockConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_portionConst returns the _portionConst rule contexts.
	Get_portionConst() IPortionConstContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_portionConst sets the _portionConst rule contexts.
	Set_portionConst(IPortionConstContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IPortionConstContext

	// GetDests returns the dests rule context list.
	GetDests() []IExpressionContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IPortionConstContext)

	// SetDests sets the dests rule context list.
	SetDests([]IExpressionContext)

	// IsAllocBlockConstContext differentiates from other interfaces.
	IsAllocBlockConstContext()
}

type AllocBlockConstContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	_portionConst IPortionConstContext
	portions      []IPortionConstContext
	_expression   IExpressionContext
	dests         []IExpressionContext
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

func (s *AllocBlockConstContext) Get_portionConst() IPortionConstContext { return s._portionConst }

func (s *AllocBlockConstContext) Get_expression() IExpressionContext { return s._expression }

func (s *AllocBlockConstContext) Set_portionConst(v IPortionConstContext) { s._portionConst = v }

func (s *AllocBlockConstContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AllocBlockConstContext) GetPortions() []IPortionConstContext { return s.portions }

func (s *AllocBlockConstContext) GetDests() []IExpressionContext { return s.dests }

func (s *AllocBlockConstContext) SetPortions(v []IPortionConstContext) { s.portions = v }

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

func (s *AllocBlockConstContext) AllPortionConst() []IPortionConstContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPortionConstContext)(nil)).Elem())
	var tst = make([]IPortionConstContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPortionConstContext)
		}
	}

	return tst
}

func (s *AllocBlockConstContext) PortionConst(i int) IPortionConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionConstContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPortionConstContext)
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
	p.EnterRule(localctx, 18, NumScriptParserRULE_allocBlockConst)
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
		p.SetState(85)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(86)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserRATIO || _la == NumScriptParserNUMBER {
		{
			p.SetState(87)

			var _x = p.PortionConst()

			localctx.(*AllocBlockConstContext)._portionConst = _x
		}
		localctx.(*AllocBlockConstContext).portions = append(localctx.(*AllocBlockConstContext).portions, localctx.(*AllocBlockConstContext)._portionConst)
		{
			p.SetState(88)
			p.Match(NumScriptParserT__2)
		}
		{
			p.SetState(89)

			var _x = p.expression(0)

			localctx.(*AllocBlockConstContext)._expression = _x
		}
		localctx.(*AllocBlockConstContext).dests = append(localctx.(*AllocBlockConstContext).dests, localctx.(*AllocBlockConstContext)._expression)
		{
			p.SetState(90)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
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
}

func NewAllocPartDynVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocPartDynVarContext {
	var p = new(AllocPartDynVarContext)

	p.AllocPartDynContext = NewEmptyAllocPartDynContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocPartDynContext))

	return p
}

func (s *AllocPartDynVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocPartDynVarContext) PortionVar() IPortionVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionVarContext)
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

func (s *AllocPartDynConstContext) PortionConst() IPortionConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionConstContext)
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

func (s *AllocPartDynRemainingContext) PortionRemaining() IPortionRemainingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionRemainingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionRemainingContext)
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
	p.EnterRule(localctx, 20, NumScriptParserRULE_allocPartDyn)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserRATIO, NumScriptParserNUMBER:
		localctx = NewAllocPartDynConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.PortionConst()
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewAllocPartDynVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.PortionVar()
		}

	case NumScriptParserT__1:
		localctx = NewAllocPartDynRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.PortionRemaining()
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
	p.EnterRule(localctx, 22, NumScriptParserRULE_allocBlockDyn)
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
		p.SetState(103)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(104)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(NumScriptParserT__1-2))|(1<<(NumScriptParserRATIO-2))|(1<<(NumScriptParserNUMBER-2))|(1<<(NumScriptParserVARIABLE_NAME-2)))) != 0) {
		{
			p.SetState(105)

			var _x = p.AllocPartDyn()

			localctx.(*AllocBlockDynContext)._allocPartDyn = _x
		}
		localctx.(*AllocBlockDynContext).portions = append(localctx.(*AllocBlockDynContext).portions, localctx.(*AllocBlockDynContext)._allocPartDyn)
		{
			p.SetState(106)
			p.Match(NumScriptParserT__2)
		}
		{
			p.SetState(107)

			var _x = p.expression(0)

			localctx.(*AllocBlockDynContext)._expression = _x
		}
		localctx.(*AllocBlockDynContext).dests = append(localctx.(*AllocBlockDynContext).dests, localctx.(*AllocBlockDynContext)._expression)
		{
			p.SetState(108)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
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
	p.EnterRule(localctx, 24, NumScriptParserRULE_allocation)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAllocConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.AllocBlockConst()
		}

	case 2:
		localctx = NewAllocDynContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.AllocBlockDyn()
		}

	case 3:
		localctx = NewAllocAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
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
	p.EnterRule(localctx, 26, NumScriptParserRULE_sourceBlock)
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
		p.SetState(121)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(122)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(NumScriptParserLBRACK-17))|(1<<(NumScriptParserNUMBER-17))|(1<<(NumScriptParserVARIABLE_NAME-17))|(1<<(NumScriptParserACCOUNT-17))|(1<<(NumScriptParserASSET-17)))) != 0) {
		{
			p.SetState(123)

			var _x = p.expression(0)

			localctx.(*SourceBlockContext)._expression = _x
		}
		localctx.(*SourceBlockContext).sources = append(localctx.(*SourceBlockContext).sources, localctx.(*SourceBlockContext)._expression)
		{
			p.SetState(124)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
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
	p.EnterRule(localctx, 28, NumScriptParserRULE_source)

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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACE:
		localctx = NewSrcBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.SourceBlock()
		}

	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
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
	p.EnterRule(localctx, 30, NumScriptParserRULE_statement)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(137)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(140)

			var _x = p.expression(0)

			localctx.(*SendContext).mon = _x
		}
		{
			p.SetState(141)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(142)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(143)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(144)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(145)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(146)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(147)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(148)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(149)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(151)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(152)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(153)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(154)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(155)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(156)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(157)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(161)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(162)
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
	p.EnterRule(localctx, 32, NumScriptParserRULE_type_)
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
		p.SetState(166)
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

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() IType_Context

	// GetName returns the name rule contexts.
	GetName() IVariableContext

	// SetTy sets the ty rule contexts.
	SetTy(IType_Context)

	// SetName sets the name rule contexts.
	SetName(IVariableContext)

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     IType_Context
	name   IVariableContext
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

func (s *VarDeclContext) SetTy(v IType_Context) { s.ty = v }

func (s *VarDeclContext) SetName(v IVariableContext) { s.name = v }

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
	p.EnterRule(localctx, 34, NumScriptParserRULE_varDecl)

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
		p.SetState(168)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(169)

		var _x = p.Variable()

		localctx.(*VarDeclContext).name = _x
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
	p.EnterRule(localctx, 36, NumScriptParserRULE_varListDecl)
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
		p.SetState(171)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(172)
		p.Match(NumScriptParserLBRACE)
	}
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
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY)|(1<<NumScriptParserTY_PORTION))) != 0) {
		{
			p.SetState(178)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(179)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(NumScriptParserRBRACE)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(189)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(192)
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
	p.EnterRule(localctx, 38, NumScriptParserRULE_script)
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
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(194)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(200)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(203)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
				{
					p.SetState(204)
					p.Match(NumScriptParserNEWLINE)
				}

				p.SetState(207)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(209)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(215)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.Match(NumScriptParserEOF)
	}

	return localctx
}

func (p *NumScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
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

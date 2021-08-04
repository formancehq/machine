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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 222,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 5, 2, 41, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 59, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12, 6, 14, 6,
	67, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 72, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 6, 8, 81, 10, 8, 13, 8, 14, 8, 82, 3, 8, 3, 8, 3, 9, 3,
	9, 5, 9, 89, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 96, 10, 10,
	13, 10, 14, 10, 97, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 6, 11, 109, 10, 11, 13, 11, 14, 11, 110, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 124, 10,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 149, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 154,
	10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 168, 10, 17, 3, 18, 3, 18, 3, 18, 6, 18, 173,
	10, 18, 13, 18, 14, 18, 174, 3, 18, 3, 18, 6, 18, 179, 10, 18, 13, 18,
	14, 18, 180, 6, 18, 183, 10, 18, 13, 18, 14, 18, 184, 3, 18, 3, 18, 6,
	18, 189, 10, 18, 13, 18, 14, 18, 190, 3, 19, 7, 19, 194, 10, 19, 12, 19,
	14, 19, 197, 11, 19, 3, 19, 5, 19, 200, 10, 19, 3, 19, 3, 19, 6, 19, 204,
	10, 19, 13, 19, 14, 19, 205, 3, 19, 7, 19, 209, 10, 19, 12, 19, 14, 19,
	212, 11, 19, 3, 19, 7, 19, 215, 10, 19, 12, 19, 14, 19, 218, 11, 19, 3,
	19, 3, 19, 3, 19, 2, 3, 10, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 2, 4, 3, 2, 19, 20, 3, 2, 29, 33, 2, 231, 2,
	40, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2,
	10, 58, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2, 14, 73, 3, 2, 2, 2, 16, 88, 3,
	2, 2, 2, 18, 90, 3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 114, 3, 2, 2, 2,
	24, 123, 3, 2, 2, 2, 26, 153, 3, 2, 2, 2, 28, 155, 3, 2, 2, 2, 30, 157,
	3, 2, 2, 2, 32, 163, 3, 2, 2, 2, 34, 169, 3, 2, 2, 2, 36, 195, 3, 2, 2,
	2, 38, 41, 7, 37, 2, 2, 39, 41, 7, 27, 2, 2, 40, 38, 3, 2, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 3, 3, 2, 2, 2, 42, 43, 7, 23, 2, 2, 43, 44, 7, 41, 2, 2,
	44, 45, 5, 2, 2, 2, 45, 46, 7, 24, 2, 2, 46, 5, 3, 2, 2, 2, 47, 52, 7,
	40, 2, 2, 48, 52, 7, 41, 2, 2, 49, 52, 7, 37, 2, 2, 50, 52, 5, 4, 3, 2,
	51, 47, 3, 2, 2, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3,
	2, 2, 2, 52, 7, 3, 2, 2, 2, 53, 54, 7, 39, 2, 2, 54, 9, 3, 2, 2, 2, 55,
	56, 8, 6, 1, 2, 56, 59, 5, 6, 4, 2, 57, 59, 5, 8, 5, 2, 58, 55, 3, 2, 2,
	2, 58, 57, 3, 2, 2, 2, 59, 65, 3, 2, 2, 2, 60, 61, 12, 5, 2, 2, 61, 62,
	9, 2, 2, 2, 62, 64, 5, 10, 6, 6, 63, 60, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 68, 72, 7, 35, 2, 2, 69, 72, 5, 8, 5, 2, 70, 72, 7, 36, 2, 2,
	71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 13, 3,
	2, 2, 2, 73, 74, 7, 25, 2, 2, 74, 80, 7, 8, 2, 2, 75, 76, 5, 12, 7, 2,
	76, 77, 7, 3, 2, 2, 77, 78, 5, 10, 6, 2, 78, 79, 7, 8, 2, 2, 79, 81, 3,
	2, 2, 2, 80, 75, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82,
	83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 26, 2, 2, 85, 15, 3, 2,
	2, 2, 86, 89, 5, 10, 6, 2, 87, 89, 5, 14, 8, 2, 88, 86, 3, 2, 2, 2, 88,
	87, 3, 2, 2, 2, 89, 17, 3, 2, 2, 2, 90, 91, 7, 25, 2, 2, 91, 95, 7, 8,
	2, 2, 92, 93, 5, 24, 13, 2, 93, 94, 7, 8, 2, 2, 94, 96, 3, 2, 2, 2, 95,
	92, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2,
	2, 98, 99, 3, 2, 2, 2, 99, 100, 7, 26, 2, 2, 100, 19, 3, 2, 2, 2, 101,
	102, 7, 25, 2, 2, 102, 108, 7, 8, 2, 2, 103, 104, 5, 12, 7, 2, 104, 105,
	7, 4, 2, 2, 105, 106, 5, 24, 13, 2, 106, 107, 7, 8, 2, 2, 107, 109, 3,
	2, 2, 2, 108, 103, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2,
	2, 110, 111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 7, 26, 2, 2, 113,
	21, 3, 2, 2, 2, 114, 115, 7, 5, 2, 2, 115, 116, 5, 10, 6, 2, 116, 117,
	7, 4, 2, 2, 117, 118, 5, 24, 13, 2, 118, 23, 3, 2, 2, 2, 119, 124, 5, 10,
	6, 2, 120, 124, 5, 18, 10, 2, 121, 124, 5, 20, 11, 2, 122, 124, 5, 22,
	12, 2, 123, 119, 3, 2, 2, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2,
	123, 122, 3, 2, 2, 2, 124, 25, 3, 2, 2, 2, 125, 126, 7, 13, 2, 2, 126,
	154, 5, 10, 6, 2, 127, 154, 7, 14, 2, 2, 128, 129, 7, 15, 2, 2, 129, 130,
	5, 10, 6, 2, 130, 131, 7, 21, 2, 2, 131, 148, 7, 8, 2, 2, 132, 133, 7,
	16, 2, 2, 133, 134, 7, 28, 2, 2, 134, 135, 5, 24, 13, 2, 135, 136, 7, 8,
	2, 2, 136, 137, 7, 17, 2, 2, 137, 138, 7, 28, 2, 2, 138, 139, 5, 16, 9,
	2, 139, 149, 3, 2, 2, 2, 140, 141, 7, 17, 2, 2, 141, 142, 7, 28, 2, 2,
	142, 143, 5, 16, 9, 2, 143, 144, 7, 8, 2, 2, 144, 145, 7, 16, 2, 2, 145,
	146, 7, 28, 2, 2, 146, 147, 5, 24, 13, 2, 147, 149, 3, 2, 2, 2, 148, 132,
	3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 7, 8,
	2, 2, 151, 152, 7, 22, 2, 2, 152, 154, 3, 2, 2, 2, 153, 125, 3, 2, 2, 2,
	153, 127, 3, 2, 2, 2, 153, 128, 3, 2, 2, 2, 154, 27, 3, 2, 2, 2, 155, 156,
	9, 3, 2, 2, 156, 29, 3, 2, 2, 2, 157, 158, 7, 6, 2, 2, 158, 159, 5, 10,
	6, 2, 159, 160, 7, 7, 2, 2, 160, 161, 7, 34, 2, 2, 161, 162, 7, 22, 2,
	2, 162, 31, 3, 2, 2, 2, 163, 164, 5, 28, 15, 2, 164, 167, 5, 8, 5, 2, 165,
	166, 7, 28, 2, 2, 166, 168, 5, 30, 16, 2, 167, 165, 3, 2, 2, 2, 167, 168,
	3, 2, 2, 2, 168, 33, 3, 2, 2, 2, 169, 170, 7, 12, 2, 2, 170, 172, 7, 25,
	2, 2, 171, 173, 7, 8, 2, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 182, 3, 2, 2, 2, 176,
	178, 5, 32, 17, 2, 177, 179, 7, 8, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2,
	2, 2, 182, 176, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2,
	184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 7, 26, 2, 2, 187,
	189, 7, 8, 2, 2, 188, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 188,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 35, 3, 2, 2, 2, 192, 194, 7, 8,
	2, 2, 193, 192, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2,
	195, 196, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198,
	200, 5, 34, 18, 2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201,
	3, 2, 2, 2, 201, 210, 5, 26, 14, 2, 202, 204, 7, 8, 2, 2, 203, 202, 3,
	2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2,
	2, 206, 207, 3, 2, 2, 2, 207, 209, 5, 26, 14, 2, 208, 203, 3, 2, 2, 2,
	209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211,
	216, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 215, 7, 8, 2, 2, 214, 213,
	3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2,
	2, 2, 217, 219, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 220, 7, 2, 2, 3,
	220, 37, 3, 2, 2, 2, 24, 40, 51, 58, 65, 71, 82, 88, 97, 110, 123, 148,
	153, 167, 174, 180, 184, 190, 195, 199, 205, 210, 216,
}
var literalNames = []string{
	"", "'to'", "'from'", "'max'", "'meta('", "','", "", "", "", "", "'vars'",
	"'print'", "'fail'", "'send'", "'source'", "'destination'", "'allocate'",
	"'+'", "'-'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'*'", "'='", "'account'",
	"'asset'", "'number'", "'monetary'", "'portion'", "", "", "'remaining'",
	"", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "PRINT", "FAIL", "SEND", "SOURCE", "DESTINATION", "ALLOCATE", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "ALL",
	"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
	"STRING", "PORTION", "PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"amount", "monetary", "literal", "variable", "expression", "allotmentPortion",
	"destinationAllotment", "destination", "sourceInOrder", "sourceAllotment",
	"sourceLimited", "source", "statement", "type_", "origin", "varDecl", "varListDecl",
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
	NumScriptParserT__2              = 3
	NumScriptParserT__3              = 4
	NumScriptParserT__4              = 5
	NumScriptParserNEWLINE           = 6
	NumScriptParserWHITESPACE        = 7
	NumScriptParserMULTILINE_COMMENT = 8
	NumScriptParserLINE_COMMENT      = 9
	NumScriptParserVARS              = 10
	NumScriptParserPRINT             = 11
	NumScriptParserFAIL              = 12
	NumScriptParserSEND              = 13
	NumScriptParserSOURCE            = 14
	NumScriptParserDESTINATION       = 15
	NumScriptParserALLOCATE          = 16
	NumScriptParserOP_ADD            = 17
	NumScriptParserOP_SUB            = 18
	NumScriptParserLPAREN            = 19
	NumScriptParserRPAREN            = 20
	NumScriptParserLBRACK            = 21
	NumScriptParserRBRACK            = 22
	NumScriptParserLBRACE            = 23
	NumScriptParserRBRACE            = 24
	NumScriptParserALL               = 25
	NumScriptParserEQ                = 26
	NumScriptParserTY_ACCOUNT        = 27
	NumScriptParserTY_ASSET          = 28
	NumScriptParserTY_NUMBER         = 29
	NumScriptParserTY_MONETARY       = 30
	NumScriptParserTY_PORTION        = 31
	NumScriptParserSTRING            = 32
	NumScriptParserPORTION           = 33
	NumScriptParserPORTION_REMAINING = 34
	NumScriptParserNUMBER            = 35
	NumScriptParserPERCENT           = 36
	NumScriptParserVARIABLE_NAME     = 37
	NumScriptParserACCOUNT           = 38
	NumScriptParserASSET             = 39
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_amount               = 0
	NumScriptParserRULE_monetary             = 1
	NumScriptParserRULE_literal              = 2
	NumScriptParserRULE_variable             = 3
	NumScriptParserRULE_expression           = 4
	NumScriptParserRULE_allotmentPortion     = 5
	NumScriptParserRULE_destinationAllotment = 6
	NumScriptParserRULE_destination          = 7
	NumScriptParserRULE_sourceInOrder        = 8
	NumScriptParserRULE_sourceAllotment      = 9
	NumScriptParserRULE_sourceLimited        = 10
	NumScriptParserRULE_source               = 11
	NumScriptParserRULE_statement            = 12
	NumScriptParserRULE_type_                = 13
	NumScriptParserRULE_origin               = 14
	NumScriptParserRULE_varDecl              = 15
	NumScriptParserRULE_varListDecl          = 16
	NumScriptParserRULE_script               = 17
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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPORTION:
		localctx = NewAllotmentPortionConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(NumScriptParserPORTION)
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewAllotmentPortionVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)

			var _x = p.Variable()

			localctx.(*AllotmentPortionVarContext).por = _x
		}

	case NumScriptParserPORTION_REMAINING:
		localctx = NewAllotmentPortionRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(NumScriptParserPORTION_REMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_allotmentPortion sets the _allotmentPortion rule contexts.
	Set_allotmentPortion(IAllotmentPortionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllotmentPortionContext

	// GetDests returns the dests rule context list.
	GetDests() []IExpressionContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetDests sets the dests rule context list.
	SetDests([]IExpressionContext)

	// IsDestinationAllotmentContext differentiates from other interfaces.
	IsDestinationAllotmentContext()
}

type DestinationAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_expression       IExpressionContext
	dests             []IExpressionContext
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

func (s *DestinationAllotmentContext) Get_expression() IExpressionContext { return s._expression }

func (s *DestinationAllotmentContext) Set_allotmentPortion(v IAllotmentPortionContext) {
	s._allotmentPortion = v
}

func (s *DestinationAllotmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DestinationAllotmentContext) GetPortions() []IAllotmentPortionContext { return s.portions }

func (s *DestinationAllotmentContext) GetDests() []IExpressionContext { return s.dests }

func (s *DestinationAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *DestinationAllotmentContext) SetDests(v []IExpressionContext) { s.dests = v }

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

func (s *DestinationAllotmentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DestinationAllotmentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	localctx = NewDestinationAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumScriptParserRULE_destinationAllotment)
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
		p.SetState(71)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(72)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(NumScriptParserPORTION-33))|(1<<(NumScriptParserPORTION_REMAINING-33))|(1<<(NumScriptParserVARIABLE_NAME-33)))) != 0) {
		{
			p.SetState(73)

			var _x = p.AllotmentPortion()

			localctx.(*DestinationAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*DestinationAllotmentContext).portions = append(localctx.(*DestinationAllotmentContext).portions, localctx.(*DestinationAllotmentContext)._allotmentPortion)
		{
			p.SetState(74)
			p.Match(NumScriptParserT__0)
		}
		{
			p.SetState(75)

			var _x = p.expression(0)

			localctx.(*DestinationAllotmentContext)._expression = _x
		}
		localctx.(*DestinationAllotmentContext).dests = append(localctx.(*DestinationAllotmentContext).dests, localctx.(*DestinationAllotmentContext)._expression)
		{
			p.SetState(76)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
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

func (p *NumScriptParser) Destination() (localctx IDestinationContext) {
	localctx = NewDestinationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumScriptParserRULE_destination)

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
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewDestAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.expression(0)
		}

	case NumScriptParserLBRACE:
		localctx = NewDestAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.DestinationAllotment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// GetSrcs returns the srcs rule context list.
	GetSrcs() []ISourceContext

	// SetSrcs sets the srcs rule context list.
	SetSrcs([]ISourceContext)

	// IsSourceInOrderContext differentiates from other interfaces.
	IsSourceInOrderContext()
}

type SourceInOrderContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_source ISourceContext
	srcs    []ISourceContext
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

func (s *SourceInOrderContext) GetSrcs() []ISourceContext { return s.srcs }

func (s *SourceInOrderContext) SetSrcs(v []ISourceContext) { s.srcs = v }

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
	localctx = NewSourceInOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumScriptParserRULE_sourceInOrder)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserT__2)|(1<<NumScriptParserLBRACK)|(1<<NumScriptParserLBRACE))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NumScriptParserNUMBER-35))|(1<<(NumScriptParserVARIABLE_NAME-35))|(1<<(NumScriptParserACCOUNT-35))|(1<<(NumScriptParserASSET-35)))) != 0) {
		{
			p.SetState(90)

			var _x = p.Source()

			localctx.(*SourceInOrderContext)._source = _x
		}
		localctx.(*SourceInOrderContext).srcs = append(localctx.(*SourceInOrderContext).srcs, localctx.(*SourceInOrderContext)._source)
		{
			p.SetState(91)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(NumScriptParserRBRACE)
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

	// GetSrcs returns the srcs rule context list.
	GetSrcs() []ISourceContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetSrcs sets the srcs rule context list.
	SetSrcs([]ISourceContext)

	// IsSourceAllotmentContext differentiates from other interfaces.
	IsSourceAllotmentContext()
}

type SourceAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_source           ISourceContext
	srcs              []ISourceContext
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

func (s *SourceAllotmentContext) GetSrcs() []ISourceContext { return s.srcs }

func (s *SourceAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *SourceAllotmentContext) SetSrcs(v []ISourceContext) { s.srcs = v }

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
	localctx = NewSourceAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NumScriptParserRULE_sourceAllotment)
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
		p.SetState(99)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(100)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(NumScriptParserPORTION-33))|(1<<(NumScriptParserPORTION_REMAINING-33))|(1<<(NumScriptParserVARIABLE_NAME-33)))) != 0) {
		{
			p.SetState(101)

			var _x = p.AllotmentPortion()

			localctx.(*SourceAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*SourceAllotmentContext).portions = append(localctx.(*SourceAllotmentContext).portions, localctx.(*SourceAllotmentContext)._allotmentPortion)
		{
			p.SetState(102)
			p.Match(NumScriptParserT__1)
		}
		{
			p.SetState(103)

			var _x = p.Source()

			localctx.(*SourceAllotmentContext)._source = _x
		}
		localctx.(*SourceAllotmentContext).srcs = append(localctx.(*SourceAllotmentContext).srcs, localctx.(*SourceAllotmentContext)._source)
		{
			p.SetState(104)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(NumScriptParserRBRACE)
	}

	return localctx
}

// ISourceLimitedContext is an interface to support dynamic dispatch.
type ISourceLimitedContext interface {
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

	// IsSourceLimitedContext differentiates from other interfaces.
	IsSourceLimitedContext()
}

type SourceLimitedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	max    IExpressionContext
	src    ISourceContext
}

func NewEmptySourceLimitedContext() *SourceLimitedContext {
	var p = new(SourceLimitedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_sourceLimited
	return p
}

func (*SourceLimitedContext) IsSourceLimitedContext() {}

func NewSourceLimitedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceLimitedContext {
	var p = new(SourceLimitedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_sourceLimited

	return p
}

func (s *SourceLimitedContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceLimitedContext) GetMax() IExpressionContext { return s.max }

func (s *SourceLimitedContext) GetSrc() ISourceContext { return s.src }

func (s *SourceLimitedContext) SetMax(v IExpressionContext) { s.max = v }

func (s *SourceLimitedContext) SetSrc(v ISourceContext) { s.src = v }

func (s *SourceLimitedContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SourceLimitedContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceLimitedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceLimitedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceLimitedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSourceLimited(s)
	}
}

func (s *SourceLimitedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSourceLimited(s)
	}
}

func (p *NumScriptParser) SourceLimited() (localctx ISourceLimitedContext) {
	localctx = NewSourceLimitedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NumScriptParserRULE_sourceLimited)

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
		p.SetState(112)
		p.Match(NumScriptParserT__2)
	}
	{
		p.SetState(113)

		var _x = p.expression(0)

		localctx.(*SourceLimitedContext).max = _x
	}
	{
		p.SetState(114)
		p.Match(NumScriptParserT__1)
	}
	{
		p.SetState(115)

		var _x = p.Source()

		localctx.(*SourceLimitedContext).src = _x
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

type SrcLimitedContext struct {
	*SourceContext
}

func NewSrcLimitedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcLimitedContext {
	var p = new(SrcLimitedContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcLimitedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcLimitedContext) SourceLimited() ISourceLimitedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceLimitedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceLimitedContext)
}

func (s *SrcLimitedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterSrcLimited(s)
	}
}

func (s *SrcLimitedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitSrcLimited(s)
	}
}

type SrcAllotmentContext struct {
	*SourceContext
}

func NewSrcAllotmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcAllotmentContext {
	var p = new(SrcAllotmentContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.expression(0)
		}

	case 2:
		localctx = NewSrcInOrderContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.SourceInOrder()
		}

	case 3:
		localctx = NewSrcAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.SourceAllotment()
		}

	case 4:
		localctx = NewSrcLimitedContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.SourceLimited()
		}

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
	dest IDestinationContext
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

func (s *SendContext) GetDest() IDestinationContext { return s.dest }

func (s *SendContext) SetMon(v IExpressionContext) { s.mon = v }

func (s *SendContext) SetSrc(v ISourceContext) { s.src = v }

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(124)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(127)

			var _x = p.expression(0)

			localctx.(*SendContext).mon = _x
		}
		{
			p.SetState(128)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(129)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(130)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(131)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(132)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(133)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(134)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(135)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(136)

				var _x = p.Destination()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(138)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(139)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(140)

				var _x = p.Destination()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(141)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(142)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(143)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(144)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(148)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(149)
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
		p.SetState(153)
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
		p.SetState(155)
		p.Match(NumScriptParserT__3)
	}
	{
		p.SetState(156)

		var _x = p.expression(0)

		localctx.(*OriginContext).acc = _x
	}
	{
		p.SetState(157)
		p.Match(NumScriptParserT__4)
	}
	{
		p.SetState(158)

		var _m = p.Match(NumScriptParserSTRING)

		localctx.(*OriginContext).key = _m
	}
	{
		p.SetState(159)
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
		p.SetState(161)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(162)

		var _x = p.Variable()

		localctx.(*VarDeclContext).name = _x
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserEQ {
		{
			p.SetState(163)
			p.Match(NumScriptParserEQ)
		}
		{
			p.SetState(164)

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
		p.SetState(167)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(168)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(169)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY)|(1<<NumScriptParserTY_PORTION))) != 0) {
		{
			p.SetState(174)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(175)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(NumScriptParserRBRACE)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(185)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(188)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(190)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(196)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(199)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
				{
					p.SetState(200)
					p.Match(NumScriptParserNEWLINE)
				}

				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(205)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(211)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
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

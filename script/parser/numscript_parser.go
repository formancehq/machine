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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 197,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 5, 2, 37, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 5, 4, 48, 10, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	57, 10, 5, 3, 6, 3, 6, 3, 6, 5, 6, 62, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 67,
	10, 6, 12, 6, 14, 6, 70, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 85, 10, 9, 13, 9, 14, 9, 86,
	3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 93, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 6, 11, 100, 10, 11, 13, 11, 14, 11, 101, 3, 11, 3, 11, 3, 12, 3,
	12, 5, 12, 108, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 133, 10, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 138, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 6, 16, 148, 10, 16, 13, 16, 14, 16, 149, 3, 16, 3, 16, 6, 16,
	154, 10, 16, 13, 16, 14, 16, 155, 6, 16, 158, 10, 16, 13, 16, 14, 16, 159,
	3, 16, 3, 16, 6, 16, 164, 10, 16, 13, 16, 14, 16, 165, 3, 17, 7, 17, 169,
	10, 17, 12, 17, 14, 17, 172, 11, 17, 3, 17, 5, 17, 175, 10, 17, 3, 17,
	3, 17, 6, 17, 179, 10, 17, 13, 17, 14, 17, 180, 3, 17, 7, 17, 184, 10,
	17, 12, 17, 14, 17, 187, 11, 17, 3, 17, 7, 17, 190, 10, 17, 12, 17, 14,
	17, 193, 11, 17, 3, 17, 3, 17, 3, 17, 2, 3, 10, 18, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 14, 15, 3, 2, 24, 27,
	2, 204, 2, 36, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 56,
	3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2,
	16, 79, 3, 2, 2, 2, 18, 92, 3, 2, 2, 2, 20, 94, 3, 2, 2, 2, 22, 107, 3,
	2, 2, 2, 24, 137, 3, 2, 2, 2, 26, 139, 3, 2, 2, 2, 28, 141, 3, 2, 2, 2,
	30, 144, 3, 2, 2, 2, 32, 170, 3, 2, 2, 2, 34, 37, 7, 29, 2, 2, 35, 37,
	7, 22, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 3, 3, 2, 2, 2,
	38, 39, 7, 18, 2, 2, 39, 40, 7, 34, 2, 2, 40, 41, 5, 2, 2, 2, 41, 42, 7,
	19, 2, 2, 42, 5, 3, 2, 2, 2, 43, 51, 7, 28, 2, 2, 44, 47, 7, 29, 2, 2,
	45, 46, 7, 3, 2, 2, 46, 48, 7, 29, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3,
	2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 7, 30, 2, 2, 50, 43, 3, 2, 2, 2, 50,
	44, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 57, 7, 33, 2, 2, 53, 57, 7, 34,
	2, 2, 54, 57, 7, 29, 2, 2, 55, 57, 5, 4, 3, 2, 56, 52, 3, 2, 2, 2, 56,
	53, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 9, 3, 2, 2,
	2, 58, 59, 8, 6, 1, 2, 59, 62, 5, 8, 5, 2, 60, 62, 7, 32, 2, 2, 61, 58,
	3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 68, 3, 2, 2, 2, 63, 64, 12, 5, 2, 2,
	64, 65, 9, 2, 2, 2, 65, 67, 5, 10, 6, 6, 66, 63, 3, 2, 2, 2, 67, 70, 3,
	2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 11, 3, 2, 2, 2, 70,
	68, 3, 2, 2, 2, 71, 72, 7, 31, 2, 2, 72, 73, 7, 23, 2, 2, 73, 74, 5, 10,
	6, 2, 74, 13, 3, 2, 2, 2, 75, 76, 5, 6, 4, 2, 76, 77, 7, 4, 2, 2, 77, 78,
	5, 10, 6, 2, 78, 15, 3, 2, 2, 2, 79, 80, 7, 20, 2, 2, 80, 84, 7, 5, 2,
	2, 81, 82, 5, 14, 8, 2, 82, 83, 7, 5, 2, 2, 83, 85, 3, 2, 2, 2, 84, 81,
	3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2,
	87, 88, 3, 2, 2, 2, 88, 89, 7, 21, 2, 2, 89, 17, 3, 2, 2, 2, 90, 93, 5,
	16, 9, 2, 91, 93, 5, 10, 6, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2,
	93, 19, 3, 2, 2, 2, 94, 95, 7, 20, 2, 2, 95, 99, 7, 5, 2, 2, 96, 97, 5,
	10, 6, 2, 97, 98, 7, 5, 2, 2, 98, 100, 3, 2, 2, 2, 99, 96, 3, 2, 2, 2,
	100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103,
	3, 2, 2, 2, 103, 104, 7, 21, 2, 2, 104, 21, 3, 2, 2, 2, 105, 108, 5, 20,
	11, 2, 106, 108, 5, 10, 6, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2,
	2, 108, 23, 3, 2, 2, 2, 109, 110, 7, 8, 2, 2, 110, 138, 5, 10, 6, 2, 111,
	138, 7, 9, 2, 2, 112, 113, 7, 10, 2, 2, 113, 114, 5, 10, 6, 2, 114, 115,
	7, 16, 2, 2, 115, 132, 7, 5, 2, 2, 116, 117, 7, 11, 2, 2, 117, 118, 7,
	23, 2, 2, 118, 119, 5, 22, 12, 2, 119, 120, 7, 5, 2, 2, 120, 121, 7, 12,
	2, 2, 121, 122, 7, 23, 2, 2, 122, 123, 5, 18, 10, 2, 123, 133, 3, 2, 2,
	2, 124, 125, 7, 12, 2, 2, 125, 126, 7, 23, 2, 2, 126, 127, 5, 18, 10, 2,
	127, 128, 7, 5, 2, 2, 128, 129, 7, 11, 2, 2, 129, 130, 7, 23, 2, 2, 130,
	131, 5, 22, 12, 2, 131, 133, 3, 2, 2, 2, 132, 116, 3, 2, 2, 2, 132, 124,
	3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 7, 5, 2, 2, 135, 136, 7, 17,
	2, 2, 136, 138, 3, 2, 2, 2, 137, 109, 3, 2, 2, 2, 137, 111, 3, 2, 2, 2,
	137, 112, 3, 2, 2, 2, 138, 25, 3, 2, 2, 2, 139, 140, 9, 3, 2, 2, 140, 27,
	3, 2, 2, 2, 141, 142, 5, 26, 14, 2, 142, 143, 7, 32, 2, 2, 143, 29, 3,
	2, 2, 2, 144, 145, 7, 7, 2, 2, 145, 147, 7, 20, 2, 2, 146, 148, 7, 5, 2,
	2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 157, 3, 2, 2, 2, 151, 153, 5, 28, 15, 2, 152, 154,
	7, 5, 2, 2, 153, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157, 151, 3, 2, 2, 2,
	158, 159, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	161, 3, 2, 2, 2, 161, 163, 7, 21, 2, 2, 162, 164, 7, 5, 2, 2, 163, 162,
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 31, 3, 2, 2, 2, 167, 169, 7, 5, 2, 2, 168, 167, 3, 2, 2, 2,
	169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 175, 5, 30, 16, 2, 174, 173,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 185, 5, 24,
	13, 2, 177, 179, 7, 5, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2,
	180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182,
	184, 5, 24, 13, 2, 183, 178, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183,
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 191, 3, 2, 2, 2, 187, 185, 3, 2,
	2, 2, 188, 190, 7, 5, 2, 2, 189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2,
	191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2, 193,
	191, 3, 2, 2, 2, 194, 195, 7, 2, 2, 3, 195, 33, 3, 2, 2, 2, 23, 36, 47,
	50, 56, 61, 68, 86, 92, 101, 107, 132, 137, 149, 155, 159, 165, 170, 174,
	180, 185, 191,
}
var literalNames = []string{
	"", "'.'", "'to'", "", "", "'vars'", "'print'", "'fail'", "'send'", "'source'",
	"'destination'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'*'", "'='", "'account'", "'asset'", "'number'", "'monetary'",
	"", "", "'%'",
}
var symbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "SOURCE",
	"DESTINATION", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN", "LBRACK",
	"RBRACK", "LBRACE", "RBRACE", "ALL", "EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER",
	"TY_MONETARY", "RATIO", "NUMBER", "PERCENT", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"amount", "monetary", "frac", "literal", "expression", "argument", "allocationPart",
	"allocationBlock", "allocation", "sourceBlock", "source", "statement",
	"type_", "varDecl", "varListDecl", "script",
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
	NumScriptParserNEWLINE       = 3
	NumScriptParserWHITESPACE    = 4
	NumScriptParserVARS          = 5
	NumScriptParserPRINT         = 6
	NumScriptParserFAIL          = 7
	NumScriptParserSEND          = 8
	NumScriptParserSOURCE        = 9
	NumScriptParserDESTINATION   = 10
	NumScriptParserALLOCATE      = 11
	NumScriptParserOP_ADD        = 12
	NumScriptParserOP_SUB        = 13
	NumScriptParserLPAREN        = 14
	NumScriptParserRPAREN        = 15
	NumScriptParserLBRACK        = 16
	NumScriptParserRBRACK        = 17
	NumScriptParserLBRACE        = 18
	NumScriptParserRBRACE        = 19
	NumScriptParserALL           = 20
	NumScriptParserEQ            = 21
	NumScriptParserTY_ACCOUNT    = 22
	NumScriptParserTY_ASSET      = 23
	NumScriptParserTY_NUMBER     = 24
	NumScriptParserTY_MONETARY   = 25
	NumScriptParserRATIO         = 26
	NumScriptParserNUMBER        = 27
	NumScriptParserPERCENT       = 28
	NumScriptParserIDENTIFIER    = 29
	NumScriptParserVARIABLE_NAME = 30
	NumScriptParserACCOUNT       = 31
	NumScriptParserASSET         = 32
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_amount          = 0
	NumScriptParserRULE_monetary        = 1
	NumScriptParserRULE_frac            = 2
	NumScriptParserRULE_literal         = 3
	NumScriptParserRULE_expression      = 4
	NumScriptParserRULE_argument        = 5
	NumScriptParserRULE_allocationPart  = 6
	NumScriptParserRULE_allocationBlock = 7
	NumScriptParserRULE_allocation      = 8
	NumScriptParserRULE_sourceBlock     = 9
	NumScriptParserRULE_source          = 10
	NumScriptParserRULE_statement       = 11
	NumScriptParserRULE_type_           = 12
	NumScriptParserRULE_varDecl         = 13
	NumScriptParserRULE_varListDecl     = 14
	NumScriptParserRULE_script          = 15
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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserNUMBER:
		localctx = NewAmountSpecificContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*AmountSpecificContext).num = _m
		}

	case NumScriptParserALL:
		localctx = NewAmountAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
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
		p.SetState(36)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(37)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(38)

		var _x = p.Amount()

		localctx.(*MonetaryContext).amt = _x
	}
	{
		p.SetState(39)
		p.Match(NumScriptParserRBRACK)
	}

	return localctx
}

// IFracContext is an interface to support dynamic dispatch.
type IFracContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFracContext differentiates from other interfaces.
	IsFracContext()
}

type FracContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFracContext() *FracContext {
	var p = new(FracContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_frac
	return p
}

func (*FracContext) IsFracContext() {}

func NewFracContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FracContext {
	var p = new(FracContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_frac

	return p
}

func (s *FracContext) GetParser() antlr.Parser { return s.parser }

func (s *FracContext) CopyFrom(ctx *FracContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FracContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FracContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PercentageContext struct {
	*FracContext
	pint  antlr.Token
	pfrac antlr.Token
}

func NewPercentageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentageContext {
	var p = new(PercentageContext)

	p.FracContext = NewEmptyFracContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FracContext))

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
	*FracContext
	r antlr.Token
}

func NewRatioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RatioContext {
	var p = new(RatioContext)

	p.FracContext = NewEmptyFracContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FracContext))

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

func (p *NumScriptParser) Frac() (localctx IFracContext) {
	localctx = NewFracContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NumScriptParserRULE_frac)
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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserRATIO:
		localctx = NewRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)

			var _m = p.Match(NumScriptParserRATIO)

			localctx.(*RatioContext).r = _m
		}

	case NumScriptParserNUMBER:
		localctx = NewPercentageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*PercentageContext).pint = _m
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NumScriptParserT__0 {
			{
				p.SetState(43)
				p.Match(NumScriptParserT__0)
			}
			{
				p.SetState(44)

				var _m = p.Match(NumScriptParserNUMBER)

				localctx.(*PercentageContext).pfrac = _m
			}

		}
		{
			p.SetState(47)
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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Monetary()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	variable antlr.Token
}

func NewExprVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprVariableContext {
	var p = new(ExprVariableContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprVariableContext) GetVariable() antlr.Token { return s.variable }

func (s *ExprVariableContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *ExprVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprVariableContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(NumScriptParserVARIABLE_NAME, 0)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(57)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)

			var _m = p.Match(NumScriptParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
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
			p.SetState(61)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(62)

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
				p.SetState(63)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetVal returns the val rule contexts.
	GetVal() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IExpressionContext)

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	val    IExpressionContext
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) GetName() antlr.Token { return s.name }

func (s *ArgumentContext) SetName(v antlr.Token) { s.name = v }

func (s *ArgumentContext) GetVal() IExpressionContext { return s.val }

func (s *ArgumentContext) SetVal(v IExpressionContext) { s.val = v }

func (s *ArgumentContext) EQ() antlr.TerminalNode {
	return s.GetToken(NumScriptParserEQ, 0)
}

func (s *ArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserIDENTIFIER, 0)
}

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *NumScriptParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NumScriptParserRULE_argument)

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
		p.SetState(69)

		var _m = p.Match(NumScriptParserIDENTIFIER)

		localctx.(*ArgumentContext).name = _m
	}
	{
		p.SetState(70)
		p.Match(NumScriptParserEQ)
	}
	{
		p.SetState(71)

		var _x = p.expression(0)

		localctx.(*ArgumentContext).val = _x
	}

	return localctx
}

// IAllocationPartContext is an interface to support dynamic dispatch.
type IAllocationPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFr returns the fr rule contexts.
	GetFr() IFracContext

	// GetDest returns the dest rule contexts.
	GetDest() IExpressionContext

	// SetFr sets the fr rule contexts.
	SetFr(IFracContext)

	// SetDest sets the dest rule contexts.
	SetDest(IExpressionContext)

	// IsAllocationPartContext differentiates from other interfaces.
	IsAllocationPartContext()
}

type AllocationPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fr     IFracContext
	dest   IExpressionContext
}

func NewEmptyAllocationPartContext() *AllocationPartContext {
	var p = new(AllocationPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocationPart
	return p
}

func (*AllocationPartContext) IsAllocationPartContext() {}

func NewAllocationPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocationPartContext {
	var p = new(AllocationPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocationPart

	return p
}

func (s *AllocationPartContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocationPartContext) GetFr() IFracContext { return s.fr }

func (s *AllocationPartContext) GetDest() IExpressionContext { return s.dest }

func (s *AllocationPartContext) SetFr(v IFracContext) { s.fr = v }

func (s *AllocationPartContext) SetDest(v IExpressionContext) { s.dest = v }

func (s *AllocationPartContext) Frac() IFracContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFracContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFracContext)
}

func (s *AllocationPartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllocationPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocationPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocationPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocationPart(s)
	}
}

func (s *AllocationPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocationPart(s)
	}
}

func (p *NumScriptParser) AllocationPart() (localctx IAllocationPartContext) {
	localctx = NewAllocationPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumScriptParserRULE_allocationPart)

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
		p.SetState(73)

		var _x = p.Frac()

		localctx.(*AllocationPartContext).fr = _x
	}
	{
		p.SetState(74)
		p.Match(NumScriptParserT__1)
	}
	{
		p.SetState(75)

		var _x = p.expression(0)

		localctx.(*AllocationPartContext).dest = _x
	}

	return localctx
}

// IAllocationBlockContext is an interface to support dynamic dispatch.
type IAllocationBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allocationPart returns the _allocationPart rule contexts.
	Get_allocationPart() IAllocationPartContext

	// Set_allocationPart sets the _allocationPart rule contexts.
	Set_allocationPart(IAllocationPartContext)

	// GetParts returns the parts rule context list.
	GetParts() []IAllocationPartContext

	// SetParts sets the parts rule context list.
	SetParts([]IAllocationPartContext)

	// IsAllocationBlockContext differentiates from other interfaces.
	IsAllocationBlockContext()
}

type AllocationBlockContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_allocationPart IAllocationPartContext
	parts           []IAllocationPartContext
}

func NewEmptyAllocationBlockContext() *AllocationBlockContext {
	var p = new(AllocationBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumScriptParserRULE_allocationBlock
	return p
}

func (*AllocationBlockContext) IsAllocationBlockContext() {}

func NewAllocationBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocationBlockContext {
	var p = new(AllocationBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumScriptParserRULE_allocationBlock

	return p
}

func (s *AllocationBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocationBlockContext) Get_allocationPart() IAllocationPartContext {
	return s._allocationPart
}

func (s *AllocationBlockContext) Set_allocationPart(v IAllocationPartContext) { s._allocationPart = v }

func (s *AllocationBlockContext) GetParts() []IAllocationPartContext { return s.parts }

func (s *AllocationBlockContext) SetParts(v []IAllocationPartContext) { s.parts = v }

func (s *AllocationBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLBRACE, 0)
}

func (s *AllocationBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumScriptParserNEWLINE)
}

func (s *AllocationBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumScriptParserNEWLINE, i)
}

func (s *AllocationBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRBRACE, 0)
}

func (s *AllocationBlockContext) AllAllocationPart() []IAllocationPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllocationPartContext)(nil)).Elem())
	var tst = make([]IAllocationPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllocationPartContext)
		}
	}

	return tst
}

func (s *AllocationBlockContext) AllocationPart(i int) IAllocationPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocationPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllocationPartContext)
}

func (s *AllocationBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocationBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocationBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocationBlock(s)
	}
}

func (s *AllocationBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocationBlock(s)
	}
}

func (p *NumScriptParser) AllocationBlock() (localctx IAllocationBlockContext) {
	localctx = NewAllocationBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumScriptParserRULE_allocationBlock)
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
		p.SetState(77)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(78)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserRATIO || _la == NumScriptParserNUMBER {
		{
			p.SetState(79)

			var _x = p.AllocationPart()

			localctx.(*AllocationBlockContext)._allocationPart = _x
		}
		localctx.(*AllocationBlockContext).parts = append(localctx.(*AllocationBlockContext).parts, localctx.(*AllocationBlockContext)._allocationPart)
		{
			p.SetState(80)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
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

type AllocBlockContext struct {
	*AllocationContext
}

func NewAllocBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocBlockContext {
	var p = new(AllocBlockContext)

	p.AllocationContext = NewEmptyAllocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllocationContext))

	return p
}

func (s *AllocBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocBlockContext) AllocationBlock() IAllocationBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllocationBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllocationBlockContext)
}

func (s *AllocBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.EnterAllocBlock(s)
	}
}

func (s *AllocBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumScriptListener); ok {
		listenerT.ExitAllocBlock(s)
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

func (p *NumScriptParser) Allocation() (localctx IAllocationContext) {
	localctx = NewAllocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumScriptParserRULE_allocation)

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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACE:
		localctx = NewAllocBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.AllocationBlock()
		}

	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewAllocAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 18, NumScriptParserRULE_sourceBlock)
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
		p.SetState(92)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(93)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(NumScriptParserLBRACK-16))|(1<<(NumScriptParserNUMBER-16))|(1<<(NumScriptParserVARIABLE_NAME-16))|(1<<(NumScriptParserACCOUNT-16))|(1<<(NumScriptParserASSET-16)))) != 0) {
		{
			p.SetState(94)

			var _x = p.expression(0)

			localctx.(*SourceBlockContext)._expression = _x
		}
		localctx.(*SourceBlockContext).sources = append(localctx.(*SourceBlockContext).sources, localctx.(*SourceBlockContext)._expression)
		{
			p.SetState(95)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
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
	p.EnterRule(localctx, 20, NumScriptParserRULE_source)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACE:
		localctx = NewSrcBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.SourceBlock()
		}

	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
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
	p.EnterRule(localctx, 22, NumScriptParserRULE_statement)

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
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(108)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(111)

			var _x = p.expression(0)

			localctx.(*SendContext).mon = _x
		}
		{
			p.SetState(112)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(113)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(114)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(115)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(116)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(117)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(118)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(119)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(120)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(122)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(123)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(124)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(125)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(126)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(127)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(128)

				var _x = p.Source()

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(132)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(133)
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
	p.EnterRule(localctx, 24, NumScriptParserRULE_type_)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY))) != 0) {
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

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTy returns the ty rule contexts.
	GetTy() IType_Context

	// SetTy sets the ty rule contexts.
	SetTy(IType_Context)

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     IType_Context
	name   antlr.Token
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

func (s *VarDeclContext) GetName() antlr.Token { return s.name }

func (s *VarDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VarDeclContext) GetTy() IType_Context { return s.ty }

func (s *VarDeclContext) SetTy(v IType_Context) { s.ty = v }

func (s *VarDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDeclContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(NumScriptParserVARIABLE_NAME, 0)
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
	p.EnterRule(localctx, 26, NumScriptParserRULE_varDecl)

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
		p.SetState(139)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(140)

		var _m = p.Match(NumScriptParserVARIABLE_NAME)

		localctx.(*VarDeclContext).name = _m
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
	p.EnterRule(localctx, 28, NumScriptParserRULE_varListDecl)
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
		p.SetState(142)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(143)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(144)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY))) != 0) {
		{
			p.SetState(149)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(150)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(159)
		p.Match(NumScriptParserRBRACE)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(160)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(163)
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
	p.EnterRule(localctx, 30, NumScriptParserRULE_script)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(165)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(171)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(174)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
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
			{
				p.SetState(180)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(186)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
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

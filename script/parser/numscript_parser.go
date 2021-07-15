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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 162,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5,
	3, 37, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 5, 3, 5, 3, 5,
	5, 5, 48, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 53, 10, 5, 12, 5, 14, 5, 56, 11,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 6, 8, 71, 10, 8, 13, 8, 14, 8, 72, 3, 8, 3, 8, 3, 9, 3, 9, 5,
	9, 79, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 104, 10, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 109, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	6, 13, 119, 10, 13, 13, 13, 14, 13, 120, 3, 13, 3, 13, 6, 13, 125, 10,
	13, 13, 13, 14, 13, 126, 6, 13, 129, 10, 13, 13, 13, 14, 13, 130, 3, 13,
	3, 13, 6, 13, 135, 10, 13, 13, 13, 14, 13, 136, 3, 14, 5, 14, 140, 10,
	14, 3, 14, 3, 14, 6, 14, 144, 10, 14, 13, 14, 14, 14, 145, 3, 14, 7, 14,
	149, 10, 14, 12, 14, 14, 14, 152, 11, 14, 3, 14, 7, 14, 155, 10, 14, 12,
	14, 14, 14, 158, 11, 14, 3, 14, 3, 14, 3, 14, 2, 3, 8, 15, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 4, 3, 2, 13, 14, 3, 2, 22, 25, 2,
	167, 2, 28, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 47, 3,
	2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 61, 3, 2, 2, 2, 14, 65, 3, 2, 2, 2, 16,
	78, 3, 2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 110, 3, 2, 2, 2, 22, 112, 3, 2,
	2, 2, 24, 115, 3, 2, 2, 2, 26, 139, 3, 2, 2, 2, 28, 29, 7, 17, 2, 2, 29,
	30, 7, 32, 2, 2, 30, 31, 7, 27, 2, 2, 31, 32, 7, 18, 2, 2, 32, 3, 3, 2,
	2, 2, 33, 37, 7, 26, 2, 2, 34, 35, 7, 27, 2, 2, 35, 37, 7, 28, 2, 2, 36,
	33, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 5, 3, 2, 2, 2, 38, 43, 7, 31, 2,
	2, 39, 43, 7, 32, 2, 2, 40, 43, 7, 27, 2, 2, 41, 43, 5, 2, 2, 2, 42, 38,
	3, 2, 2, 2, 42, 39, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2,
	43, 7, 3, 2, 2, 2, 44, 45, 8, 5, 1, 2, 45, 48, 5, 6, 4, 2, 46, 48, 7, 30,
	2, 2, 47, 44, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 54, 3, 2, 2, 2, 49, 50,
	12, 5, 2, 2, 50, 51, 9, 2, 2, 2, 51, 53, 5, 8, 5, 6, 52, 49, 3, 2, 2, 2,
	53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 9, 3, 2,
	2, 2, 56, 54, 3, 2, 2, 2, 57, 58, 7, 29, 2, 2, 58, 59, 7, 21, 2, 2, 59,
	60, 5, 8, 5, 2, 60, 11, 3, 2, 2, 2, 61, 62, 5, 4, 3, 2, 62, 63, 7, 3, 2,
	2, 63, 64, 5, 8, 5, 2, 64, 13, 3, 2, 2, 2, 65, 66, 7, 19, 2, 2, 66, 70,
	7, 4, 2, 2, 67, 68, 5, 12, 7, 2, 68, 69, 7, 4, 2, 2, 69, 71, 3, 2, 2, 2,
	70, 67, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 7, 20, 2, 2, 75, 15, 3, 2, 2, 2, 76,
	79, 5, 14, 8, 2, 77, 79, 5, 8, 5, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2,
	2, 2, 79, 17, 3, 2, 2, 2, 80, 81, 7, 7, 2, 2, 81, 109, 5, 8, 5, 2, 82,
	109, 7, 8, 2, 2, 83, 84, 7, 9, 2, 2, 84, 85, 5, 8, 5, 2, 85, 86, 7, 15,
	2, 2, 86, 103, 7, 4, 2, 2, 87, 88, 7, 10, 2, 2, 88, 89, 7, 21, 2, 2, 89,
	90, 5, 8, 5, 2, 90, 91, 7, 4, 2, 2, 91, 92, 7, 11, 2, 2, 92, 93, 7, 21,
	2, 2, 93, 94, 5, 16, 9, 2, 94, 104, 3, 2, 2, 2, 95, 96, 7, 11, 2, 2, 96,
	97, 7, 21, 2, 2, 97, 98, 5, 16, 9, 2, 98, 99, 7, 4, 2, 2, 99, 100, 7, 10,
	2, 2, 100, 101, 7, 21, 2, 2, 101, 102, 5, 8, 5, 2, 102, 104, 3, 2, 2, 2,
	103, 87, 3, 2, 2, 2, 103, 95, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106,
	7, 4, 2, 2, 106, 107, 7, 16, 2, 2, 107, 109, 3, 2, 2, 2, 108, 80, 3, 2,
	2, 2, 108, 82, 3, 2, 2, 2, 108, 83, 3, 2, 2, 2, 109, 19, 3, 2, 2, 2, 110,
	111, 9, 3, 2, 2, 111, 21, 3, 2, 2, 2, 112, 113, 5, 20, 11, 2, 113, 114,
	7, 30, 2, 2, 114, 23, 3, 2, 2, 2, 115, 116, 7, 6, 2, 2, 116, 118, 7, 19,
	2, 2, 117, 119, 7, 4, 2, 2, 118, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 128, 3, 2, 2, 2, 122,
	124, 5, 22, 12, 2, 123, 125, 7, 4, 2, 2, 124, 123, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2,
	2, 2, 128, 122, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2,
	130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 7, 20, 2, 2, 133,
	135, 7, 4, 2, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 25, 3, 2, 2, 2, 138, 140, 5, 24,
	13, 2, 139, 138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2,
	141, 150, 5, 18, 10, 2, 142, 144, 7, 4, 2, 2, 143, 142, 3, 2, 2, 2, 144,
	145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147,
	3, 2, 2, 2, 147, 149, 5, 18, 10, 2, 148, 143, 3, 2, 2, 2, 149, 152, 3,
	2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 156, 3, 2, 2,
	2, 152, 150, 3, 2, 2, 2, 153, 155, 7, 4, 2, 2, 154, 153, 3, 2, 2, 2, 155,
	158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159,
	3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 7, 2, 2, 3, 160, 27, 3, 2,
	2, 2, 18, 36, 42, 47, 54, 72, 78, 103, 108, 120, 126, 130, 136, 139, 145,
	150, 156,
}
var literalNames = []string{
	"", "'to'", "", "", "'vars'", "'print'", "'fail'", "'send'", "'source'",
	"'destination'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'='", "'account'", "'asset'", "'number'", "'monetary'",
	"", "", "'%'",
}
var symbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "VARS", "PRINT", "FAIL", "SEND", "SOURCE",
	"DESTINATION", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN", "LBRACK",
	"RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER",
	"TY_MONETARY", "RATIO", "NUMBER", "PERCENT", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var ruleNames = []string{
	"monetary", "frac", "literal", "expression", "argument", "allocationPart",
	"allocationBlock", "allocation", "statement", "type_", "varDecl", "varListDecl",
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
	NumScriptParserEOF           = antlr.TokenEOF
	NumScriptParserT__0          = 1
	NumScriptParserNEWLINE       = 2
	NumScriptParserWHITESPACE    = 3
	NumScriptParserVARS          = 4
	NumScriptParserPRINT         = 5
	NumScriptParserFAIL          = 6
	NumScriptParserSEND          = 7
	NumScriptParserSOURCE        = 8
	NumScriptParserDESTINATION   = 9
	NumScriptParserALLOCATE      = 10
	NumScriptParserOP_ADD        = 11
	NumScriptParserOP_SUB        = 12
	NumScriptParserLPAREN        = 13
	NumScriptParserRPAREN        = 14
	NumScriptParserLBRACK        = 15
	NumScriptParserRBRACK        = 16
	NumScriptParserLBRACE        = 17
	NumScriptParserRBRACE        = 18
	NumScriptParserEQ            = 19
	NumScriptParserTY_ACCOUNT    = 20
	NumScriptParserTY_ASSET      = 21
	NumScriptParserTY_NUMBER     = 22
	NumScriptParserTY_MONETARY   = 23
	NumScriptParserRATIO         = 24
	NumScriptParserNUMBER        = 25
	NumScriptParserPERCENT       = 26
	NumScriptParserIDENTIFIER    = 27
	NumScriptParserVARIABLE_NAME = 28
	NumScriptParserACCOUNT       = 29
	NumScriptParserASSET         = 30
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_monetary        = 0
	NumScriptParserRULE_frac            = 1
	NumScriptParserRULE_literal         = 2
	NumScriptParserRULE_expression      = 3
	NumScriptParserRULE_argument        = 4
	NumScriptParserRULE_allocationPart  = 5
	NumScriptParserRULE_allocationBlock = 6
	NumScriptParserRULE_allocation      = 7
	NumScriptParserRULE_statement       = 8
	NumScriptParserRULE_type_           = 9
	NumScriptParserRULE_varDecl         = 10
	NumScriptParserRULE_varListDecl     = 11
	NumScriptParserRULE_script          = 12
)

// IMonetaryContext is an interface to support dynamic dispatch.
type IMonetaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsset returns the asset token.
	GetAsset() antlr.Token

	// GetAmount returns the amount token.
	GetAmount() antlr.Token

	// SetAsset sets the asset token.
	SetAsset(antlr.Token)

	// SetAmount sets the amount token.
	SetAmount(antlr.Token)

	// IsMonetaryContext differentiates from other interfaces.
	IsMonetaryContext()
}

type MonetaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	asset  antlr.Token
	amount antlr.Token
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

func (s *MonetaryContext) GetAmount() antlr.Token { return s.amount }

func (s *MonetaryContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryContext) SetAmount(v antlr.Token) { s.amount = v }

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
		p.SetState(26)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(27)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(28)

		var _m = p.Match(NumScriptParserNUMBER)

		localctx.(*MonetaryContext).amount = _m
	}
	{
		p.SetState(29)
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
	p antlr.Token
}

func NewPercentageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentageContext {
	var p = new(PercentageContext)

	p.FracContext = NewEmptyFracContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FracContext))

	return p
}

func (s *PercentageContext) GetP() antlr.Token { return s.p }

func (s *PercentageContext) SetP(v antlr.Token) { s.p = v }

func (s *PercentageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentageContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(NumScriptParserPERCENT, 0)
}

func (s *PercentageContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumScriptParserNUMBER, 0)
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
	p.EnterRule(localctx, 2, NumScriptParserRULE_frac)

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
	case NumScriptParserRATIO:
		localctx = NewRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)

			var _m = p.Match(NumScriptParserRATIO)

			localctx.(*RatioContext).r = _m
		}

	case NumScriptParserNUMBER:
		localctx = NewPercentageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)

			var _m = p.Match(NumScriptParserNUMBER)

			localctx.(*PercentageContext).p = _m
		}
		{
			p.SetState(33)
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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, NumScriptParserRULE_expression, _p)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(43)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)

			var _m = p.Match(NumScriptParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(52)
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
			p.SetState(47)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(48)

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
				p.SetState(49)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 8, NumScriptParserRULE_argument)

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
		p.SetState(55)

		var _m = p.Match(NumScriptParserIDENTIFIER)

		localctx.(*ArgumentContext).name = _m
	}
	{
		p.SetState(56)
		p.Match(NumScriptParserEQ)
	}
	{
		p.SetState(57)

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
	p.EnterRule(localctx, 10, NumScriptParserRULE_allocationPart)

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
		p.SetState(59)

		var _x = p.Frac()

		localctx.(*AllocationPartContext).fr = _x
	}
	{
		p.SetState(60)
		p.Match(NumScriptParserT__0)
	}
	{
		p.SetState(61)

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
	p.EnterRule(localctx, 12, NumScriptParserRULE_allocationBlock)
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
		p.SetState(63)
		p.Match(NumScriptParserLBRACE)
	}
	{
		p.SetState(64)
		p.Match(NumScriptParserNEWLINE)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserRATIO || _la == NumScriptParserNUMBER {
		{
			p.SetState(65)

			var _x = p.AllocationPart()

			localctx.(*AllocationBlockContext)._allocationPart = _x
		}
		localctx.(*AllocationBlockContext).parts = append(localctx.(*AllocationBlockContext).parts, localctx.(*AllocationBlockContext)._allocationPart)
		{
			p.SetState(66)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
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
	p.EnterRule(localctx, 14, NumScriptParserRULE_allocation)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACE:
		localctx = NewAllocBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.AllocationBlock()
		}

	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserVARIABLE_NAME, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewAllocAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
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
	src  IExpressionContext
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

func (s *SendContext) GetSrc() IExpressionContext { return s.src }

func (s *SendContext) GetDest() IAllocationContext { return s.dest }

func (s *SendContext) SetMon(v IExpressionContext) { s.mon = v }

func (s *SendContext) SetSrc(v IExpressionContext) { s.src = v }

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

func (s *SendContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SendContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 16, NumScriptParserRULE_statement)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(79)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(82)

			var _x = p.expression(0)

			localctx.(*SendContext).mon = _x
		}
		{
			p.SetState(83)
			p.Match(NumScriptParserLPAREN)
		}
		{
			p.SetState(84)
			p.Match(NumScriptParserNEWLINE)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NumScriptParserSOURCE:
			{
				p.SetState(85)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(86)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(87)

				var _x = p.expression(0)

				localctx.(*SendContext).src = _x
			}
			{
				p.SetState(88)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(89)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(90)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(91)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}

		case NumScriptParserDESTINATION:
			{
				p.SetState(93)
				p.Match(NumScriptParserDESTINATION)
			}
			{
				p.SetState(94)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(95)

				var _x = p.Allocation()

				localctx.(*SendContext).dest = _x
			}
			{
				p.SetState(96)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(97)
				p.Match(NumScriptParserSOURCE)
			}
			{
				p.SetState(98)
				p.Match(NumScriptParserEQ)
			}
			{
				p.SetState(99)

				var _x = p.expression(0)

				localctx.(*SendContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(103)
			p.Match(NumScriptParserNEWLINE)
		}
		{
			p.SetState(104)
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
	p.EnterRule(localctx, 18, NumScriptParserRULE_type_)
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
		p.SetState(108)
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
	p.EnterRule(localctx, 20, NumScriptParserRULE_varDecl)

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
		p.SetState(110)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(111)

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
	p.EnterRule(localctx, 22, NumScriptParserRULE_varListDecl)
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
		p.SetState(113)
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(114)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(115)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY))) != 0) {
		{
			p.SetState(120)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
			{
				p.SetState(121)
				p.Match(NumScriptParserNEWLINE)
			}

			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(NumScriptParserRBRACE)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
		{
			p.SetState(131)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(134)
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
	p.EnterRule(localctx, 24, NumScriptParserRULE_script)
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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(136)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(139)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NumScriptParserNEWLINE {
				{
					p.SetState(140)
					p.Match(NumScriptParserNEWLINE)
				}

				p.SetState(143)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(145)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NumScriptParserNEWLINE {
		{
			p.SetState(151)
			p.Match(NumScriptParserNEWLINE)
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(NumScriptParserEOF)
	}

	return localctx
}

func (p *NumScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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

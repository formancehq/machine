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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 98, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 35, 10, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 57, 10,
	6, 13, 6, 14, 6, 58, 3, 6, 5, 6, 62, 10, 6, 3, 6, 3, 6, 5, 6, 66, 10, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 6, 9, 76, 10, 9, 13, 9,
	14, 9, 77, 3, 9, 3, 9, 3, 10, 5, 10, 83, 10, 10, 3, 10, 3, 10, 3, 10, 7,
	10, 88, 10, 10, 12, 10, 14, 10, 91, 11, 10, 3, 10, 5, 10, 94, 10, 10, 3,
	10, 3, 10, 3, 10, 2, 3, 6, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 3,
	2, 9, 10, 3, 2, 18, 21, 2, 101, 2, 20, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6,
	34, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 67, 3, 2, 2,
	2, 14, 69, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2, 18, 82, 3, 2, 2, 2, 20, 21,
	7, 13, 2, 2, 21, 22, 7, 26, 2, 2, 22, 23, 7, 22, 2, 2, 23, 24, 7, 14, 2,
	2, 24, 3, 3, 2, 2, 2, 25, 30, 7, 25, 2, 2, 26, 30, 7, 26, 2, 2, 27, 30,
	7, 22, 2, 2, 28, 30, 5, 2, 2, 2, 29, 25, 3, 2, 2, 2, 29, 26, 3, 2, 2, 2,
	29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31, 32, 8, 4,
	1, 2, 32, 35, 5, 4, 3, 2, 33, 35, 7, 24, 2, 2, 34, 31, 3, 2, 2, 2, 34,
	33, 3, 2, 2, 2, 35, 41, 3, 2, 2, 2, 36, 37, 12, 5, 2, 2, 37, 38, 9, 2,
	2, 2, 38, 40, 5, 6, 4, 6, 39, 36, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 7, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2,
	44, 45, 7, 23, 2, 2, 45, 46, 7, 17, 2, 2, 46, 47, 5, 6, 4, 2, 47, 9, 3,
	2, 2, 2, 48, 49, 7, 6, 2, 2, 49, 66, 5, 6, 4, 2, 50, 66, 7, 7, 2, 2, 51,
	52, 7, 8, 2, 2, 52, 56, 7, 11, 2, 2, 53, 54, 5, 8, 5, 2, 54, 55, 7, 3,
	2, 2, 55, 57, 3, 2, 2, 2, 56, 53, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 62, 5, 8, 5, 2,
	61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7,
	12, 2, 2, 64, 66, 3, 2, 2, 2, 65, 48, 3, 2, 2, 2, 65, 50, 3, 2, 2, 2, 65,
	51, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68, 9, 3, 2, 2, 68, 13, 3, 2, 2,
	2, 69, 70, 5, 12, 7, 2, 70, 71, 7, 24, 2, 2, 71, 15, 3, 2, 2, 2, 72, 73,
	7, 5, 2, 2, 73, 75, 7, 15, 2, 2, 74, 76, 5, 14, 8, 2, 75, 74, 3, 2, 2,
	2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79,
	3, 2, 2, 2, 79, 80, 7, 16, 2, 2, 80, 17, 3, 2, 2, 2, 81, 83, 5, 16, 9,
	2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 89,
	5, 10, 6, 2, 85, 86, 7, 4, 2, 2, 86, 88, 5, 10, 6, 2, 87, 85, 3, 2, 2,
	2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 93,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 94, 7, 4, 2, 2, 93, 92, 3, 2, 2, 2,
	93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 7, 2, 2, 3, 96, 19, 3,
	2, 2, 2, 12, 29, 34, 41, 58, 61, 65, 77, 82, 89, 93,
}
var literalNames = []string{
	"", "','", "", "'vars'", "'print'", "'fail'", "'send'", "'+'", "'-'", "'('",
	"')'", "'['", "']'", "'{'", "'}'", "'='", "'account'", "'asset'", "'number'",
	"'monetary'", "", "", "", "", "", "';'",
}
var symbolicNames = []string{
	"", "", "NEWLINE", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT",
	"TY_ASSET", "TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "SEP", "WHITESPACE",
}

var ruleNames = []string{
	"monetary", "literal", "expression", "argument", "statement", "type_",
	"varDecl", "varListDecl", "script",
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
	NumScriptParserVARS          = 3
	NumScriptParserPRINT         = 4
	NumScriptParserFAIL          = 5
	NumScriptParserSEND          = 6
	NumScriptParserOP_ADD        = 7
	NumScriptParserOP_SUB        = 8
	NumScriptParserLPAREN        = 9
	NumScriptParserRPAREN        = 10
	NumScriptParserLBRACK        = 11
	NumScriptParserRBRACK        = 12
	NumScriptParserLBRACE        = 13
	NumScriptParserRBRACE        = 14
	NumScriptParserEQ            = 15
	NumScriptParserTY_ACCOUNT    = 16
	NumScriptParserTY_ASSET      = 17
	NumScriptParserTY_NUMBER     = 18
	NumScriptParserTY_MONETARY   = 19
	NumScriptParserNUMBER        = 20
	NumScriptParserIDENTIFIER    = 21
	NumScriptParserVARIABLE_NAME = 22
	NumScriptParserACCOUNT       = 23
	NumScriptParserASSET         = 24
	NumScriptParserSEP           = 25
	NumScriptParserWHITESPACE    = 26
)

// NumScriptParser rules.
const (
	NumScriptParserRULE_monetary    = 0
	NumScriptParserRULE_literal     = 1
	NumScriptParserRULE_expression  = 2
	NumScriptParserRULE_argument    = 3
	NumScriptParserRULE_statement   = 4
	NumScriptParserRULE_type_       = 5
	NumScriptParserRULE_varDecl     = 6
	NumScriptParserRULE_varListDecl = 7
	NumScriptParserRULE_script      = 8
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
		p.SetState(18)
		p.Match(NumScriptParserLBRACK)
	}
	{
		p.SetState(19)

		var _m = p.Match(NumScriptParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(20)

		var _m = p.Match(NumScriptParserNUMBER)

		localctx.(*MonetaryContext).amount = _m
	}
	{
		p.SetState(21)
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
	p.EnterRule(localctx, 2, NumScriptParserRULE_literal)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(NumScriptParserACCOUNT)
		}

	case NumScriptParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(NumScriptParserASSET)
		}

	case NumScriptParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.Match(NumScriptParserNUMBER)
		}

	case NumScriptParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, NumScriptParserRULE_expression, _p)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserLBRACK, NumScriptParserNUMBER, NumScriptParserACCOUNT, NumScriptParserASSET:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(30)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case NumScriptParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)

			var _m = p.Match(NumScriptParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
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
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(35)

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
				p.SetState(36)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, NumScriptParserRULE_argument)

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
		p.SetState(42)

		var _m = p.Match(NumScriptParserIDENTIFIER)

		localctx.(*ArgumentContext).name = _m
	}
	{
		p.SetState(43)
		p.Match(NumScriptParserEQ)
	}
	{
		p.SetState(44)

		var _x = p.expression(0)

		localctx.(*ArgumentContext).val = _x
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
	_argument IArgumentContext
	args      []IArgumentContext
}

func NewSendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendContext {
	var p = new(SendContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SendContext) Get_argument() IArgumentContext { return s._argument }

func (s *SendContext) Set_argument(v IArgumentContext) { s._argument = v }

func (s *SendContext) GetArgs() []IArgumentContext { return s.args }

func (s *SendContext) SetArgs(v []IArgumentContext) { s.args = v }

func (s *SendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendContext) SEND() antlr.TerminalNode {
	return s.GetToken(NumScriptParserSEND, 0)
}

func (s *SendContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserLPAREN, 0)
}

func (s *SendContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(NumScriptParserRPAREN, 0)
}

func (s *SendContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *SendContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
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
	p.EnterRule(localctx, 8, NumScriptParserRULE_statement)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NumScriptParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(NumScriptParserPRINT)
		}
		{
			p.SetState(47)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case NumScriptParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(NumScriptParserFAIL)
		}

	case NumScriptParserSEND:
		localctx = NewSendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(NumScriptParserSEND)
		}
		{
			p.SetState(50)
			p.Match(NumScriptParserLPAREN)
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(51)

					var _x = p.Argument()

					localctx.(*SendContext)._argument = _x
				}
				localctx.(*SendContext).args = append(localctx.(*SendContext).args, localctx.(*SendContext)._argument)
				{
					p.SetState(52)
					p.Match(NumScriptParserT__0)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NumScriptParserIDENTIFIER {
			{
				p.SetState(58)

				var _x = p.Argument()

				localctx.(*SendContext)._argument = _x
			}
			localctx.(*SendContext).args = append(localctx.(*SendContext).args, localctx.(*SendContext)._argument)

		}

		{
			p.SetState(61)
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
	p.EnterRule(localctx, 10, NumScriptParserRULE_type_)
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
		p.SetState(65)
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
	p.EnterRule(localctx, 12, NumScriptParserRULE_varDecl)

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
		p.SetState(67)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(68)

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
	p.EnterRule(localctx, 14, NumScriptParserRULE_varListDecl)
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
		p.Match(NumScriptParserVARS)
	}
	{
		p.SetState(71)
		p.Match(NumScriptParserLBRACE)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NumScriptParserTY_ACCOUNT)|(1<<NumScriptParserTY_ASSET)|(1<<NumScriptParserTY_NUMBER)|(1<<NumScriptParserTY_MONETARY))) != 0) {
		{
			p.SetState(72)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(NumScriptParserRBRACE)
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
	p.EnterRule(localctx, 16, NumScriptParserRULE_script)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserVARS {
		{
			p.SetState(79)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(82)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(83)
				p.Match(NumScriptParserNEWLINE)
			}
			{
				p.SetState(84)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NumScriptParserNEWLINE {
		{
			p.SetState(90)
			p.Match(NumScriptParserNEWLINE)
		}

	}
	{
		p.SetState(93)
		p.Match(NumScriptParserEOF)
	}

	return localctx
}

func (p *NumScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
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

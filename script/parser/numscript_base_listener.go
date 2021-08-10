// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // NumScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNumScriptListener is a complete listener for a parse tree produced by NumScriptParser.
type BaseNumScriptListener struct{}

var _ NumScriptListener = &BaseNumScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNumScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNumScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNumScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNumScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMonetary is called when production monetary is entered.
func (s *BaseNumScriptListener) EnterMonetary(ctx *MonetaryContext) {}

// ExitMonetary is called when production monetary is exited.
func (s *BaseNumScriptListener) ExitMonetary(ctx *MonetaryContext) {}

// EnterMonetaryAll is called when production monetaryAll is entered.
func (s *BaseNumScriptListener) EnterMonetaryAll(ctx *MonetaryAllContext) {}

// ExitMonetaryAll is called when production monetaryAll is exited.
func (s *BaseNumScriptListener) ExitMonetaryAll(ctx *MonetaryAllContext) {}

// EnterLitAccount is called when production LitAccount is entered.
func (s *BaseNumScriptListener) EnterLitAccount(ctx *LitAccountContext) {}

// ExitLitAccount is called when production LitAccount is exited.
func (s *BaseNumScriptListener) ExitLitAccount(ctx *LitAccountContext) {}

// EnterLitAsset is called when production LitAsset is entered.
func (s *BaseNumScriptListener) EnterLitAsset(ctx *LitAssetContext) {}

// ExitLitAsset is called when production LitAsset is exited.
func (s *BaseNumScriptListener) ExitLitAsset(ctx *LitAssetContext) {}

// EnterLitNumber is called when production LitNumber is entered.
func (s *BaseNumScriptListener) EnterLitNumber(ctx *LitNumberContext) {}

// ExitLitNumber is called when production LitNumber is exited.
func (s *BaseNumScriptListener) ExitLitNumber(ctx *LitNumberContext) {}

// EnterLitMonetary is called when production LitMonetary is entered.
func (s *BaseNumScriptListener) EnterLitMonetary(ctx *LitMonetaryContext) {}

// ExitLitMonetary is called when production LitMonetary is exited.
func (s *BaseNumScriptListener) ExitLitMonetary(ctx *LitMonetaryContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseNumScriptListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseNumScriptListener) ExitVariable(ctx *VariableContext) {}

// EnterExprAddSub is called when production ExprAddSub is entered.
func (s *BaseNumScriptListener) EnterExprAddSub(ctx *ExprAddSubContext) {}

// ExitExprAddSub is called when production ExprAddSub is exited.
func (s *BaseNumScriptListener) ExitExprAddSub(ctx *ExprAddSubContext) {}

// EnterExprLiteral is called when production ExprLiteral is entered.
func (s *BaseNumScriptListener) EnterExprLiteral(ctx *ExprLiteralContext) {}

// ExitExprLiteral is called when production ExprLiteral is exited.
func (s *BaseNumScriptListener) ExitExprLiteral(ctx *ExprLiteralContext) {}

// EnterExprVariable is called when production ExprVariable is entered.
func (s *BaseNumScriptListener) EnterExprVariable(ctx *ExprVariableContext) {}

// ExitExprVariable is called when production ExprVariable is exited.
func (s *BaseNumScriptListener) ExitExprVariable(ctx *ExprVariableContext) {}

// EnterAllotmentPortionConst is called when production allotmentPortionConst is entered.
func (s *BaseNumScriptListener) EnterAllotmentPortionConst(ctx *AllotmentPortionConstContext) {}

// ExitAllotmentPortionConst is called when production allotmentPortionConst is exited.
func (s *BaseNumScriptListener) ExitAllotmentPortionConst(ctx *AllotmentPortionConstContext) {}

// EnterAllotmentPortionVar is called when production allotmentPortionVar is entered.
func (s *BaseNumScriptListener) EnterAllotmentPortionVar(ctx *AllotmentPortionVarContext) {}

// ExitAllotmentPortionVar is called when production allotmentPortionVar is exited.
func (s *BaseNumScriptListener) ExitAllotmentPortionVar(ctx *AllotmentPortionVarContext) {}

// EnterAllotmentPortionRemaining is called when production allotmentPortionRemaining is entered.
func (s *BaseNumScriptListener) EnterAllotmentPortionRemaining(ctx *AllotmentPortionRemainingContext) {
}

// ExitAllotmentPortionRemaining is called when production allotmentPortionRemaining is exited.
func (s *BaseNumScriptListener) ExitAllotmentPortionRemaining(ctx *AllotmentPortionRemainingContext) {
}

// EnterDestinationAllotment is called when production destinationAllotment is entered.
func (s *BaseNumScriptListener) EnterDestinationAllotment(ctx *DestinationAllotmentContext) {}

// ExitDestinationAllotment is called when production destinationAllotment is exited.
func (s *BaseNumScriptListener) ExitDestinationAllotment(ctx *DestinationAllotmentContext) {}

// EnterDestAccount is called when production DestAccount is entered.
func (s *BaseNumScriptListener) EnterDestAccount(ctx *DestAccountContext) {}

// ExitDestAccount is called when production DestAccount is exited.
func (s *BaseNumScriptListener) ExitDestAccount(ctx *DestAccountContext) {}

// EnterDestAllotment is called when production DestAllotment is entered.
func (s *BaseNumScriptListener) EnterDestAllotment(ctx *DestAllotmentContext) {}

// ExitDestAllotment is called when production DestAllotment is exited.
func (s *BaseNumScriptListener) ExitDestAllotment(ctx *DestAllotmentContext) {}

// EnterSourceBlock is called when production sourceBlock is entered.
func (s *BaseNumScriptListener) EnterSourceBlock(ctx *SourceBlockContext) {}

// ExitSourceBlock is called when production sourceBlock is exited.
func (s *BaseNumScriptListener) ExitSourceBlock(ctx *SourceBlockContext) {}

// EnterSourceAllotment is called when production sourceAllotment is entered.
func (s *BaseNumScriptListener) EnterSourceAllotment(ctx *SourceAllotmentContext) {}

// ExitSourceAllotment is called when production sourceAllotment is exited.
func (s *BaseNumScriptListener) ExitSourceAllotment(ctx *SourceAllotmentContext) {}

// EnterSrcAccount is called when production SrcAccount is entered.
func (s *BaseNumScriptListener) EnterSrcAccount(ctx *SrcAccountContext) {}

// ExitSrcAccount is called when production SrcAccount is exited.
func (s *BaseNumScriptListener) ExitSrcAccount(ctx *SrcAccountContext) {}

// EnterSrcBlock is called when production SrcBlock is entered.
func (s *BaseNumScriptListener) EnterSrcBlock(ctx *SrcBlockContext) {}

// ExitSrcBlock is called when production SrcBlock is exited.
func (s *BaseNumScriptListener) ExitSrcBlock(ctx *SrcBlockContext) {}

// EnterSrcAllotment is called when production SrcAllotment is entered.
func (s *BaseNumScriptListener) EnterSrcAllotment(ctx *SrcAllotmentContext) {}

// ExitSrcAllotment is called when production SrcAllotment is exited.
func (s *BaseNumScriptListener) ExitSrcAllotment(ctx *SrcAllotmentContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseNumScriptListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseNumScriptListener) ExitPrint(ctx *PrintContext) {}

// EnterFail is called when production Fail is entered.
func (s *BaseNumScriptListener) EnterFail(ctx *FailContext) {}

// ExitFail is called when production Fail is exited.
func (s *BaseNumScriptListener) ExitFail(ctx *FailContext) {}

// EnterSend is called when production Send is entered.
func (s *BaseNumScriptListener) EnterSend(ctx *SendContext) {}

// ExitSend is called when production Send is exited.
func (s *BaseNumScriptListener) ExitSend(ctx *SendContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseNumScriptListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseNumScriptListener) ExitType_(ctx *Type_Context) {}

// EnterOrigin is called when production origin is entered.
func (s *BaseNumScriptListener) EnterOrigin(ctx *OriginContext) {}

// ExitOrigin is called when production origin is exited.
func (s *BaseNumScriptListener) ExitOrigin(ctx *OriginContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseNumScriptListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseNumScriptListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarListDecl is called when production varListDecl is entered.
func (s *BaseNumScriptListener) EnterVarListDecl(ctx *VarListDeclContext) {}

// ExitVarListDecl is called when production varListDecl is exited.
func (s *BaseNumScriptListener) ExitVarListDecl(ctx *VarListDeclContext) {}

// EnterScript is called when production script is entered.
func (s *BaseNumScriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseNumScriptListener) ExitScript(ctx *ScriptContext) {}

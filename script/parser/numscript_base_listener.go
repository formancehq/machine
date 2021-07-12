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

// EnterRatio is called when production Ratio is entered.
func (s *BaseNumScriptListener) EnterRatio(ctx *RatioContext) {}

// ExitRatio is called when production Ratio is exited.
func (s *BaseNumScriptListener) ExitRatio(ctx *RatioContext) {}

// EnterPercentage is called when production Percentage is entered.
func (s *BaseNumScriptListener) EnterPercentage(ctx *PercentageContext) {}

// ExitPercentage is called when production Percentage is exited.
func (s *BaseNumScriptListener) ExitPercentage(ctx *PercentageContext) {}

// EnterAllocationPart is called when production allocationPart is entered.
func (s *BaseNumScriptListener) EnterAllocationPart(ctx *AllocationPartContext) {}

// ExitAllocationPart is called when production allocationPart is exited.
func (s *BaseNumScriptListener) ExitAllocationPart(ctx *AllocationPartContext) {}

// EnterAllocation is called when production allocation is entered.
func (s *BaseNumScriptListener) EnterAllocation(ctx *AllocationContext) {}

// ExitAllocation is called when production allocation is exited.
func (s *BaseNumScriptListener) ExitAllocation(ctx *AllocationContext) {}

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

// EnterLitAllocation is called when production LitAllocation is entered.
func (s *BaseNumScriptListener) EnterLitAllocation(ctx *LitAllocationContext) {}

// ExitLitAllocation is called when production LitAllocation is exited.
func (s *BaseNumScriptListener) ExitLitAllocation(ctx *LitAllocationContext) {}

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

// EnterArgument is called when production argument is entered.
func (s *BaseNumScriptListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseNumScriptListener) ExitArgument(ctx *ArgumentContext) {}

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

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

// EnterAmountSpecific is called when production AmountSpecific is entered.
func (s *BaseNumScriptListener) EnterAmountSpecific(ctx *AmountSpecificContext) {}

// ExitAmountSpecific is called when production AmountSpecific is exited.
func (s *BaseNumScriptListener) ExitAmountSpecific(ctx *AmountSpecificContext) {}

// EnterAmountAll is called when production AmountAll is entered.
func (s *BaseNumScriptListener) EnterAmountAll(ctx *AmountAllContext) {}

// ExitAmountAll is called when production AmountAll is exited.
func (s *BaseNumScriptListener) ExitAmountAll(ctx *AmountAllContext) {}

// EnterMonetary is called when production monetary is entered.
func (s *BaseNumScriptListener) EnterMonetary(ctx *MonetaryContext) {}

// ExitMonetary is called when production monetary is exited.
func (s *BaseNumScriptListener) ExitMonetary(ctx *MonetaryContext) {}

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

// EnterAllocPartConstConst is called when production allocPartConstConst is entered.
func (s *BaseNumScriptListener) EnterAllocPartConstConst(ctx *AllocPartConstConstContext) {}

// ExitAllocPartConstConst is called when production allocPartConstConst is exited.
func (s *BaseNumScriptListener) ExitAllocPartConstConst(ctx *AllocPartConstConstContext) {}

// EnterAllocPartConstRemaining is called when production allocPartConstRemaining is entered.
func (s *BaseNumScriptListener) EnterAllocPartConstRemaining(ctx *AllocPartConstRemainingContext) {}

// ExitAllocPartConstRemaining is called when production allocPartConstRemaining is exited.
func (s *BaseNumScriptListener) ExitAllocPartConstRemaining(ctx *AllocPartConstRemainingContext) {}

// EnterAllocBlockConst is called when production allocBlockConst is entered.
func (s *BaseNumScriptListener) EnterAllocBlockConst(ctx *AllocBlockConstContext) {}

// ExitAllocBlockConst is called when production allocBlockConst is exited.
func (s *BaseNumScriptListener) ExitAllocBlockConst(ctx *AllocBlockConstContext) {}

// EnterAllocPartDynConst is called when production allocPartDynConst is entered.
func (s *BaseNumScriptListener) EnterAllocPartDynConst(ctx *AllocPartDynConstContext) {}

// ExitAllocPartDynConst is called when production allocPartDynConst is exited.
func (s *BaseNumScriptListener) ExitAllocPartDynConst(ctx *AllocPartDynConstContext) {}

// EnterAllocPartDynVar is called when production allocPartDynVar is entered.
func (s *BaseNumScriptListener) EnterAllocPartDynVar(ctx *AllocPartDynVarContext) {}

// ExitAllocPartDynVar is called when production allocPartDynVar is exited.
func (s *BaseNumScriptListener) ExitAllocPartDynVar(ctx *AllocPartDynVarContext) {}

// EnterAllocPartDynRemaining is called when production allocPartDynRemaining is entered.
func (s *BaseNumScriptListener) EnterAllocPartDynRemaining(ctx *AllocPartDynRemainingContext) {}

// ExitAllocPartDynRemaining is called when production allocPartDynRemaining is exited.
func (s *BaseNumScriptListener) ExitAllocPartDynRemaining(ctx *AllocPartDynRemainingContext) {}

// EnterAllocBlockDyn is called when production allocBlockDyn is entered.
func (s *BaseNumScriptListener) EnterAllocBlockDyn(ctx *AllocBlockDynContext) {}

// ExitAllocBlockDyn is called when production allocBlockDyn is exited.
func (s *BaseNumScriptListener) ExitAllocBlockDyn(ctx *AllocBlockDynContext) {}

// EnterAllocConst is called when production AllocConst is entered.
func (s *BaseNumScriptListener) EnterAllocConst(ctx *AllocConstContext) {}

// ExitAllocConst is called when production AllocConst is exited.
func (s *BaseNumScriptListener) ExitAllocConst(ctx *AllocConstContext) {}

// EnterAllocDyn is called when production AllocDyn is entered.
func (s *BaseNumScriptListener) EnterAllocDyn(ctx *AllocDynContext) {}

// ExitAllocDyn is called when production AllocDyn is exited.
func (s *BaseNumScriptListener) ExitAllocDyn(ctx *AllocDynContext) {}

// EnterAllocAccount is called when production AllocAccount is entered.
func (s *BaseNumScriptListener) EnterAllocAccount(ctx *AllocAccountContext) {}

// ExitAllocAccount is called when production AllocAccount is exited.
func (s *BaseNumScriptListener) ExitAllocAccount(ctx *AllocAccountContext) {}

// EnterSourceBlock is called when production sourceBlock is entered.
func (s *BaseNumScriptListener) EnterSourceBlock(ctx *SourceBlockContext) {}

// ExitSourceBlock is called when production sourceBlock is exited.
func (s *BaseNumScriptListener) ExitSourceBlock(ctx *SourceBlockContext) {}

// EnterSrcBlock is called when production SrcBlock is entered.
func (s *BaseNumScriptListener) EnterSrcBlock(ctx *SrcBlockContext) {}

// ExitSrcBlock is called when production SrcBlock is exited.
func (s *BaseNumScriptListener) ExitSrcBlock(ctx *SrcBlockContext) {}

// EnterSrcAccount is called when production SrcAccount is entered.
func (s *BaseNumScriptListener) EnterSrcAccount(ctx *SrcAccountContext) {}

// ExitSrcAccount is called when production SrcAccount is exited.
func (s *BaseNumScriptListener) ExitSrcAccount(ctx *SrcAccountContext) {}

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

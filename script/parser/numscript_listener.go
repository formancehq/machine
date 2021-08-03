// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // NumScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NumScriptListener is a complete listener for a parse tree produced by NumScriptParser.
type NumScriptListener interface {
	antlr.ParseTreeListener

	// EnterAmountSpecific is called when entering the AmountSpecific production.
	EnterAmountSpecific(c *AmountSpecificContext)

	// EnterAmountAll is called when entering the AmountAll production.
	EnterAmountAll(c *AmountAllContext)

	// EnterMonetary is called when entering the monetary production.
	EnterMonetary(c *MonetaryContext)

	// EnterLitAccount is called when entering the LitAccount production.
	EnterLitAccount(c *LitAccountContext)

	// EnterLitAsset is called when entering the LitAsset production.
	EnterLitAsset(c *LitAssetContext)

	// EnterLitNumber is called when entering the LitNumber production.
	EnterLitNumber(c *LitNumberContext)

	// EnterLitMonetary is called when entering the LitMonetary production.
	EnterLitMonetary(c *LitMonetaryContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExprAddSub is called when entering the ExprAddSub production.
	EnterExprAddSub(c *ExprAddSubContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprVariable is called when entering the ExprVariable production.
	EnterExprVariable(c *ExprVariableContext)

	// EnterAllocPartConstConst is called when entering the allocPartConstConst production.
	EnterAllocPartConstConst(c *AllocPartConstConstContext)

	// EnterAllocPartConstRemaining is called when entering the allocPartConstRemaining production.
	EnterAllocPartConstRemaining(c *AllocPartConstRemainingContext)

	// EnterAllocBlockConst is called when entering the allocBlockConst production.
	EnterAllocBlockConst(c *AllocBlockConstContext)

	// EnterAllocPartDynConst is called when entering the allocPartDynConst production.
	EnterAllocPartDynConst(c *AllocPartDynConstContext)

	// EnterAllocPartDynVar is called when entering the allocPartDynVar production.
	EnterAllocPartDynVar(c *AllocPartDynVarContext)

	// EnterAllocPartDynRemaining is called when entering the allocPartDynRemaining production.
	EnterAllocPartDynRemaining(c *AllocPartDynRemainingContext)

	// EnterAllocBlockDyn is called when entering the allocBlockDyn production.
	EnterAllocBlockDyn(c *AllocBlockDynContext)

	// EnterAllocConst is called when entering the AllocConst production.
	EnterAllocConst(c *AllocConstContext)

	// EnterAllocDyn is called when entering the AllocDyn production.
	EnterAllocDyn(c *AllocDynContext)

	// EnterAllocAccount is called when entering the AllocAccount production.
	EnterAllocAccount(c *AllocAccountContext)

	// EnterSourceBlock is called when entering the sourceBlock production.
	EnterSourceBlock(c *SourceBlockContext)

	// EnterSrcBlock is called when entering the SrcBlock production.
	EnterSrcBlock(c *SrcBlockContext)

	// EnterSrcAccount is called when entering the SrcAccount production.
	EnterSrcAccount(c *SrcAccountContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterFail is called when entering the Fail production.
	EnterFail(c *FailContext)

	// EnterSend is called when entering the Send production.
	EnterSend(c *SendContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterOrigin is called when entering the origin production.
	EnterOrigin(c *OriginContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterVarListDecl is called when entering the varListDecl production.
	EnterVarListDecl(c *VarListDeclContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitAmountSpecific is called when exiting the AmountSpecific production.
	ExitAmountSpecific(c *AmountSpecificContext)

	// ExitAmountAll is called when exiting the AmountAll production.
	ExitAmountAll(c *AmountAllContext)

	// ExitMonetary is called when exiting the monetary production.
	ExitMonetary(c *MonetaryContext)

	// ExitLitAccount is called when exiting the LitAccount production.
	ExitLitAccount(c *LitAccountContext)

	// ExitLitAsset is called when exiting the LitAsset production.
	ExitLitAsset(c *LitAssetContext)

	// ExitLitNumber is called when exiting the LitNumber production.
	ExitLitNumber(c *LitNumberContext)

	// ExitLitMonetary is called when exiting the LitMonetary production.
	ExitLitMonetary(c *LitMonetaryContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExprAddSub is called when exiting the ExprAddSub production.
	ExitExprAddSub(c *ExprAddSubContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprVariable is called when exiting the ExprVariable production.
	ExitExprVariable(c *ExprVariableContext)

	// ExitAllocPartConstConst is called when exiting the allocPartConstConst production.
	ExitAllocPartConstConst(c *AllocPartConstConstContext)

	// ExitAllocPartConstRemaining is called when exiting the allocPartConstRemaining production.
	ExitAllocPartConstRemaining(c *AllocPartConstRemainingContext)

	// ExitAllocBlockConst is called when exiting the allocBlockConst production.
	ExitAllocBlockConst(c *AllocBlockConstContext)

	// ExitAllocPartDynConst is called when exiting the allocPartDynConst production.
	ExitAllocPartDynConst(c *AllocPartDynConstContext)

	// ExitAllocPartDynVar is called when exiting the allocPartDynVar production.
	ExitAllocPartDynVar(c *AllocPartDynVarContext)

	// ExitAllocPartDynRemaining is called when exiting the allocPartDynRemaining production.
	ExitAllocPartDynRemaining(c *AllocPartDynRemainingContext)

	// ExitAllocBlockDyn is called when exiting the allocBlockDyn production.
	ExitAllocBlockDyn(c *AllocBlockDynContext)

	// ExitAllocConst is called when exiting the AllocConst production.
	ExitAllocConst(c *AllocConstContext)

	// ExitAllocDyn is called when exiting the AllocDyn production.
	ExitAllocDyn(c *AllocDynContext)

	// ExitAllocAccount is called when exiting the AllocAccount production.
	ExitAllocAccount(c *AllocAccountContext)

	// ExitSourceBlock is called when exiting the sourceBlock production.
	ExitSourceBlock(c *SourceBlockContext)

	// ExitSrcBlock is called when exiting the SrcBlock production.
	ExitSrcBlock(c *SrcBlockContext)

	// ExitSrcAccount is called when exiting the SrcAccount production.
	ExitSrcAccount(c *SrcAccountContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitFail is called when exiting the Fail production.
	ExitFail(c *FailContext)

	// ExitSend is called when exiting the Send production.
	ExitSend(c *SendContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitOrigin is called when exiting the origin production.
	ExitOrigin(c *OriginContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVarListDecl is called when exiting the varListDecl production.
	ExitVarListDecl(c *VarListDeclContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}

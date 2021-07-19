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

	// EnterRatio is called when entering the Ratio production.
	EnterRatio(c *RatioContext)

	// EnterPercentage is called when entering the Percentage production.
	EnterPercentage(c *PercentageContext)

	// EnterLitAccount is called when entering the LitAccount production.
	EnterLitAccount(c *LitAccountContext)

	// EnterLitAsset is called when entering the LitAsset production.
	EnterLitAsset(c *LitAssetContext)

	// EnterLitNumber is called when entering the LitNumber production.
	EnterLitNumber(c *LitNumberContext)

	// EnterLitMonetary is called when entering the LitMonetary production.
	EnterLitMonetary(c *LitMonetaryContext)

	// EnterExprAddSub is called when entering the ExprAddSub production.
	EnterExprAddSub(c *ExprAddSubContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprVariable is called when entering the ExprVariable production.
	EnterExprVariable(c *ExprVariableContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterAllocationPart is called when entering the allocationPart production.
	EnterAllocationPart(c *AllocationPartContext)

	// EnterAllocationBlock is called when entering the allocationBlock production.
	EnterAllocationBlock(c *AllocationBlockContext)

	// EnterAllocBlock is called when entering the AllocBlock production.
	EnterAllocBlock(c *AllocBlockContext)

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

	// ExitRatio is called when exiting the Ratio production.
	ExitRatio(c *RatioContext)

	// ExitPercentage is called when exiting the Percentage production.
	ExitPercentage(c *PercentageContext)

	// ExitLitAccount is called when exiting the LitAccount production.
	ExitLitAccount(c *LitAccountContext)

	// ExitLitAsset is called when exiting the LitAsset production.
	ExitLitAsset(c *LitAssetContext)

	// ExitLitNumber is called when exiting the LitNumber production.
	ExitLitNumber(c *LitNumberContext)

	// ExitLitMonetary is called when exiting the LitMonetary production.
	ExitLitMonetary(c *LitMonetaryContext)

	// ExitExprAddSub is called when exiting the ExprAddSub production.
	ExitExprAddSub(c *ExprAddSubContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprVariable is called when exiting the ExprVariable production.
	ExitExprVariable(c *ExprVariableContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitAllocationPart is called when exiting the allocationPart production.
	ExitAllocationPart(c *AllocationPartContext)

	// ExitAllocationBlock is called when exiting the allocationBlock production.
	ExitAllocationBlock(c *AllocationBlockContext)

	// ExitAllocBlock is called when exiting the AllocBlock production.
	ExitAllocBlock(c *AllocBlockContext)

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

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVarListDecl is called when exiting the varListDecl production.
	ExitVarListDecl(c *VarListDeclContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}

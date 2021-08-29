// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // NumScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NumScriptListener is a complete listener for a parse tree produced by NumScriptParser.
type NumScriptListener interface {
	antlr.ParseTreeListener

	// EnterMonetary is called when entering the monetary production.
	EnterMonetary(c *MonetaryContext)

	// EnterMonetaryAll is called when entering the monetaryAll production.
	EnterMonetaryAll(c *MonetaryAllContext)

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

	// EnterAllotmentPortionConst is called when entering the allotmentPortionConst production.
	EnterAllotmentPortionConst(c *AllotmentPortionConstContext)

	// EnterAllotmentPortionVar is called when entering the allotmentPortionVar production.
	EnterAllotmentPortionVar(c *AllotmentPortionVarContext)

	// EnterAllotmentPortionRemaining is called when entering the allotmentPortionRemaining production.
	EnterAllotmentPortionRemaining(c *AllotmentPortionRemainingContext)

	// EnterDestinationInOrder is called when entering the destinationInOrder production.
	EnterDestinationInOrder(c *DestinationInOrderContext)

	// EnterDestinationMaxed is called when entering the destinationMaxed production.
	EnterDestinationMaxed(c *DestinationMaxedContext)

	// EnterDestinationAllotment is called when entering the destinationAllotment production.
	EnterDestinationAllotment(c *DestinationAllotmentContext)

	// EnterDestAccount is called when entering the DestAccount production.
	EnterDestAccount(c *DestAccountContext)

	// EnterDestMaxed is called when entering the DestMaxed production.
	EnterDestMaxed(c *DestMaxedContext)

	// EnterDestInOrder is called when entering the DestInOrder production.
	EnterDestInOrder(c *DestInOrderContext)

	// EnterDestAllotment is called when entering the DestAllotment production.
	EnterDestAllotment(c *DestAllotmentContext)

	// EnterSourceInOrder is called when entering the sourceInOrder production.
	EnterSourceInOrder(c *SourceInOrderContext)

	// EnterSourceMaxed is called when entering the sourceMaxed production.
	EnterSourceMaxed(c *SourceMaxedContext)

	// EnterSrcAccount is called when entering the SrcAccount production.
	EnterSrcAccount(c *SrcAccountContext)

	// EnterSrcMaxed is called when entering the SrcMaxed production.
	EnterSrcMaxed(c *SrcMaxedContext)

	// EnterSrcInOrder is called when entering the SrcInOrder production.
	EnterSrcInOrder(c *SrcInOrderContext)

	// EnterSourceAllotment is called when entering the sourceAllotment production.
	EnterSourceAllotment(c *SourceAllotmentContext)

	// EnterSrc is called when entering the Src production.
	EnterSrc(c *SrcContext)

	// EnterSrcAllotment is called when entering the SrcAllotment production.
	EnterSrcAllotment(c *SrcAllotmentContext)

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

	// ExitMonetary is called when exiting the monetary production.
	ExitMonetary(c *MonetaryContext)

	// ExitMonetaryAll is called when exiting the monetaryAll production.
	ExitMonetaryAll(c *MonetaryAllContext)

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

	// ExitAllotmentPortionConst is called when exiting the allotmentPortionConst production.
	ExitAllotmentPortionConst(c *AllotmentPortionConstContext)

	// ExitAllotmentPortionVar is called when exiting the allotmentPortionVar production.
	ExitAllotmentPortionVar(c *AllotmentPortionVarContext)

	// ExitAllotmentPortionRemaining is called when exiting the allotmentPortionRemaining production.
	ExitAllotmentPortionRemaining(c *AllotmentPortionRemainingContext)

	// ExitDestinationInOrder is called when exiting the destinationInOrder production.
	ExitDestinationInOrder(c *DestinationInOrderContext)

	// ExitDestinationMaxed is called when exiting the destinationMaxed production.
	ExitDestinationMaxed(c *DestinationMaxedContext)

	// ExitDestinationAllotment is called when exiting the destinationAllotment production.
	ExitDestinationAllotment(c *DestinationAllotmentContext)

	// ExitDestAccount is called when exiting the DestAccount production.
	ExitDestAccount(c *DestAccountContext)

	// ExitDestMaxed is called when exiting the DestMaxed production.
	ExitDestMaxed(c *DestMaxedContext)

	// ExitDestInOrder is called when exiting the DestInOrder production.
	ExitDestInOrder(c *DestInOrderContext)

	// ExitDestAllotment is called when exiting the DestAllotment production.
	ExitDestAllotment(c *DestAllotmentContext)

	// ExitSourceInOrder is called when exiting the sourceInOrder production.
	ExitSourceInOrder(c *SourceInOrderContext)

	// ExitSourceMaxed is called when exiting the sourceMaxed production.
	ExitSourceMaxed(c *SourceMaxedContext)

	// ExitSrcAccount is called when exiting the SrcAccount production.
	ExitSrcAccount(c *SrcAccountContext)

	// ExitSrcMaxed is called when exiting the SrcMaxed production.
	ExitSrcMaxed(c *SrcMaxedContext)

	// ExitSrcInOrder is called when exiting the SrcInOrder production.
	ExitSrcInOrder(c *SrcInOrderContext)

	// ExitSourceAllotment is called when exiting the sourceAllotment production.
	ExitSourceAllotment(c *SourceAllotmentContext)

	// ExitSrc is called when exiting the Src production.
	ExitSrc(c *SrcContext)

	// ExitSrcAllotment is called when exiting the SrcAllotment production.
	ExitSrcAllotment(c *SrcAllotmentContext)

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

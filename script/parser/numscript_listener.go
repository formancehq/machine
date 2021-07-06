// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // NumScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NumScriptListener is a complete listener for a parse tree produced by NumScriptParser.
type NumScriptListener interface {
	antlr.ParseTreeListener

	// EnterMonetary is called when entering the monetary production.
	EnterMonetary(c *MonetaryContext)

	// EnterLitAddress is called when entering the LitAddress production.
	EnterLitAddress(c *LitAddressContext)

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

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterFail is called when entering the Fail production.
	EnterFail(c *FailContext)

	// EnterSend is called when entering the Send production.
	EnterSend(c *SendContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitMonetary is called when exiting the monetary production.
	ExitMonetary(c *MonetaryContext)

	// ExitLitAddress is called when exiting the LitAddress production.
	ExitLitAddress(c *LitAddressContext)

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

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitFail is called when exiting the Fail production.
	ExitFail(c *FailContext)

	// ExitSend is called when exiting the Send production.
	ExitSend(c *SendContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}

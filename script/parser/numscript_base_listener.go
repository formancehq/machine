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

// EnterNumber is called when production Number is entered.
func (s *BaseNumScriptListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseNumScriptListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseNumScriptListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseNumScriptListener) ExitAddSub(ctx *AddSubContext) {}

// EnterScript is called when production script is entered.
func (s *BaseNumScriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseNumScriptListener) ExitScript(ctx *ScriptContext) {}

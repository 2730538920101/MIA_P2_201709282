// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCmdListener is a complete listener for a parse tree produced by CmdParser.
type BaseCmdListener struct{}

var _ CmdListener = &BaseCmdListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCmdListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCmdListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCmdListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCmdListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCmdListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCmdListener) ExitStart(ctx *StartContext) {}

// EnterComandList is called when production comandList is entered.
func (s *BaseCmdListener) EnterComandList(ctx *ComandListContext) {}

// ExitComandList is called when production comandList is exited.
func (s *BaseCmdListener) ExitComandList(ctx *ComandListContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseCmdListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseCmdListener) ExitComando(ctx *ComandoContext) {}

// EnterParam_list is called when production param_list is entered.
func (s *BaseCmdListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BaseCmdListener) ExitParam_list(ctx *Param_listContext) {}

// EnterParam is called when production param is entered.
func (s *BaseCmdListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseCmdListener) ExitParam(ctx *ParamContext) {}

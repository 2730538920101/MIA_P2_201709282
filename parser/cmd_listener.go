// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CmdListener is a complete listener for a parse tree produced by CmdParser.
type CmdListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterComandList is called when entering the comandList production.
	EnterComandList(c *ComandListContext)

	// EnterComando is called when entering the comando production.
	EnterComando(c *ComandoContext)

	// EnterParam_list is called when entering the param_list production.
	EnterParam_list(c *Param_listContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitComandList is called when exiting the comandList production.
	ExitComandList(c *ComandListContext)

	// ExitComando is called when exiting the comando production.
	ExitComando(c *ComandoContext)

	// ExitParam_list is called when exiting the param_list production.
	ExitParam_list(c *Param_listContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)
}

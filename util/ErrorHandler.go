package util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorSintactico struct{
	line, column int 
	msg string
}

type ErrorListener struct{
	*antlr.DefaultErrorListener
	Errors []ErrorSintactico
}

func ( c *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException){
	c.Errors = append(c.Errors, ErrorSintactico{
		line: line,
		column: column,
		msg: msg,
	})
}
package util

import (
	"../ast"
	"../parser"
)

type TreeShapeListener struct {
	*parser.BaseCmdListener
	Ast ast.Ast
}

func NewTreeShapeListener() *TreeShapeListener{
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext){
	this.Ast = ctx.GetAst()
}

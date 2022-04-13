package main

import (
	"fmt"
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type cmdListener struct{
	*parser.BaseCmdListener
	stack []string
}

func (l *cmdListener) push(i string){
	l.stack = append(l.stack, i)
}

func (l *cmdListener) pop() string{
	if len(l.stack) < 1 {
		panic("Error en el comando")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func Analizar(input string) string{
	is := antlr.NewInputStream(input)
	lexer := parser.NewCmdLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCmdParser(stream)
	var listener cmdListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	return listener.pop()
}

func main(){
	fmt.Println("HOLA MUNDO")
}


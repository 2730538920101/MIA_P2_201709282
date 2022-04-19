package main

import (
	"fmt"
	"bufio"
	"os"
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os/exec"
	"./util"
	"./ast"
)



func Analizar(input string){
	errores := &util.ErrorListener{}
	is := antlr.NewInputStream(input)
	lexer := parser.NewCmdLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errores)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCmdParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errores)
	p.BuildParseTrees = true
	var listener *util.TreeShapeListener = util.NewTreeShapeListener()
	if len(errores.Errors)==0 {
		antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	}
	AST := listener.Ast
	for i := 0; i < AST.ComList.Len(); i++{
		r := AST.ComList.GetValue(i)
		if r != nil{
			r.(ast.Command).Ejecutar()
		}
	}
}

func main(){
	for true{
		fmt.Print("Ingrese un comando: ")
		prueba := bufio.NewReader(os.Stdin)
		input, _ := prueba.ReadString('\n')
		if input !="\n"{
			switch input {
				case "exit\n":
					{
						os.Exit(0)
					}
				case "cls\n":
					{
						cmd := exec.Command("clear")
						cmd.Stdout = os.Stdout
						cmd.Run()
					}
				default:
					{

						Analizar(input)
						fmt.Println("ANALISIS REALIZADO")	
					}					
			}
		}
	}
}


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
	lexerErrors := &util.ErrorListener{}
	is := antlr.NewInputStream(input)
	lexer := parser.NewCmdLex(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCmdParser(stream)
	p.RemoveErrorListeners()
	parserErrors := &util.ErrorListener{}
	p.AddErrorListener(parserErrors)
	p.BuildParseTrees = true
	var listener *util.TreeShapeListener = util.NewTreeShapeListener()
	if len(lexerErrors.Errors)==0 && len(parserErrors.Errors)==0{
		antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())
	}
	if len(lexerErrors.Errors) > 0 {
        fmt.Printf("Lexer %d errors found", len(lexerErrors.Errors))
        for _, e := range lexerErrors.Errors {
            fmt.Println("", e)
        }
    }

    if len(parserErrors.Errors) > 0 {
        fmt.Printf("Parser %d errors found", len(parserErrors.Errors))
        for _, e := range parserErrors.Errors {
            fmt.Println("", e)
        }
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


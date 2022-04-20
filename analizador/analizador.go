package analizador

import(
	"fmt"
	"../parser"
	"../ast"
	"../util"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"bufio"
	"os"
	"log"
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
			res := r.(ast.Command).Ejecutar()
			if res != nil{
				LeerArch(res.(string))
			}
		}
	}
}


func LeerArch(text string){
	file, err := os.Open(text)
	if err != nil {
        log.Fatalf("Error al abrir el archivo: %s", err)
    }
	fileScanner := bufio.NewScanner(file)
    // read line by line
    for fileScanner.Scan() {
		Analizar(fileScanner.Text())
    }
    // handle first encountered error while reading
    if err := fileScanner.Err(); err != nil {
        log.Fatalf("Error while reading file: %s", err)
    }
    file.Close()
}
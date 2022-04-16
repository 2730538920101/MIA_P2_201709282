package main

import (
	"fmt"
	"bufio"
	"os"
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

func (l *cmdListener) ExitMkdisk(c *parser.MkdiskContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRmdisk(c *parser.RmdiskContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitFdisk(c *parser.FdiskContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMount(c *parser.MountContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMkfs(c *parser.MkfsContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitLogin(c *parser.LoginContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitLogout(c *parser.LogoutContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMkgrp(c *parser.MkgrpContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRmgrp(c *parser.RmgrpContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMkusr(c *parser.MkusrContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRmusr(c *parser.RmusrContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMkfile(c *parser.MkfileContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitMkdir(c *parser.MkdirContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitExec(c *parser.ExecContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitPause(c *parser.PauseContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRep(c *parser.RepContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitAddCommand(c *parser.AddCommandContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitSize(c *parser.SizeContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitPath_R(c *parser.Path_RContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitPath_S(c *parser.Path_SContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitFitBF(c *parser.FitBFContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitFitFF(c *parser.FitFFContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitFitWF(c *parser.FitWFContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitUnit_K(c *parser.Unit_KContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitUnit_M(c *parser.Unit_MContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitUnit_B(c *parser.Unit_BContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitName_P(c *parser.Name_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitName_S(c *parser.Name_SContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitUsr_P(c *parser.Usr_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitUsr_S(c *parser.Usr_SContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitGrp_P(c *parser.Grp_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitGrp_S(c *parser.Grp_SContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitPass_P(c *parser.Pass_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitPwd_P(c *parser.Pwd_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitType_P(c *parser.Type_PContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitType_E(c *parser.Type_EContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitType_L(c *parser.Type_LContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitType_Fast(c *parser.Type_FastContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitType_Full(c *parser.Type_FullContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitId(c *parser.IdContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitCont_R(c *parser.Cont_RContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitCont_S(c *parser.Cont_SContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRuta_R(c *parser.Ruta_RContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRuta_S(c *parser.Ruta_SContext){
	i := c.GetText()
	l.push(i)
}


func (l *cmdListener) ExitPp(c *parser.PpContext){
	i := c.GetText()
	l.push(i)
}

func (l *cmdListener) ExitRr(c *parser.RrContext){
	i := c.GetText()
	l.push(i)
}

func main(){
	fmt.Print("Ingrese un comando: ")
	prueba := bufio.NewReader(os.Stdin)
	input, _ := prueba.ReadString('\n')
	if input != "\n"{
		res :=  Analizar(input)
		fmt.Printf("RESPUESTA: %s\n",res)
	}
	
}


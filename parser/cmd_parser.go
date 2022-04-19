// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CmdParser struct {
	*antlr.BaseParser
}

var cmdParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cmdParserInit() {
	staticData := &cmdParserStaticData
	staticData.symbolicNames = []string{
		"", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", "TOK_MKFS",
		"TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR",
		"TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", "TOK_PATH",
		"TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", "TOK_ID",
		"TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", "TOK_RUTA",
		"TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", "TOK_KB", "TOK_MB",
		"TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FAST",
		"TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", "TOK_CAMINO",
		"TOK_PALABRA", "TOK_IGUAL", "COMENTARIOS", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "comando", "comando_estado", "param_list", "param",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 131, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 18, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 34, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 41, 8, 3, 10, 3, 12,
		3, 44, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 129, 8, 4, 1, 4, 0, 1, 6, 5, 0, 2, 4, 6, 8, 0, 0, 169, 0, 10,
		1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8,
		128, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 1, 1, 0, 0, 0, 12, 13, 3, 4, 2,
		0, 13, 14, 3, 6, 3, 0, 14, 18, 1, 0, 0, 0, 15, 18, 5, 15, 0, 0, 16, 18,
		5, 7, 0, 0, 17, 12, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16, 1, 0, 0, 0,
		18, 3, 1, 0, 0, 0, 19, 34, 5, 1, 0, 0, 20, 34, 5, 2, 0, 0, 21, 34, 5, 3,
		0, 0, 22, 34, 5, 4, 0, 0, 23, 34, 5, 5, 0, 0, 24, 34, 5, 6, 0, 0, 25, 34,
		5, 8, 0, 0, 26, 34, 5, 9, 0, 0, 27, 34, 5, 10, 0, 0, 28, 34, 5, 11, 0,
		0, 29, 34, 5, 12, 0, 0, 30, 34, 5, 13, 0, 0, 31, 34, 5, 14, 0, 0, 32, 34,
		5, 16, 0, 0, 33, 19, 1, 0, 0, 0, 33, 20, 1, 0, 0, 0, 33, 21, 1, 0, 0, 0,
		33, 22, 1, 0, 0, 0, 33, 23, 1, 0, 0, 0, 33, 24, 1, 0, 0, 0, 33, 25, 1,
		0, 0, 0, 33, 26, 1, 0, 0, 0, 33, 27, 1, 0, 0, 0, 33, 28, 1, 0, 0, 0, 33,
		29, 1, 0, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 32, 1, 0, 0,
		0, 34, 5, 1, 0, 0, 0, 35, 36, 6, 3, -1, 0, 36, 37, 3, 8, 4, 0, 37, 42,
		1, 0, 0, 0, 38, 39, 10, 2, 0, 0, 39, 41, 3, 8, 4, 0, 40, 38, 1, 0, 0, 0,
		41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 7, 1, 0,
		0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 5, 19, 0, 0, 46, 47, 5, 48, 0, 0, 47,
		129, 5, 44, 0, 0, 48, 49, 5, 17, 0, 0, 49, 50, 5, 48, 0, 0, 50, 129, 5,
		46, 0, 0, 51, 52, 5, 17, 0, 0, 52, 53, 5, 48, 0, 0, 53, 129, 5, 43, 0,
		0, 54, 55, 5, 18, 0, 0, 55, 56, 5, 48, 0, 0, 56, 129, 5, 32, 0, 0, 57,
		58, 5, 18, 0, 0, 58, 59, 5, 48, 0, 0, 59, 129, 5, 33, 0, 0, 60, 61, 5,
		18, 0, 0, 61, 62, 5, 48, 0, 0, 62, 129, 5, 34, 0, 0, 63, 64, 5, 20, 0,
		0, 64, 65, 5, 48, 0, 0, 65, 129, 5, 37, 0, 0, 66, 67, 5, 20, 0, 0, 67,
		68, 5, 48, 0, 0, 68, 129, 5, 35, 0, 0, 69, 70, 5, 20, 0, 0, 70, 71, 5,
		48, 0, 0, 71, 129, 5, 36, 0, 0, 72, 73, 5, 21, 0, 0, 73, 74, 5, 48, 0,
		0, 74, 129, 5, 47, 0, 0, 75, 76, 5, 21, 0, 0, 76, 77, 5, 48, 0, 0, 77,
		129, 5, 43, 0, 0, 78, 79, 5, 24, 0, 0, 79, 80, 5, 48, 0, 0, 80, 129, 5,
		47, 0, 0, 81, 82, 5, 24, 0, 0, 82, 83, 5, 48, 0, 0, 83, 129, 5, 43, 0,
		0, 84, 85, 5, 28, 0, 0, 85, 86, 5, 48, 0, 0, 86, 129, 5, 47, 0, 0, 87,
		88, 5, 28, 0, 0, 88, 89, 5, 48, 0, 0, 89, 129, 5, 43, 0, 0, 90, 91, 5,
		25, 0, 0, 91, 92, 5, 48, 0, 0, 92, 129, 5, 47, 0, 0, 93, 94, 5, 26, 0,
		0, 94, 95, 5, 48, 0, 0, 95, 129, 5, 47, 0, 0, 96, 97, 5, 22, 0, 0, 97,
		98, 5, 48, 0, 0, 98, 129, 5, 38, 0, 0, 99, 100, 5, 22, 0, 0, 100, 101,
		5, 48, 0, 0, 101, 129, 5, 40, 0, 0, 102, 103, 5, 22, 0, 0, 103, 104, 5,
		48, 0, 0, 104, 129, 5, 39, 0, 0, 105, 106, 5, 22, 0, 0, 106, 107, 5, 48,
		0, 0, 107, 129, 5, 41, 0, 0, 108, 109, 5, 22, 0, 0, 109, 110, 5, 48, 0,
		0, 110, 129, 5, 42, 0, 0, 111, 112, 5, 23, 0, 0, 112, 113, 5, 48, 0, 0,
		113, 129, 5, 45, 0, 0, 114, 115, 5, 27, 0, 0, 115, 116, 5, 48, 0, 0, 116,
		129, 5, 43, 0, 0, 117, 118, 5, 27, 0, 0, 118, 119, 5, 48, 0, 0, 119, 129,
		5, 46, 0, 0, 120, 121, 5, 29, 0, 0, 121, 122, 5, 48, 0, 0, 122, 129, 5,
		43, 0, 0, 123, 124, 5, 29, 0, 0, 124, 125, 5, 48, 0, 0, 125, 129, 5, 46,
		0, 0, 126, 129, 5, 31, 0, 0, 127, 129, 5, 30, 0, 0, 128, 45, 1, 0, 0, 0,
		128, 48, 1, 0, 0, 0, 128, 51, 1, 0, 0, 0, 128, 54, 1, 0, 0, 0, 128, 57,
		1, 0, 0, 0, 128, 60, 1, 0, 0, 0, 128, 63, 1, 0, 0, 0, 128, 66, 1, 0, 0,
		0, 128, 69, 1, 0, 0, 0, 128, 72, 1, 0, 0, 0, 128, 75, 1, 0, 0, 0, 128,
		78, 1, 0, 0, 0, 128, 81, 1, 0, 0, 0, 128, 84, 1, 0, 0, 0, 128, 87, 1, 0,
		0, 0, 128, 90, 1, 0, 0, 0, 128, 93, 1, 0, 0, 0, 128, 96, 1, 0, 0, 0, 128,
		99, 1, 0, 0, 0, 128, 102, 1, 0, 0, 0, 128, 105, 1, 0, 0, 0, 128, 108, 1,
		0, 0, 0, 128, 111, 1, 0, 0, 0, 128, 114, 1, 0, 0, 0, 128, 117, 1, 0, 0,
		0, 128, 120, 1, 0, 0, 0, 128, 123, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		127, 1, 0, 0, 0, 129, 9, 1, 0, 0, 0, 4, 17, 33, 42, 128,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CmdParserInit initializes any static state used to implement CmdParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCmdParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CmdParserInit() {
	staticData := &cmdParserStaticData
	staticData.once.Do(cmdParserInit)
}

// NewCmdParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCmdParser(input antlr.TokenStream) *CmdParser {
	CmdParserInit()
	this := new(CmdParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cmdParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cmd.g4"

	return this
}

// CmdParser tokens.
const (
	CmdParserEOF               = antlr.TokenEOF
	CmdParserTOK_MKDISK        = 1
	CmdParserTOK_RMDISK        = 2
	CmdParserTOK_FDISK         = 3
	CmdParserTOK_MOUNT         = 4
	CmdParserTOK_MKFS          = 5
	CmdParserTOK_LOGIN         = 6
	CmdParserTOK_LOGOUT        = 7
	CmdParserTOK_MKGRP         = 8
	CmdParserTOK_RMGRP         = 9
	CmdParserTOK_MKUSR         = 10
	CmdParserTOK_RMUSR         = 11
	CmdParserTOK_MKFILE        = 12
	CmdParserTOK_MKDIR         = 13
	CmdParserTOK_EXEC          = 14
	CmdParserTOK_PAUSE         = 15
	CmdParserTOK_REP           = 16
	CmdParserTOK_PATH          = 17
	CmdParserTOK_FIT           = 18
	CmdParserTOK_SIZE          = 19
	CmdParserTOK_UNIT          = 20
	CmdParserTOK_NAME          = 21
	CmdParserTOK_TYPE          = 22
	CmdParserTOK_ID            = 23
	CmdParserTOK_USUARIO       = 24
	CmdParserTOK_PASSWORD      = 25
	CmdParserTOK_PWD           = 26
	CmdParserTOK_CONT          = 27
	CmdParserTOK_GRP           = 28
	CmdParserTOK_RUTA          = 29
	CmdParserTOK_R             = 30
	CmdParserTOK_P             = 31
	CmdParserTOK_FIRST         = 32
	CmdParserTOK_WORST         = 33
	CmdParserTOK_BEST          = 34
	CmdParserTOK_KB            = 35
	CmdParserTOK_MB            = 36
	CmdParserTOK_BYTES         = 37
	CmdParserTOK_PRIMARIA      = 38
	CmdParserTOK_EXTENDIDA     = 39
	CmdParserTOK_LOGICA        = 40
	CmdParserTOK_FAST          = 41
	CmdParserTOK_FULL          = 42
	CmdParserTOK_CADENA        = 43
	CmdParserTOK_NUMERO        = 44
	CmdParserTOK_IDENTIFICADOR = 45
	CmdParserTOK_CAMINO        = 46
	CmdParserTOK_PALABRA       = 47
	CmdParserTOK_IGUAL         = 48
	CmdParserCOMENTARIOS       = 49
	CmdParserWHITESPACE        = 50
)

// CmdParser rules.
const (
	CmdParserRULE_start          = 0
	CmdParserRULE_comando        = 1
	CmdParserRULE_comando_estado = 2
	CmdParserRULE_param_list     = 3
	CmdParserRULE_param          = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Comando() IComandoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *CmdParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CmdParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Comando()
	}

	return localctx
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoContext() *ComandoContext {
	var p = new(ComandoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_comando
	return p
}

func (*ComandoContext) IsComandoContext() {}

func NewComandoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoContext {
	var p = new(ComandoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_comando

	return p
}

func (s *ComandoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoContext) CopyFrom(ctx *ComandoContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PauseContext struct {
	*ComandoContext
}

func NewPauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PauseContext {
	var p = new(PauseContext)

	p.ComandoContext = NewEmptyComandoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComandoContext))

	return p
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) TOK_PAUSE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PAUSE, 0)
}

func (s *PauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPause(s)
	}
}

func (s *PauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPause(s)
	}
}

type LogoutContext struct {
	*ComandoContext
}

func NewLogoutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogoutContext {
	var p = new(LogoutContext)

	p.ComandoContext = NewEmptyComandoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComandoContext))

	return p
}

func (s *LogoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogoutContext) TOK_LOGOUT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGOUT, 0)
}

func (s *LogoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterLogout(s)
	}
}

func (s *LogoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitLogout(s)
	}
}

type AddCommandContext struct {
	*ComandoContext
}

func NewAddCommandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddCommandContext {
	var p = new(AddCommandContext)

	p.ComandoContext = NewEmptyComandoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComandoContext))

	return p
}

func (s *AddCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddCommandContext) Comando_estado() IComando_estadoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComando_estadoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComando_estadoContext)
}

func (s *AddCommandContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *AddCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterAddCommand(s)
	}
}

func (s *AddCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitAddCommand(s)
	}
}

func (p *CmdParser) Comando() (localctx IComandoContext) {
	this := p
	_ = this

	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CmdParserRULE_comando)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CmdParserTOK_MKDISK, CmdParserTOK_RMDISK, CmdParserTOK_FDISK, CmdParserTOK_MOUNT, CmdParserTOK_MKFS, CmdParserTOK_LOGIN, CmdParserTOK_MKGRP, CmdParserTOK_RMGRP, CmdParserTOK_MKUSR, CmdParserTOK_RMUSR, CmdParserTOK_MKFILE, CmdParserTOK_MKDIR, CmdParserTOK_EXEC, CmdParserTOK_REP:
		localctx = NewAddCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Comando_estado()
		}
		{
			p.SetState(13)
			p.param_list(0)
		}

	case CmdParserTOK_PAUSE:
		localctx = NewPauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Match(CmdParserTOK_PAUSE)
		}

	case CmdParserTOK_LOGOUT:
		localctx = NewLogoutContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(16)
			p.Match(CmdParserTOK_LOGOUT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComando_estadoContext is an interface to support dynamic dispatch.
type IComando_estadoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComando_estadoContext differentiates from other interfaces.
	IsComando_estadoContext()
}

type Comando_estadoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComando_estadoContext() *Comando_estadoContext {
	var p = new(Comando_estadoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_comando_estado
	return p
}

func (*Comando_estadoContext) IsComando_estadoContext() {}

func NewComando_estadoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comando_estadoContext {
	var p = new(Comando_estadoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_comando_estado

	return p
}

func (s *Comando_estadoContext) GetParser() antlr.Parser { return s.parser }

func (s *Comando_estadoContext) CopyFrom(ctx *Comando_estadoContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Comando_estadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comando_estadoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MkdirContext struct {
	*Comando_estadoContext
}

func NewMkdirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkdirContext {
	var p = new(MkdirContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirContext) TOK_MKDIR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKDIR, 0)
}

func (s *MkdirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkdir(s)
	}
}

func (s *MkdirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkdir(s)
	}
}

type MkfileContext struct {
	*Comando_estadoContext
}

func NewMkfileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkfileContext {
	var p = new(MkfileContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileContext) TOK_MKFILE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKFILE, 0)
}

func (s *MkfileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkfile(s)
	}
}

func (s *MkfileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkfile(s)
	}
}

type MkdiskContext struct {
	*Comando_estadoContext
}

func NewMkdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkdiskContext {
	var p = new(MkdiskContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) TOK_MKDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKDISK, 0)
}

func (s *MkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkdisk(s)
	}
}

func (s *MkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkdisk(s)
	}
}

type RmdiskContext struct {
	*Comando_estadoContext
}

func NewRmdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RmdiskContext {
	var p = new(RmdiskContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *RmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmdiskContext) TOK_RMDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMDISK, 0)
}

func (s *RmdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRmdisk(s)
	}
}

func (s *RmdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRmdisk(s)
	}
}

type MountContext struct {
	*Comando_estadoContext
}

func NewMountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MountContext {
	var p = new(MountContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountContext) TOK_MOUNT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MOUNT, 0)
}

func (s *MountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMount(s)
	}
}

func (s *MountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMount(s)
	}
}

type MkgrpContext struct {
	*Comando_estadoContext
}

func NewMkgrpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkgrpContext {
	var p = new(MkgrpContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpContext) TOK_MKGRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKGRP, 0)
}

func (s *MkgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkgrp(s)
	}
}

func (s *MkgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkgrp(s)
	}
}

type FdiskContext struct {
	*Comando_estadoContext
}

func NewFdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FdiskContext {
	var p = new(FdiskContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *FdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskContext) TOK_FDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FDISK, 0)
}

func (s *FdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterFdisk(s)
	}
}

func (s *FdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitFdisk(s)
	}
}

type LoginContext struct {
	*Comando_estadoContext
}

func NewLoginContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoginContext {
	var p = new(LoginContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *LoginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginContext) TOK_LOGIN() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGIN, 0)
}

func (s *LoginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterLogin(s)
	}
}

func (s *LoginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitLogin(s)
	}
}

type MkfsContext struct {
	*Comando_estadoContext
}

func NewMkfsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkfsContext {
	var p = new(MkfsContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsContext) TOK_MKFS() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKFS, 0)
}

func (s *MkfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkfs(s)
	}
}

func (s *MkfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkfs(s)
	}
}

type RmusrContext struct {
	*Comando_estadoContext
}

func NewRmusrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RmusrContext {
	var p = new(RmusrContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *RmusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmusrContext) TOK_RMUSR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMUSR, 0)
}

func (s *RmusrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRmusr(s)
	}
}

func (s *RmusrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRmusr(s)
	}
}

type ExecContext struct {
	*Comando_estadoContext
}

func NewExecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExecContext {
	var p = new(ExecContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *ExecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecContext) TOK_EXEC() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_EXEC, 0)
}

func (s *ExecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterExec(s)
	}
}

func (s *ExecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitExec(s)
	}
}

type MkusrContext struct {
	*Comando_estadoContext
}

func NewMkusrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MkusrContext {
	var p = new(MkusrContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *MkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrContext) TOK_MKUSR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKUSR, 0)
}

func (s *MkusrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterMkusr(s)
	}
}

func (s *MkusrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitMkusr(s)
	}
}

type RmgrpContext struct {
	*Comando_estadoContext
}

func NewRmgrpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RmgrpContext {
	var p = new(RmgrpContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *RmgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmgrpContext) TOK_RMGRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMGRP, 0)
}

func (s *RmgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRmgrp(s)
	}
}

func (s *RmgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRmgrp(s)
	}
}

type RepContext struct {
	*Comando_estadoContext
}

func NewRepContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepContext {
	var p = new(RepContext)

	p.Comando_estadoContext = NewEmptyComando_estadoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Comando_estadoContext))

	return p
}

func (s *RepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepContext) TOK_REP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_REP, 0)
}

func (s *RepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRep(s)
	}
}

func (s *RepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRep(s)
	}
}

func (p *CmdParser) Comando_estado() (localctx IComando_estadoContext) {
	this := p
	_ = this

	localctx = NewComando_estadoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CmdParserRULE_comando_estado)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CmdParserTOK_MKDISK:
		localctx = NewMkdiskContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Match(CmdParserTOK_MKDISK)
		}

	case CmdParserTOK_RMDISK:
		localctx = NewRmdiskContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Match(CmdParserTOK_RMDISK)
		}

	case CmdParserTOK_FDISK:
		localctx = NewFdiskContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(21)
			p.Match(CmdParserTOK_FDISK)
		}

	case CmdParserTOK_MOUNT:
		localctx = NewMountContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(22)
			p.Match(CmdParserTOK_MOUNT)
		}

	case CmdParserTOK_MKFS:
		localctx = NewMkfsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(23)
			p.Match(CmdParserTOK_MKFS)
		}

	case CmdParserTOK_LOGIN:
		localctx = NewLoginContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(24)
			p.Match(CmdParserTOK_LOGIN)
		}

	case CmdParserTOK_MKGRP:
		localctx = NewMkgrpContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(25)
			p.Match(CmdParserTOK_MKGRP)
		}

	case CmdParserTOK_RMGRP:
		localctx = NewRmgrpContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(26)
			p.Match(CmdParserTOK_RMGRP)
		}

	case CmdParserTOK_MKUSR:
		localctx = NewMkusrContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(27)
			p.Match(CmdParserTOK_MKUSR)
		}

	case CmdParserTOK_RMUSR:
		localctx = NewRmusrContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(28)
			p.Match(CmdParserTOK_RMUSR)
		}

	case CmdParserTOK_MKFILE:
		localctx = NewMkfileContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(29)
			p.Match(CmdParserTOK_MKFILE)
		}

	case CmdParserTOK_MKDIR:
		localctx = NewMkdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(30)
			p.Match(CmdParserTOK_MKDIR)
		}

	case CmdParserTOK_EXEC:
		localctx = NewExecContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(31)
			p.Match(CmdParserTOK_EXEC)
		}

	case CmdParserTOK_REP:
		localctx = NewRepContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(32)
			p.Match(CmdParserTOK_REP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_param_list
	return p
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) Param() IParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Param_listContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterParam_list(s)
	}
}

func (s *Param_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitParam_list(s)
	}
}

func (p *CmdParser) Param_list() (localctx IParam_listContext) {
	return p.param_list(0)
}

func (p *CmdParser) param_list(_p int) (localctx IParam_listContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParam_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, CmdParserRULE_param_list, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Param()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParam_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CmdParserRULE_param_list)
			p.SetState(38)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(39)
				p.Param()
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) CopyFrom(ctx *ParamContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Cont_RContext struct {
	*ParamContext
}

func NewCont_RContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cont_RContext {
	var p = new(Cont_RContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Cont_RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cont_RContext) TOK_CONT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CONT, 0)
}

func (s *Cont_RContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Cont_RContext) TOK_CAMINO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CAMINO, 0)
}

func (s *Cont_RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterCont_R(s)
	}
}

func (s *Cont_RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitCont_R(s)
	}
}

type Cont_SContext struct {
	*ParamContext
}

func NewCont_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cont_SContext {
	var p = new(Cont_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Cont_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cont_SContext) TOK_CONT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CONT, 0)
}

func (s *Cont_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Cont_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Cont_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterCont_S(s)
	}
}

func (s *Cont_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitCont_S(s)
	}
}

type SizeContext struct {
	*ParamContext
}

func NewSizeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SizeContext {
	var p = new(SizeContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) TOK_SIZE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_SIZE, 0)
}

func (s *SizeContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *SizeContext) TOK_NUMERO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_NUMERO, 0)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitSize(s)
	}
}

type Grp_SContext struct {
	*ParamContext
}

func NewGrp_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Grp_SContext {
	var p = new(Grp_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Grp_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Grp_SContext) TOK_GRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_GRP, 0)
}

func (s *Grp_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Grp_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Grp_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterGrp_S(s)
	}
}

func (s *Grp_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitGrp_S(s)
	}
}

type Path_RContext struct {
	*ParamContext
}

func NewPath_RContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Path_RContext {
	var p = new(Path_RContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Path_RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Path_RContext) TOK_PATH() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PATH, 0)
}

func (s *Path_RContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Path_RContext) TOK_CAMINO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CAMINO, 0)
}

func (s *Path_RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPath_R(s)
	}
}

func (s *Path_RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPath_R(s)
	}
}

type Grp_PContext struct {
	*ParamContext
}

func NewGrp_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Grp_PContext {
	var p = new(Grp_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Grp_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Grp_PContext) TOK_GRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_GRP, 0)
}

func (s *Grp_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Grp_PContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *Grp_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterGrp_P(s)
	}
}

func (s *Grp_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitGrp_P(s)
	}
}

type Path_SContext struct {
	*ParamContext
}

func NewPath_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Path_SContext {
	var p = new(Path_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Path_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Path_SContext) TOK_PATH() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PATH, 0)
}

func (s *Path_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Path_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Path_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPath_S(s)
	}
}

func (s *Path_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPath_S(s)
	}
}

type FitWFContext struct {
	*ParamContext
}

func NewFitWFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FitWFContext {
	var p = new(FitWFContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *FitWFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FitWFContext) TOK_FIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIT, 0)
}

func (s *FitWFContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *FitWFContext) TOK_WORST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_WORST, 0)
}

func (s *FitWFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterFitWF(s)
	}
}

func (s *FitWFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitFitWF(s)
	}
}

type Pass_PContext struct {
	*ParamContext
}

func NewPass_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Pass_PContext {
	var p = new(Pass_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Pass_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pass_PContext) TOK_PASSWORD() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PASSWORD, 0)
}

func (s *Pass_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Pass_PContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *Pass_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPass_P(s)
	}
}

func (s *Pass_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPass_P(s)
	}
}

type Type_LContext struct {
	*ParamContext
}

func NewType_LContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Type_LContext {
	var p = new(Type_LContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Type_LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_LContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *Type_LContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Type_LContext) TOK_LOGICA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGICA, 0)
}

func (s *Type_LContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterType_L(s)
	}
}

func (s *Type_LContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitType_L(s)
	}
}

type Type_FastContext struct {
	*ParamContext
}

func NewType_FastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Type_FastContext {
	var p = new(Type_FastContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Type_FastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_FastContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *Type_FastContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Type_FastContext) TOK_FAST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FAST, 0)
}

func (s *Type_FastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterType_Fast(s)
	}
}

func (s *Type_FastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitType_Fast(s)
	}
}

type Type_EContext struct {
	*ParamContext
}

func NewType_EContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Type_EContext {
	var p = new(Type_EContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Type_EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_EContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *Type_EContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Type_EContext) TOK_EXTENDIDA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_EXTENDIDA, 0)
}

func (s *Type_EContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterType_E(s)
	}
}

func (s *Type_EContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitType_E(s)
	}
}

type Name_SContext struct {
	*ParamContext
}

func NewName_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Name_SContext {
	var p = new(Name_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Name_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_SContext) TOK_NAME() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_NAME, 0)
}

func (s *Name_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Name_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Name_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterName_S(s)
	}
}

func (s *Name_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitName_S(s)
	}
}

type Name_PContext struct {
	*ParamContext
}

func NewName_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Name_PContext {
	var p = new(Name_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Name_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_PContext) TOK_NAME() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_NAME, 0)
}

func (s *Name_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Name_PContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *Name_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterName_P(s)
	}
}

func (s *Name_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitName_P(s)
	}
}

type Type_PContext struct {
	*ParamContext
}

func NewType_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Type_PContext {
	var p = new(Type_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Type_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_PContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *Type_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Type_PContext) TOK_PRIMARIA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PRIMARIA, 0)
}

func (s *Type_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterType_P(s)
	}
}

func (s *Type_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitType_P(s)
	}
}

type PpContext struct {
	*ParamContext
}

func NewPpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PpContext {
	var p = new(PpContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *PpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PpContext) TOK_P() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_P, 0)
}

func (s *PpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPp(s)
	}
}

func (s *PpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPp(s)
	}
}

type RrContext struct {
	*ParamContext
}

func NewRrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrContext {
	var p = new(RrContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *RrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrContext) TOK_R() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_R, 0)
}

func (s *RrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRr(s)
	}
}

func (s *RrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRr(s)
	}
}

type Usr_PContext struct {
	*ParamContext
}

func NewUsr_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Usr_PContext {
	var p = new(Usr_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Usr_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Usr_PContext) TOK_USUARIO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_USUARIO, 0)
}

func (s *Usr_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Usr_PContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *Usr_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterUsr_P(s)
	}
}

func (s *Usr_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitUsr_P(s)
	}
}

type Type_FullContext struct {
	*ParamContext
}

func NewType_FullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Type_FullContext {
	var p = new(Type_FullContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Type_FullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_FullContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *Type_FullContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Type_FullContext) TOK_FULL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FULL, 0)
}

func (s *Type_FullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterType_Full(s)
	}
}

func (s *Type_FullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitType_Full(s)
	}
}

type Usr_SContext struct {
	*ParamContext
}

func NewUsr_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Usr_SContext {
	var p = new(Usr_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Usr_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Usr_SContext) TOK_USUARIO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_USUARIO, 0)
}

func (s *Usr_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Usr_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Usr_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterUsr_S(s)
	}
}

func (s *Usr_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitUsr_S(s)
	}
}

type Unit_BContext struct {
	*ParamContext
}

func NewUnit_BContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Unit_BContext {
	var p = new(Unit_BContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Unit_BContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unit_BContext) TOK_UNIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_UNIT, 0)
}

func (s *Unit_BContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Unit_BContext) TOK_BYTES() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_BYTES, 0)
}

func (s *Unit_BContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterUnit_B(s)
	}
}

func (s *Unit_BContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitUnit_B(s)
	}
}

type Pwd_PContext struct {
	*ParamContext
}

func NewPwd_PContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Pwd_PContext {
	var p = new(Pwd_PContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Pwd_PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pwd_PContext) TOK_PWD() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PWD, 0)
}

func (s *Pwd_PContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Pwd_PContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *Pwd_PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterPwd_P(s)
	}
}

func (s *Pwd_PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitPwd_P(s)
	}
}

type Ruta_SContext struct {
	*ParamContext
}

func NewRuta_SContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ruta_SContext {
	var p = new(Ruta_SContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Ruta_SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ruta_SContext) TOK_RUTA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RUTA, 0)
}

func (s *Ruta_SContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Ruta_SContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *Ruta_SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRuta_S(s)
	}
}

func (s *Ruta_SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRuta_S(s)
	}
}

type Ruta_RContext struct {
	*ParamContext
}

func NewRuta_RContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ruta_RContext {
	var p = new(Ruta_RContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Ruta_RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ruta_RContext) TOK_RUTA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RUTA, 0)
}

func (s *Ruta_RContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Ruta_RContext) TOK_CAMINO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CAMINO, 0)
}

func (s *Ruta_RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterRuta_R(s)
	}
}

func (s *Ruta_RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitRuta_R(s)
	}
}

type Unit_MContext struct {
	*ParamContext
}

func NewUnit_MContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Unit_MContext {
	var p = new(Unit_MContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Unit_MContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unit_MContext) TOK_UNIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_UNIT, 0)
}

func (s *Unit_MContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Unit_MContext) TOK_MB() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MB, 0)
}

func (s *Unit_MContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterUnit_M(s)
	}
}

func (s *Unit_MContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitUnit_M(s)
	}
}

type FitFFContext struct {
	*ParamContext
}

func NewFitFFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FitFFContext {
	var p = new(FitFFContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *FitFFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FitFFContext) TOK_FIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIT, 0)
}

func (s *FitFFContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *FitFFContext) TOK_FIRST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIRST, 0)
}

func (s *FitFFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterFitFF(s)
	}
}

func (s *FitFFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitFitFF(s)
	}
}

type Unit_KContext struct {
	*ParamContext
}

func NewUnit_KContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Unit_KContext {
	var p = new(Unit_KContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *Unit_KContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unit_KContext) TOK_UNIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_UNIT, 0)
}

func (s *Unit_KContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *Unit_KContext) TOK_KB() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_KB, 0)
}

func (s *Unit_KContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterUnit_K(s)
	}
}

func (s *Unit_KContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitUnit_K(s)
	}
}

type IdContext struct {
	*ParamContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) TOK_ID() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_ID, 0)
}

func (s *IdContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *IdContext) TOK_IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IDENTIFICADOR, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitId(s)
	}
}

type FitBFContext struct {
	*ParamContext
}

func NewFitBFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FitBFContext {
	var p = new(FitBFContext)

	p.ParamContext = NewEmptyParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParamContext))

	return p
}

func (s *FitBFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FitBFContext) TOK_FIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIT, 0)
}

func (s *FitBFContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *FitBFContext) TOK_BEST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_BEST, 0)
}

func (s *FitBFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterFitBF(s)
	}
}

func (s *FitBFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitFitBF(s)
	}
}

func (p *CmdParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CmdParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(CmdParserTOK_SIZE)
		}
		{
			p.SetState(46)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(47)
			p.Match(CmdParserTOK_NUMERO)
		}

	case 2:
		localctx = NewPath_RContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(CmdParserTOK_PATH)
		}
		{
			p.SetState(49)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(50)
			p.Match(CmdParserTOK_CAMINO)
		}

	case 3:
		localctx = NewPath_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Match(CmdParserTOK_PATH)
		}
		{
			p.SetState(52)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(53)
			p.Match(CmdParserTOK_CADENA)
		}

	case 4:
		localctx = NewFitFFContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(55)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(56)
			p.Match(CmdParserTOK_FIRST)
		}

	case 5:
		localctx = NewFitWFContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(58)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(59)
			p.Match(CmdParserTOK_WORST)
		}

	case 6:
		localctx = NewFitBFContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(61)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(62)
			p.Match(CmdParserTOK_BEST)
		}

	case 7:
		localctx = NewUnit_BContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(63)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(64)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(65)
			p.Match(CmdParserTOK_BYTES)
		}

	case 8:
		localctx = NewUnit_KContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(66)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(67)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(68)
			p.Match(CmdParserTOK_KB)
		}

	case 9:
		localctx = NewUnit_MContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(69)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(70)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(71)
			p.Match(CmdParserTOK_MB)
		}

	case 10:
		localctx = NewName_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(72)
			p.Match(CmdParserTOK_NAME)
		}
		{
			p.SetState(73)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(74)
			p.Match(CmdParserTOK_PALABRA)
		}

	case 11:
		localctx = NewName_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(75)
			p.Match(CmdParserTOK_NAME)
		}
		{
			p.SetState(76)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(77)
			p.Match(CmdParserTOK_CADENA)
		}

	case 12:
		localctx = NewUsr_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(78)
			p.Match(CmdParserTOK_USUARIO)
		}
		{
			p.SetState(79)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(80)
			p.Match(CmdParserTOK_PALABRA)
		}

	case 13:
		localctx = NewUsr_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(81)
			p.Match(CmdParserTOK_USUARIO)
		}
		{
			p.SetState(82)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(83)
			p.Match(CmdParserTOK_CADENA)
		}

	case 14:
		localctx = NewGrp_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(84)
			p.Match(CmdParserTOK_GRP)
		}
		{
			p.SetState(85)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(86)
			p.Match(CmdParserTOK_PALABRA)
		}

	case 15:
		localctx = NewGrp_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(87)
			p.Match(CmdParserTOK_GRP)
		}
		{
			p.SetState(88)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(89)
			p.Match(CmdParserTOK_CADENA)
		}

	case 16:
		localctx = NewPass_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(90)
			p.Match(CmdParserTOK_PASSWORD)
		}
		{
			p.SetState(91)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(92)
			p.Match(CmdParserTOK_PALABRA)
		}

	case 17:
		localctx = NewPwd_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(93)
			p.Match(CmdParserTOK_PWD)
		}
		{
			p.SetState(94)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(95)
			p.Match(CmdParserTOK_PALABRA)
		}

	case 18:
		localctx = NewType_PContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(96)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(97)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(98)
			p.Match(CmdParserTOK_PRIMARIA)
		}

	case 19:
		localctx = NewType_LContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(99)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(100)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(101)
			p.Match(CmdParserTOK_LOGICA)
		}

	case 20:
		localctx = NewType_EContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(102)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(103)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(104)
			p.Match(CmdParserTOK_EXTENDIDA)
		}

	case 21:
		localctx = NewType_FastContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(105)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(106)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(107)
			p.Match(CmdParserTOK_FAST)
		}

	case 22:
		localctx = NewType_FullContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(108)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(109)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(110)
			p.Match(CmdParserTOK_FULL)
		}

	case 23:
		localctx = NewIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(111)
			p.Match(CmdParserTOK_ID)
		}
		{
			p.SetState(112)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(113)
			p.Match(CmdParserTOK_IDENTIFICADOR)
		}

	case 24:
		localctx = NewCont_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(114)
			p.Match(CmdParserTOK_CONT)
		}
		{
			p.SetState(115)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(116)
			p.Match(CmdParserTOK_CADENA)
		}

	case 25:
		localctx = NewCont_RContext(p, localctx)
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(117)
			p.Match(CmdParserTOK_CONT)
		}
		{
			p.SetState(118)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(119)
			p.Match(CmdParserTOK_CAMINO)
		}

	case 26:
		localctx = NewRuta_SContext(p, localctx)
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(120)
			p.Match(CmdParserTOK_RUTA)
		}
		{
			p.SetState(121)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(122)
			p.Match(CmdParserTOK_CADENA)
		}

	case 27:
		localctx = NewRuta_RContext(p, localctx)
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(123)
			p.Match(CmdParserTOK_RUTA)
		}
		{
			p.SetState(124)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(125)
			p.Match(CmdParserTOK_CAMINO)
		}

	case 28:
		localctx = NewPpContext(p, localctx)
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(126)
			p.Match(CmdParserTOK_P)
		}

	case 29:
		localctx = NewRrContext(p, localctx)
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(127)
			p.Match(CmdParserTOK_R)
		}

	}

	return localctx
}

func (p *CmdParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *Param_listContext = nil
		if localctx != nil {
			t = localctx.(*Param_listContext)
		}
		return p.Param_list_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CmdParser) Param_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "../ast"
import "../ast/parametros"
import arrayList "github.com/colegno/arraylist"
import "strings"

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
		"", "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT",
		"TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR",
		"TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP",
		"TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE",
		"TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP",
		"TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB", "TOK_BYTES", "TOK_PRIMARIA",
		"TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST", "TOK_WORST", "TOK_BEST",
		"TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR",
		"TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "comandList", "comando", "param_list", "param",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 216, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 87, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 97, 8, 3, 10, 3, 12, 3, 100, 9, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 214, 8, 4, 1, 4, 0, 2, 2, 6, 5, 0, 2, 4, 6,
		8, 0, 0, 255, 0, 10, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 86, 1, 0, 0, 0,
		6, 88, 1, 0, 0, 0, 8, 213, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 12, 6, 0,
		-1, 0, 12, 1, 1, 0, 0, 0, 13, 14, 6, 1, -1, 0, 14, 15, 3, 4, 2, 0, 15,
		16, 6, 1, -1, 0, 16, 23, 1, 0, 0, 0, 17, 18, 10, 2, 0, 0, 18, 19, 3, 4,
		2, 0, 19, 20, 6, 1, -1, 0, 20, 22, 1, 0, 0, 0, 21, 17, 1, 0, 0, 0, 22,
		25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 3, 1, 0, 0,
		0, 25, 23, 1, 0, 0, 0, 26, 27, 5, 2, 0, 0, 27, 28, 3, 6, 3, 0, 28, 29,
		6, 2, -1, 0, 29, 87, 1, 0, 0, 0, 30, 31, 5, 3, 0, 0, 31, 32, 3, 6, 3, 0,
		32, 33, 6, 2, -1, 0, 33, 87, 1, 0, 0, 0, 34, 35, 5, 4, 0, 0, 35, 36, 3,
		6, 3, 0, 36, 37, 6, 2, -1, 0, 37, 87, 1, 0, 0, 0, 38, 39, 5, 5, 0, 0, 39,
		40, 3, 6, 3, 0, 40, 41, 6, 2, -1, 0, 41, 87, 1, 0, 0, 0, 42, 43, 5, 6,
		0, 0, 43, 44, 3, 6, 3, 0, 44, 45, 6, 2, -1, 0, 45, 87, 1, 0, 0, 0, 46,
		47, 5, 7, 0, 0, 47, 48, 3, 6, 3, 0, 48, 49, 6, 2, -1, 0, 49, 87, 1, 0,
		0, 0, 50, 51, 5, 9, 0, 0, 51, 52, 3, 6, 3, 0, 52, 53, 6, 2, -1, 0, 53,
		87, 1, 0, 0, 0, 54, 55, 5, 10, 0, 0, 55, 56, 3, 6, 3, 0, 56, 57, 6, 2,
		-1, 0, 57, 87, 1, 0, 0, 0, 58, 59, 5, 11, 0, 0, 59, 60, 3, 6, 3, 0, 60,
		61, 6, 2, -1, 0, 61, 87, 1, 0, 0, 0, 62, 63, 5, 12, 0, 0, 63, 64, 3, 6,
		3, 0, 64, 65, 6, 2, -1, 0, 65, 87, 1, 0, 0, 0, 66, 67, 5, 13, 0, 0, 67,
		68, 3, 6, 3, 0, 68, 69, 6, 2, -1, 0, 69, 87, 1, 0, 0, 0, 70, 71, 5, 14,
		0, 0, 71, 72, 3, 6, 3, 0, 72, 73, 6, 2, -1, 0, 73, 87, 1, 0, 0, 0, 74,
		75, 5, 15, 0, 0, 75, 76, 3, 6, 3, 0, 76, 77, 6, 2, -1, 0, 77, 87, 1, 0,
		0, 0, 78, 79, 5, 17, 0, 0, 79, 80, 3, 6, 3, 0, 80, 81, 6, 2, -1, 0, 81,
		87, 1, 0, 0, 0, 82, 83, 5, 16, 0, 0, 83, 87, 6, 2, -1, 0, 84, 85, 5, 8,
		0, 0, 85, 87, 6, 2, -1, 0, 86, 26, 1, 0, 0, 0, 86, 30, 1, 0, 0, 0, 86,
		34, 1, 0, 0, 0, 86, 38, 1, 0, 0, 0, 86, 42, 1, 0, 0, 0, 86, 46, 1, 0, 0,
		0, 86, 50, 1, 0, 0, 0, 86, 54, 1, 0, 0, 0, 86, 58, 1, 0, 0, 0, 86, 62,
		1, 0, 0, 0, 86, 66, 1, 0, 0, 0, 86, 70, 1, 0, 0, 0, 86, 74, 1, 0, 0, 0,
		86, 78, 1, 0, 0, 0, 86, 82, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 5, 1, 0,
		0, 0, 88, 89, 6, 3, -1, 0, 89, 90, 3, 8, 4, 0, 90, 91, 6, 3, -1, 0, 91,
		98, 1, 0, 0, 0, 92, 93, 10, 2, 0, 0, 93, 94, 3, 8, 4, 0, 94, 95, 6, 3,
		-1, 0, 95, 97, 1, 0, 0, 0, 96, 92, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 7, 1, 0, 0, 0, 100, 98, 1, 0, 0,
		0, 101, 102, 5, 20, 0, 0, 102, 103, 5, 1, 0, 0, 103, 104, 5, 45, 0, 0,
		104, 214, 6, 4, -1, 0, 105, 106, 5, 18, 0, 0, 106, 107, 5, 1, 0, 0, 107,
		108, 5, 47, 0, 0, 108, 214, 6, 4, -1, 0, 109, 110, 5, 18, 0, 0, 110, 111,
		5, 1, 0, 0, 111, 112, 5, 44, 0, 0, 112, 214, 6, 4, -1, 0, 113, 114, 5,
		19, 0, 0, 114, 115, 5, 1, 0, 0, 115, 116, 5, 39, 0, 0, 116, 214, 6, 4,
		-1, 0, 117, 118, 5, 19, 0, 0, 118, 119, 5, 1, 0, 0, 119, 120, 5, 40, 0,
		0, 120, 214, 6, 4, -1, 0, 121, 122, 5, 19, 0, 0, 122, 123, 5, 1, 0, 0,
		123, 124, 5, 41, 0, 0, 124, 214, 6, 4, -1, 0, 125, 126, 5, 21, 0, 0, 126,
		127, 5, 1, 0, 0, 127, 128, 5, 35, 0, 0, 128, 214, 6, 4, -1, 0, 129, 130,
		5, 21, 0, 0, 130, 131, 5, 1, 0, 0, 131, 132, 5, 33, 0, 0, 132, 214, 6,
		4, -1, 0, 133, 134, 5, 21, 0, 0, 134, 135, 5, 1, 0, 0, 135, 136, 5, 34,
		0, 0, 136, 214, 6, 4, -1, 0, 137, 138, 5, 22, 0, 0, 138, 139, 5, 1, 0,
		0, 139, 140, 5, 48, 0, 0, 140, 214, 6, 4, -1, 0, 141, 142, 5, 22, 0, 0,
		142, 143, 5, 1, 0, 0, 143, 144, 5, 44, 0, 0, 144, 214, 6, 4, -1, 0, 145,
		146, 5, 25, 0, 0, 146, 147, 5, 1, 0, 0, 147, 148, 5, 48, 0, 0, 148, 214,
		6, 4, -1, 0, 149, 150, 5, 25, 0, 0, 150, 151, 5, 1, 0, 0, 151, 152, 5,
		44, 0, 0, 152, 214, 6, 4, -1, 0, 153, 154, 5, 29, 0, 0, 154, 155, 5, 1,
		0, 0, 155, 156, 5, 48, 0, 0, 156, 214, 6, 4, -1, 0, 157, 158, 5, 29, 0,
		0, 158, 159, 5, 1, 0, 0, 159, 160, 5, 44, 0, 0, 160, 214, 6, 4, -1, 0,
		161, 162, 5, 26, 0, 0, 162, 163, 5, 1, 0, 0, 163, 164, 5, 48, 0, 0, 164,
		214, 6, 4, -1, 0, 165, 166, 5, 27, 0, 0, 166, 167, 5, 1, 0, 0, 167, 168,
		5, 48, 0, 0, 168, 214, 6, 4, -1, 0, 169, 170, 5, 23, 0, 0, 170, 171, 5,
		1, 0, 0, 171, 172, 5, 36, 0, 0, 172, 214, 6, 4, -1, 0, 173, 174, 5, 23,
		0, 0, 174, 175, 5, 1, 0, 0, 175, 176, 5, 38, 0, 0, 176, 214, 6, 4, -1,
		0, 177, 178, 5, 23, 0, 0, 178, 179, 5, 1, 0, 0, 179, 180, 5, 37, 0, 0,
		180, 214, 6, 4, -1, 0, 181, 182, 5, 23, 0, 0, 182, 183, 5, 1, 0, 0, 183,
		184, 5, 42, 0, 0, 184, 214, 6, 4, -1, 0, 185, 186, 5, 23, 0, 0, 186, 187,
		5, 1, 0, 0, 187, 188, 5, 43, 0, 0, 188, 214, 6, 4, -1, 0, 189, 190, 5,
		24, 0, 0, 190, 191, 5, 1, 0, 0, 191, 192, 5, 46, 0, 0, 192, 214, 6, 4,
		-1, 0, 193, 194, 5, 28, 0, 0, 194, 195, 5, 1, 0, 0, 195, 196, 5, 44, 0,
		0, 196, 214, 6, 4, -1, 0, 197, 198, 5, 28, 0, 0, 198, 199, 5, 1, 0, 0,
		199, 200, 5, 47, 0, 0, 200, 214, 6, 4, -1, 0, 201, 202, 5, 30, 0, 0, 202,
		203, 5, 1, 0, 0, 203, 204, 5, 44, 0, 0, 204, 214, 6, 4, -1, 0, 205, 206,
		5, 30, 0, 0, 206, 207, 5, 1, 0, 0, 207, 208, 5, 47, 0, 0, 208, 214, 6,
		4, -1, 0, 209, 210, 5, 32, 0, 0, 210, 214, 6, 4, -1, 0, 211, 212, 5, 31,
		0, 0, 212, 214, 6, 4, -1, 0, 213, 101, 1, 0, 0, 0, 213, 105, 1, 0, 0, 0,
		213, 109, 1, 0, 0, 0, 213, 113, 1, 0, 0, 0, 213, 117, 1, 0, 0, 0, 213,
		121, 1, 0, 0, 0, 213, 125, 1, 0, 0, 0, 213, 129, 1, 0, 0, 0, 213, 133,
		1, 0, 0, 0, 213, 137, 1, 0, 0, 0, 213, 141, 1, 0, 0, 0, 213, 145, 1, 0,
		0, 0, 213, 149, 1, 0, 0, 0, 213, 153, 1, 0, 0, 0, 213, 157, 1, 0, 0, 0,
		213, 161, 1, 0, 0, 0, 213, 165, 1, 0, 0, 0, 213, 169, 1, 0, 0, 0, 213,
		173, 1, 0, 0, 0, 213, 177, 1, 0, 0, 0, 213, 181, 1, 0, 0, 0, 213, 185,
		1, 0, 0, 0, 213, 189, 1, 0, 0, 0, 213, 193, 1, 0, 0, 0, 213, 197, 1, 0,
		0, 0, 213, 201, 1, 0, 0, 0, 213, 205, 1, 0, 0, 0, 213, 209, 1, 0, 0, 0,
		213, 211, 1, 0, 0, 0, 214, 9, 1, 0, 0, 0, 4, 23, 86, 98, 213,
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
	CmdParserTOK_IGUAL         = 1
	CmdParserTOK_MKDISK        = 2
	CmdParserTOK_RMDISK        = 3
	CmdParserTOK_FDISK         = 4
	CmdParserTOK_MOUNT         = 5
	CmdParserTOK_MKFS          = 6
	CmdParserTOK_LOGIN         = 7
	CmdParserTOK_LOGOUT        = 8
	CmdParserTOK_MKGRP         = 9
	CmdParserTOK_RMGRP         = 10
	CmdParserTOK_MKUSR         = 11
	CmdParserTOK_RMUSR         = 12
	CmdParserTOK_MKFILE        = 13
	CmdParserTOK_MKDIR         = 14
	CmdParserTOK_EXEC          = 15
	CmdParserTOK_PAUSE         = 16
	CmdParserTOK_REP           = 17
	CmdParserTOK_PATH          = 18
	CmdParserTOK_FIT           = 19
	CmdParserTOK_SIZE          = 20
	CmdParserTOK_UNIT          = 21
	CmdParserTOK_NAME          = 22
	CmdParserTOK_TYPE          = 23
	CmdParserTOK_ID            = 24
	CmdParserTOK_USUARIO       = 25
	CmdParserTOK_PASSWORD      = 26
	CmdParserTOK_PWD           = 27
	CmdParserTOK_CONT          = 28
	CmdParserTOK_GRP           = 29
	CmdParserTOK_RUTA          = 30
	CmdParserTOK_R             = 31
	CmdParserTOK_P             = 32
	CmdParserTOK_KB            = 33
	CmdParserTOK_MB            = 34
	CmdParserTOK_BYTES         = 35
	CmdParserTOK_PRIMARIA      = 36
	CmdParserTOK_EXTENDIDA     = 37
	CmdParserTOK_LOGICA        = 38
	CmdParserTOK_FIRST         = 39
	CmdParserTOK_WORST         = 40
	CmdParserTOK_BEST          = 41
	CmdParserTOK_FAST          = 42
	CmdParserTOK_FULL          = 43
	CmdParserTOK_CADENA        = 44
	CmdParserTOK_NUMERO        = 45
	CmdParserTOK_IDENTIFICADOR = 46
	CmdParserTOK_CAMINO        = 47
	CmdParserTOK_PALABRA       = 48
	CmdParserCOMENTARIOS       = 49
	CmdParserWHITESPACE        = 50
)

// CmdParser rules.
const (
	CmdParserRULE_start      = 0
	CmdParserRULE_comandList = 1
	CmdParserRULE_comando    = 2
	CmdParserRULE_param_list = 3
	CmdParserRULE_param      = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_comandList returns the _comandList rule contexts.
	Get_comandList() IComandListContext

	// Set_comandList sets the _comandList rule contexts.
	Set_comandList(IComandListContext)

	// GetAst returns the ast attribute.
	GetAst() ast.Ast

	// SetAst sets the ast attribute.
	SetAst(ast.Ast)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ast         ast.Ast
	_comandList IComandListContext
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

func (s *StartContext) Get_comandList() IComandListContext { return s._comandList }

func (s *StartContext) Set_comandList(v IComandListContext) { s._comandList = v }

func (s *StartContext) GetAst() ast.Ast { return s.ast }

func (s *StartContext) SetAst(v ast.Ast) { s.ast = v }

func (s *StartContext) ComandList() IComandListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandListContext)
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

		var _x = p.comandList(0)

		localctx.(*StartContext)._comandList = _x
	}
	localctx.(*StartContext).ast = ast.NewAst(localctx.(*StartContext).Get_comandList().GetLista())

	return localctx
}

// IComandListContext is an interface to support dynamic dispatch.
type IComandListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAUXCLIST returns the AUXCLIST rule contexts.
	GetAUXCLIST() IComandListContext

	// Get_comando returns the _comando rule contexts.
	Get_comando() IComandoContext

	// SetAUXCLIST sets the AUXCLIST rule contexts.
	SetAUXCLIST(IComandListContext)

	// Set_comando sets the _comando rule contexts.
	Set_comando(IComandoContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsComandListContext differentiates from other interfaces.
	IsComandListContext()
}

type ComandListContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	AUXCLIST IComandListContext
	_comando IComandoContext
}

func NewEmptyComandListContext() *ComandListContext {
	var p = new(ComandListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CmdParserRULE_comandList
	return p
}

func (*ComandListContext) IsComandListContext() {}

func NewComandListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandListContext {
	var p = new(ComandListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CmdParserRULE_comandList

	return p
}

func (s *ComandListContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandListContext) GetAUXCLIST() IComandListContext { return s.AUXCLIST }

func (s *ComandListContext) Get_comando() IComandoContext { return s._comando }

func (s *ComandListContext) SetAUXCLIST(v IComandListContext) { s.AUXCLIST = v }

func (s *ComandListContext) Set_comando(v IComandoContext) { s._comando = v }

func (s *ComandListContext) GetLista() *arrayList.List { return s.lista }

func (s *ComandListContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ComandListContext) Comando() IComandoContext {
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

func (s *ComandListContext) ComandList() IComandListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandListContext)
}

func (s *ComandListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterComandList(s)
	}
}

func (s *ComandListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitComandList(s)
	}
}

func (p *CmdParser) ComandList() (localctx IComandListContext) {
	return p.comandList(0)
}

func (p *CmdParser) comandList(_p int) (localctx IComandListContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewComandListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComandListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CmdParserRULE_comandList, _p)
	localctx.(*ComandListContext).lista = arrayList.New()

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
		p.SetState(14)

		var _x = p.Comando()

		localctx.(*ComandListContext)._comando = _x
	}

	localctx.(*ComandListContext).lista.Add(localctx.(*ComandListContext).Get_comando().GetCom())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComandListContext(p, _parentctx, _parentState)
			localctx.(*ComandListContext).AUXCLIST = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CmdParserRULE_comandList)
			p.SetState(17)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(18)

				var _x = p.Comando()

				localctx.(*ComandListContext)._comando = _x
			}

			localctx.(*ComandListContext).GetAUXCLIST().GetLista().Add(localctx.(*ComandListContext).Get_comando().GetCom())
			localctx.(*ComandListContext).lista = localctx.(*ComandListContext).GetAUXCLIST().GetLista()

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_param_list returns the _param_list rule contexts.
	Get_param_list() IParam_listContext

	// Set_param_list sets the _param_list rule contexts.
	Set_param_list(IParam_listContext)

	// GetCom returns the com attribute.
	GetCom() ast.Command

	// SetCom sets the com attribute.
	SetCom(ast.Command)

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	com         ast.Command
	_param_list IParam_listContext
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

func (s *ComandoContext) Get_param_list() IParam_listContext { return s._param_list }

func (s *ComandoContext) Set_param_list(v IParam_listContext) { s._param_list = v }

func (s *ComandoContext) GetCom() ast.Command { return s.com }

func (s *ComandoContext) SetCom(v ast.Command) { s.com = v }

func (s *ComandoContext) TOK_MKDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKDISK, 0)
}

func (s *ComandoContext) Param_list() IParam_listContext {
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

func (s *ComandoContext) TOK_RMDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMDISK, 0)
}

func (s *ComandoContext) TOK_FDISK() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FDISK, 0)
}

func (s *ComandoContext) TOK_MOUNT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MOUNT, 0)
}

func (s *ComandoContext) TOK_MKFS() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKFS, 0)
}

func (s *ComandoContext) TOK_LOGIN() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGIN, 0)
}

func (s *ComandoContext) TOK_MKGRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKGRP, 0)
}

func (s *ComandoContext) TOK_RMGRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMGRP, 0)
}

func (s *ComandoContext) TOK_MKUSR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKUSR, 0)
}

func (s *ComandoContext) TOK_RMUSR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RMUSR, 0)
}

func (s *ComandoContext) TOK_MKFILE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKFILE, 0)
}

func (s *ComandoContext) TOK_MKDIR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MKDIR, 0)
}

func (s *ComandoContext) TOK_EXEC() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_EXEC, 0)
}

func (s *ComandoContext) TOK_REP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_REP, 0)
}

func (s *ComandoContext) TOK_PAUSE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PAUSE, 0)
}

func (s *ComandoContext) TOK_LOGOUT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGOUT, 0)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterComando(s)
	}
}

func (s *ComandoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitComando(s)
	}
}

func (p *CmdParser) Comando() (localctx IComandoContext) {
	this := p
	_ = this

	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CmdParserRULE_comando)

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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CmdParserTOK_MKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(CmdParserTOK_MKDISK)
		}
		{
			p.SetState(27)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkdisk(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_RMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(CmdParserTOK_RMDISK)
		}
		{
			p.SetState(31)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewRmdisk(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_FDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(CmdParserTOK_FDISK)
		}
		{
			p.SetState(35)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewFdisk(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.Match(CmdParserTOK_MOUNT)
		}
		{
			p.SetState(39)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMount(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MKFS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.Match(CmdParserTOK_MKFS)
		}
		{
			p.SetState(43)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkfs(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_LOGIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(46)
			p.Match(CmdParserTOK_LOGIN)
		}
		{
			p.SetState(47)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewLogin(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MKGRP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(50)
			p.Match(CmdParserTOK_MKGRP)
		}
		{
			p.SetState(51)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkgrp(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_RMGRP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(54)
			p.Match(CmdParserTOK_RMGRP)
		}
		{
			p.SetState(55)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewRmgrp(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MKUSR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(58)
			p.Match(CmdParserTOK_MKUSR)
		}
		{
			p.SetState(59)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkusr(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_RMUSR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(62)
			p.Match(CmdParserTOK_RMUSR)
		}
		{
			p.SetState(63)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewRmusr(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MKFILE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(66)
			p.Match(CmdParserTOK_MKFILE)
		}
		{
			p.SetState(67)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkfile(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_MKDIR:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(70)
			p.Match(CmdParserTOK_MKDIR)
		}
		{
			p.SetState(71)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewMkdir(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_EXEC:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(74)
			p.Match(CmdParserTOK_EXEC)
		}
		{
			p.SetState(75)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewExec(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_REP:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(78)
			p.Match(CmdParserTOK_REP)
		}
		{
			p.SetState(79)

			var _x = p.param_list(0)

			localctx.(*ComandoContext)._param_list = _x
		}

		localctx.(*ComandoContext).com = ast.NewRep(localctx.(*ComandoContext).Get_param_list().GetLista())

	case CmdParserTOK_PAUSE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(82)
			p.Match(CmdParserTOK_PAUSE)
		}

		localctx.(*ComandoContext).com = ast.NewPause(true)

	case CmdParserTOK_LOGOUT:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(84)
			p.Match(CmdParserTOK_LOGOUT)
		}

		localctx.(*ComandoContext).com = ast.NewLogout(true)

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

	// GetAUXPLIST returns the AUXPLIST rule contexts.
	GetAUXPLIST() IParam_listContext

	// Get_param returns the _param rule contexts.
	Get_param() IParamContext

	// SetAUXPLIST sets the AUXPLIST rule contexts.
	SetAUXPLIST(IParam_listContext)

	// Set_param sets the _param rule contexts.
	Set_param(IParamContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	AUXPLIST IParam_listContext
	_param   IParamContext
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

func (s *Param_listContext) GetAUXPLIST() IParam_listContext { return s.AUXPLIST }

func (s *Param_listContext) Get_param() IParamContext { return s._param }

func (s *Param_listContext) SetAUXPLIST(v IParam_listContext) { s.AUXPLIST = v }

func (s *Param_listContext) Set_param(v IParamContext) { s._param = v }

func (s *Param_listContext) GetLista() *arrayList.List { return s.lista }

func (s *Param_listContext) SetLista(v *arrayList.List) { s.lista = v }

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
	localctx.(*Param_listContext).lista = arrayList.New()

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
		p.SetState(89)

		var _x = p.Param()

		localctx.(*Param_listContext)._param = _x
	}

	localctx.(*Param_listContext).lista.Add(localctx.(*Param_listContext).Get_param().GetPar())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParam_listContext(p, _parentctx, _parentState)
			localctx.(*Param_listContext).AUXPLIST = _prevctx
			p.PushNewRecursionContext(localctx, _startState, CmdParserRULE_param_list)
			p.SetState(92)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(93)

				var _x = p.Param()

				localctx.(*Param_listContext)._param = _x
			}

			localctx.(*Param_listContext).GetAUXPLIST().GetLista().Add(localctx.(*Param_listContext).Get_param().GetPar())
			localctx.(*Param_listContext).lista = localctx.(*Param_listContext).GetAUXPLIST().GetLista()

		}
		p.SetState(100)
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

	// Get_TOK_NUMERO returns the _TOK_NUMERO token.
	Get_TOK_NUMERO() antlr.Token

	// Get_TOK_CAMINO returns the _TOK_CAMINO token.
	Get_TOK_CAMINO() antlr.Token

	// Get_TOK_CADENA returns the _TOK_CADENA token.
	Get_TOK_CADENA() antlr.Token

	// Get_TOK_PALABRA returns the _TOK_PALABRA token.
	Get_TOK_PALABRA() antlr.Token

	// Get_TOK_IDENTIFICADOR returns the _TOK_IDENTIFICADOR token.
	Get_TOK_IDENTIFICADOR() antlr.Token

	// Set_TOK_NUMERO sets the _TOK_NUMERO token.
	Set_TOK_NUMERO(antlr.Token)

	// Set_TOK_CAMINO sets the _TOK_CAMINO token.
	Set_TOK_CAMINO(antlr.Token)

	// Set_TOK_CADENA sets the _TOK_CADENA token.
	Set_TOK_CADENA(antlr.Token)

	// Set_TOK_PALABRA sets the _TOK_PALABRA token.
	Set_TOK_PALABRA(antlr.Token)

	// Set_TOK_IDENTIFICADOR sets the _TOK_IDENTIFICADOR token.
	Set_TOK_IDENTIFICADOR(antlr.Token)

	// GetPar returns the par attribute.
	GetPar() parametros.Parametro

	// SetPar sets the par attribute.
	SetPar(parametros.Parametro)

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	par                parametros.Parametro
	_TOK_NUMERO        antlr.Token
	_TOK_CAMINO        antlr.Token
	_TOK_CADENA        antlr.Token
	_TOK_PALABRA       antlr.Token
	_TOK_IDENTIFICADOR antlr.Token
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

func (s *ParamContext) Get_TOK_NUMERO() antlr.Token { return s._TOK_NUMERO }

func (s *ParamContext) Get_TOK_CAMINO() antlr.Token { return s._TOK_CAMINO }

func (s *ParamContext) Get_TOK_CADENA() antlr.Token { return s._TOK_CADENA }

func (s *ParamContext) Get_TOK_PALABRA() antlr.Token { return s._TOK_PALABRA }

func (s *ParamContext) Get_TOK_IDENTIFICADOR() antlr.Token { return s._TOK_IDENTIFICADOR }

func (s *ParamContext) Set_TOK_NUMERO(v antlr.Token) { s._TOK_NUMERO = v }

func (s *ParamContext) Set_TOK_CAMINO(v antlr.Token) { s._TOK_CAMINO = v }

func (s *ParamContext) Set_TOK_CADENA(v antlr.Token) { s._TOK_CADENA = v }

func (s *ParamContext) Set_TOK_PALABRA(v antlr.Token) { s._TOK_PALABRA = v }

func (s *ParamContext) Set_TOK_IDENTIFICADOR(v antlr.Token) { s._TOK_IDENTIFICADOR = v }

func (s *ParamContext) GetPar() parametros.Parametro { return s.par }

func (s *ParamContext) SetPar(v parametros.Parametro) { s.par = v }

func (s *ParamContext) TOK_SIZE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_SIZE, 0)
}

func (s *ParamContext) TOK_IGUAL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IGUAL, 0)
}

func (s *ParamContext) TOK_NUMERO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_NUMERO, 0)
}

func (s *ParamContext) TOK_PATH() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PATH, 0)
}

func (s *ParamContext) TOK_CAMINO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CAMINO, 0)
}

func (s *ParamContext) TOK_CADENA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CADENA, 0)
}

func (s *ParamContext) TOK_FIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIT, 0)
}

func (s *ParamContext) TOK_FIRST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FIRST, 0)
}

func (s *ParamContext) TOK_WORST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_WORST, 0)
}

func (s *ParamContext) TOK_BEST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_BEST, 0)
}

func (s *ParamContext) TOK_UNIT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_UNIT, 0)
}

func (s *ParamContext) TOK_BYTES() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_BYTES, 0)
}

func (s *ParamContext) TOK_KB() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_KB, 0)
}

func (s *ParamContext) TOK_MB() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_MB, 0)
}

func (s *ParamContext) TOK_NAME() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_NAME, 0)
}

func (s *ParamContext) TOK_PALABRA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PALABRA, 0)
}

func (s *ParamContext) TOK_USUARIO() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_USUARIO, 0)
}

func (s *ParamContext) TOK_GRP() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_GRP, 0)
}

func (s *ParamContext) TOK_PASSWORD() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PASSWORD, 0)
}

func (s *ParamContext) TOK_PWD() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PWD, 0)
}

func (s *ParamContext) TOK_TYPE() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_TYPE, 0)
}

func (s *ParamContext) TOK_PRIMARIA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_PRIMARIA, 0)
}

func (s *ParamContext) TOK_LOGICA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_LOGICA, 0)
}

func (s *ParamContext) TOK_EXTENDIDA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_EXTENDIDA, 0)
}

func (s *ParamContext) TOK_FAST() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FAST, 0)
}

func (s *ParamContext) TOK_FULL() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_FULL, 0)
}

func (s *ParamContext) TOK_ID() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_ID, 0)
}

func (s *ParamContext) TOK_IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_IDENTIFICADOR, 0)
}

func (s *ParamContext) TOK_CONT() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_CONT, 0)
}

func (s *ParamContext) TOK_RUTA() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_RUTA, 0)
}

func (s *ParamContext) TOK_P() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_P, 0)
}

func (s *ParamContext) TOK_R() antlr.TerminalNode {
	return s.GetToken(CmdParserTOK_R, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CmdListener); ok {
		listenerT.ExitParam(s)
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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(CmdParserTOK_SIZE)
		}
		{
			p.SetState(102)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(103)

			var _m = p.Match(CmdParserTOK_NUMERO)

			localctx.(*ParamContext)._TOK_NUMERO = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.SIZE, (func() string {
			if localctx.(*ParamContext).Get_TOK_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_NUMERO().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(CmdParserTOK_PATH)
		}
		{
			p.SetState(106)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(107)

			var _m = p.Match(CmdParserTOK_CAMINO)

			localctx.(*ParamContext)._TOK_CAMINO = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.PATH, (func() string {
			if localctx.(*ParamContext).Get_TOK_CAMINO() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CAMINO().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Match(CmdParserTOK_PATH)
		}
		{
			p.SetState(110)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(111)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.PATH, str)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(114)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(115)
			p.Match(CmdParserTOK_FIRST)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.FIT, ast.FF)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(118)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(119)
			p.Match(CmdParserTOK_WORST)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.FIT, ast.WF)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Match(CmdParserTOK_FIT)
		}
		{
			p.SetState(122)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(123)
			p.Match(CmdParserTOK_BEST)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.FIT, ast.BF)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(125)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(126)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(127)
			p.Match(CmdParserTOK_BYTES)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.UNIT, ast.B)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(129)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(130)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(131)
			p.Match(CmdParserTOK_KB)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.UNIT, ast.K)

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(133)
			p.Match(CmdParserTOK_UNIT)
		}
		{
			p.SetState(134)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(135)
			p.Match(CmdParserTOK_MB)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.UNIT, ast.M)

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(137)
			p.Match(CmdParserTOK_NAME)
		}
		{
			p.SetState(138)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(139)

			var _m = p.Match(CmdParserTOK_PALABRA)

			localctx.(*ParamContext)._TOK_PALABRA = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.NAME, (func() string {
			if localctx.(*ParamContext).Get_TOK_PALABRA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_PALABRA().GetText()
			}
		}()))

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(141)
			p.Match(CmdParserTOK_NAME)
		}
		{
			p.SetState(142)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(143)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.NAME, str)

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(145)
			p.Match(CmdParserTOK_USUARIO)
		}
		{
			p.SetState(146)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(147)

			var _m = p.Match(CmdParserTOK_PALABRA)

			localctx.(*ParamContext)._TOK_PALABRA = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.USUARIO, (func() string {
			if localctx.(*ParamContext).Get_TOK_PALABRA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_PALABRA().GetText()
			}
		}()))

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(149)
			p.Match(CmdParserTOK_USUARIO)
		}
		{
			p.SetState(150)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(151)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.USUARIO, str)

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(153)
			p.Match(CmdParserTOK_GRP)
		}
		{
			p.SetState(154)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(155)

			var _m = p.Match(CmdParserTOK_PALABRA)

			localctx.(*ParamContext)._TOK_PALABRA = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.GRP, (func() string {
			if localctx.(*ParamContext).Get_TOK_PALABRA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_PALABRA().GetText()
			}
		}()))

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(157)
			p.Match(CmdParserTOK_GRP)
		}
		{
			p.SetState(158)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(159)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.GRP, str)

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(161)
			p.Match(CmdParserTOK_PASSWORD)
		}
		{
			p.SetState(162)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(163)

			var _m = p.Match(CmdParserTOK_PALABRA)

			localctx.(*ParamContext)._TOK_PALABRA = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.PASSWORD, (func() string {
			if localctx.(*ParamContext).Get_TOK_PALABRA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_PALABRA().GetText()
			}
		}()))

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(165)
			p.Match(CmdParserTOK_PWD)
		}
		{
			p.SetState(166)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(167)

			var _m = p.Match(CmdParserTOK_PALABRA)

			localctx.(*ParamContext)._TOK_PALABRA = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.PWD, (func() string {
			if localctx.(*ParamContext).Get_TOK_PALABRA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_PALABRA().GetText()
			}
		}()))

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(169)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(170)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(171)
			p.Match(CmdParserTOK_PRIMARIA)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.TYPE, ast.PP)

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(173)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(174)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(175)
			p.Match(CmdParserTOK_LOGICA)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.TYPE, ast.L)

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(177)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(178)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(179)
			p.Match(CmdParserTOK_EXTENDIDA)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.TYPE, ast.E)

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(181)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(182)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(183)
			p.Match(CmdParserTOK_FAST)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.TYPE, ast.FAST)

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(185)
			p.Match(CmdParserTOK_TYPE)
		}
		{
			p.SetState(186)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(187)
			p.Match(CmdParserTOK_FULL)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.TYPE, ast.FULL)

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(189)
			p.Match(CmdParserTOK_ID)
		}
		{
			p.SetState(190)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(191)

			var _m = p.Match(CmdParserTOK_IDENTIFICADOR)

			localctx.(*ParamContext)._TOK_IDENTIFICADOR = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.ID, (func() string {
			if localctx.(*ParamContext).Get_TOK_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_IDENTIFICADOR().GetText()
			}
		}()))

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(193)
			p.Match(CmdParserTOK_CONT)
		}
		{
			p.SetState(194)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(195)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.CONT, str)

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(197)
			p.Match(CmdParserTOK_CONT)
		}
		{
			p.SetState(198)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(199)

			var _m = p.Match(CmdParserTOK_CAMINO)

			localctx.(*ParamContext)._TOK_CAMINO = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.CONT, (func() string {
			if localctx.(*ParamContext).Get_TOK_CAMINO() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CAMINO().GetText()
			}
		}()))

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(201)
			p.Match(CmdParserTOK_RUTA)
		}
		{
			p.SetState(202)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(203)

			var _m = p.Match(CmdParserTOK_CADENA)

			localctx.(*ParamContext)._TOK_CADENA = _m
		}

		str := strings.Trim((func() string {
			if localctx.(*ParamContext).Get_TOK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CADENA().GetText()
			}
		}()), "\"")
		localctx.(*ParamContext).par = parametros.NewParametro(ast.RUTA, str)

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(205)
			p.Match(CmdParserTOK_RUTA)
		}
		{
			p.SetState(206)
			p.Match(CmdParserTOK_IGUAL)
		}
		{
			p.SetState(207)

			var _m = p.Match(CmdParserTOK_CAMINO)

			localctx.(*ParamContext)._TOK_CAMINO = _m
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.RUTA, (func() string {
			if localctx.(*ParamContext).Get_TOK_CAMINO() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_TOK_CAMINO().GetText()
			}
		}()))

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(209)
			p.Match(CmdParserTOK_P)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.P, "true")

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(211)
			p.Match(CmdParserTOK_R)
		}

		localctx.(*ParamContext).par = parametros.NewParametro(ast.R, "true")

	}

	return localctx
}

func (p *CmdParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ComandListContext = nil
		if localctx != nil {
			t = localctx.(*ComandListContext)
		}
		return p.ComandList_Sempred(t, predIndex)

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

func (p *CmdParser) ComandList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CmdParser) Param_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

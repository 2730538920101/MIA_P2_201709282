// Code generated from CmdLex.g4 by ANTLR 4.10. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CmdLex struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cmdlexLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cmdlexLexerInit() {
	staticData := &cmdlexLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"SLASH", "COMDOB", "HASHTAG", "DASH", "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK",
		"TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP",
		"TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC",
		"TOK_PAUSE", "TOK_REP", "TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT",
		"TOK_NAME", "TOK_TYPE", "TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD",
		"TOK_CONT", "TOK_GRP", "TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB",
		"TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST",
		"TOK_WORST", "TOK_BEST", "TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO",
		"TOK_IDENTIFICADOR", "TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 390, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 47, 1, 47, 5, 47, 335, 8, 47, 10, 47, 12, 47, 338, 9, 47, 1,
		47, 1, 47, 1, 48, 4, 48, 343, 8, 48, 11, 48, 12, 48, 344, 1, 49, 1, 49,
		1, 49, 1, 49, 3, 49, 351, 8, 49, 1, 50, 1, 50, 5, 50, 355, 8, 50, 10, 50,
		12, 50, 358, 9, 50, 4, 50, 360, 8, 50, 11, 50, 12, 50, 361, 1, 51, 4, 51,
		365, 8, 51, 11, 51, 12, 51, 366, 1, 52, 1, 52, 5, 52, 371, 8, 52, 10, 52,
		12, 52, 374, 9, 52, 1, 52, 5, 52, 377, 8, 52, 10, 52, 12, 52, 380, 9, 52,
		1, 52, 1, 52, 1, 53, 4, 53, 385, 8, 53, 11, 53, 12, 53, 386, 1, 53, 1,
		53, 0, 0, 54, 1, 0, 3, 0, 5, 0, 7, 0, 9, 1, 11, 2, 13, 3, 15, 4, 17, 5,
		19, 6, 21, 7, 23, 8, 25, 9, 27, 10, 29, 11, 31, 12, 33, 13, 35, 14, 37,
		15, 39, 16, 41, 17, 43, 18, 45, 19, 47, 20, 49, 21, 51, 22, 53, 23, 55,
		24, 57, 25, 59, 26, 61, 27, 63, 28, 65, 29, 67, 30, 69, 31, 71, 32, 73,
		33, 75, 34, 77, 35, 79, 36, 81, 37, 83, 38, 85, 39, 87, 40, 89, 41, 91,
		42, 93, 43, 95, 44, 97, 45, 99, 46, 101, 47, 103, 48, 105, 49, 107, 50,
		1, 0, 31, 1, 0, 61, 61, 2, 0, 77, 77, 109, 109, 2, 0, 75, 75, 107, 107,
		2, 0, 68, 68, 100, 100, 2, 0, 73, 73, 105, 105, 2, 0, 83, 83, 115, 115,
		2, 0, 82, 82, 114, 114, 2, 0, 70, 70, 102, 102, 2, 0, 79, 79, 111, 111,
		2, 0, 85, 85, 117, 117, 2, 0, 78, 78, 110, 110, 2, 0, 84, 84, 116, 116,
		2, 0, 76, 76, 108, 108, 2, 0, 71, 71, 103, 103, 2, 0, 80, 80, 112, 112,
		2, 0, 69, 69, 101, 101, 2, 0, 88, 88, 120, 120, 2, 0, 67, 67, 99, 99, 2,
		0, 65, 65, 97, 97, 2, 0, 72, 72, 104, 104, 2, 0, 90, 90, 122, 122, 2, 0,
		89, 89, 121, 121, 2, 0, 87, 87, 119, 119, 2, 0, 66, 66, 98, 98, 2, 0, 10,
		10, 34, 34, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 10, 10, 32, 32,
		34, 34, 5, 0, 45, 46, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 10, 10, 3,
		0, 9, 10, 13, 13, 32, 32, 393, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0,
		0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1,
		0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97,
		1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0,
		0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 1, 109, 1, 0, 0, 0, 3, 111, 1,
		0, 0, 0, 5, 113, 1, 0, 0, 0, 7, 115, 1, 0, 0, 0, 9, 117, 1, 0, 0, 0, 11,
		119, 1, 0, 0, 0, 13, 126, 1, 0, 0, 0, 15, 133, 1, 0, 0, 0, 17, 139, 1,
		0, 0, 0, 19, 145, 1, 0, 0, 0, 21, 150, 1, 0, 0, 0, 23, 156, 1, 0, 0, 0,
		25, 163, 1, 0, 0, 0, 27, 169, 1, 0, 0, 0, 29, 175, 1, 0, 0, 0, 31, 181,
		1, 0, 0, 0, 33, 187, 1, 0, 0, 0, 35, 194, 1, 0, 0, 0, 37, 200, 1, 0, 0,
		0, 39, 205, 1, 0, 0, 0, 41, 211, 1, 0, 0, 0, 43, 215, 1, 0, 0, 0, 45, 221,
		1, 0, 0, 0, 47, 226, 1, 0, 0, 0, 49, 232, 1, 0, 0, 0, 51, 238, 1, 0, 0,
		0, 53, 244, 1, 0, 0, 0, 55, 250, 1, 0, 0, 0, 57, 254, 1, 0, 0, 0, 59, 263,
		1, 0, 0, 0, 61, 273, 1, 0, 0, 0, 63, 278, 1, 0, 0, 0, 65, 284, 1, 0, 0,
		0, 67, 289, 1, 0, 0, 0, 69, 295, 1, 0, 0, 0, 71, 298, 1, 0, 0, 0, 73, 301,
		1, 0, 0, 0, 75, 303, 1, 0, 0, 0, 77, 305, 1, 0, 0, 0, 79, 307, 1, 0, 0,
		0, 81, 309, 1, 0, 0, 0, 83, 311, 1, 0, 0, 0, 85, 313, 1, 0, 0, 0, 87, 316,
		1, 0, 0, 0, 89, 319, 1, 0, 0, 0, 91, 322, 1, 0, 0, 0, 93, 327, 1, 0, 0,
		0, 95, 332, 1, 0, 0, 0, 97, 342, 1, 0, 0, 0, 99, 346, 1, 0, 0, 0, 101,
		359, 1, 0, 0, 0, 103, 364, 1, 0, 0, 0, 105, 368, 1, 0, 0, 0, 107, 384,
		1, 0, 0, 0, 109, 110, 5, 47, 0, 0, 110, 2, 1, 0, 0, 0, 111, 112, 5, 34,
		0, 0, 112, 4, 1, 0, 0, 0, 113, 114, 5, 35, 0, 0, 114, 6, 1, 0, 0, 0, 115,
		116, 5, 45, 0, 0, 116, 8, 1, 0, 0, 0, 117, 118, 7, 0, 0, 0, 118, 10, 1,
		0, 0, 0, 119, 120, 7, 1, 0, 0, 120, 121, 7, 2, 0, 0, 121, 122, 7, 3, 0,
		0, 122, 123, 7, 4, 0, 0, 123, 124, 7, 5, 0, 0, 124, 125, 7, 2, 0, 0, 125,
		12, 1, 0, 0, 0, 126, 127, 7, 6, 0, 0, 127, 128, 7, 1, 0, 0, 128, 129, 7,
		3, 0, 0, 129, 130, 7, 4, 0, 0, 130, 131, 7, 5, 0, 0, 131, 132, 7, 2, 0,
		0, 132, 14, 1, 0, 0, 0, 133, 134, 7, 7, 0, 0, 134, 135, 7, 3, 0, 0, 135,
		136, 7, 4, 0, 0, 136, 137, 7, 5, 0, 0, 137, 138, 7, 2, 0, 0, 138, 16, 1,
		0, 0, 0, 139, 140, 7, 1, 0, 0, 140, 141, 7, 8, 0, 0, 141, 142, 7, 9, 0,
		0, 142, 143, 7, 10, 0, 0, 143, 144, 7, 11, 0, 0, 144, 18, 1, 0, 0, 0, 145,
		146, 7, 1, 0, 0, 146, 147, 7, 2, 0, 0, 147, 148, 7, 7, 0, 0, 148, 149,
		7, 5, 0, 0, 149, 20, 1, 0, 0, 0, 150, 151, 7, 12, 0, 0, 151, 152, 7, 8,
		0, 0, 152, 153, 7, 13, 0, 0, 153, 154, 7, 4, 0, 0, 154, 155, 7, 10, 0,
		0, 155, 22, 1, 0, 0, 0, 156, 157, 7, 12, 0, 0, 157, 158, 7, 8, 0, 0, 158,
		159, 7, 13, 0, 0, 159, 160, 7, 8, 0, 0, 160, 161, 7, 9, 0, 0, 161, 162,
		7, 11, 0, 0, 162, 24, 1, 0, 0, 0, 163, 164, 7, 1, 0, 0, 164, 165, 7, 2,
		0, 0, 165, 166, 7, 13, 0, 0, 166, 167, 7, 6, 0, 0, 167, 168, 7, 14, 0,
		0, 168, 26, 1, 0, 0, 0, 169, 170, 7, 6, 0, 0, 170, 171, 7, 1, 0, 0, 171,
		172, 7, 13, 0, 0, 172, 173, 7, 6, 0, 0, 173, 174, 7, 14, 0, 0, 174, 28,
		1, 0, 0, 0, 175, 176, 7, 1, 0, 0, 176, 177, 7, 2, 0, 0, 177, 178, 7, 9,
		0, 0, 178, 179, 7, 5, 0, 0, 179, 180, 7, 6, 0, 0, 180, 30, 1, 0, 0, 0,
		181, 182, 7, 6, 0, 0, 182, 183, 7, 1, 0, 0, 183, 184, 7, 9, 0, 0, 184,
		185, 7, 5, 0, 0, 185, 186, 7, 6, 0, 0, 186, 32, 1, 0, 0, 0, 187, 188, 7,
		1, 0, 0, 188, 189, 7, 2, 0, 0, 189, 190, 7, 7, 0, 0, 190, 191, 7, 4, 0,
		0, 191, 192, 7, 12, 0, 0, 192, 193, 7, 15, 0, 0, 193, 34, 1, 0, 0, 0, 194,
		195, 7, 1, 0, 0, 195, 196, 7, 2, 0, 0, 196, 197, 7, 3, 0, 0, 197, 198,
		7, 4, 0, 0, 198, 199, 7, 6, 0, 0, 199, 36, 1, 0, 0, 0, 200, 201, 7, 15,
		0, 0, 201, 202, 7, 16, 0, 0, 202, 203, 7, 15, 0, 0, 203, 204, 7, 17, 0,
		0, 204, 38, 1, 0, 0, 0, 205, 206, 7, 14, 0, 0, 206, 207, 7, 18, 0, 0, 207,
		208, 7, 9, 0, 0, 208, 209, 7, 5, 0, 0, 209, 210, 7, 15, 0, 0, 210, 40,
		1, 0, 0, 0, 211, 212, 7, 6, 0, 0, 212, 213, 7, 15, 0, 0, 213, 214, 7, 14,
		0, 0, 214, 42, 1, 0, 0, 0, 215, 216, 3, 7, 3, 0, 216, 217, 7, 14, 0, 0,
		217, 218, 7, 18, 0, 0, 218, 219, 7, 11, 0, 0, 219, 220, 7, 19, 0, 0, 220,
		44, 1, 0, 0, 0, 221, 222, 3, 7, 3, 0, 222, 223, 7, 7, 0, 0, 223, 224, 7,
		4, 0, 0, 224, 225, 7, 11, 0, 0, 225, 46, 1, 0, 0, 0, 226, 227, 3, 7, 3,
		0, 227, 228, 7, 5, 0, 0, 228, 229, 7, 4, 0, 0, 229, 230, 7, 20, 0, 0, 230,
		231, 7, 15, 0, 0, 231, 48, 1, 0, 0, 0, 232, 233, 3, 7, 3, 0, 233, 234,
		7, 9, 0, 0, 234, 235, 7, 10, 0, 0, 235, 236, 7, 4, 0, 0, 236, 237, 7, 11,
		0, 0, 237, 50, 1, 0, 0, 0, 238, 239, 3, 7, 3, 0, 239, 240, 7, 10, 0, 0,
		240, 241, 7, 18, 0, 0, 241, 242, 7, 1, 0, 0, 242, 243, 7, 15, 0, 0, 243,
		52, 1, 0, 0, 0, 244, 245, 3, 7, 3, 0, 245, 246, 7, 11, 0, 0, 246, 247,
		7, 21, 0, 0, 247, 248, 7, 14, 0, 0, 248, 249, 7, 15, 0, 0, 249, 54, 1,
		0, 0, 0, 250, 251, 3, 7, 3, 0, 251, 252, 7, 4, 0, 0, 252, 253, 7, 3, 0,
		0, 253, 56, 1, 0, 0, 0, 254, 255, 3, 7, 3, 0, 255, 256, 7, 9, 0, 0, 256,
		257, 7, 5, 0, 0, 257, 258, 7, 9, 0, 0, 258, 259, 7, 18, 0, 0, 259, 260,
		7, 6, 0, 0, 260, 261, 7, 4, 0, 0, 261, 262, 7, 8, 0, 0, 262, 58, 1, 0,
		0, 0, 263, 264, 3, 7, 3, 0, 264, 265, 7, 14, 0, 0, 265, 266, 7, 18, 0,
		0, 266, 267, 7, 5, 0, 0, 267, 268, 7, 5, 0, 0, 268, 269, 7, 22, 0, 0, 269,
		270, 7, 8, 0, 0, 270, 271, 7, 6, 0, 0, 271, 272, 7, 3, 0, 0, 272, 60, 1,
		0, 0, 0, 273, 274, 3, 7, 3, 0, 274, 275, 7, 14, 0, 0, 275, 276, 7, 22,
		0, 0, 276, 277, 7, 3, 0, 0, 277, 62, 1, 0, 0, 0, 278, 279, 3, 7, 3, 0,
		279, 280, 7, 17, 0, 0, 280, 281, 7, 8, 0, 0, 281, 282, 7, 10, 0, 0, 282,
		283, 7, 11, 0, 0, 283, 64, 1, 0, 0, 0, 284, 285, 3, 7, 3, 0, 285, 286,
		7, 13, 0, 0, 286, 287, 7, 6, 0, 0, 287, 288, 7, 14, 0, 0, 288, 66, 1, 0,
		0, 0, 289, 290, 3, 7, 3, 0, 290, 291, 7, 6, 0, 0, 291, 292, 7, 9, 0, 0,
		292, 293, 7, 11, 0, 0, 293, 294, 7, 18, 0, 0, 294, 68, 1, 0, 0, 0, 295,
		296, 3, 7, 3, 0, 296, 297, 7, 6, 0, 0, 297, 70, 1, 0, 0, 0, 298, 299, 3,
		7, 3, 0, 299, 300, 7, 14, 0, 0, 300, 72, 1, 0, 0, 0, 301, 302, 7, 2, 0,
		0, 302, 74, 1, 0, 0, 0, 303, 304, 7, 1, 0, 0, 304, 76, 1, 0, 0, 0, 305,
		306, 7, 23, 0, 0, 306, 78, 1, 0, 0, 0, 307, 308, 7, 14, 0, 0, 308, 80,
		1, 0, 0, 0, 309, 310, 7, 15, 0, 0, 310, 82, 1, 0, 0, 0, 311, 312, 7, 12,
		0, 0, 312, 84, 1, 0, 0, 0, 313, 314, 7, 7, 0, 0, 314, 315, 7, 7, 0, 0,
		315, 86, 1, 0, 0, 0, 316, 317, 7, 22, 0, 0, 317, 318, 7, 7, 0, 0, 318,
		88, 1, 0, 0, 0, 319, 320, 7, 23, 0, 0, 320, 321, 7, 7, 0, 0, 321, 90, 1,
		0, 0, 0, 322, 323, 7, 7, 0, 0, 323, 324, 7, 18, 0, 0, 324, 325, 7, 5, 0,
		0, 325, 326, 7, 11, 0, 0, 326, 92, 1, 0, 0, 0, 327, 328, 7, 7, 0, 0, 328,
		329, 7, 9, 0, 0, 329, 330, 7, 12, 0, 0, 330, 331, 7, 12, 0, 0, 331, 94,
		1, 0, 0, 0, 332, 336, 3, 3, 1, 0, 333, 335, 8, 24, 0, 0, 334, 333, 1, 0,
		0, 0, 335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0,
		337, 339, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 340, 3, 3, 1, 0, 340,
		96, 1, 0, 0, 0, 341, 343, 7, 25, 0, 0, 342, 341, 1, 0, 0, 0, 343, 344,
		1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 98, 1, 0,
		0, 0, 346, 347, 7, 25, 0, 0, 347, 348, 7, 25, 0, 0, 348, 350, 7, 25, 0,
		0, 349, 351, 7, 26, 0, 0, 350, 349, 1, 0, 0, 0, 351, 100, 1, 0, 0, 0, 352,
		356, 3, 1, 0, 0, 353, 355, 8, 27, 0, 0, 354, 353, 1, 0, 0, 0, 355, 358,
		1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 360, 1, 0,
		0, 0, 358, 356, 1, 0, 0, 0, 359, 352, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0,
		361, 359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 102, 1, 0, 0, 0, 363,
		365, 7, 28, 0, 0, 364, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 364,
		1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 104, 1, 0, 0, 0, 368, 372, 3, 5,
		2, 0, 369, 371, 8, 29, 0, 0, 370, 369, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0,
		372, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 378, 1, 0, 0, 0, 374,
		372, 1, 0, 0, 0, 375, 377, 7, 29, 0, 0, 376, 375, 1, 0, 0, 0, 377, 380,
		1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 381, 1, 0,
		0, 0, 380, 378, 1, 0, 0, 0, 381, 382, 6, 52, 0, 0, 382, 106, 1, 0, 0, 0,
		383, 385, 7, 30, 0, 0, 384, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389,
		6, 53, 0, 0, 389, 108, 1, 0, 0, 0, 11, 0, 336, 344, 350, 356, 361, 364,
		366, 372, 378, 386, 1, 6, 0, 0,
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

// CmdLexInit initializes any static state used to implement CmdLex. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCmdLex(). You can call this function if you wish to initialize the static state ahead
// of time.
func CmdLexInit() {
	staticData := &cmdlexLexerStaticData
	staticData.once.Do(cmdlexLexerInit)
}

// NewCmdLex produces a new lexer instance for the optional input antlr.CharStream.
func NewCmdLex(input antlr.CharStream) *CmdLex {
	CmdLexInit()
	l := new(CmdLex)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cmdlexLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CmdLex.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CmdLex tokens.
const (
	CmdLexTOK_IGUAL         = 1
	CmdLexTOK_MKDISK        = 2
	CmdLexTOK_RMDISK        = 3
	CmdLexTOK_FDISK         = 4
	CmdLexTOK_MOUNT         = 5
	CmdLexTOK_MKFS          = 6
	CmdLexTOK_LOGIN         = 7
	CmdLexTOK_LOGOUT        = 8
	CmdLexTOK_MKGRP         = 9
	CmdLexTOK_RMGRP         = 10
	CmdLexTOK_MKUSR         = 11
	CmdLexTOK_RMUSR         = 12
	CmdLexTOK_MKFILE        = 13
	CmdLexTOK_MKDIR         = 14
	CmdLexTOK_EXEC          = 15
	CmdLexTOK_PAUSE         = 16
	CmdLexTOK_REP           = 17
	CmdLexTOK_PATH          = 18
	CmdLexTOK_FIT           = 19
	CmdLexTOK_SIZE          = 20
	CmdLexTOK_UNIT          = 21
	CmdLexTOK_NAME          = 22
	CmdLexTOK_TYPE          = 23
	CmdLexTOK_ID            = 24
	CmdLexTOK_USUARIO       = 25
	CmdLexTOK_PASSWORD      = 26
	CmdLexTOK_PWD           = 27
	CmdLexTOK_CONT          = 28
	CmdLexTOK_GRP           = 29
	CmdLexTOK_RUTA          = 30
	CmdLexTOK_R             = 31
	CmdLexTOK_P             = 32
	CmdLexTOK_KB            = 33
	CmdLexTOK_MB            = 34
	CmdLexTOK_BYTES         = 35
	CmdLexTOK_PRIMARIA      = 36
	CmdLexTOK_EXTENDIDA     = 37
	CmdLexTOK_LOGICA        = 38
	CmdLexTOK_FIRST         = 39
	CmdLexTOK_WORST         = 40
	CmdLexTOK_BEST          = 41
	CmdLexTOK_FAST          = 42
	CmdLexTOK_FULL          = 43
	CmdLexTOK_CADENA        = 44
	CmdLexTOK_NUMERO        = 45
	CmdLexTOK_IDENTIFICADOR = 46
	CmdLexTOK_CAMINO        = 47
	CmdLexTOK_PALABRA       = 48
	CmdLexCOMENTARIOS       = 49
	CmdLexWHITESPACE        = 50
)

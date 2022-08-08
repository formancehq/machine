// Code generated from NumScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type NumScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var numscriptlexerLexerStaticData struct {
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

func numscriptlexerLexerInit() {
	staticData := &numscriptlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'*'", "'allowing overdraft up to'", "'allowing overdraft up to its limit'",
		"'allowing unbounded overdraft'", "','", "", "", "", "", "'vars'", "'meta'",
		"'set_tx_meta'", "'print'", "'fail'", "'send'", "'source'", "'from'",
		"'max'", "'destination'", "'to'", "'allocate'", "'+'", "'-'", "'('",
		"')'", "'['", "']'", "'{'", "'}'", "'='", "'account'", "'asset'", "'number'",
		"'monetary'", "'portion'", "'string'", "", "", "'remaining'", "'kept'",
		"", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT",
		"LINE_COMMENT", "VARS", "META", "SET_TX_META", "PRINT", "FAIL", "SEND",
		"SOURCE", "FROM", "MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD",
		"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
		"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
		"TY_STRING", "STRING", "PORTION", "REMAINING", "KEPT", "NUMBER", "PERCENT",
		"VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT",
		"LINE_COMMENT", "VARS", "META", "SET_TX_META", "PRINT", "FAIL", "SEND",
		"SOURCE", "FROM", "MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD",
		"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
		"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
		"TY_STRING", "STRING", "PORTION", "REMAINING", "KEPT", "NUMBER", "PERCENT",
		"VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 458, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4, 5, 186, 8, 5, 11, 5, 12, 5, 187, 1,
		6, 4, 6, 191, 8, 6, 11, 6, 12, 6, 192, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 202, 8, 7, 10, 7, 12, 7, 205, 9, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 216, 8, 8, 10, 8, 12, 8, 219, 9,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 5,
		36, 368, 8, 36, 10, 36, 12, 36, 371, 9, 36, 1, 36, 1, 36, 1, 37, 4, 37,
		376, 8, 37, 11, 37, 12, 37, 377, 1, 37, 3, 37, 381, 8, 37, 1, 37, 1, 37,
		3, 37, 385, 8, 37, 1, 37, 4, 37, 388, 8, 37, 11, 37, 12, 37, 389, 1, 37,
		4, 37, 393, 8, 37, 11, 37, 12, 37, 394, 1, 37, 1, 37, 4, 37, 399, 8, 37,
		11, 37, 12, 37, 400, 3, 37, 403, 8, 37, 1, 37, 3, 37, 406, 8, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 40, 4, 40, 424, 8, 40, 11, 40, 12, 40, 425,
		1, 41, 1, 41, 1, 42, 1, 42, 4, 42, 432, 8, 42, 11, 42, 12, 42, 433, 1,
		42, 5, 42, 437, 8, 42, 10, 42, 12, 42, 440, 9, 42, 1, 43, 1, 43, 4, 43,
		444, 8, 43, 11, 43, 12, 43, 445, 1, 43, 5, 43, 449, 8, 43, 10, 43, 12,
		43, 452, 9, 43, 1, 44, 4, 44, 455, 8, 44, 11, 44, 12, 44, 456, 2, 203,
		217, 0, 45, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 1,
		0, 10, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 6, 0, 32, 32, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 32, 32, 2, 0, 95,
		95, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 58, 65, 90, 95, 95, 97, 122, 2, 0, 47, 57, 65, 90, 477, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 91, 1, 0, 0, 0, 3, 93, 1,
		0, 0, 0, 5, 118, 1, 0, 0, 0, 7, 153, 1, 0, 0, 0, 9, 182, 1, 0, 0, 0, 11,
		185, 1, 0, 0, 0, 13, 190, 1, 0, 0, 0, 15, 196, 1, 0, 0, 0, 17, 211, 1,
		0, 0, 0, 19, 224, 1, 0, 0, 0, 21, 229, 1, 0, 0, 0, 23, 234, 1, 0, 0, 0,
		25, 246, 1, 0, 0, 0, 27, 252, 1, 0, 0, 0, 29, 257, 1, 0, 0, 0, 31, 262,
		1, 0, 0, 0, 33, 269, 1, 0, 0, 0, 35, 274, 1, 0, 0, 0, 37, 278, 1, 0, 0,
		0, 39, 290, 1, 0, 0, 0, 41, 293, 1, 0, 0, 0, 43, 302, 1, 0, 0, 0, 45, 304,
		1, 0, 0, 0, 47, 306, 1, 0, 0, 0, 49, 308, 1, 0, 0, 0, 51, 310, 1, 0, 0,
		0, 53, 312, 1, 0, 0, 0, 55, 314, 1, 0, 0, 0, 57, 316, 1, 0, 0, 0, 59, 318,
		1, 0, 0, 0, 61, 320, 1, 0, 0, 0, 63, 328, 1, 0, 0, 0, 65, 334, 1, 0, 0,
		0, 67, 341, 1, 0, 0, 0, 69, 350, 1, 0, 0, 0, 71, 358, 1, 0, 0, 0, 73, 365,
		1, 0, 0, 0, 75, 405, 1, 0, 0, 0, 77, 407, 1, 0, 0, 0, 79, 417, 1, 0, 0,
		0, 81, 423, 1, 0, 0, 0, 83, 427, 1, 0, 0, 0, 85, 429, 1, 0, 0, 0, 87, 441,
		1, 0, 0, 0, 89, 454, 1, 0, 0, 0, 91, 92, 5, 42, 0, 0, 92, 2, 1, 0, 0, 0,
		93, 94, 5, 97, 0, 0, 94, 95, 5, 108, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97,
		5, 111, 0, 0, 97, 98, 5, 119, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100, 5, 110,
		0, 0, 100, 101, 5, 103, 0, 0, 101, 102, 5, 32, 0, 0, 102, 103, 5, 111,
		0, 0, 103, 104, 5, 118, 0, 0, 104, 105, 5, 101, 0, 0, 105, 106, 5, 114,
		0, 0, 106, 107, 5, 100, 0, 0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 97,
		0, 0, 109, 110, 5, 102, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 32,
		0, 0, 112, 113, 5, 117, 0, 0, 113, 114, 5, 112, 0, 0, 114, 115, 5, 32,
		0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 111, 0, 0, 117, 4, 1, 0, 0,
		0, 118, 119, 5, 97, 0, 0, 119, 120, 5, 108, 0, 0, 120, 121, 5, 108, 0,
		0, 121, 122, 5, 111, 0, 0, 122, 123, 5, 119, 0, 0, 123, 124, 5, 105, 0,
		0, 124, 125, 5, 110, 0, 0, 125, 126, 5, 103, 0, 0, 126, 127, 5, 32, 0,
		0, 127, 128, 5, 111, 0, 0, 128, 129, 5, 118, 0, 0, 129, 130, 5, 101, 0,
		0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 100, 0, 0, 132, 133, 5, 114, 0,
		0, 133, 134, 5, 97, 0, 0, 134, 135, 5, 102, 0, 0, 135, 136, 5, 116, 0,
		0, 136, 137, 5, 32, 0, 0, 137, 138, 5, 117, 0, 0, 138, 139, 5, 112, 0,
		0, 139, 140, 5, 32, 0, 0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 111, 0,
		0, 142, 143, 5, 32, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 116, 0,
		0, 145, 146, 5, 115, 0, 0, 146, 147, 5, 32, 0, 0, 147, 148, 5, 108, 0,
		0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 109, 0, 0, 150, 151, 5, 105, 0,
		0, 151, 152, 5, 116, 0, 0, 152, 6, 1, 0, 0, 0, 153, 154, 5, 97, 0, 0, 154,
		155, 5, 108, 0, 0, 155, 156, 5, 108, 0, 0, 156, 157, 5, 111, 0, 0, 157,
		158, 5, 119, 0, 0, 158, 159, 5, 105, 0, 0, 159, 160, 5, 110, 0, 0, 160,
		161, 5, 103, 0, 0, 161, 162, 5, 32, 0, 0, 162, 163, 5, 117, 0, 0, 163,
		164, 5, 110, 0, 0, 164, 165, 5, 98, 0, 0, 165, 166, 5, 111, 0, 0, 166,
		167, 5, 117, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169, 5, 100, 0, 0, 169,
		170, 5, 101, 0, 0, 170, 171, 5, 100, 0, 0, 171, 172, 5, 32, 0, 0, 172,
		173, 5, 111, 0, 0, 173, 174, 5, 118, 0, 0, 174, 175, 5, 101, 0, 0, 175,
		176, 5, 114, 0, 0, 176, 177, 5, 100, 0, 0, 177, 178, 5, 114, 0, 0, 178,
		179, 5, 97, 0, 0, 179, 180, 5, 102, 0, 0, 180, 181, 5, 116, 0, 0, 181,
		8, 1, 0, 0, 0, 182, 183, 5, 44, 0, 0, 183, 10, 1, 0, 0, 0, 184, 186, 7,
		0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 185, 1, 0, 0,
		0, 187, 188, 1, 0, 0, 0, 188, 12, 1, 0, 0, 0, 189, 191, 7, 1, 0, 0, 190,
		189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 6, 6, 0, 0, 195, 14, 1, 0,
		0, 0, 196, 197, 5, 47, 0, 0, 197, 198, 5, 42, 0, 0, 198, 203, 1, 0, 0,
		0, 199, 202, 3, 15, 7, 0, 200, 202, 9, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201,
		200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 203, 201,
		1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207, 5, 42,
		0, 0, 207, 208, 5, 47, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 6, 7, 0, 0,
		210, 16, 1, 0, 0, 0, 211, 212, 5, 47, 0, 0, 212, 213, 5, 47, 0, 0, 213,
		217, 1, 0, 0, 0, 214, 216, 9, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219,
		1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 220, 1, 0,
		0, 0, 219, 217, 1, 0, 0, 0, 220, 221, 3, 11, 5, 0, 221, 222, 1, 0, 0, 0,
		222, 223, 6, 8, 0, 0, 223, 18, 1, 0, 0, 0, 224, 225, 5, 118, 0, 0, 225,
		226, 5, 97, 0, 0, 226, 227, 5, 114, 0, 0, 227, 228, 5, 115, 0, 0, 228,
		20, 1, 0, 0, 0, 229, 230, 5, 109, 0, 0, 230, 231, 5, 101, 0, 0, 231, 232,
		5, 116, 0, 0, 232, 233, 5, 97, 0, 0, 233, 22, 1, 0, 0, 0, 234, 235, 5,
		115, 0, 0, 235, 236, 5, 101, 0, 0, 236, 237, 5, 116, 0, 0, 237, 238, 5,
		95, 0, 0, 238, 239, 5, 116, 0, 0, 239, 240, 5, 120, 0, 0, 240, 241, 5,
		95, 0, 0, 241, 242, 5, 109, 0, 0, 242, 243, 5, 101, 0, 0, 243, 244, 5,
		116, 0, 0, 244, 245, 5, 97, 0, 0, 245, 24, 1, 0, 0, 0, 246, 247, 5, 112,
		0, 0, 247, 248, 5, 114, 0, 0, 248, 249, 5, 105, 0, 0, 249, 250, 5, 110,
		0, 0, 250, 251, 5, 116, 0, 0, 251, 26, 1, 0, 0, 0, 252, 253, 5, 102, 0,
		0, 253, 254, 5, 97, 0, 0, 254, 255, 5, 105, 0, 0, 255, 256, 5, 108, 0,
		0, 256, 28, 1, 0, 0, 0, 257, 258, 5, 115, 0, 0, 258, 259, 5, 101, 0, 0,
		259, 260, 5, 110, 0, 0, 260, 261, 5, 100, 0, 0, 261, 30, 1, 0, 0, 0, 262,
		263, 5, 115, 0, 0, 263, 264, 5, 111, 0, 0, 264, 265, 5, 117, 0, 0, 265,
		266, 5, 114, 0, 0, 266, 267, 5, 99, 0, 0, 267, 268, 5, 101, 0, 0, 268,
		32, 1, 0, 0, 0, 269, 270, 5, 102, 0, 0, 270, 271, 5, 114, 0, 0, 271, 272,
		5, 111, 0, 0, 272, 273, 5, 109, 0, 0, 273, 34, 1, 0, 0, 0, 274, 275, 5,
		109, 0, 0, 275, 276, 5, 97, 0, 0, 276, 277, 5, 120, 0, 0, 277, 36, 1, 0,
		0, 0, 278, 279, 5, 100, 0, 0, 279, 280, 5, 101, 0, 0, 280, 281, 5, 115,
		0, 0, 281, 282, 5, 116, 0, 0, 282, 283, 5, 105, 0, 0, 283, 284, 5, 110,
		0, 0, 284, 285, 5, 97, 0, 0, 285, 286, 5, 116, 0, 0, 286, 287, 5, 105,
		0, 0, 287, 288, 5, 111, 0, 0, 288, 289, 5, 110, 0, 0, 289, 38, 1, 0, 0,
		0, 290, 291, 5, 116, 0, 0, 291, 292, 5, 111, 0, 0, 292, 40, 1, 0, 0, 0,
		293, 294, 5, 97, 0, 0, 294, 295, 5, 108, 0, 0, 295, 296, 5, 108, 0, 0,
		296, 297, 5, 111, 0, 0, 297, 298, 5, 99, 0, 0, 298, 299, 5, 97, 0, 0, 299,
		300, 5, 116, 0, 0, 300, 301, 5, 101, 0, 0, 301, 42, 1, 0, 0, 0, 302, 303,
		5, 43, 0, 0, 303, 44, 1, 0, 0, 0, 304, 305, 5, 45, 0, 0, 305, 46, 1, 0,
		0, 0, 306, 307, 5, 40, 0, 0, 307, 48, 1, 0, 0, 0, 308, 309, 5, 41, 0, 0,
		309, 50, 1, 0, 0, 0, 310, 311, 5, 91, 0, 0, 311, 52, 1, 0, 0, 0, 312, 313,
		5, 93, 0, 0, 313, 54, 1, 0, 0, 0, 314, 315, 5, 123, 0, 0, 315, 56, 1, 0,
		0, 0, 316, 317, 5, 125, 0, 0, 317, 58, 1, 0, 0, 0, 318, 319, 5, 61, 0,
		0, 319, 60, 1, 0, 0, 0, 320, 321, 5, 97, 0, 0, 321, 322, 5, 99, 0, 0, 322,
		323, 5, 99, 0, 0, 323, 324, 5, 111, 0, 0, 324, 325, 5, 117, 0, 0, 325,
		326, 5, 110, 0, 0, 326, 327, 5, 116, 0, 0, 327, 62, 1, 0, 0, 0, 328, 329,
		5, 97, 0, 0, 329, 330, 5, 115, 0, 0, 330, 331, 5, 115, 0, 0, 331, 332,
		5, 101, 0, 0, 332, 333, 5, 116, 0, 0, 333, 64, 1, 0, 0, 0, 334, 335, 5,
		110, 0, 0, 335, 336, 5, 117, 0, 0, 336, 337, 5, 109, 0, 0, 337, 338, 5,
		98, 0, 0, 338, 339, 5, 101, 0, 0, 339, 340, 5, 114, 0, 0, 340, 66, 1, 0,
		0, 0, 341, 342, 5, 109, 0, 0, 342, 343, 5, 111, 0, 0, 343, 344, 5, 110,
		0, 0, 344, 345, 5, 101, 0, 0, 345, 346, 5, 116, 0, 0, 346, 347, 5, 97,
		0, 0, 347, 348, 5, 114, 0, 0, 348, 349, 5, 121, 0, 0, 349, 68, 1, 0, 0,
		0, 350, 351, 5, 112, 0, 0, 351, 352, 5, 111, 0, 0, 352, 353, 5, 114, 0,
		0, 353, 354, 5, 116, 0, 0, 354, 355, 5, 105, 0, 0, 355, 356, 5, 111, 0,
		0, 356, 357, 5, 110, 0, 0, 357, 70, 1, 0, 0, 0, 358, 359, 5, 115, 0, 0,
		359, 360, 5, 116, 0, 0, 360, 361, 5, 114, 0, 0, 361, 362, 5, 105, 0, 0,
		362, 363, 5, 110, 0, 0, 363, 364, 5, 103, 0, 0, 364, 72, 1, 0, 0, 0, 365,
		369, 5, 34, 0, 0, 366, 368, 7, 2, 0, 0, 367, 366, 1, 0, 0, 0, 368, 371,
		1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 372, 1, 0,
		0, 0, 371, 369, 1, 0, 0, 0, 372, 373, 5, 34, 0, 0, 373, 74, 1, 0, 0, 0,
		374, 376, 7, 3, 0, 0, 375, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377,
		375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 380, 1, 0, 0, 0, 379, 381,
		7, 4, 0, 0, 380, 379, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 1, 0,
		0, 0, 382, 384, 5, 47, 0, 0, 383, 385, 7, 4, 0, 0, 384, 383, 1, 0, 0, 0,
		384, 385, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 388, 7, 3, 0, 0, 387,
		386, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 390,
		1, 0, 0, 0, 390, 406, 1, 0, 0, 0, 391, 393, 7, 3, 0, 0, 392, 391, 1, 0,
		0, 0, 393, 394, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0,
		395, 402, 1, 0, 0, 0, 396, 398, 5, 46, 0, 0, 397, 399, 7, 3, 0, 0, 398,
		397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 400, 401,
		1, 0, 0, 0, 401, 403, 1, 0, 0, 0, 402, 396, 1, 0, 0, 0, 402, 403, 1, 0,
		0, 0, 403, 404, 1, 0, 0, 0, 404, 406, 5, 37, 0, 0, 405, 375, 1, 0, 0, 0,
		405, 392, 1, 0, 0, 0, 406, 76, 1, 0, 0, 0, 407, 408, 5, 114, 0, 0, 408,
		409, 5, 101, 0, 0, 409, 410, 5, 109, 0, 0, 410, 411, 5, 97, 0, 0, 411,
		412, 5, 105, 0, 0, 412, 413, 5, 110, 0, 0, 413, 414, 5, 105, 0, 0, 414,
		415, 5, 110, 0, 0, 415, 416, 5, 103, 0, 0, 416, 78, 1, 0, 0, 0, 417, 418,
		5, 107, 0, 0, 418, 419, 5, 101, 0, 0, 419, 420, 5, 112, 0, 0, 420, 421,
		5, 116, 0, 0, 421, 80, 1, 0, 0, 0, 422, 424, 7, 3, 0, 0, 423, 422, 1, 0,
		0, 0, 424, 425, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0,
		426, 82, 1, 0, 0, 0, 427, 428, 5, 37, 0, 0, 428, 84, 1, 0, 0, 0, 429, 431,
		5, 36, 0, 0, 430, 432, 7, 5, 0, 0, 431, 430, 1, 0, 0, 0, 432, 433, 1, 0,
		0, 0, 433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 438, 1, 0, 0, 0,
		435, 437, 7, 6, 0, 0, 436, 435, 1, 0, 0, 0, 437, 440, 1, 0, 0, 0, 438,
		436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 86, 1, 0, 0, 0, 440, 438, 1,
		0, 0, 0, 441, 443, 5, 64, 0, 0, 442, 444, 7, 7, 0, 0, 443, 442, 1, 0, 0,
		0, 444, 445, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446,
		450, 1, 0, 0, 0, 447, 449, 7, 8, 0, 0, 448, 447, 1, 0, 0, 0, 449, 452,
		1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 88, 1, 0,
		0, 0, 452, 450, 1, 0, 0, 0, 453, 455, 7, 9, 0, 0, 454, 453, 1, 0, 0, 0,
		455, 456, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457,
		90, 1, 0, 0, 0, 21, 0, 187, 192, 201, 203, 217, 369, 377, 380, 384, 389,
		394, 400, 402, 405, 425, 433, 438, 445, 450, 456, 1, 6, 0, 0,
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

// NumScriptLexerInit initializes any static state used to implement NumScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumScriptLexerInit() {
	staticData := &numscriptlexerLexerStaticData
	staticData.once.Do(numscriptlexerLexerInit)
}

// NewNumScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumScriptLexer(input antlr.CharStream) *NumScriptLexer {
	NumScriptLexerInit()
	l := new(NumScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &numscriptlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "NumScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumScriptLexer tokens.
const (
	NumScriptLexerT__0              = 1
	NumScriptLexerT__1              = 2
	NumScriptLexerT__2              = 3
	NumScriptLexerT__3              = 4
	NumScriptLexerT__4              = 5
	NumScriptLexerNEWLINE           = 6
	NumScriptLexerWHITESPACE        = 7
	NumScriptLexerMULTILINE_COMMENT = 8
	NumScriptLexerLINE_COMMENT      = 9
	NumScriptLexerVARS              = 10
	NumScriptLexerMETA              = 11
	NumScriptLexerSET_TX_META       = 12
	NumScriptLexerPRINT             = 13
	NumScriptLexerFAIL              = 14
	NumScriptLexerSEND              = 15
	NumScriptLexerSOURCE            = 16
	NumScriptLexerFROM              = 17
	NumScriptLexerMAX               = 18
	NumScriptLexerDESTINATION       = 19
	NumScriptLexerTO                = 20
	NumScriptLexerALLOCATE          = 21
	NumScriptLexerOP_ADD            = 22
	NumScriptLexerOP_SUB            = 23
	NumScriptLexerLPAREN            = 24
	NumScriptLexerRPAREN            = 25
	NumScriptLexerLBRACK            = 26
	NumScriptLexerRBRACK            = 27
	NumScriptLexerLBRACE            = 28
	NumScriptLexerRBRACE            = 29
	NumScriptLexerEQ                = 30
	NumScriptLexerTY_ACCOUNT        = 31
	NumScriptLexerTY_ASSET          = 32
	NumScriptLexerTY_NUMBER         = 33
	NumScriptLexerTY_MONETARY       = 34
	NumScriptLexerTY_PORTION        = 35
	NumScriptLexerTY_STRING         = 36
	NumScriptLexerSTRING            = 37
	NumScriptLexerPORTION           = 38
	NumScriptLexerREMAINING         = 39
	NumScriptLexerKEPT              = 40
	NumScriptLexerNUMBER            = 41
	NumScriptLexerPERCENT           = 42
	NumScriptLexerVARIABLE_NAME     = 43
	NumScriptLexerACCOUNT           = 44
	NumScriptLexerASSET             = 45
)

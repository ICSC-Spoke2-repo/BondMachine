// Code generated from Nnef_flat.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type Nnef_flatLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var nnef_flatlexerLexerStaticData struct {
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

func nnef_flatlexerLexerInit() {
	staticData := &nnef_flatlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'version'", "';'", "'graph'", "'('", "')'", "'->'", "','", "'{'",
		"'}'", "'='", "'<'", "'>'", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TYPE_NAME",
		"IDENTIFIER", "FLOAT", "STRING_LITERAL", "LOGICAL_LITERAL", "NUMERIC_LITERAL",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "TYPE_NAME", "IDENTIFIER",
		"FLOAT", "STRING_LITERAL", "LOGICAL_LITERAL", "NUMERIC_LITERAL", "WHITESPACE",
		"LETTER", "ALPHA", "DIGIT", "EXPONENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 201, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 103, 8, 14, 1, 15, 1, 15, 5, 15, 107, 8,
		15, 10, 15, 12, 15, 110, 9, 15, 1, 16, 4, 16, 113, 8, 16, 11, 16, 12, 16,
		114, 1, 16, 1, 16, 4, 16, 119, 8, 16, 11, 16, 12, 16, 120, 1, 17, 1, 17,
		4, 17, 125, 8, 17, 11, 17, 12, 17, 126, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 140, 8, 18, 1, 19,
		4, 19, 143, 8, 19, 11, 19, 12, 19, 144, 1, 19, 1, 19, 5, 19, 149, 8, 19,
		10, 19, 12, 19, 152, 9, 19, 1, 19, 3, 19, 155, 8, 19, 1, 19, 1, 19, 4,
		19, 159, 8, 19, 11, 19, 12, 19, 160, 1, 19, 3, 19, 164, 8, 19, 1, 19, 4,
		19, 167, 8, 19, 11, 19, 12, 19, 168, 1, 19, 1, 19, 1, 19, 4, 19, 174, 8,
		19, 11, 19, 12, 19, 175, 3, 19, 178, 8, 19, 1, 20, 4, 20, 181, 8, 20, 11,
		20, 12, 20, 182, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 3, 24, 195, 8, 24, 1, 24, 4, 24, 198, 8, 24, 11, 24, 12,
		24, 199, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 0, 45, 0, 47, 0, 49, 0, 1, 0, 5, 3, 0,
		9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 215,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5,
		61, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0,
		0, 13, 74, 1, 0, 0, 0, 15, 76, 1, 0, 0, 0, 17, 78, 1, 0, 0, 0, 19, 80,
		1, 0, 0, 0, 21, 82, 1, 0, 0, 0, 23, 84, 1, 0, 0, 0, 25, 86, 1, 0, 0, 0,
		27, 88, 1, 0, 0, 0, 29, 102, 1, 0, 0, 0, 31, 104, 1, 0, 0, 0, 33, 112,
		1, 0, 0, 0, 35, 122, 1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 177, 1, 0, 0,
		0, 41, 180, 1, 0, 0, 0, 43, 186, 1, 0, 0, 0, 45, 188, 1, 0, 0, 0, 47, 190,
		1, 0, 0, 0, 49, 192, 1, 0, 0, 0, 51, 52, 5, 118, 0, 0, 52, 53, 5, 101,
		0, 0, 53, 54, 5, 114, 0, 0, 54, 55, 5, 115, 0, 0, 55, 56, 5, 105, 0, 0,
		56, 57, 5, 111, 0, 0, 57, 58, 5, 110, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60,
		5, 59, 0, 0, 60, 4, 1, 0, 0, 0, 61, 62, 5, 103, 0, 0, 62, 63, 5, 114, 0,
		0, 63, 64, 5, 97, 0, 0, 64, 65, 5, 112, 0, 0, 65, 66, 5, 104, 0, 0, 66,
		6, 1, 0, 0, 0, 67, 68, 5, 40, 0, 0, 68, 8, 1, 0, 0, 0, 69, 70, 5, 41, 0,
		0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 45, 0, 0, 72, 73, 5, 62, 0, 0, 73, 12,
		1, 0, 0, 0, 74, 75, 5, 44, 0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 5, 123, 0,
		0, 77, 16, 1, 0, 0, 0, 78, 79, 5, 125, 0, 0, 79, 18, 1, 0, 0, 0, 80, 81,
		5, 61, 0, 0, 81, 20, 1, 0, 0, 0, 82, 83, 5, 60, 0, 0, 83, 22, 1, 0, 0,
		0, 84, 85, 5, 62, 0, 0, 85, 24, 1, 0, 0, 0, 86, 87, 5, 91, 0, 0, 87, 26,
		1, 0, 0, 0, 88, 89, 5, 93, 0, 0, 89, 28, 1, 0, 0, 0, 90, 91, 5, 115, 0,
		0, 91, 92, 5, 99, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5, 108, 0, 0, 94,
		95, 5, 97, 0, 0, 95, 103, 5, 114, 0, 0, 96, 97, 5, 116, 0, 0, 97, 98, 5,
		101, 0, 0, 98, 99, 5, 110, 0, 0, 99, 100, 5, 115, 0, 0, 100, 101, 5, 111,
		0, 0, 101, 103, 5, 114, 0, 0, 102, 90, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0,
		103, 30, 1, 0, 0, 0, 104, 108, 3, 43, 21, 0, 105, 107, 3, 45, 22, 0, 106,
		105, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 32, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 113, 3, 47,
		23, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 5, 46, 0, 0, 117,
		119, 3, 47, 23, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 34, 1, 0, 0, 0, 122, 124, 5, 39,
		0, 0, 123, 125, 3, 45, 22, 0, 124, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0,
		0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 5, 39, 0, 0, 129, 36, 1, 0, 0, 0, 130, 131, 5, 116, 0, 0, 131, 132,
		5, 114, 0, 0, 132, 133, 5, 117, 0, 0, 133, 140, 5, 101, 0, 0, 134, 135,
		5, 102, 0, 0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138,
		5, 115, 0, 0, 138, 140, 5, 101, 0, 0, 139, 130, 1, 0, 0, 0, 139, 134, 1,
		0, 0, 0, 140, 38, 1, 0, 0, 0, 141, 143, 3, 47, 23, 0, 142, 141, 1, 0, 0,
		0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 150, 5, 46, 0, 0, 147, 149, 3, 47, 23, 0, 148, 147,
		1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 3, 49, 24,
		0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 178, 1, 0, 0, 0, 156,
		158, 5, 46, 0, 0, 157, 159, 3, 47, 23, 0, 158, 157, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0,
		0, 0, 162, 164, 3, 49, 24, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0,
		0, 164, 178, 1, 0, 0, 0, 165, 167, 3, 47, 23, 0, 166, 165, 1, 0, 0, 0,
		167, 168, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 171, 3, 49, 24, 0, 171, 178, 1, 0, 0, 0, 172, 174,
		3, 47, 23, 0, 173, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1,
		0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 142, 1, 0, 0,
		0, 177, 156, 1, 0, 0, 0, 177, 166, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 178,
		40, 1, 0, 0, 0, 179, 181, 7, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 182, 1,
		0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0,
		0, 184, 185, 6, 20, 0, 0, 185, 42, 1, 0, 0, 0, 186, 187, 7, 1, 0, 0, 187,
		44, 1, 0, 0, 0, 188, 189, 7, 2, 0, 0, 189, 46, 1, 0, 0, 0, 190, 191, 2,
		48, 57, 0, 191, 48, 1, 0, 0, 0, 192, 194, 7, 3, 0, 0, 193, 195, 7, 4, 0,
		0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196,
		198, 2, 48, 57, 0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197,
		1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 50, 1, 0, 0, 0, 18, 0, 102, 108,
		114, 120, 126, 139, 144, 150, 154, 160, 163, 168, 175, 177, 182, 194, 199,
		1, 6, 0, 0,
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

// Nnef_flatLexerInit initializes any static state used to implement Nnef_flatLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNnef_flatLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Nnef_flatLexerInit() {
	staticData := &nnef_flatlexerLexerStaticData
	staticData.once.Do(nnef_flatlexerLexerInit)
}

// NewNnef_flatLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNnef_flatLexer(input antlr.CharStream) *Nnef_flatLexer {
	Nnef_flatLexerInit()
	l := new(Nnef_flatLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &nnef_flatlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Nnef_flat.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Nnef_flatLexer tokens.
const (
	Nnef_flatLexerT__0            = 1
	Nnef_flatLexerT__1            = 2
	Nnef_flatLexerT__2            = 3
	Nnef_flatLexerT__3            = 4
	Nnef_flatLexerT__4            = 5
	Nnef_flatLexerT__5            = 6
	Nnef_flatLexerT__6            = 7
	Nnef_flatLexerT__7            = 8
	Nnef_flatLexerT__8            = 9
	Nnef_flatLexerT__9            = 10
	Nnef_flatLexerT__10           = 11
	Nnef_flatLexerT__11           = 12
	Nnef_flatLexerT__12           = 13
	Nnef_flatLexerT__13           = 14
	Nnef_flatLexerTYPE_NAME       = 15
	Nnef_flatLexerIDENTIFIER      = 16
	Nnef_flatLexerFLOAT           = 17
	Nnef_flatLexerSTRING_LITERAL  = 18
	Nnef_flatLexerLOGICAL_LITERAL = 19
	Nnef_flatLexerNUMERIC_LITERAL = 20
	Nnef_flatLexerWHITESPACE      = 21
)

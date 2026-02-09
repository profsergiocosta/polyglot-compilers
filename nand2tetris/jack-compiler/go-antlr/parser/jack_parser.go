// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 277,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 67,
	10, 2, 12, 2, 14, 2, 70, 11, 2, 3, 2, 7, 2, 73, 10, 2, 12, 2, 14, 2, 76,
	11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 86, 10, 3,
	12, 3, 14, 3, 89, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 105, 10, 6, 3, 7, 3, 7, 3, 7,
	7, 7, 110, 10, 7, 12, 7, 14, 7, 113, 11, 7, 5, 7, 115, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 123, 10, 9, 12, 9, 14, 9, 126, 11, 9,
	3, 9, 3, 9, 3, 10, 7, 10, 131, 10, 10, 12, 10, 14, 10, 134, 11, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 141, 10, 11, 12, 11, 14, 11, 144,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 153, 10,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 166, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 175, 10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 180, 10, 17, 12, 17,
	14, 17, 183, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 190, 10,
	18, 12, 18, 14, 18, 193, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 205, 10, 20, 12, 20, 14, 20, 208,
	11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22,
	218, 10, 22, 3, 22, 3, 22, 3, 23, 5, 23, 223, 10, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 7, 25, 236,
	10, 25, 12, 25, 14, 25, 239, 11, 25, 5, 25, 241, 10, 25, 3, 26, 3, 26,
	7, 26, 245, 10, 26, 12, 26, 14, 26, 248, 11, 26, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 269, 10, 28, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 2, 2, 32, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 2, 8, 4, 2, 23, 23, 32, 32, 5, 2, 24, 25, 34, 34, 43, 43, 5,
	2, 22, 22, 31, 31, 41, 41, 4, 2, 3, 8, 10, 12, 5, 2, 26, 27, 37, 37, 42,
	42, 4, 2, 4, 4, 9, 9, 2, 276, 2, 62, 3, 2, 2, 2, 4, 80, 3, 2, 2, 2, 6,
	92, 3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10, 104, 3, 2, 2, 2, 12, 114, 3, 2,
	2, 2, 14, 116, 3, 2, 2, 2, 16, 119, 3, 2, 2, 2, 18, 132, 3, 2, 2, 2, 20,
	135, 3, 2, 2, 2, 22, 152, 3, 2, 2, 2, 24, 154, 3, 2, 2, 2, 26, 160, 3,
	2, 2, 2, 28, 167, 3, 2, 2, 2, 30, 169, 3, 2, 2, 2, 32, 176, 3, 2, 2, 2,
	34, 186, 3, 2, 2, 2, 36, 196, 3, 2, 2, 2, 38, 201, 3, 2, 2, 2, 40, 211,
	3, 2, 2, 2, 42, 215, 3, 2, 2, 2, 44, 222, 3, 2, 2, 2, 46, 229, 3, 2, 2,
	2, 48, 240, 3, 2, 2, 2, 50, 242, 3, 2, 2, 2, 52, 249, 3, 2, 2, 2, 54, 268,
	3, 2, 2, 2, 56, 270, 3, 2, 2, 2, 58, 272, 3, 2, 2, 2, 60, 274, 3, 2, 2,
	2, 62, 63, 7, 36, 2, 2, 63, 64, 5, 56, 29, 2, 64, 68, 7, 18, 2, 2, 65,
	67, 5, 4, 3, 2, 66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2,
	2, 68, 69, 3, 2, 2, 2, 69, 74, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 73,
	5, 8, 5, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 78, 7,
	19, 2, 2, 78, 79, 7, 2, 2, 3, 79, 3, 3, 2, 2, 2, 80, 81, 9, 2, 2, 2, 81,
	82, 5, 6, 4, 2, 82, 87, 5, 58, 30, 2, 83, 84, 7, 14, 2, 2, 84, 86, 5, 58,
	30, 2, 85, 83, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87,
	88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 91, 7, 15,
	2, 2, 91, 5, 3, 2, 2, 2, 92, 93, 9, 3, 2, 2, 93, 7, 3, 2, 2, 2, 94, 95,
	9, 4, 2, 2, 95, 96, 5, 10, 6, 2, 96, 97, 5, 60, 31, 2, 97, 98, 7, 16, 2,
	2, 98, 99, 5, 12, 7, 2, 99, 100, 7, 17, 2, 2, 100, 101, 5, 16, 9, 2, 101,
	9, 3, 2, 2, 2, 102, 105, 7, 35, 2, 2, 103, 105, 5, 6, 4, 2, 104, 102, 3,
	2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 11, 3, 2, 2, 2, 106, 111, 5, 14, 8,
	2, 107, 108, 7, 14, 2, 2, 108, 110, 5, 14, 8, 2, 109, 107, 3, 2, 2, 2,
	110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 106, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 117, 5, 6, 4, 2, 117, 118, 5, 58,
	30, 2, 118, 15, 3, 2, 2, 2, 119, 120, 7, 18, 2, 2, 120, 124, 5, 18, 10,
	2, 121, 123, 5, 22, 12, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2,
	124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 128, 7, 19, 2, 2, 128, 17, 3, 2, 2, 2, 129, 131,
	5, 20, 11, 2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3,
	2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 19, 3, 2, 2, 2, 134, 132, 3, 2, 2,
	2, 135, 136, 7, 33, 2, 2, 136, 137, 5, 6, 4, 2, 137, 142, 5, 58, 30, 2,
	138, 139, 7, 14, 2, 2, 139, 141, 5, 58, 30, 2, 140, 138, 3, 2, 2, 2, 141,
	144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145,
	3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 15, 2, 2, 146, 21, 3, 2,
	2, 2, 147, 153, 5, 24, 13, 2, 148, 153, 5, 30, 16, 2, 149, 153, 5, 36,
	19, 2, 150, 153, 5, 40, 21, 2, 151, 153, 5, 42, 22, 2, 152, 147, 3, 2,
	2, 2, 152, 148, 3, 2, 2, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2,
	152, 151, 3, 2, 2, 2, 153, 23, 3, 2, 2, 2, 154, 155, 7, 28, 2, 2, 155,
	156, 5, 26, 14, 2, 156, 157, 7, 12, 2, 2, 157, 158, 5, 28, 15, 2, 158,
	159, 7, 15, 2, 2, 159, 25, 3, 2, 2, 2, 160, 165, 5, 58, 30, 2, 161, 162,
	7, 20, 2, 2, 162, 163, 5, 50, 26, 2, 163, 164, 7, 21, 2, 2, 164, 166, 3,
	2, 2, 2, 165, 161, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 27, 3, 2, 2,
	2, 167, 168, 5, 50, 26, 2, 168, 29, 3, 2, 2, 2, 169, 170, 7, 29, 2, 2,
	170, 171, 7, 16, 2, 2, 171, 172, 5, 50, 26, 2, 172, 174, 5, 32, 17, 2,
	173, 175, 5, 34, 18, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175,
	31, 3, 2, 2, 2, 176, 177, 7, 17, 2, 2, 177, 181, 7, 18, 2, 2, 178, 180,
	5, 22, 12, 2, 179, 178, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3,
	2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184, 3, 2, 2, 2, 183, 181, 3, 2, 2,
	2, 184, 185, 7, 19, 2, 2, 185, 33, 3, 2, 2, 2, 186, 187, 7, 39, 2, 2, 187,
	191, 7, 18, 2, 2, 188, 190, 5, 22, 12, 2, 189, 188, 3, 2, 2, 2, 190, 193,
	3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2,
	2, 2, 193, 191, 3, 2, 2, 2, 194, 195, 7, 19, 2, 2, 195, 35, 3, 2, 2, 2,
	196, 197, 7, 30, 2, 2, 197, 198, 7, 16, 2, 2, 198, 199, 5, 50, 26, 2, 199,
	200, 5, 38, 20, 2, 200, 37, 3, 2, 2, 2, 201, 202, 7, 17, 2, 2, 202, 206,
	7, 18, 2, 2, 203, 205, 5, 22, 12, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3,
	2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2,
	2, 208, 206, 3, 2, 2, 2, 209, 210, 7, 19, 2, 2, 210, 39, 3, 2, 2, 2, 211,
	212, 7, 38, 2, 2, 212, 213, 5, 44, 23, 2, 213, 214, 7, 15, 2, 2, 214, 41,
	3, 2, 2, 2, 215, 217, 7, 40, 2, 2, 216, 218, 5, 50, 26, 2, 217, 216, 3,
	2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 7, 15, 2,
	2, 220, 43, 3, 2, 2, 2, 221, 223, 5, 46, 24, 2, 222, 221, 3, 2, 2, 2, 222,
	223, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 5, 60, 31, 2, 225, 226,
	7, 16, 2, 2, 226, 227, 5, 48, 25, 2, 227, 228, 7, 17, 2, 2, 228, 45, 3,
	2, 2, 2, 229, 230, 7, 43, 2, 2, 230, 231, 7, 13, 2, 2, 231, 47, 3, 2, 2,
	2, 232, 237, 5, 50, 26, 2, 233, 234, 7, 14, 2, 2, 234, 236, 5, 50, 26,
	2, 235, 233, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237,
	238, 3, 2, 2, 2, 238, 241, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 232,
	3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 49, 3, 2, 2, 2, 242, 246, 5, 54,
	28, 2, 243, 245, 5, 52, 27, 2, 244, 243, 3, 2, 2, 2, 245, 248, 3, 2, 2,
	2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 51, 3, 2, 2, 2, 248,
	246, 3, 2, 2, 2, 249, 250, 9, 5, 2, 2, 250, 251, 5, 54, 28, 2, 251, 53,
	3, 2, 2, 2, 252, 269, 7, 45, 2, 2, 253, 269, 7, 44, 2, 2, 254, 269, 9,
	6, 2, 2, 255, 269, 5, 58, 30, 2, 256, 257, 5, 58, 30, 2, 257, 258, 7, 20,
	2, 2, 258, 259, 5, 50, 26, 2, 259, 260, 7, 21, 2, 2, 260, 269, 3, 2, 2,
	2, 261, 269, 5, 44, 23, 2, 262, 263, 7, 16, 2, 2, 263, 264, 5, 50, 26,
	2, 264, 265, 7, 17, 2, 2, 265, 269, 3, 2, 2, 2, 266, 267, 9, 7, 2, 2, 267,
	269, 5, 54, 28, 2, 268, 252, 3, 2, 2, 2, 268, 253, 3, 2, 2, 2, 268, 254,
	3, 2, 2, 2, 268, 255, 3, 2, 2, 2, 268, 256, 3, 2, 2, 2, 268, 261, 3, 2,
	2, 2, 268, 262, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 55, 3, 2, 2, 2,
	270, 271, 7, 43, 2, 2, 271, 57, 3, 2, 2, 2, 272, 273, 7, 43, 2, 2, 273,
	59, 3, 2, 2, 2, 274, 275, 7, 43, 2, 2, 275, 61, 3, 2, 2, 2, 23, 68, 74,
	87, 104, 111, 114, 124, 132, 142, 152, 165, 174, 181, 191, 206, 217, 222,
	237, 240, 246, 268,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'~'", "'<'", "'>'", "'='",
	"'.'", "','", "';'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'method'",
	"'static'", "'int'", "'boolean'", "'true'", "'null'", "'let'", "'if'",
	"'while'", "'constructor'", "'field'", "'var'", "'char'", "'void'", "'class'",
	"'false'", "'do'", "'else'", "'return'", "'function'", "'this'",
}
var symbolicNames = []string{
	"", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT", "LT", "GT",
	"EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
	"LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN", "TRUE", "NULL",
	"LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", "CHAR", "VOID", "CLASS",
	"FALSE", "DO", "ELSE", "RETURN", "FUNCTION", "THIS", "ID", "STRING", "INTEGER",
	"COMMENT", "LINE_COMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"classdef", "classvardec", "atype", "subrotinedec", "subroutinetype", "parameterList",
	"parameter", "subroutinebody", "vardecs", "vardec", "statement", "letStatement",
	"lvalue", "rvalue", "ifStatement", "ifBlock", "elseBlock", "whileStatement",
	"whileBlock", "doStatement", "returnStatement", "subroutinecall", "classObject",
	"expressionlist", "expression", "binaryoperation", "term", "classname",
	"varname", "subroutinename",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JackParser struct {
	*antlr.BaseParser
}

func NewJackParser(input antlr.TokenStream) *JackParser {
	this := new(JackParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jack.g4"

	return this
}

// JackParser tokens.
const (
	JackParserEOF          = antlr.TokenEOF
	JackParserPLUS         = 1
	JackParserMINUS        = 2
	JackParserASTERISK     = 3
	JackParserSLASH        = 4
	JackParserAND          = 5
	JackParserOR           = 6
	JackParserNOT          = 7
	JackParserLT           = 8
	JackParserGT           = 9
	JackParserEQ           = 10
	JackParserDOT          = 11
	JackParserCOMMA        = 12
	JackParserSEMICOLON    = 13
	JackParserLPAREN       = 14
	JackParserRPAREN       = 15
	JackParserLBRACE       = 16
	JackParserRBRACE       = 17
	JackParserLBRACKET     = 18
	JackParserRBRACKET     = 19
	JackParserMETHOD       = 20
	JackParserSTATIC       = 21
	JackParserINT          = 22
	JackParserBOOLEAN      = 23
	JackParserTRUE         = 24
	JackParserNULL         = 25
	JackParserLET          = 26
	JackParserIF           = 27
	JackParserWHILE        = 28
	JackParserCONSTRUCTOR  = 29
	JackParserFIELD        = 30
	JackParserVAR          = 31
	JackParserCHAR         = 32
	JackParserVOID         = 33
	JackParserCLASS        = 34
	JackParserFALSE        = 35
	JackParserDO           = 36
	JackParserELSE         = 37
	JackParserRETURN       = 38
	JackParserFUNCTION     = 39
	JackParserTHIS         = 40
	JackParserID           = 41
	JackParserSTRING       = 42
	JackParserINTEGER      = 43
	JackParserCOMMENT      = 44
	JackParserLINE_COMMENT = 45
	JackParserWHITESPACE   = 46
)

// JackParser rules.
const (
	JackParserRULE_classdef        = 0
	JackParserRULE_classvardec     = 1
	JackParserRULE_atype           = 2
	JackParserRULE_subrotinedec    = 3
	JackParserRULE_subroutinetype  = 4
	JackParserRULE_parameterList   = 5
	JackParserRULE_parameter       = 6
	JackParserRULE_subroutinebody  = 7
	JackParserRULE_vardecs         = 8
	JackParserRULE_vardec          = 9
	JackParserRULE_statement       = 10
	JackParserRULE_letStatement    = 11
	JackParserRULE_lvalue          = 12
	JackParserRULE_rvalue          = 13
	JackParserRULE_ifStatement     = 14
	JackParserRULE_ifBlock         = 15
	JackParserRULE_elseBlock       = 16
	JackParserRULE_whileStatement  = 17
	JackParserRULE_whileBlock      = 18
	JackParserRULE_doStatement     = 19
	JackParserRULE_returnStatement = 20
	JackParserRULE_subroutinecall  = 21
	JackParserRULE_classObject     = 22
	JackParserRULE_expressionlist  = 23
	JackParserRULE_expression      = 24
	JackParserRULE_binaryoperation = 25
	JackParserRULE_term            = 26
	JackParserRULE_classname       = 27
	JackParserRULE_varname         = 28
	JackParserRULE_subroutinename  = 29
)

// IClassdefContext is an interface to support dynamic dispatch.
type IClassdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassdefContext differentiates from other interfaces.
	IsClassdefContext()
}

type ClassdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassdefContext() *ClassdefContext {
	var p = new(ClassdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_classdef
	return p
}

func (*ClassdefContext) IsClassdefContext() {}

func NewClassdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassdefContext {
	var p = new(ClassdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_classdef

	return p
}

func (s *ClassdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassdefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(JackParserCLASS, 0)
}

func (s *ClassdefContext) Classname() IClassnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassnameContext)
}

func (s *ClassdefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *ClassdefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
}

func (s *ClassdefContext) EOF() antlr.TerminalNode {
	return s.GetToken(JackParserEOF, 0)
}

func (s *ClassdefContext) AllClassvardec() []IClassvardecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassvardecContext)(nil)).Elem())
	var tst = make([]IClassvardecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassvardecContext)
		}
	}

	return tst
}

func (s *ClassdefContext) Classvardec(i int) IClassvardecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassvardecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassvardecContext)
}

func (s *ClassdefContext) AllSubrotinedec() []ISubrotinedecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubrotinedecContext)(nil)).Elem())
	var tst = make([]ISubrotinedecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubrotinedecContext)
		}
	}

	return tst
}

func (s *ClassdefContext) Subrotinedec(i int) ISubrotinedecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubrotinedecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubrotinedecContext)
}

func (s *ClassdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterClassdef(s)
	}
}

func (s *ClassdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitClassdef(s)
	}
}

func (p *JackParser) Classdef() (localctx IClassdefContext) {
	localctx = NewClassdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JackParserRULE_classdef)
	var _la int

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
		p.SetState(60)
		p.Match(JackParserCLASS)
	}
	{
		p.SetState(61)
		p.Classname()
	}
	{
		p.SetState(62)
		p.Match(JackParserLBRACE)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JackParserSTATIC || _la == JackParserFIELD {
		{
			p.SetState(63)
			p.Classvardec()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(JackParserMETHOD-20))|(1<<(JackParserCONSTRUCTOR-20))|(1<<(JackParserFUNCTION-20)))) != 0 {
		{
			p.SetState(69)
			p.Subrotinedec()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(JackParserRBRACE)
	}
	{
		p.SetState(76)
		p.Match(JackParserEOF)
	}

	return localctx
}

// IClassvardecContext is an interface to support dynamic dispatch.
type IClassvardecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// IsClassvardecContext differentiates from other interfaces.
	IsClassvardecContext()
}

type ClassvardecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
}

func NewEmptyClassvardecContext() *ClassvardecContext {
	var p = new(ClassvardecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_classvardec
	return p
}

func (*ClassvardecContext) IsClassvardecContext() {}

func NewClassvardecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassvardecContext {
	var p = new(ClassvardecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_classvardec

	return p
}

func (s *ClassvardecContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassvardecContext) GetKind() antlr.Token { return s.kind }

func (s *ClassvardecContext) SetKind(v antlr.Token) { s.kind = v }

func (s *ClassvardecContext) Atype() IAtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *ClassvardecContext) AllVarname() []IVarnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarnameContext)(nil)).Elem())
	var tst = make([]IVarnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarnameContext)
		}
	}

	return tst
}

func (s *ClassvardecContext) Varname(i int) IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *ClassvardecContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *ClassvardecContext) STATIC() antlr.TerminalNode {
	return s.GetToken(JackParserSTATIC, 0)
}

func (s *ClassvardecContext) FIELD() antlr.TerminalNode {
	return s.GetToken(JackParserFIELD, 0)
}

func (s *ClassvardecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JackParserCOMMA)
}

func (s *ClassvardecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JackParserCOMMA, i)
}

func (s *ClassvardecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassvardecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassvardecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterClassvardec(s)
	}
}

func (s *ClassvardecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitClassvardec(s)
	}
}

func (p *JackParser) Classvardec() (localctx IClassvardecContext) {
	localctx = NewClassvardecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JackParserRULE_classvardec)
	var _la int

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
	p.SetState(78)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*ClassvardecContext).kind = _lt

	_la = p.GetTokenStream().LA(1)

	if !(_la == JackParserSTATIC || _la == JackParserFIELD) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*ClassvardecContext).kind = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(79)
		p.Atype()
	}
	{
		p.SetState(80)
		p.Varname()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JackParserCOMMA {
		{
			p.SetState(81)
			p.Match(JackParserCOMMA)
		}
		{
			p.SetState(82)
			p.Varname()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}

// IAtypeContext is an interface to support dynamic dispatch.
type IAtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtypeContext differentiates from other interfaces.
	IsAtypeContext()
}

type AtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtypeContext() *AtypeContext {
	var p = new(AtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_atype
	return p
}

func (*AtypeContext) IsAtypeContext() {}

func NewAtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtypeContext {
	var p = new(AtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_atype

	return p
}

func (s *AtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AtypeContext) INT() antlr.TerminalNode {
	return s.GetToken(JackParserINT, 0)
}

func (s *AtypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(JackParserCHAR, 0)
}

func (s *AtypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JackParserBOOLEAN, 0)
}

func (s *AtypeContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *AtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterAtype(s)
	}
}

func (s *AtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitAtype(s)
	}
}

func (p *JackParser) Atype() (localctx IAtypeContext) {
	localctx = NewAtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JackParserRULE_atype)
	var _la int

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
	p.SetState(90)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(JackParserINT-22))|(1<<(JackParserBOOLEAN-22))|(1<<(JackParserCHAR-22))|(1<<(JackParserID-22)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ISubrotinedecContext is an interface to support dynamic dispatch.
type ISubrotinedecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// IsSubrotinedecContext differentiates from other interfaces.
	IsSubrotinedecContext()
}

type SubrotinedecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
}

func NewEmptySubrotinedecContext() *SubrotinedecContext {
	var p = new(SubrotinedecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_subrotinedec
	return p
}

func (*SubrotinedecContext) IsSubrotinedecContext() {}

func NewSubrotinedecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubrotinedecContext {
	var p = new(SubrotinedecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_subrotinedec

	return p
}

func (s *SubrotinedecContext) GetParser() antlr.Parser { return s.parser }

func (s *SubrotinedecContext) GetKind() antlr.Token { return s.kind }

func (s *SubrotinedecContext) SetKind(v antlr.Token) { s.kind = v }

func (s *SubrotinedecContext) Subroutinetype() ISubroutinetypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinetypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinetypeContext)
}

func (s *SubrotinedecContext) Subroutinename() ISubroutinenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinenameContext)
}

func (s *SubrotinedecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserLPAREN, 0)
}

func (s *SubrotinedecContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *SubrotinedecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserRPAREN, 0)
}

func (s *SubrotinedecContext) Subroutinebody() ISubroutinebodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinebodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinebodyContext)
}

func (s *SubrotinedecContext) CONSTRUCTOR() antlr.TerminalNode {
	return s.GetToken(JackParserCONSTRUCTOR, 0)
}

func (s *SubrotinedecContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(JackParserFUNCTION, 0)
}

func (s *SubrotinedecContext) METHOD() antlr.TerminalNode {
	return s.GetToken(JackParserMETHOD, 0)
}

func (s *SubrotinedecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubrotinedecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubrotinedecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubrotinedec(s)
	}
}

func (s *SubrotinedecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubrotinedec(s)
	}
}

func (p *JackParser) Subrotinedec() (localctx ISubrotinedecContext) {
	localctx = NewSubrotinedecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JackParserRULE_subrotinedec)
	var _la int

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
	p.SetState(92)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*SubrotinedecContext).kind = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(JackParserMETHOD-20))|(1<<(JackParserCONSTRUCTOR-20))|(1<<(JackParserFUNCTION-20)))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*SubrotinedecContext).kind = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(93)
		p.Subroutinetype()
	}
	{
		p.SetState(94)
		p.Subroutinename()
	}
	{
		p.SetState(95)
		p.Match(JackParserLPAREN)
	}
	{
		p.SetState(96)
		p.ParameterList()
	}
	{
		p.SetState(97)
		p.Match(JackParserRPAREN)
	}
	{
		p.SetState(98)
		p.Subroutinebody()
	}

	return localctx
}

// ISubroutinetypeContext is an interface to support dynamic dispatch.
type ISubroutinetypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinetypeContext differentiates from other interfaces.
	IsSubroutinetypeContext()
}

type SubroutinetypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinetypeContext() *SubroutinetypeContext {
	var p = new(SubroutinetypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_subroutinetype
	return p
}

func (*SubroutinetypeContext) IsSubroutinetypeContext() {}

func NewSubroutinetypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinetypeContext {
	var p = new(SubroutinetypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_subroutinetype

	return p
}

func (s *SubroutinetypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinetypeContext) VOID() antlr.TerminalNode {
	return s.GetToken(JackParserVOID, 0)
}

func (s *SubroutinetypeContext) Atype() IAtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *SubroutinetypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinetypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinetypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubroutinetype(s)
	}
}

func (s *SubroutinetypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubroutinetype(s)
	}
}

func (p *JackParser) Subroutinetype() (localctx ISubroutinetypeContext) {
	localctx = NewSubroutinetypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JackParserRULE_subroutinetype)

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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JackParserVOID:
		{
			p.SetState(100)
			p.Match(JackParserVOID)
		}

	case JackParserINT, JackParserBOOLEAN, JackParserCHAR, JackParserID:
		{
			p.SetState(101)
			p.Atype()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JackParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JackParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *JackParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JackParserRULE_parameterList)
	var _la int

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
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(JackParserINT-22))|(1<<(JackParserBOOLEAN-22))|(1<<(JackParserCHAR-22))|(1<<(JackParserID-22)))) != 0 {
		{
			p.SetState(104)
			p.Parameter()
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JackParserCOMMA {
			{
				p.SetState(105)
				p.Match(JackParserCOMMA)
			}
			{
				p.SetState(106)
				p.Parameter()
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Atype() IAtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *ParameterContext) Varname() IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *JackParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JackParserRULE_parameter)

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
		p.SetState(114)
		p.Atype()
	}
	{
		p.SetState(115)
		p.Varname()
	}

	return localctx
}

// ISubroutinebodyContext is an interface to support dynamic dispatch.
type ISubroutinebodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinebodyContext differentiates from other interfaces.
	IsSubroutinebodyContext()
}

type SubroutinebodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinebodyContext() *SubroutinebodyContext {
	var p = new(SubroutinebodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_subroutinebody
	return p
}

func (*SubroutinebodyContext) IsSubroutinebodyContext() {}

func NewSubroutinebodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinebodyContext {
	var p = new(SubroutinebodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_subroutinebody

	return p
}

func (s *SubroutinebodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinebodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *SubroutinebodyContext) Vardecs() IVardecsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVardecsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVardecsContext)
}

func (s *SubroutinebodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
}

func (s *SubroutinebodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SubroutinebodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SubroutinebodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinebodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinebodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubroutinebody(s)
	}
}

func (s *SubroutinebodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubroutinebody(s)
	}
}

func (p *JackParser) Subroutinebody() (localctx ISubroutinebodyContext) {
	localctx = NewSubroutinebodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JackParserRULE_subroutinebody)
	var _la int

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
		p.SetState(117)
		p.Match(JackParserLBRACE)
	}
	{
		p.SetState(118)
		p.Vardecs()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(JackParserLET-26))|(1<<(JackParserIF-26))|(1<<(JackParserWHILE-26))|(1<<(JackParserDO-26))|(1<<(JackParserRETURN-26)))) != 0 {
		{
			p.SetState(119)
			p.Statement()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(JackParserRBRACE)
	}

	return localctx
}

// IVardecsContext is an interface to support dynamic dispatch.
type IVardecsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVardecsContext differentiates from other interfaces.
	IsVardecsContext()
}

type VardecsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVardecsContext() *VardecsContext {
	var p = new(VardecsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_vardecs
	return p
}

func (*VardecsContext) IsVardecsContext() {}

func NewVardecsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VardecsContext {
	var p = new(VardecsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_vardecs

	return p
}

func (s *VardecsContext) GetParser() antlr.Parser { return s.parser }

func (s *VardecsContext) AllVardec() []IVardecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVardecContext)(nil)).Elem())
	var tst = make([]IVardecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVardecContext)
		}
	}

	return tst
}

func (s *VardecsContext) Vardec(i int) IVardecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVardecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVardecContext)
}

func (s *VardecsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VardecsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VardecsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterVardecs(s)
	}
}

func (s *VardecsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitVardecs(s)
	}
}

func (p *JackParser) Vardecs() (localctx IVardecsContext) {
	localctx = NewVardecsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JackParserRULE_vardecs)
	var _la int

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
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JackParserVAR {
		{
			p.SetState(127)
			p.Vardec()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVardecContext is an interface to support dynamic dispatch.
type IVardecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVardecContext differentiates from other interfaces.
	IsVardecContext()
}

type VardecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVardecContext() *VardecContext {
	var p = new(VardecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_vardec
	return p
}

func (*VardecContext) IsVardecContext() {}

func NewVardecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VardecContext {
	var p = new(VardecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_vardec

	return p
}

func (s *VardecContext) GetParser() antlr.Parser { return s.parser }

func (s *VardecContext) VAR() antlr.TerminalNode {
	return s.GetToken(JackParserVAR, 0)
}

func (s *VardecContext) Atype() IAtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *VardecContext) AllVarname() []IVarnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarnameContext)(nil)).Elem())
	var tst = make([]IVarnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarnameContext)
		}
	}

	return tst
}

func (s *VardecContext) Varname(i int) IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *VardecContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *VardecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JackParserCOMMA)
}

func (s *VardecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JackParserCOMMA, i)
}

func (s *VardecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VardecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VardecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterVardec(s)
	}
}

func (s *VardecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitVardec(s)
	}
}

func (p *JackParser) Vardec() (localctx IVardecContext) {
	localctx = NewVardecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JackParserRULE_vardec)
	var _la int

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
		p.SetState(133)
		p.Match(JackParserVAR)
	}
	{
		p.SetState(134)
		p.Atype()
	}
	{
		p.SetState(135)
		p.Varname()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JackParserCOMMA {
		{
			p.SetState(136)
			p.Match(JackParserCOMMA)
		}
		{
			p.SetState(137)
			p.Varname()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) LetStatement() ILetStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILetStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) DoStatement() IDoStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *JackParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JackParserRULE_statement)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JackParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.LetStatement()
		}

	case JackParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.IfStatement()
		}

	case JackParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.WhileStatement()
		}

	case JackParserDO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.DoStatement()
		}

	case JackParserRETURN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.ReturnStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILetStatementContext is an interface to support dynamic dispatch.
type ILetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetStatementContext differentiates from other interfaces.
	IsLetStatementContext()
}

type LetStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetStatementContext() *LetStatementContext {
	var p = new(LetStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_letStatement
	return p
}

func (*LetStatementContext) IsLetStatementContext() {}

func NewLetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetStatementContext {
	var p = new(LetStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_letStatement

	return p
}

func (s *LetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LetStatementContext) LET() antlr.TerminalNode {
	return s.GetToken(JackParserLET, 0)
}

func (s *LetStatementContext) Lvalue() ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LetStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(JackParserEQ, 0)
}

func (s *LetStatementContext) Rvalue() IRvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *LetStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *LetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterLetStatement(s)
	}
}

func (s *LetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitLetStatement(s)
	}
}

func (p *JackParser) LetStatement() (localctx ILetStatementContext) {
	localctx = NewLetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JackParserRULE_letStatement)

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
		p.SetState(152)
		p.Match(JackParserLET)
	}
	{
		p.SetState(153)
		p.Lvalue()
	}
	{
		p.SetState(154)
		p.Match(JackParserEQ)
	}
	{
		p.SetState(155)
		p.Rvalue()
	}
	{
		p.SetState(156)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdent returns the ident rule contexts.
	GetIdent() IVarnameContext

	// SetIdent sets the ident rule contexts.
	SetIdent(IVarnameContext)

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ident  IVarnameContext
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) GetIdent() IVarnameContext { return s.ident }

func (s *LvalueContext) SetIdent(v IVarnameContext) { s.ident = v }

func (s *LvalueContext) Varname() IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *LvalueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACKET, 0)
}

func (s *LvalueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LvalueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACKET, 0)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *JackParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JackParserRULE_lvalue)
	var _la int

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
		p.SetState(158)

		var _x = p.Varname()

		localctx.(*LvalueContext).ident = _x
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JackParserLBRACKET {
		{
			p.SetState(159)
			p.Match(JackParserLBRACKET)
		}
		{
			p.SetState(160)
			p.Expression()
		}
		{
			p.SetState(161)
			p.Match(JackParserRBRACKET)
		}

	}

	return localctx
}

// IRvalueContext is an interface to support dynamic dispatch.
type IRvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRvalueContext differentiates from other interfaces.
	IsRvalueContext()
}

type RvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRvalueContext() *RvalueContext {
	var p = new(RvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_rvalue
	return p
}

func (*RvalueContext) IsRvalueContext() {}

func NewRvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RvalueContext {
	var p = new(RvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_rvalue

	return p
}

func (s *RvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *RvalueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterRvalue(s)
	}
}

func (s *RvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitRvalue(s)
	}
}

func (p *JackParser) Rvalue() (localctx IRvalueContext) {
	localctx = NewRvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JackParserRULE_rvalue)

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
		p.SetState(165)
		p.Expression()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(JackParserIF, 0)
}

func (s *IfStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserLPAREN, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) IfBlock() IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfStatementContext) ElseBlock() IElseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *JackParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JackParserRULE_ifStatement)
	var _la int

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
		p.SetState(167)
		p.Match(JackParserIF)
	}
	{
		p.SetState(168)
		p.Match(JackParserLPAREN)
	}
	{
		p.SetState(169)
		p.Expression()
	}
	{
		p.SetState(170)
		p.IfBlock()
	}

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JackParserELSE {
		{
			p.SetState(171)
			p.ElseBlock()
		}

	}

	return localctx
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_ifBlock
	return p
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserRPAREN, 0)
}

func (s *IfBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *IfBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
}

func (s *IfBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (p *JackParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JackParserRULE_ifBlock)
	var _la int

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
		p.SetState(174)
		p.Match(JackParserRPAREN)
	}
	{
		p.SetState(175)
		p.Match(JackParserLBRACE)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(JackParserLET-26))|(1<<(JackParserIF-26))|(1<<(JackParserWHILE-26))|(1<<(JackParserDO-26))|(1<<(JackParserRETURN-26)))) != 0 {
		{
			p.SetState(176)
			p.Statement()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(JackParserRBRACE)
	}

	return localctx
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_elseBlock
	return p
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(JackParserELSE, 0)
}

func (s *ElseBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *ElseBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
}

func (s *ElseBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ElseBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *JackParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JackParserRULE_elseBlock)
	var _la int

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
		p.SetState(184)
		p.Match(JackParserELSE)
	}
	{
		p.SetState(185)
		p.Match(JackParserLBRACE)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(JackParserLET-26))|(1<<(JackParserIF-26))|(1<<(JackParserWHILE-26))|(1<<(JackParserDO-26))|(1<<(JackParserRETURN-26)))) != 0 {
		{
			p.SetState(186)
			p.Statement()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
		p.Match(JackParserRBRACE)
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(JackParserWHILE, 0)
}

func (s *WhileStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserLPAREN, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) WhileBlock() IWhileBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (p *JackParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JackParserRULE_whileStatement)

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
		p.SetState(194)
		p.Match(JackParserWHILE)
	}
	{
		p.SetState(195)
		p.Match(JackParserLPAREN)
	}
	{
		p.SetState(196)
		p.Expression()
	}
	{
		p.SetState(197)
		p.WhileBlock()
	}

	return localctx
}

// IWhileBlockContext is an interface to support dynamic dispatch.
type IWhileBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileBlockContext differentiates from other interfaces.
	IsWhileBlockContext()
}

type WhileBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileBlockContext() *WhileBlockContext {
	var p = new(WhileBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_whileBlock
	return p
}

func (*WhileBlockContext) IsWhileBlockContext() {}

func NewWhileBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileBlockContext {
	var p = new(WhileBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_whileBlock

	return p
}

func (s *WhileBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileBlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserRPAREN, 0)
}

func (s *WhileBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *WhileBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
}

func (s *WhileBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *WhileBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterWhileBlock(s)
	}
}

func (s *WhileBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitWhileBlock(s)
	}
}

func (p *JackParser) WhileBlock() (localctx IWhileBlockContext) {
	localctx = NewWhileBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JackParserRULE_whileBlock)
	var _la int

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
		p.SetState(199)
		p.Match(JackParserRPAREN)
	}
	{
		p.SetState(200)
		p.Match(JackParserLBRACE)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(JackParserLET-26))|(1<<(JackParserIF-26))|(1<<(JackParserWHILE-26))|(1<<(JackParserDO-26))|(1<<(JackParserRETURN-26)))) != 0 {
		{
			p.SetState(201)
			p.Statement()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(JackParserRBRACE)
	}

	return localctx
}

// IDoStatementContext is an interface to support dynamic dispatch.
type IDoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoStatementContext differentiates from other interfaces.
	IsDoStatementContext()
}

type DoStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoStatementContext() *DoStatementContext {
	var p = new(DoStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_doStatement
	return p
}

func (*DoStatementContext) IsDoStatementContext() {}

func NewDoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoStatementContext {
	var p = new(DoStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_doStatement

	return p
}

func (s *DoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DoStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(JackParserDO, 0)
}

func (s *DoStatementContext) Subroutinecall() ISubroutinecallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinecallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinecallContext)
}

func (s *DoStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *DoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterDoStatement(s)
	}
}

func (s *DoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitDoStatement(s)
	}
}

func (p *JackParser) DoStatement() (localctx IDoStatementContext) {
	localctx = NewDoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JackParserRULE_doStatement)

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
		p.SetState(209)
		p.Match(JackParserDO)
	}
	{
		p.SetState(210)
		p.Subroutinecall()
	}
	{
		p.SetState(211)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(JackParserRETURN, 0)
}

func (s *ReturnStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *JackParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JackParserRULE_returnStatement)
	var _la int

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
		p.SetState(213)
		p.Match(JackParserRETURN)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JackParserMINUS)|(1<<JackParserNOT)|(1<<JackParserLPAREN)|(1<<JackParserTRUE)|(1<<JackParserNULL))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(JackParserFALSE-35))|(1<<(JackParserTHIS-35))|(1<<(JackParserID-35))|(1<<(JackParserSTRING-35))|(1<<(JackParserINTEGER-35)))) != 0) {
		{
			p.SetState(214)
			p.Expression()
		}

	}
	{
		p.SetState(217)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}

// ISubroutinecallContext is an interface to support dynamic dispatch.
type ISubroutinecallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinecallContext differentiates from other interfaces.
	IsSubroutinecallContext()
}

type SubroutinecallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinecallContext() *SubroutinecallContext {
	var p = new(SubroutinecallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_subroutinecall
	return p
}

func (*SubroutinecallContext) IsSubroutinecallContext() {}

func NewSubroutinecallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinecallContext {
	var p = new(SubroutinecallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_subroutinecall

	return p
}

func (s *SubroutinecallContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinecallContext) Subroutinename() ISubroutinenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinenameContext)
}

func (s *SubroutinecallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserLPAREN, 0)
}

func (s *SubroutinecallContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *SubroutinecallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserRPAREN, 0)
}

func (s *SubroutinecallContext) ClassObject() IClassObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassObjectContext)
}

func (s *SubroutinecallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinecallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinecallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubroutinecall(s)
	}
}

func (s *SubroutinecallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubroutinecall(s)
	}
}

func (p *JackParser) Subroutinecall() (localctx ISubroutinecallContext) {
	localctx = NewSubroutinecallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JackParserRULE_subroutinecall)

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
	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(219)
			p.ClassObject()
		}

	}
	{
		p.SetState(222)
		p.Subroutinename()
	}
	{
		p.SetState(223)
		p.Match(JackParserLPAREN)
	}
	{
		p.SetState(224)
		p.Expressionlist()
	}
	{
		p.SetState(225)
		p.Match(JackParserRPAREN)
	}

	return localctx
}

// IClassObjectContext is an interface to support dynamic dispatch.
type IClassObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassObjectContext differentiates from other interfaces.
	IsClassObjectContext()
}

type ClassObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassObjectContext() *ClassObjectContext {
	var p = new(ClassObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_classObject
	return p
}

func (*ClassObjectContext) IsClassObjectContext() {}

func NewClassObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassObjectContext {
	var p = new(ClassObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_classObject

	return p
}

func (s *ClassObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassObjectContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *ClassObjectContext) DOT() antlr.TerminalNode {
	return s.GetToken(JackParserDOT, 0)
}

func (s *ClassObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterClassObject(s)
	}
}

func (s *ClassObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitClassObject(s)
	}
}

func (p *JackParser) ClassObject() (localctx IClassObjectContext) {
	localctx = NewClassObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JackParserRULE_classObject)

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
		p.SetState(227)
		p.Match(JackParserID)
	}
	{
		p.SetState(228)
		p.Match(JackParserDOT)
	}

	return localctx
}

// IExpressionlistContext is an interface to support dynamic dispatch.
type IExpressionlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionlistContext differentiates from other interfaces.
	IsExpressionlistContext()
}

type ExpressionlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionlistContext() *ExpressionlistContext {
	var p = new(ExpressionlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_expressionlist
	return p
}

func (*ExpressionlistContext) IsExpressionlistContext() {}

func NewExpressionlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionlistContext {
	var p = new(ExpressionlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_expressionlist

	return p
}

func (s *ExpressionlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionlistContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionlistContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JackParserCOMMA)
}

func (s *ExpressionlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JackParserCOMMA, i)
}

func (s *ExpressionlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterExpressionlist(s)
	}
}

func (s *ExpressionlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitExpressionlist(s)
	}
}

func (p *JackParser) Expressionlist() (localctx IExpressionlistContext) {
	localctx = NewExpressionlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JackParserRULE_expressionlist)
	var _la int

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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JackParserMINUS)|(1<<JackParserNOT)|(1<<JackParserLPAREN)|(1<<JackParserTRUE)|(1<<JackParserNULL))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(JackParserFALSE-35))|(1<<(JackParserTHIS-35))|(1<<(JackParserID-35))|(1<<(JackParserSTRING-35))|(1<<(JackParserINTEGER-35)))) != 0) {
		{
			p.SetState(230)
			p.Expression()
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JackParserCOMMA {
			{
				p.SetState(231)
				p.Match(JackParserCOMMA)
			}
			{
				p.SetState(232)
				p.Expression()
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) AllBinaryoperation() []IBinaryoperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBinaryoperationContext)(nil)).Elem())
	var tst = make([]IBinaryoperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBinaryoperationContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Binaryoperation(i int) IBinaryoperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryoperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBinaryoperationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JackParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JackParserRULE_expression)
	var _la int

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
		p.SetState(240)
		p.Term()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JackParserPLUS)|(1<<JackParserMINUS)|(1<<JackParserASTERISK)|(1<<JackParserSLASH)|(1<<JackParserAND)|(1<<JackParserOR)|(1<<JackParserLT)|(1<<JackParserGT)|(1<<JackParserEQ))) != 0 {
		{
			p.SetState(241)
			p.Binaryoperation()
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBinaryoperationContext is an interface to support dynamic dispatch.
type IBinaryoperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// IsBinaryoperationContext differentiates from other interfaces.
	IsBinaryoperationContext()
}

type BinaryoperationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
}

func NewEmptyBinaryoperationContext() *BinaryoperationContext {
	var p = new(BinaryoperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_binaryoperation
	return p
}

func (*BinaryoperationContext) IsBinaryoperationContext() {}

func NewBinaryoperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryoperationContext {
	var p = new(BinaryoperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_binaryoperation

	return p
}

func (s *BinaryoperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryoperationContext) GetOperator() antlr.Token { return s.operator }

func (s *BinaryoperationContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BinaryoperationContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BinaryoperationContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JackParserPLUS, 0)
}

func (s *BinaryoperationContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JackParserMINUS, 0)
}

func (s *BinaryoperationContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(JackParserASTERISK, 0)
}

func (s *BinaryoperationContext) SLASH() antlr.TerminalNode {
	return s.GetToken(JackParserSLASH, 0)
}

func (s *BinaryoperationContext) AND() antlr.TerminalNode {
	return s.GetToken(JackParserAND, 0)
}

func (s *BinaryoperationContext) OR() antlr.TerminalNode {
	return s.GetToken(JackParserOR, 0)
}

func (s *BinaryoperationContext) LT() antlr.TerminalNode {
	return s.GetToken(JackParserLT, 0)
}

func (s *BinaryoperationContext) GT() antlr.TerminalNode {
	return s.GetToken(JackParserGT, 0)
}

func (s *BinaryoperationContext) EQ() antlr.TerminalNode {
	return s.GetToken(JackParserEQ, 0)
}

func (s *BinaryoperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryoperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryoperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterBinaryoperation(s)
	}
}

func (s *BinaryoperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitBinaryoperation(s)
	}
}

func (p *JackParser) Binaryoperation() (localctx IBinaryoperationContext) {
	localctx = NewBinaryoperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, JackParserRULE_binaryoperation)
	var _la int

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
	p.SetState(247)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*BinaryoperationContext).operator = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JackParserPLUS)|(1<<JackParserMINUS)|(1<<JackParserASTERISK)|(1<<JackParserSLASH)|(1<<JackParserAND)|(1<<JackParserOR)|(1<<JackParserLT)|(1<<JackParserGT)|(1<<JackParserEQ))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*BinaryoperationContext).operator = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(248)
		p.Term()
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParsTermContext struct {
	*TermContext
}

func NewParsTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParsTermContext {
	var p = new(ParsTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ParsTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParsTermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserLPAREN, 0)
}

func (s *ParsTermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParsTermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JackParserRPAREN, 0)
}

func (s *ParsTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterParsTerm(s)
	}
}

func (s *ParsTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitParsTerm(s)
	}
}

type IntegerTermContext struct {
	*TermContext
}

func NewIntegerTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerTermContext {
	var p = new(IntegerTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *IntegerTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerTermContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(JackParserINTEGER, 0)
}

func (s *IntegerTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterIntegerTerm(s)
	}
}

func (s *IntegerTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitIntegerTerm(s)
	}
}

type StringTermContext struct {
	*TermContext
}

func NewStringTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTermContext {
	var p = new(StringTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *StringTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTermContext) STRING() antlr.TerminalNode {
	return s.GetToken(JackParserSTRING, 0)
}

func (s *StringTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterStringTerm(s)
	}
}

func (s *StringTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitStringTerm(s)
	}
}

type SubroutineTermContext struct {
	*TermContext
}

func NewSubroutineTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubroutineTermContext {
	var p = new(SubroutineTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *SubroutineTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutineTermContext) Subroutinecall() ISubroutinecallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinecallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutinecallContext)
}

func (s *SubroutineTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubroutineTerm(s)
	}
}

func (s *SubroutineTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubroutineTerm(s)
	}
}

type KeywordTermContext struct {
	*TermContext
	keywordConst antlr.Token
}

func NewKeywordTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeywordTermContext {
	var p = new(KeywordTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *KeywordTermContext) GetKeywordConst() antlr.Token { return s.keywordConst }

func (s *KeywordTermContext) SetKeywordConst(v antlr.Token) { s.keywordConst = v }

func (s *KeywordTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordTermContext) TRUE() antlr.TerminalNode {
	return s.GetToken(JackParserTRUE, 0)
}

func (s *KeywordTermContext) FALSE() antlr.TerminalNode {
	return s.GetToken(JackParserFALSE, 0)
}

func (s *KeywordTermContext) NULL() antlr.TerminalNode {
	return s.GetToken(JackParserNULL, 0)
}

func (s *KeywordTermContext) THIS() antlr.TerminalNode {
	return s.GetToken(JackParserTHIS, 0)
}

func (s *KeywordTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterKeywordTerm(s)
	}
}

func (s *KeywordTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitKeywordTerm(s)
	}
}

type ArrayTermContext struct {
	*TermContext
	ident IVarnameContext
}

func NewArrayTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTermContext {
	var p = new(ArrayTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ArrayTermContext) GetIdent() IVarnameContext { return s.ident }

func (s *ArrayTermContext) SetIdent(v IVarnameContext) { s.ident = v }

func (s *ArrayTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTermContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACKET, 0)
}

func (s *ArrayTermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayTermContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACKET, 0)
}

func (s *ArrayTermContext) Varname() IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *ArrayTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterArrayTerm(s)
	}
}

func (s *ArrayTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitArrayTerm(s)
	}
}

type VarnameTermContext struct {
	*TermContext
}

func NewVarnameTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarnameTermContext {
	var p = new(VarnameTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *VarnameTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarnameTermContext) Varname() IVarnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarnameContext)
}

func (s *VarnameTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterVarnameTerm(s)
	}
}

func (s *VarnameTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitVarnameTerm(s)
	}
}

type UnaryopTermContext struct {
	*TermContext
	unaryop antlr.Token
}

func NewUnaryopTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryopTermContext {
	var p = new(UnaryopTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *UnaryopTermContext) GetUnaryop() antlr.Token { return s.unaryop }

func (s *UnaryopTermContext) SetUnaryop(v antlr.Token) { s.unaryop = v }

func (s *UnaryopTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryopTermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *UnaryopTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JackParserMINUS, 0)
}

func (s *UnaryopTermContext) NOT() antlr.TerminalNode {
	return s.GetToken(JackParserNOT, 0)
}

func (s *UnaryopTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterUnaryopTerm(s)
	}
}

func (s *UnaryopTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitUnaryopTerm(s)
	}
}

func (p *JackParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, JackParserRULE_term)
	var _la int

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

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntegerTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(JackParserINTEGER)
		}

	case 2:
		localctx = NewStringTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(JackParserSTRING)
		}

	case 3:
		localctx = NewKeywordTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(252)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*KeywordTermContext).keywordConst = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(JackParserTRUE-24))|(1<<(JackParserNULL-24))|(1<<(JackParserFALSE-24))|(1<<(JackParserTHIS-24)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*KeywordTermContext).keywordConst = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 4:
		localctx = NewVarnameTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(253)
			p.Varname()
		}

	case 5:
		localctx = NewArrayTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(254)

			var _x = p.Varname()

			localctx.(*ArrayTermContext).ident = _x
		}
		{
			p.SetState(255)
			p.Match(JackParserLBRACKET)
		}
		{
			p.SetState(256)
			p.Expression()
		}
		{
			p.SetState(257)
			p.Match(JackParserRBRACKET)
		}

	case 6:
		localctx = NewSubroutineTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(259)
			p.Subroutinecall()
		}

	case 7:
		localctx = NewParsTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(260)
			p.Match(JackParserLPAREN)
		}
		{
			p.SetState(261)
			p.Expression()
		}
		{
			p.SetState(262)
			p.Match(JackParserRPAREN)
		}

	case 8:
		localctx = NewUnaryopTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(264)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UnaryopTermContext).unaryop = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == JackParserMINUS || _la == JackParserNOT) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UnaryopTermContext).unaryop = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(265)
			p.Term()
		}

	}

	return localctx
}

// IClassnameContext is an interface to support dynamic dispatch.
type IClassnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassnameContext differentiates from other interfaces.
	IsClassnameContext()
}

type ClassnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassnameContext() *ClassnameContext {
	var p = new(ClassnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_classname
	return p
}

func (*ClassnameContext) IsClassnameContext() {}

func NewClassnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassnameContext {
	var p = new(ClassnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_classname

	return p
}

func (s *ClassnameContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassnameContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *ClassnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterClassname(s)
	}
}

func (s *ClassnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitClassname(s)
	}
}

func (p *JackParser) Classname() (localctx IClassnameContext) {
	localctx = NewClassnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, JackParserRULE_classname)

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
		p.SetState(268)
		p.Match(JackParserID)
	}

	return localctx
}

// IVarnameContext is an interface to support dynamic dispatch.
type IVarnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarnameContext differentiates from other interfaces.
	IsVarnameContext()
}

type VarnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarnameContext() *VarnameContext {
	var p = new(VarnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_varname
	return p
}

func (*VarnameContext) IsVarnameContext() {}

func NewVarnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarnameContext {
	var p = new(VarnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_varname

	return p
}

func (s *VarnameContext) GetParser() antlr.Parser { return s.parser }

func (s *VarnameContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *VarnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterVarname(s)
	}
}

func (s *VarnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitVarname(s)
	}
}

func (p *JackParser) Varname() (localctx IVarnameContext) {
	localctx = NewVarnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, JackParserRULE_varname)

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
		p.SetState(270)
		p.Match(JackParserID)
	}

	return localctx
}

// ISubroutinenameContext is an interface to support dynamic dispatch.
type ISubroutinenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinenameContext differentiates from other interfaces.
	IsSubroutinenameContext()
}

type SubroutinenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinenameContext() *SubroutinenameContext {
	var p = new(SubroutinenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_subroutinename
	return p
}

func (*SubroutinenameContext) IsSubroutinenameContext() {}

func NewSubroutinenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinenameContext {
	var p = new(SubroutinenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_subroutinename

	return p
}

func (s *SubroutinenameContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinenameContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *SubroutinenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterSubroutinename(s)
	}
}

func (s *SubroutinenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitSubroutinename(s)
	}
}

func (p *JackParser) Subroutinename() (localctx ISubroutinenameContext) {
	localctx = NewSubroutinenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, JackParserRULE_subroutinename)

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
		p.SetState(272)
		p.Match(JackParserID)
	}

	return localctx
}

package compiler

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"jackcompiler-antlr/parser"

	"jackcompiler-antlr/listener"
)

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

//s.vm = vmwriter.New(filenameWithoutExtension(pathName) + ".vm")

func Compile(pathName string) {

	input, err := ioutil.ReadFile(pathName)
	if err != nil {
		panic("error in opening file")
	}

	// Setup the input
	is := antlr.NewInputStream(string(input))

	// Create the Lexer
	lexer := parser.NewJackLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJackParser(stream)

	// generate AST
	tree := p.Classdef()

	vmfilename := filenameWithoutExtension(pathName) + ".vm"
	// visit tree
	antlr.ParseTreeWalkerDefault.Walk(listener.New(vmfilename), tree)

}

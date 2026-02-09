
# gopath
export GOPATH=$HOME/dev/go



# Download

wget http://www.antlr.org/download/antlr-4.7-complete.jar

# Criando alias
alias antlr='java -jar ~/Applications/antlr-4.7-complete.jar'

# baixando o antlr para o go

go get github.com/antlr/antlr4/runtime/Go/antlr@4.7.2

# Executando o antlr

antlr  -Dlanguage=Go -o parser Jack.g4 

# codigo base

```go
// example1.go
package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"./parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
```

parser

```go
type jackListener struct {
	*parser.BaseJackListener
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewJackLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJackParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&jackListener{}, p.Start())
}
```

# Run

go run . testes/Main.jack 

# Referencias

references

[antlr/antlr4](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)

[Parsing with ANTLR 4 and Go](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/)

[Learning ANTLR 4: building grammar](http://kedizheng.com/2019/03/31/learning-antlr-4-building-grammar/)

[potigol/Potigol](https://github.com/potigol/Potigol/blob/master/src/main/antlr4/br/edu/ifrn/potigol/parser/potigol.g4)

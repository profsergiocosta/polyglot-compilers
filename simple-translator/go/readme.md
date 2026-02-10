
# Dragon Compiler Go

An educational implementation in Go of a complete compilation pipeline, inspired by the "Dragon Book" (Compilers: Principles, Techniques, and Tools) and the Nand2Tetris project.

## Overview

This project demonstrates the main stages of a compiler:

1. **Lexical Analysis** - Source code tokenization
2. **Syntax Analysis** - Parsing and AST construction
3. **Intermediate Code Generation** - VM bytecode production
4. **Stack-based Virtual Machine** - Intermediate code execution
5. **Assembly Compiler** - Translation to machine code

## Architecture

```
┌─────────────┐      ┌──────────┐      ┌──────────┐      ┌─────────┐
│  Source     │ -->  │  Parser  │ -->  │   VM     │ -->  │  Stack  │
│  (.jack)    │      │   (AST)  │      │ (bytecode)│      │   VM    │
└─────────────┘      └──────────┘      └──────────┘      └─────────┘
                                              │
                                              v
                                        ┌──────────┐
                                        │ Assembly │
                                        │  (.asm)  │
                                        └──────────┘
```

## Components

### Jack Language (Simplified)
High-level language with simple syntax for demonstration:
```jack
let x = 10;
let y = x + 5;
print y;
```

### VM Code (Intermediate)
Stack-based bytecode:
```
push 10
pop x
push x
push 5
add
pop y
push y
print
```

### Assembly Generator
Translates VM bytecode to Hack architecture assembly language.

## Installation

```bash
export GOPATH=$HOME/dev/go
go mod download
```

## Usage

### Interpreter Mode (VM)
Execute VM bytecode files directly:
```bash
go run . -mode interpreter -input teste.vm
```

### Compiler Mode (Jack → VM)
Compile Jack code to VM bytecode:
```bash
go run . -input teste.jack
```

### Interactive REPL Mode
Experiment with the language interactively:
```bash
go run . -mode it
```

### Assembly Compiler
Translate VM to assembly code:
```bash
go run . -mode compiler -input teste.vm
```

Or compile Jack directly to assembly:
```bash
go run . -mode compiler -input teste.jack
```

## Project Structure

```
├── lexer/        # Lexical analysis
├── parser/       # Syntax analysis
├── token/        # Token definitions
├── compiler/     # Intermediate code generation
├── gen/          # Assembly generation
├── vm/           # Stack-based virtual machine
├── translator/   # VM → Assembly translation
├── interpret/    # Bytecode interpreter
├── repl/         # Interactive REPL
└── main.go       # Entry point
```

## Educational Goals

This project is ideal for:
- Studying compiler concepts practically
- Understanding intermediate code generation
- Exploring stack-based virtual machines
- Learning about machine language translation

## Implemented Concepts

- **Lexical Analysis**: Tokenization and pattern recognition
- **Syntax Analysis**: Recursive descent parsing
- **Symbol Table**: Variable management
- **Code Generation**: Syntax-directed translation
- **Stack-Based VM**: Stack architecture for execution
- **Assembly Backend**: Machine code generation

## Inspiration

Based on concepts from:
- **Compilers: Principles, Techniques, and Tools** (Aho, Sethi, Ullman)
- **Nand2Tetris** - From Nand to Tetris: Building a Modern Computer

## License

MIT License

## Author

Prof. Sergio Costa


Pr
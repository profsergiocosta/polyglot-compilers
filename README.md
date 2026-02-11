Perfect! I can see it's actually **eva-cpp-vm** (C++, not C) and **eva-js-interpreter**. Here's the complete README with all corrections:

***

# 🔧 Polyglot Compilers Lab

> **A multi-language laboratory for compiler and interpreter implementations**

A unified repository consolidating multiple compiler, interpreter, and translator implementations developed across different programming languages (**C, C++, Go, Python, Java, Scala, Haskell, Rust, JavaScript**). This project demonstrates translation strategies ranging from hardware-level assemblers to high-level functional interpreters across various programming paradigms.


***

## 📋 Table of Contents

- [Overview](#-overview)
- [Migration Status](#-migration-status)
- [Project Structure](#-project-structure)
- [Implementations](#-implementations)
  - [Nand2Tetris Stack](#-nand2tetris-stack)
  - [Language Interpreters](#-language-interpreters--compilers)
- [Execution Models Comparison](#-execution-models-comparison)
- [Parser Technologies Showcase](#-parser-technologies-showcase)
- [Getting Started](#-getting-started)
- [Engineering Standards](#-engineering-standards)
- [Roadmap](#-roadmap)
- [Contributing](#-contributing)
- [Author](#-author)
- [License](#-license)

***

## 🎯 Overview

Welcome to my central laboratory for language engineering. This repository represents a **unified effort to integrate** several compiler and interpreter implementations I previously developed in separate environments.

**Core Objectives:**

- **Educational Resource**: Comprehensive examples for teaching compiler construction and language design
- **Technical Portfolio**: Demonstrating expertise across multiple programming paradigms and languages
- **Comparative Analysis**: Enabling performance benchmarking and design pattern comparison across implementations
- **Execution Model Showcase**: Examples of tree-walking interpreters, bytecode VMs, and IR interpreters
- **Parser Technology Showcase**: Examples of hand-written parsers, parser combinators, parser generators (ANTLR), and functional parsers
- **From Hardware to High-Level**: Complete translation stack from NAND gates to functional programming languages

***

## 🚀 Migration Status

> **⚠️ Important**: This repository is currently under **active migration and integration**. Code is being consolidated from legacy repositories, refactored for consistency, and enhanced with unified testing frameworks.

### ✅ Completed Migrations

| Component | Languages/Tools | Status |
|-----------|-----------------|---------|
| **Hack Assembler** | Go | ✅ Migrated |
| **Jack Compiler** | C++, Go, Go+ANTLR, Java, Scala, Rust | ✅ Migrated (6 implementations) |
| **Lox Interpreters** | Java (jlox), C (clox), Scala (slox) | ✅ Migrated |
| **Eva Interpreter** | JavaScript (tree-walk), C++ (VM) | ✅ Migrated (2 implementations) |
| **Lispy Interpreter** | C + MPC library | ✅ Migrated |
| **Scheme Interpreter** | Haskell (early stage) | ✅ Migrated |
| **SubC Interpreter** | Three-address code | ✅ Migrated |

### 🔄 In Progress

- [ ] **VM Translator**: Migration and multi-language implementations
- [ ] **Additional parser examples**: Extending parser technology demonstrations
- [ ] **C Compiler**: Integration and documentation

### 📝 Pending Updates

Each migrated component requires:

- [ ] **Dedicated README.md**: Detailed documentation with build instructions
- [ ] **Build Automation**: Standardized build scripts (Make, CMake, Go modules)
- [ ] **Test Integration**: Automated test suites with validation scripts
- [ ] **Performance Benchmarks**: Cross-language comparison metrics
- [ ] **Code Examples**: Sample programs demonstrating features

***

## 📂 Project Structure

```
polyglot-compilers/
├── nand2tetris/              # Nand2Tetris Translation Stack
│   ├── assembler/            
│   │   └── go/              # ✅ Go implementation (MIGRATED)
│   │
│   ├── vm-translator/        # 🔄 VM → Assembly translator (IN PROGRESS)
│   │
│   ├── jack-compiler/        # ✅ Jack language compiler (MIGRATED - 6 versions)
│   │   ├── cpp/             # C++ implementation
│   │   ├── go/              # Go hand-written parser
│   │   ├── go-antlr/        # Go + ANTLR parser generator
│   │   ├── java/            # Java implementation
│   │   ├── scala/           # Scala with unique AST design
│   │   └── rust/            # Rust implementation
│   │
│   └── examples/             # Test programs and examples
│
├── interpreters/             # Language Interpreters
│   ├── lox/                 # ✅ Lox language (3 implementations)
│   │   ├── jlox/           # Tree-walking (Java + Visitor Pattern)
│   │   ├── clox/           # Bytecode VM (C)
│   │   └── slox/           # Scala with ADTs + Pattern Matching
│   │
│   ├── eva/                 # ✅ Eva - Lisp-style (2 implementations)
│   │   ├── eva-js-interpreter/    # JavaScript tree-walking
│   │   └── eva-cpp-vm/            # C++ bytecode virtual machine
│   │
│   ├── lisp/                # ✅ Lisp interpreters
│   │   ├── lispy/          # C with MPC parser combinators
│   │   └── scheme-hs/      # Haskell with functional parsers (early stage)
│   │
│   └── subc/                # ✅ Three-address code interpreter
│
├── compiler/                 # Additional Compilers
│   └── ccompiler/           # 🔄 C compiler (NEEDS DOCUMENTATION)
│
└── simple-translator/        # ⏳ Educational translators (PENDING)
```

**Legend:**
- ✅ **Migrated**: Code moved and functional
- 🔄 **In Progress**: Migration underway
- ⏳ **Pending**: Not yet migrated
- 📝 **Needs Update**: Requires documentation/testing

***

## 🛠 Implementations

### 🎮 Nand2Tetris Stack

Complete implementation of the [Nand2Tetris](https://www.nand2tetris.org/) project, covering the entire Hack computer translation pipeline:

#### **Hack Assembler** ✅
- **Language**: Go
- **Status**: Migrated and functional
- **Function**: Translates Hack Assembly (`.asm`) to binary code (`.hack`)
- **Features**: Two-pass parser, symbol table management, label resolution
- **Parser**: Hand-written lexer and parser
- **Needs**: Test suite integration, performance benchmarks

***

#### **VM Translator** 🔄
- **Status**: Migration in progress
- **Function**: Translates stack-based VM bytecode to Hack Assembly
- **Support**: Arithmetic commands, memory access, control flow, function calls
- **Planned**: Multi-language implementations (Python, C++, Go)

***

#### **Jack Compiler** ✅ - Six Implementations!

The Jack compiler showcases **six different implementations** across various languages and parsing technologies:

| Language | Parser Technology | AST Design | Status |
|----------|------------------|------------|---------|
| **C++** | Hand-written recursive descent | Traditional OOP | ✅ Migrated |
| **Go** | Hand-written recursive descent | Procedural | ✅ Migrated |
| **Go + ANTLR** | **ANTLR4 parser generator** | Generated from grammar | ✅ Migrated |
| **Java** | Hand-written recursive descent | OOP with visitors | ✅ Migrated |
| **Scala** | Hand-written with ADTs | **Unique functional AST design** | ✅ Migrated |
| **Rust** | Hand-written with pattern matching | Ownership-based | ✅ Migrated |

**Key Highlights:**

- **ANTLR Example** (`go-antlr`): Demonstrates using a **parser generator** with declarative grammar specification instead of hand-written parsers
- **Scala Version**: Features a **unique AST design** using algebraic data types, different from the traditional OOP approach used in other implementations
- **Educational Value**: Compare hand-written parsers vs. generated parsers, and imperative vs. functional AST representations

**Function**: Full compiler for Jack language (Java-like OOP language)  
**Pipeline**: Tokenizer → Parser → AST Builder → Code Generator  
**Features**: Classes, methods, static typing, expressions, arrays, objects  
**Target**: Hack VM bytecode

**Needs**: Comprehensive test suite, performance benchmarks across implementations

***

### 🚀 Language Interpreters & Compilers

#### **Lox Implementations** ✅ (from *Crafting Interpreters*)

Three distinct implementations of the Lox programming language, each demonstrating different design approaches and programming paradigms:

| Implementation | Language | Type | Design Approach | Status |
|---------------|----------|------|-----------------|---------|
| **jlox** | Java | Tree-walking interpreter | Visitor Pattern (OOP) | ✅ Migrated |
| **clox** | C | Bytecode VM | Manual memory mgmt | ✅ Migrated |
| **slox** | Scala | Tree-walking interpreter | Algebraic Data Types (FP) | ✅ Migrated |

**Design Paradigm Comparison:**

- **jlox (Java)**: Classic **Visitor Pattern** for AST traversal, demonstrating object-oriented design principles from *Crafting Interpreters*
- **clox (C)**: Bytecode compiler with VM, focusing on performance through manual memory management and optimized instruction dispatch
- **slox (Scala)**: **Algebraic Data Types (ADTs)** with **pattern matching**, showcasing functional programming approach that eliminates the need for visitor pattern

**Lox Language Features**: Dynamic typing, first-class functions, closures, classes with inheritance, automatic garbage collection

**Educational Value**: These three implementations showcase how the same language can be built using radically different paradigms and techniques.

**Needs**: Standardized test suite, performance benchmarks comparing all three implementations

***

#### **Eva Interpreter** ✅ - Two Implementations

Eva is a Lisp-style educational language focusing on functional programming concepts, implemented in two different execution models:

| Implementation | Language | Type | Status |
|---------------|----------|------|---------|
| **eva-js-interpreter** | JavaScript | Tree-walking (AST interpreter) | ✅ Migrated |
| **eva-cpp-vm** | C++ | Bytecode Virtual Machine | ✅ Migrated |

**eva-js-interpreter - JavaScript Tree-Walking Interpreter**

- **Language**: JavaScript
- **Architecture**: Direct AST interpretation
- **Input**: JavaScript objects representing the AST
- **Execution**: Recursive evaluation of AST nodes

**Key Features:**
- **AST as JavaScript Objects**: The abstract syntax tree is represented using native JavaScript objects
  ```javascript
  // Example AST representation
  ['def', 'x', 10]           // Variable definition
  ['+', ['var', 'x'], 5]     // Expression
  ['if', condition, then, else]  // Control flow
  ```
- **Direct Interpretation**: Walks and evaluates the AST directly without compilation
- **S-Expression Support**: Lisp-style syntax with prefix notation
- **Dynamic Execution**: Leverages JavaScript's dynamic nature

***

**eva-cpp-vm - C++ Virtual Machine Implementation**

- **Language**: C++
- **Architecture**: Bytecode compiler + virtual machine
- **Pipeline**: AST → Bytecode → VM Execution

**Key Features:**
- **Bytecode Compilation**: Compiles AST to bytecode instructions
- **Stack-Based VM**: Virtual machine with stack-based execution model
- **Performance**: Significantly faster execution than tree-walking approach
- **OOP Design**: Uses C++ classes for VM components

**VM Architecture:**
```
Source → Parser → AST → Compiler → Bytecode → VM → Result
```

**Instruction Set Examples:**
```
PUSH <value>     # Push value onto stack
ADD              # Pop two values, push sum
LOAD <var>       # Load variable
STORE <var>      # Store to variable
CALL <addr>      # Function call
JMP <addr>       # Unconditional jump
```

**Educational Value**:
- Demonstrates bytecode compilation techniques
- Shows VM implementation in modern C++
- Compares interpreted vs. VM execution models
- Illustrates instruction encoding and dispatch

***

**Comparison: eva-js vs. eva-cpp-vm**

| Aspect | eva-js-interpreter (JavaScript) | eva-cpp-vm (C++) |
|--------|----------------------------------|-------------------|
| **Execution Model** | Tree-walking | Bytecode VM |
| **Speed** | Slower (AST traversal) | Faster (compiled bytecode) |
| **Memory** | Automatic GC | Manual/RAII management |
| **Complexity** | Simpler implementation | More complex (compiler + VM) |
| **Portability** | Requires JS runtime | Standalone binary |
| **Code Size** | More concise | More verbose |

**Eva Language Features** (both implementations):
- S-expressions and prefix notation
- Dynamic typing
- First-class functions
- Closures and lexical scoping
- Higher-order functions
- Functional programming emphasis

**Needs**: Comprehensive examples, performance benchmarks comparing both implementations, extended standard library

***

#### **Lisp Family Interpreters** ✅

**lispy - C Implementation with Parser Combinators**

- **Language**: C
- **Parser**: [Micro Parser Combinators (MPC)](https://github.com/orangeduck/mpc) library
- **Status**: Migrated
- **Architecture**: Hand-written in C using a lightweight parser combinator library

**Key Features:**
- **Parser Combinators in C**: Demonstrates functional parsing techniques in a procedural language
- **MPC Library**: Uses combinators like `mpc_or`, `mpc_and`, `mpc_many` to build the parser declaratively
- **S-Expression Support**: Full Lisp-style syntax parsing
- **Minimal Dependencies**: Lightweight implementation suitable for educational purposes

**Example MPC Usage:**
```c
mpc_parser_t* Number = mpc_new("number");
mpc_parser_t* Symbol = mpc_new("symbol");
mpc_parser_t* Sexpr  = mpc_new("sexpr");

// Combine parsers using combinators
mpca_lang(MPCA_LANG_DEFAULT,
  " number  : /-?[0-9]+/ ;           \
    symbol  : '+' | '-' | '*' | '/' ; \
    sexpr   : '(' <expr>* ')' ;      \
    expr    : <number> | <symbol> | <sexpr> ; ",
  Number, Symbol, Sexpr, ...);
```

**Educational Value**: Shows how functional parser combinator patterns can be applied in C, bridging functional and procedural worlds.

**Needs**: Extended Lisp features, more built-in functions

***

**scheme-hs - Haskell Scheme Interpreter**

- **Language**: Haskell
- **Parser**: Functional parser (Parsec or custom parser combinators)
- **Status**: Early stage implementation
- **Paradigm**: Pure functional programming

**Key Features:**
- **Functional Parsers**: Demonstrates idiomatic Haskell parsing techniques
- **Type Safety**: Leverages Haskell's strong type system
- **Monadic Composition**: Parser built using monadic parser combinators
- **Early Stage**: Foundation for exploring advanced functional programming concepts

**Example Functional Parser:**
```haskell
-- Parser combinator style
parseExpr :: Parser Expr
parseExpr = parseAtom
        <|> parseList
        <|> parseQuoted

parseList :: Parser Expr
parseList = do
  char '('
  exprs <- many parseExpr
  char ')'
  return $ List exprs
```

**Educational Value**: 
- Example of **functional parser design** in a pure functional language
- Contrasts with imperative parsing approaches
- Demonstrates how parser combinators naturally fit functional languages
- Shows type-driven development

**Planned Features:**
- Complete Scheme R5RS subset
- Tail-call optimization
- Macro system
- Advanced functional features

**Needs**: Complete implementation, comprehensive examples, REPL

***

#### **SubC - Three-Address Code Interpreter** ✅
- **Type**: Intermediate representation (IR) interpreter
- **Status**: Migrated
- **Architecture**: Three-address code (TAC) based execution

**Key Characteristics:**
- **Three-Address Code Format**: Each instruction has at most three operands
  ```
  t1 = a + b      # Binary operation
  t2 = t1 * c     # Use temporary variables
  result = t2     # Assignment
  ```
- **IR-Level Execution**: Directly interprets intermediate representation
- **Simplified Semantics**: One operation per instruction for clear execution model
- **Low-Level Focus**: Closer to assembly than high-level languages

**Pipeline**: 
1. Source Code Parsing
2. Three-Address Code Generation
3. Direct Interpretation of TAC

**Educational Value**: 
- Demonstrates the **intermediate representation layer** used in real compilers
- Bridges gap between high-level constructs and machine code
- Shows how complex expressions are linearized
- Prepares foundation for optimization passes

**Example TAC:**
```
# High-level: result = (a + b) * (c - d)

# Three-address code:
t1 = a + b
t2 = c - d
t3 = t1 * t2
result = t3
```

**Needs**: Extended documentation, more example programs, TAC optimization demonstrations

***

## 🎯 Execution Models Comparison

This repository showcases **three major execution strategies** across different languages, with multiple implementations of the same strategy for comparison:

| Execution Model | Examples | Language | Description |
|-----------------|----------|----------|-------------|
| **Tree-Walking Interpreters** | jlox, slox, eva-js-interpreter | Java, Scala, JavaScript | Direct AST evaluation |
| **Bytecode Virtual Machines** | clox, eva-cpp-vm | C, C++ | Compile to bytecode, execute on VM |
| **Intermediate Representation** | SubC | (Three-address code) | Direct interpretation of IR |

### Side-by-Side Comparison

**Lox Language** (3 implementations):
- **jlox**: Tree-walking with Visitor Pattern (Java)
- **clox**: Bytecode VM (C)
- **slox**: Tree-walking with ADTs (Scala)

**Eva Language** (2 implementations):
- **eva-js-interpreter**: Tree-walking (JavaScript)
- **eva-cpp-vm**: Bytecode VM (C++)

**Educational Value**: 
- Compare performance characteristics across execution models
- Analyze implementation complexity trade-offs
- Understand when to choose each approach
- Study how language paradigms affect interpreter design

***

## 🔍 Parser Technologies Showcase

This repository demonstrates **four major parsing approaches** with practical implementations:

| Technology | Example | Language | Description |
|------------|---------|----------|-------------|
| **Hand-Written Recursive Descent** | Jack (C++, Go, Java, Rust) | Multiple | Traditional manual parser implementation |
| **Parser Generators (ANTLR)** | Jack (go-antlr) | Go | Declarative grammar → generated parser |
| **Parser Combinators (MPC)** | lispy | C | Functional combinators in procedural language |
| **Functional Parsers** | scheme-hs | Haskell | Monadic parser combinators |
| **ADT Pattern Matching** | slox, Jack (Scala) | Scala | Functional parsing with algebraic types |

**Educational Value**: Compare and contrast different parsing philosophies, from imperative to functional approaches, and evaluate trade-offs in maintainability, performance, and expressiveness.

***

## 🚀 Getting Started

### Prerequisites

- **Go** 1.18+ (for Go projects)
- **Python** 3.8+ (for Python implementations)
- **GCC/Clang** (for C/C++ projects)
- **Java JDK** 11+ (for Java projects)
- **Scala** 2.13+ with sbt (for Scala projects)
- **GHC** 8.10+ (for Haskell projects)
- **Rust** 1.60+ (for Rust projects)
- **Node.js** 14+ (for JavaScript projects)
- **ANTLR4** (for ANTLR-based parsers)
- **Make** or **CMake** (for build automation)

### Quick Start Examples

#### Running the Hack Assembler (Go)

```bash
cd nand2tetris/assembler/go
go build -o assembler
./assembler ../../examples/Add.asm
```

#### Running Jack Compiler with ANTLR (Go)

```bash
cd nand2tetris/jack-compiler/go-antlr
# Follow local README for ANTLR setup
make build
./jackc program.jack
```

#### Running Jack Compiler (Scala with unique AST)

```bash
cd nand2tetris/jack-compiler/scala
sbt compile
sbt "run program.jack"
```

#### Running jlox (Tree-walking Lox Interpreter - Java)

```bash
cd interpreters/lox/jlox
javac Lox.java
java Lox your_script.lox
```

#### Running clox (Bytecode Lox VM - C)

```bash
cd interpreters/lox/clox
make
./clox your_script.lox
```

#### Running slox (Scala Lox with ADTs)

```bash
cd interpreters/lox/slox
sbt run your_script.lox
```

#### Running eva-js-interpreter (JavaScript)

```bash
cd interpreters/eva/eva-js-interpreter
node eva.js
# Or run with a source file
node eva.js program.eva
```

#### Running eva-cpp-vm (C++ VM)

```bash
cd interpreters/eva/eva-cpp-vm
make
./eva-vm program.eva
```

#### Running lispy (C with MPC)

```bash
cd interpreters/lisp/lispy
make
./lispy
> (+ 1 2 3)
6
> (def x 10)
> (+ x 5)
15
```

#### Running scheme-hs (Haskell)

```bash
cd interpreters/lisp/scheme-hs
stack build  # or cabal build
stack exec scheme-hs
```

#### Running SubC (Three-address code interpreter)

```bash
cd interpreters/subc
# Follow local README for build instructions
./subc program.subc
```

> **Note**: Each subdirectory will contain its own README with specific build and execution instructions (documentation in progress).

***

## 📊 Engineering Standards

To maintain a high-quality polyglot environment, each implementation is being refactored to meet these standards:

| Standard | Description | Migration Status |
|----------|-------------|------------------|
| **Isolated Builds** | Each project contains its own build manifest (`go.mod`, `CMakeLists.txt`, `Makefile`, `build.sbt`, `package.json`) | 🔄 In progress |
| **Local Documentation** | Each subfolder includes a README explaining design decisions and usage | ⏳ Pending for most |
| **Automated Validation** | Scripts to run standard test programs against each tool | ⏳ Planned |
| **Performance Metrics** | Comparative benchmarks across languages and execution models | ⏳ Planned |
| **Consistent Code Style** | Language-specific style guides and formatters | 🔄 Applying |
| **Test Coverage** | Unit and integration tests for each implementation | ⏳ Planned |

***

## 🗺 Roadmap

### Phase 1: Migration & Consolidation (Current)
- [x] Repository structure established
- [x] Core implementations migrated
  - [x] Assembler (Go)
  - [x] Jack Compilers (6 languages)
  - [x] Lox interpreters (3 variants)
  - [x] Eva interpreters (2 variants)
  - [x] Lisp interpreters (2 variants)
  - [x] SubC interpreter
- [ ] All legacy code consolidated
- [ ] Build systems standardized across all projects

### Phase 2: Documentation & Testing
- [ ] README.md for each implementation with build instructions
- [ ] Integrated test suites
  - [ ] Nand2Tetris `.tst` files
  - [ ] Lox test suite for all three implementations
  - [ ] Eva test programs
- [ ] Usage examples and tutorials for each tool
- [ ] Parser technology comparison guide
- [ ] Execution model comparison guide
- [ ] API documentation where applicable

### Phase 3: Quality & Performance
- [ ] Automated CI/CD pipeline (GitHub Actions)
- [ ] Cross-language performance benchmarks
  - [ ] Jack compiler comparison (6 implementations)
  - [ ] Lox interpreter comparison (jlox vs clox vs slox)
  - [ ] Eva interpreter comparison (JS vs C++ VM)
- [ ] Memory profiling and optimization
- [ ] Code coverage analysis
- [ ] Parser performance comparison

### Phase 4: Expansion
- [ ] VM Translator in multiple languages (Go, Python, C++)
- [ ] Additional parser examples (PEG, Earley, LR)
- [ ] Complete scheme-hs implementation (R5RS subset)
- [ ] Optimization passes for compilers
- [ ] Interactive debugging tools
- [ ] REPL for all interpreters
- [ ] Standard libraries for Eva and SubC
- [ ] JIT compilation examples

***

## 🤝 Contributing

Contributions are welcome! This is both a personal project and an educational resource.

### How to Contribute

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/NewImplementation`)
3. **Commit** your changes (`git commit -m 'Add Python VM Translator'`)
4. **Push** to the branch (`git push origin feature/NewImplementation`)
5. **Open** a Pull Request

### Areas of Interest

- 🔧 **New implementations** in different languages (Python, Elixir, OCaml, Zig)
- 📝 **Documentation** improvements and examples
- ✅ **Test suites** and validation scripts
- ⚡ **Performance optimizations** for existing implementations
- 📊 **Benchmarking infrastructure** and comparative analysis
- 🐛 **Bug fixes** and code improvements
- 🎓 **Educational materials** and tutorials
- 🔍 **Parser comparisons** and parser technology examples
- 🎯 **Optimization passes** (dead code elimination, constant folding, etc.)

### Current Priorities

Since the project is in migration phase, these contributions are especially valuable:

1. **Completing README documentation** for migrated implementations
2. **Creating build automation scripts** for consistent builds
3. **Developing test suites** with standard test programs
4. **Writing usage examples** and tutorials
5. **Parser comparison documentation** with performance analysis
6. **Execution model comparison** with benchmarks

***

## 👨‍🏫 About the Author

I am a **Computer Engineering Professor** with deep interests in:
- Software Engineering
- Functional Programming
- Compiler Design and Implementation
- Programming Language Theory
- Parser Technologies
- Virtual Machine Design

This repository serves as both my **personal engineering portfolio** and a **pedagogical resource** for my students. It demonstrates practical applications of theoretical concepts taught in computer science curricula, with special emphasis on:

- Comparing different implementation approaches and programming paradigms
- Understanding execution model trade-offs
- Exploring parser design patterns
- Analyzing performance characteristics across languages

The repository showcases how fundamental concepts can be expressed in different ways, helping students understand that there's rarely "one right way" to build a compiler or interpreter.

***

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

***

## 📚 References & Resources

### Books
- [Nand2Tetris: Building a Modern Computer from First Principles](https://www.nand2tetris.org/)
- [Crafting Interpreters](https://craftinginterpreters.com/) by Robert Nystrom
- [Engineering a Compiler](https://www.elsevier.com/books/engineering-a-compiler/cooper/978-0-12-088478-0) by Cooper & Torczon
- [Modern Compiler Implementation](https://www.cs.princeton.edu/~appel/modern/) by Andrew Appel
- [Types and Programming Languages](https://www.cis.upenn.edu/~bcpierce/tapl/) by Benjamin C. Pierce

### Tools & Libraries
- [ANTLR4](https://www.antlr.org/) - Parser generator for reading, processing, executing structured text
- [Micro Parser Combinators (MPC)](https://github.com/orangeduck/mpc) - Lightweight parser combinator library for C
- [Parsec](https://hackage.haskell.org/package/parsec) - Monadic parser combinators for Haskell

### Courses & Tutorials
- [Stanford CS143: Compilers](https://web.stanford.edu/class/cs143/)
- [MIT 6.035: Computer Language Engineering](http://web.mit.edu/6.035/)
- [Write Yourself a Scheme in 48 Hours](https://en.wikibooks.org/wiki/Write_Yourself_a_Scheme_in_48_Hours)

***

<div align="center">

**⭐ If this project helps you learn about compilers and interpreters, consider giving it a star!**

[Report Bug](https://github.com/profsergiocosta/polyglot-compilers/issues) · [Request Feature](https://github.com/profsergiocosta/polyglot-compilers/issues) · [Documentation](https://github.com/profsergiocosta/polyglot-compilers/wiki)

</div>

***


This README now accurately reflects the entire repository structure with all corrections incorporated!

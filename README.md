## KSM Language

KSM Language is a simple, interpreted language implemented in Go. The project includes a lexer, parser, and interpreter that work together to read, understand, and execute code written in the KSM language. It is a great starting point for understanding the components of a programming language.

## Project Structure

The KSM language project is composed of three main components:

1. **Lexer**: Responsible for breaking the source code into tokens.
2. **Parser**: Organizes the tokens into a meaningful structure (Abstract Syntax Tree - AST).
3. **Interpreter**: Executes the parsed instructions.

### Files:

- `lexer.go`: Contains the lexical analyzer (lexer) that converts the input code into tokens.
- `parser.go`: Contains the parser that processes the tokens to build an abstract syntax tree (AST).
- `interpreter.go`: Contains the interpreter that executes the AST.
- `main.go`: The entry point for running KSM either in REPL mode or by executing a .ksm file.

## Features

- **Basic Syntax**: Supports variable declarations, control structures (like `if` and `case`), basic arithmetic operations, and more.
- **Tokenization**: Tokenizes keywords, identifiers, numbers, strings, and operators.
- **Parsing**: Builds a syntax tree from the tokens.
- **Interpreting**: Executes the parsed code.

## Installation
```bash
git clone https://github.com/yourusername/ksm-language.git
```
cd ksm-language

## Usage
- You can run KSM in two ways:

    1. Interactive REPL: Starts an interactive session where you can type KSM code line by line.

```bash
go run .
```

    2. Run a File: Executes a .ksm file.

```bash
go run . core.ksm
```

## Example on how the program works
declare x = "Hello World"
displayln(x)

* Output will be
. Hello World

## Features

The project supports the following data structures:

    Number: Supports both integers.
    String: A sequence of characters.
    Boolean: Represents true and false values.

# Data Types

    int -> digit
    string -> sent
    rune -> char
    float64 -> deci64
    float32 -> deci32
    bool -> booly
    byte -> bite

# Control Structures

    main -> core
    if -> if case
    else -> otherwise
    switch -> option
    case -> option case
    default -> default case
    return -> return
    break -> exit

# Functions and Methods

    func -> method
    print -> display
    println -> displayln
    printf -> displayf

# Declarations

    var -> declare
    const -> constant
    type -> define

## Contributors
- [hanapiko](https://github.com/hanapiko)
- [Kevwasonga](https://github.com/kevwasonga)
- [krodgers](https://github.com/karodgers)
- [Somulo](https://github.com/samuelomulo)







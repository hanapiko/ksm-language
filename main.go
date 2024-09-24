package main

import (
	"bufio"
	"fmt"
	"os"

	"ksm/interpreter"
	"ksm/lexer"
	"ksm/parser"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	i := interpreter.NewInterpreter()

	for {
		fmt.Print(green + ">> " + reset)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		ast, err := p.Parse()
		if err != nil {
			fmt.Printf("Parsing error: %v\n", err)
			continue
		}

		err = i.Interpret(ast)
		if err != nil {
			fmt.Printf("Interpretation error: %v\n", err)
		}
	}
}

func runFile(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err)
		return
	}

	i := interpreter.NewInterpreter()
	l := lexer.NewLexer(string(content))
	p := parser.NewParser(l)

	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("Parsing error: %v\n", err)
		return
	}

	err = i.Interpret(ast)
	if err != nil {
		fmt.Printf("Interpretation error: %v\n", err)
	}
}

const (
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	reset   = "\033[0m"
)

func main() {
	if len(os.Args) > 1 {
		// If a filename is provided, run the file
		runFile(os.Args[1])
	} else {
		// Otherwise, start REPL
		fmt.Println(yellow + "Welcome to the KSM REPL!" + reset)
		fmt.Println(cyan + "Type 'exit' to quit." + reset)
		repl()
	}
}

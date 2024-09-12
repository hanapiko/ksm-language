package main

import (
	"fmt"
	"ksm/lexer"
)

func main() {
	input := `declare x = 10
displayln("Hello, World!")
if x > 5 {
    displayln("x is greater than 5")
} else {
    displayln("x is 5 or less")
}`

	// Create a new lexer instance
	l := lexer.NewLexer(input)

	// Iterate through the tokens generated by the lexer
	for tok := l.NextToken(); tok.Type != lexer.TokenEOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}

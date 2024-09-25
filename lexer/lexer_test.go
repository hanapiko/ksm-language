package lexer

import (
	"testing"
)

// TestNextToken tests the lexer for various inputs and expected token outputs.
func TestNextToken(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []Token
	}{
		{
			input: "declare x = 5",
			expectedTokens: []Token{
				{Type: TokenKeyword, Literal: "declare", Line: 1, Column: 1},
				{Type: TokenIdent, Literal: "x", Line: 1, Column: 9},
				{Type: TokenOperator, Literal: "=", Line: 1, Column: 11},
				{Type: TokenNumber, Literal: "5", Line: 1, Column: 13},
				{Type: TokenEOF, Literal: "", Line: 1, Column: 14},
			},
		},
		{
			input: `displayln("Hello World")`,
			expectedTokens: []Token{
				{Type: TokenKeyword, Literal: "displayln", Line: 1, Column: 1},
				{Type: TokenOperator, Literal: "(", Line: 1, Column: 10},
				{Type: TokenString, Literal: "Hello World", Line: 1, Column: 11},
				{Type: TokenOperator, Literal: ")", Line: 1, Column: 24},
				{Type: TokenEOF, Literal: "", Line: 1, Column: 25},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)

		for i, expectedToken := range tt.expectedTokens {
			actualToken := l.NextToken()
			if actualToken != expectedToken {
				t.Errorf("Test %d: Expected token %v, got %v", i, expectedToken, actualToken)
			}
		}
	}
}

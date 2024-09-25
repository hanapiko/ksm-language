package lexer

import (
	"unicode"
)

// TokenType represents the type of token
type TokenType string

const (
	TokenKeyword  TokenType = "KEYWORD"  // Keywords like 'if', 'else'
	TokenOperator TokenType = "OPERATOR" // Operators like '=', '+', '=='
	TokenIdent    TokenType = "IDENT"    // Identifiers like variable names
	TokenNumber   TokenType = "NUMBER"   // Numeric literals
	TokenString   TokenType = "STRING"   // String literals
	TokenEOF      TokenType = "EOF"      // End of file/input
	TokenError    TokenType = "ERROR"    // Error token for unrecognized characters
)

// Token represents a single token with its type, value, and position
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Lexer struct holds the state of the lexer
type Lexer struct {
	input        string 
	position     int    
	readPosition int    
	ch           byte   
	line         int    
	column       int    
}

// NewLexer initializes a new lexer with the input string
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar() // Read the first character
	return l
}

// readChar reads the next character from the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Set to null byte if at the end
	} else {
		l.ch = l.input[l.readPosition] // current character
	}
	l.position = l.readPosition // Update position
	l.readPosition++            // Move to the next character
	l.column++                  // Increment column count
}

// peekChar returns the next character without advancing the position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition] // return the next character
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	tok := Token{Line: l.line, Column: l.column} // initialize a new token

	// Match characters to produce tokens
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok.Literal = "=="
			tok.Type = TokenOperator
			l.readChar() // Read the next character
		} else {
			tok = l.newToken(TokenOperator, string(l.ch))
		}
	case '>':
		tok = l.newToken(TokenOperator, string(l.ch))
	case '<':
		tok = l.newToken(TokenOperator, string(l.ch))
	case '(':
		tok = l.newToken(TokenOperator, string(l.ch))
	case ')':
		tok = l.newToken(TokenOperator, string(l.ch))
	case '{':
		tok = l.newToken(TokenOperator, string(l.ch))
	case '}':
		tok = l.newToken(TokenOperator, string(l.ch))
	case '"':
		tok.Type = TokenString
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = TokenEOF
	default:
		// Identify letters and numbers
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.identifierType(tok.Literal) // Determine if it's a keyword or identifier
			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tok.Literal = l.readNumber()
			tok.Type = TokenNumber
			return tok
		} else {
			tok = l.newToken(TokenError, string(l.ch)) // unrecognized character
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads identifier from the input
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || unicode.IsDigit(rune(l.ch)) {
		l.readChar() // Read characters for the identifier
	}
	return l.input[position:l.position] // Return the identified string
}

// readNumber reads a number from the input
func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return l.input[position:l.position] // Return the identified number
}

// readString reads a string literal from the input
func (l *Lexer) readString() string {
	position := l.position + 1 // Start after the opening quote
	for {
		l.readChar() // Read until closing quote or EOF
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position] // Return the string content
}

// skipWhitespace skips over whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

// newToken creates a new token with the specified type and literal value
func (l *Lexer) newToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal, Line: l.line, Column: l.column}
}

// checks if a character is a valid letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// identifierType determines if an identifier is a keyword or an identifier
func (l *Lexer) identifierType(ident string) TokenType {
	keywords := map[string]bool{
		"declare": true, "displayln": true, "if": true, "case": true, "otherwise": true,
	}
	if keywords[ident] {
		return TokenKeyword
	}
	return TokenIdent
}

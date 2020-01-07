// lexer.go - lexer for monkey language
package lexer

import (
	token "github.com/Yepez1997/interpreterGo/src/token"
)

// token info at lexer level
type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // peek the next position
	ch           byte // current character examining
}

// create a lexer input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// start advancing ch by 1
	l.readChar()
	return l
}

// readChar - helper method
func (l *Lexer) readChar() {
	// finished reading @ the end
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// update the current position to be the position read
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken - get the next token in the lexer input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// look at the characer l.ch and return a token depending on the characer
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

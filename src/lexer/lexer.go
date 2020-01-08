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

// skipWhiteSpace - skip white space every time it is encounted
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken - get the next token in the lexer input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// look at the characer l.ch and return a token depending on the characer

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		// check for double ==
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(ch)
			tok = token.Token{Type: token.EQ, Literal: literal}

		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
		// check for not equal
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// once the literal has been set, need to determain the right token type
			// ie could be language syntax like function, let or variable name
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			// do not need to further read characters here so returning here is nec...
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok
}

// readIdenifier - read the entire identifier and advance iff the next char is not
// illegal
func (l *Lexer) readIdentifier() string {
	position := l.position
	// continue reading the word, adva
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber - keep reading the input for every valid digit
func (l *Lexer) readNumber() string {
	// very simplified as we can add floats, hex, octal notation etc ...
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// peekChar - peek the next character in the input sequence

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// isLetter - return if the byte is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '='
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken - return a newly formed token given input from the lexer
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

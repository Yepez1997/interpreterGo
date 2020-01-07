// lexer.go - lexer for monkey language
package lexer

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

// lexer.go - lexer for monkey language
package lexer

// token info at lexer level
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// create a lexer input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

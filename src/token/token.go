// token.go - lexical analysis
package token

type TokenType string

// Token shoudl store the token type and the literal it holds
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookUpIdent - look in the set of key words for the right token type
func LookUpIdent(ident string) TokenType {
	// comma, ok idiom
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// set of token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// key words
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

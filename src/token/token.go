// token.go - lexical analysis
package token

type TokenType string

// Token shoudl store the token type and the literal it holds
type Token struct {
	Type    TokenType
	Literal string
}

// keywords read from the lexer and mapped to the correct token
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"false":  FALSE,
	"true":   TRUE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	LT = "<"
	GT = ">"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	NOT_EQ = "!="
	EQ     = "=="
)

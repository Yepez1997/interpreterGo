// parser
package parser

import (
	ast "github.com/Yepez1997/interpreterGo/src/ast"
	lexer "github.com/Yepez1997/interpreterGo/src/lexer"
	token "github.com/Yepez1997/interpreterGo/src/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// set both tokens for the parser
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

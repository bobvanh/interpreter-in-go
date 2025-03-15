package parser

import (
	"calc/ast"
	"calc/lexer"
	"calc/token"
	"fmt"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken    token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read 2 tokens so curToken and peekToken are set
	//
	// in calc a Statement could consist of only 1 digit
	// in that case the current token would be that digit
	// and the peek token would be EOF
	//
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	if p.currentToken.Type != token.EOF {
		program.Expression = p.parseExpression()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseExpression() *ast.ExpressionStatement {
	if p.currentTokenIs(token.INT) {
		if p.peekToken.Type == token.EOF {
			return p.parseIntegerLiteral()
		} else {
			return p.parseOperatorExpression()
		}
	} else if p.currentTokenIs(token.LPAREN) {
		return p.parseGroupedExpression()
	} else {
		return nil
	}
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	return nil
}

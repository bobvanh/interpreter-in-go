package lexer

import "calc/token"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func (l *Lexer) nextChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.nextChar()
			literal := string(ch) + string(l.ch)
			tok = *token.New(token.ADD_ONE, literal)
		} else {
			tok = newToken(token.ADD, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.nextChar()
			literal := string(ch) + string(l.ch)
			tok = *token.New(token.SUB_ONE, literal)
		} else {
			tok = newToken(token.SUB, l.ch)
		}
	case '*':
		tok = newToken(token.PRODUCT, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.nextChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.nextChar()
	}
	return l.input[position:l.position]
}

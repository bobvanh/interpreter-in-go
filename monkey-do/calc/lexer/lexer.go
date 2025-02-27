package lexer

import "monkey-do/calc/token"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	currentChar  byte
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func New(input string) *Lexer {
	var l := &Lexer(input: input)
	l.readChar();
	return l;
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		if (l.peekChar() == '+') {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok := New(token.ADD_ONE, literal)
		} else {
			tok := New(token.ADD, string(l.ch))
		}
	case '-':
		if (l.peekChar() == '-') {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok := New(token.SUB_ONE, literal)
		} else {
			tok := New(token.SUB, string(l.ch))
		}
	case '*':
		tok := New(token.PRODUCT, string(l.ch))
	case '/':
		tok := New(token.DIVIDE, string(l.ch))
	case '(':
		tok := New(token.LPAREN, string(l.ch))
	case ')':
		tok := New(token.RPAREN, string(l.ch))
	default:
		if isDigit(l.ch) {
			literal := l.readNumber()
			tok := New(token.INT, literal)
		} else {
			return nil
		}
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

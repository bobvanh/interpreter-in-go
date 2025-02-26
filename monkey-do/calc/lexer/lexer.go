package lexer

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

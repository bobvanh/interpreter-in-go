package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT" // 1343456

	// Operators
	LPAREN  = "("
	RPAREN  = ")"
	ADD     = "+"
	SUB     = "-"
	PRODUCT = "*"
	DIVIDE  = "/"
	ADD_ONE = "++"
	SUB_ONE = "--"
)

func New(tokenType TokenType, literal string) *Token {
	return &Token{Type: tokenType, Literal: literal}
}

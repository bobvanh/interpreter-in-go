package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LPAREN    = "("
	RPAREN    = ")"
	PLUS      = "+"
	MINUS     = "-"
	PRODUCT   = "*"
	DIVIDE    = "/"
	ADD_ONE   = "++"
	MINUS_ONE = "--"
)

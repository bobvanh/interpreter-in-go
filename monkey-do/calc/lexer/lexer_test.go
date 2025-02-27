package lexer

import (
	"testing"

	"calc/token"
)

func TestNextToken(t *testing.T) {
	input := `1 + 3-- * (834 -1)   / ( 1 + ++56)`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "1"},
		{token.ADD, "+"},
		{token.INT, "3"},
		{token.SUB_ONE, "--"},
		{token.PRODUCT, "*"},
		{token.LPAREN, "("},
		{token.INT, "834"},
		{token.SUB, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.DIVIDE, "/"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.ADD, "+"},
		{token.ADD_ONE, "++"},
		{token.INT, "56"},
		{token.RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

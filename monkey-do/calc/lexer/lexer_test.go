package lexer

import (
	"testing"
	"monkey-do/calc/token"
)

func TestNextToken(t *testing.T) {
	input := `1 + 3 * (834 - 1) / (1 + 1)`

	tests := []struct {
		expectedType token.TokenType,
		expectedLiteral string
	}{
		{token.INT, "1"},
		{token.ADD, "+"},
		{token.PRODUCT, "*"},
		{token.LPAREN, "("},
		{token.INT, "834"},
		{token.SUB, "1"},
		{token.RPAREN, ")"},
		{token.DIVIDE, "/"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.ADD, "+"},
		{token.INT, "1"}
	}
}
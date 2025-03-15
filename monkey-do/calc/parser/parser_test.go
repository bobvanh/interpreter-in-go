package parser

import (
	"calc/ast"
	"calc/lexer"
	"testing"
)

func testExpression(t *testing.T, ex ast.Expression, name string) bool {
	input := `++5
	+ 3 * ((7 + 2 / -3)
	- 2) + --1
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()
}

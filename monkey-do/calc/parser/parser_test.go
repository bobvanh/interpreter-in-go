package parser

import (
	"calc/ast"
	"calc/lexer"
	"testing"
)

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"++1", "++", 1},
		{"--1", "--", 1},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if program.Statement == nil {
			t.Fatalf("program.Statement does not contain a Statement")
		}

		stmt, ok := program.Statement.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statement is not an ast.ExpressionStatement, got=%T",
				program.Statement)
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression, got %T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got %s",
				tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if program.Statement == nil {
			t.Fatalf("program.Statement is empty")
		}

		stmt, ok := program.Statement.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statement in not an ast.ExpressionStatement, got %T",
				program.Statement)
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is nog ast.InfixExpression, got %T",
				stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s', got%s",
				tt.operator, exp.Operator)
		}
	}
}

func TestExpression(t *testing.T) {
	input := `++5
	+ 3 * ((7 + 2 / -3)
	- 2) + --1
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	} else {
		// TODO : Test content
		return
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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral, got %T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d, got=%s",
			value, integ.TokenLiteral())
		return false
	}

	return true
}

package ast

import (
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Expression: nil,
	}

	if program.String() != "" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

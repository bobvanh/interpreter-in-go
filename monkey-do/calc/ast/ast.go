package ast

import (
	"bytes"
	"calc/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// the calc language is for one-liners
// hence a single statment
type Program struct {
	Expression Expression
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) expressionNode() {
	// noop
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.TokenLiteral()
	}
	return ""
}

func (p *Program) TokenLiteral() string {
	if p.Expression != nil {
		return p.Expression.TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	out.WriteString(p.Expression.String())
	return out.String()
}

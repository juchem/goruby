package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// PrefixExpression represents a prefix operator
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (pe *PrefixExpression) Pos() int { return pe.Token.Pos }

// End returns the end of the right expression
func (pe *PrefixExpression) End() int { return pe.Right.End() }

// TokenLiteral returns the literal from the prefix operator token
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// An InfixExpression represents an infix operator in the AST
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// MustEvaluateRight returns true if it is mandatory to evaluate the right side
// of the operator, false otherwise
func (oe *InfixExpression) MustEvaluateRight() bool {
	return oe.Token.Type != token.LOGICALOR
}

// IsControlExpression returns true if the infix is used for control flow,
// false otherwise
func (oe *InfixExpression) IsControlExpression() bool {
	return oe.Token.Type == token.LOGICALOR || oe.Token.Type == token.LOGICALAND
}

func (oe *InfixExpression) expressionNode() {}

// Pos returns the position of first character belonging to the left node
func (oe *InfixExpression) Pos() int { return oe.Left.Pos() }

// End returns the position of last character belonging to the right node
func (oe *InfixExpression) End() int { return oe.Right.End() }

// TokenLiteral returns the literal from the infix operator token
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

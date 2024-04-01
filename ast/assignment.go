package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// Assignment represents a generic assignment
type Assignment struct {
	Token token.Token
	Left  Expression
	Right Expression
}

func (a *Assignment) String() string {
	var out bytes.Buffer
	out.WriteString(encloseInParensIfNeeded(a.Left))
	out.WriteString(" = ")
	out.WriteString(encloseInParensIfNeeded(a.Right))
	return out.String()
}
func (a *Assignment) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (a *Assignment) Pos() int { return a.Left.Pos() }

// End returns the position of first character immediately after the node
func (a *Assignment) End() int { return a.Right.End() }

// TokenLiteral returns the literal of the ASSIGN token
func (a *Assignment) TokenLiteral() string { return a.Token.Literal }

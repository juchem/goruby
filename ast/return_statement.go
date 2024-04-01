package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// A ReturnStatement represents a return node which yields another Expression.
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	return out.String()
}
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the 'return' token literal
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Pos returns the position of first character belonging to the node
func (rs *ReturnStatement) Pos() int { return rs.Token.Pos }

// End returns the position of first character immediately after the node
func (rs *ReturnStatement) End() int { return rs.ReturnValue.End() }

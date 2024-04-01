package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// A ArrayExpansion represents an array expansion node which yields an ExpressionList.
type ArrayExpansion struct {
	Token token.Token // the 'return' token
	Value Expression
}

func (ae *ArrayExpansion) String() string {
	var out bytes.Buffer
	out.WriteString(ae.TokenLiteral())
	if ae.Value != nil {
		out.WriteString(ae.Value.String())
	}
	return out.String()
}

func (ae *ArrayExpansion) expressionNode() {}

// TokenLiteral returns the 'return' token literal
func (ae *ArrayExpansion) TokenLiteral() string { return ae.Token.Literal }

// Pos returns the position of first character belonging to the node
func (ae *ArrayExpansion) Pos() int { return ae.Token.Pos }

// End returns the position of first character immediately after the node
func (ae *ArrayExpansion) End() int { return ae.Value.End() }

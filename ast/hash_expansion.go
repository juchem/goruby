package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// A ArrayExpansion represents an array expansion node which yields an ExpressionList.
type HashExpansion struct {
	Token token.Token // the 'return' token
	Value Expression
}

func (he *HashExpansion) String() string {
	var out bytes.Buffer
	out.WriteString(he.TokenLiteral())
	if he.Value != nil {
		out.WriteString(he.Value.String())
	}
	return out.String()
}

func (he *HashExpansion) expressionNode() {}

// TokenLiteral returns the 'return' token literal
func (he *HashExpansion) TokenLiteral() string { return he.Token.Literal }

// Pos returns the position of first character belonging to the node
func (he *HashExpansion) Pos() int { return he.Token.Pos }

// End returns the position of first character immediately after the node
func (he *HashExpansion) End() int { return he.Value.End() }

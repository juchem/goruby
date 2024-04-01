package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// ScopedIdentifier represents a scoped Constant declaration
type ScopedIdentifier struct {
	Token token.Token // the token.SCOPE
	Outer *Identifier
	Inner Expression
}

func (i *ScopedIdentifier) String() string {
	var out bytes.Buffer
	out.WriteString(i.Outer.String())
	out.WriteString(i.Token.Literal)
	out.WriteString(i.Inner.String())
	return out.String()
}
func (i *ScopedIdentifier) expressionNode() {}
func (i *ScopedIdentifier) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (i *ScopedIdentifier) Pos() int { return i.Outer.Pos() }

// End returns the position of first character immediately after the node
func (i *ScopedIdentifier) End() int { return i.Inner.End() }

// TokenLiteral returns the literal of the token.SCOPE token
func (i *ScopedIdentifier) TokenLiteral() string { return i.Token.Literal }

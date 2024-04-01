package ast

import (
	"fmt"

	"github.com/goruby/goruby/token"
)

// Boolean represents a boolean in the AST
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (b *Boolean) Pos() int { return b.Token.Pos }

// End returns the position of first character immediately after the node
func (b *Boolean) End() int { return b.Token.Pos + len(fmt.Sprintf("%t", b.Value)) }

// TokenLiteral returns the literal from the token token.BOOLEAN
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return fmt.Sprintf("%t", b.Value) }

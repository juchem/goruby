package ast

import (
	"github.com/goruby/goruby/token"
)

// Nil represents the 'nil' keyword
type Nil struct {
	Token token.Token
}

func (n *Nil) expressionNode() {}
func (n *Nil) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (n *Nil) Pos() int { return n.Token.Pos }

// End returns the position of first character immediately after the node
func (n *Nil) End() int { return n.Token.Pos + 3 }

// TokenLiteral returns the literal from the token token.NIL
func (n *Nil) TokenLiteral() string { return n.Token.Literal }
func (n *Nil) String() string       { return "nil" }

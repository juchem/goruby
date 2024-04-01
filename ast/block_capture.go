package ast

import (
	"github.com/goruby/goruby/token"
)

// A BlockCapture represents a function scoped variable capturing a block
type BlockCapture struct {
	Token token.Token // the `&`
	Name  *Identifier
}

func (b *BlockCapture) expressionNode() {}
func (b *BlockCapture) literalNode()    {}

// Pos returns the position of the ampersand
func (b *BlockCapture) Pos() int { return b.Token.Pos }

// End returns the position of the last character of Name
func (b *BlockCapture) End() int { return b.Name.End() }
func (b *BlockCapture) String() string {
	return "&" + b.Name.Value
}

// TokenLiteral returns the literal of the token
func (b *BlockCapture) TokenLiteral() string { return b.Token.Literal }

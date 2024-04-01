package ast

import (
	"github.com/goruby/goruby/token"
)

// An Identifier represents an identifier in the program
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) String() string  { return i.Value }
func (i *Identifier) expressionNode() {}
func (i *Identifier) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (i *Identifier) Pos() int { return i.Token.Pos }

// End returns the position of first character immediately after the node
func (i *Identifier) End() int { return i.Token.Pos + len(i.Value) }

// IsConstant returns true if the Identifier represents a Constant, false otherwise
func (i *Identifier) IsConstant() bool { return i.Token.Type == token.CONST }

// TokenLiteral returns the literal of the token.IDENT token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

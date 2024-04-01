package ast

import (
	"github.com/goruby/goruby/token"
)

// Global represents a global in the AST
type Global struct {
	Token token.Token // the token.GLOBAL token
	Value string
}

func (g *Global) String() string  { return g.Value }
func (g *Global) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (g *Global) Pos() int { return g.Token.Pos }

// End returns the position of first character immediately after the node
func (g *Global) End() int     { return g.Token.Pos + len(g.Value) }
func (g *Global) literalNode() {}

// TokenLiteral returns the literal of the token.GLOBAL token
func (g *Global) TokenLiteral() string { return g.Token.Literal }

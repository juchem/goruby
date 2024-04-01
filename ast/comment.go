package ast

import (
	"github.com/goruby/goruby/token"
)

// Comment represents a double quoted string in the AST
type Comment struct {
	Token token.Token // the #
	Value string
}

func (c *Comment) statementNode() {}
func (c *Comment) literalNode()   {}

// Pos returns the position of first character belonging to the node
func (c *Comment) Pos() int { return c.Token.Pos }

// End returns the position of first character immediately after the node
func (c *Comment) End() int { return c.Token.Pos + len(c.Value) }

// TokenLiteral returns the literal from token token.STRING
func (c *Comment) TokenLiteral() string { return c.Token.Literal }
func (c *Comment) String() string       { return c.Value }

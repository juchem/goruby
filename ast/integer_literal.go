package ast

import (
	"fmt"

	"github.com/goruby/goruby/token"
)

// IntegerLiteral represents an integer in the AST
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (il *IntegerLiteral) Pos() int { return il.Token.Pos }

// End returns the position of first character immediately after the node
func (il *IntegerLiteral) End() int { return il.Token.Pos + len(fmt.Sprintf("%d", il.Value)) }

// TokenLiteral returns the literal from the token.INT token
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return fmt.Sprintf("%d", il.Value) }

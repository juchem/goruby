package ast

import (
	"github.com/goruby/goruby/token"
)

// StringLiteral represents a double quoted string in the AST
type StringLiteral struct {
	Token token.Token // the '"'
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (sl *StringLiteral) Pos() int { return sl.Token.Pos }

// End returns the position of first character immediately after the node
func (sl *StringLiteral) End() int { return sl.Token.Pos + len(sl.Value) }

// TokenLiteral returns the literal from token token.STRING
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Value }

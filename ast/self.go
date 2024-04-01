package ast

import (

	"github.com/goruby/goruby/token"
)

// Self represents self in the current context in the program
type Self struct {
	Token token.Token // the token.SELF token
}

func (s *Self) String() string  { return s.Token.Literal }
func (s *Self) expressionNode() {}
func (s *Self) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (s *Self) Pos() int { return s.Token.Pos }

// End returns the position of first character immediately after the node
func (s *Self) End() int { return s.Token.Pos + 4 }

// TokenLiteral returns the literal of the token.SELF token
func (s *Self) TokenLiteral() string { return s.Token.Literal }

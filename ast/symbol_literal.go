package ast

import (
	"github.com/goruby/goruby/token"
)

// SymbolLiteral represents a symbol within the AST
type SymbolLiteral struct {
	Token token.Token // the ':'
	Value Expression
}

func (s *SymbolLiteral) expressionNode() {}
func (s *SymbolLiteral) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (s *SymbolLiteral) Pos() int { return s.Token.Pos }

// End returns the position of first character immediately after the node
func (s *SymbolLiteral) End() int { return s.Value.End() }

// TokenLiteral returns the literal from token token.SYMBOL
func (s *SymbolLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *SymbolLiteral) String() string       { return ":" + s.Value.String() }

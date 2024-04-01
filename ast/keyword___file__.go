package ast

import (
	"github.com/goruby/goruby/token"
)

// Keyword__FILE__ represents __FILE__ in the AST
type Keyword__FILE__ struct {
	Token    token.Token // the token.FILE__ token
	Filename string
}

func (f *Keyword__FILE__) String() string  { return f.Token.Literal }
func (f *Keyword__FILE__) expressionNode() {}
func (f *Keyword__FILE__) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (f *Keyword__FILE__) Pos() int { return f.Token.Pos }

// End returns the position of first character immediately after the node
func (f *Keyword__FILE__) End() int { return f.Token.Pos + 8 }

// TokenLiteral returns the literal of the token.FILE__ token
func (f *Keyword__FILE__) TokenLiteral() string { return f.Token.Literal }

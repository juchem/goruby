package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// An InstanceVariable represents an instance variable in the AST
type InstanceVariable struct {
	Token token.Token
	Name  *Identifier
}

func (i *InstanceVariable) String() string {
	var out bytes.Buffer
	out.WriteString(i.Token.Literal)
	out.WriteString(i.Name.String())
	return out.String()
}
func (i *InstanceVariable) literalNode()    {}
func (i *InstanceVariable) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (i *InstanceVariable) Pos() int { return i.Token.Pos }

// End returns the position of first character immediately after the node
func (i *InstanceVariable) End() int { return i.Name.End() }

// TokenLiteral returns the literal of the AT token
func (i *InstanceVariable) TokenLiteral() string { return i.Token.Literal }

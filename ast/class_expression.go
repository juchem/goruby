package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// ClassExpression represents a module definition
type ClassExpression struct {
	Token      token.Token // The class keyword
	EndToken   token.Token // The end token
	Name       *Identifier // The class name, will always be a const
	SuperClass *Identifier // The superclass, if any
	Body       *BlockStatement
}

func (m *ClassExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (m *ClassExpression) Pos() int { return m.Token.Pos }

// End returns the position of the `end` token
func (m *ClassExpression) End() int { return m.EndToken.Pos }

// TokenLiteral returns the literal from token.CLASS
func (m *ClassExpression) TokenLiteral() string { return m.Token.Literal }
func (m *ClassExpression) String() string {
	var out bytes.Buffer
	out.WriteString(m.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(m.Name.String())
	if m.SuperClass != nil {
		out.WriteString(" ")
		out.WriteString("<")
		out.WriteString(" ")
		out.WriteString(m.SuperClass.String())
	}
	out.WriteString("\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n")
	out.WriteString(" end")
	return out.String()
}

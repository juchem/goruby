package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// ModuleExpression represents a module definition
type ModuleExpression struct {
	Token    token.Token // The module keyword
	EndToken token.Token // The end token
	Name     *Identifier // The module name, will always be a const
	Body     *BlockStatement
}

func (m *ModuleExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (m *ModuleExpression) Pos() int { return m.Token.Pos }

// End returns the position of the `end` token
func (m *ModuleExpression) End() int { return m.EndToken.Pos }

// TokenLiteral returns the literal from token.MODULE
func (m *ModuleExpression) TokenLiteral() string { return m.Token.Literal }
func (m *ModuleExpression) String() string {
	var out bytes.Buffer
	out.WriteString(m.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(m.Name.String())
	out.WriteString("\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n")
	out.WriteString("end")
	return out.String()
}

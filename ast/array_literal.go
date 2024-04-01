package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// ArrayLiteral represents an Array literal within the AST
type ArrayLiteral struct {
	Token    token.Token // the '['
	Rbracket token.Token // the ']'
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) literalNode()    {}

// Pos returns the position of first character belonging to the node
func (al *ArrayLiteral) Pos() int { return al.Token.Pos }

// End returns the position of first character immediately after the node
func (al *ArrayLiteral) End() int {
	return al.Rbracket.Pos
}

// TokenLiteral returns the literal of the token token.LBRACKET
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

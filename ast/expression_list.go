package ast

import (
	"bytes"
	"strings"
)

// ExpressionList represents a list of expressions within the AST divided by commas
type ExpressionList []Expression

func (el ExpressionList) expressionNode() {}
func (el ExpressionList) literalNode()    {}

// Pos returns the position of first character from the first expression
func (el ExpressionList) Pos() int {
	if len(el) == 0 {
		return 0
	}
	return el[0].End()
}

// End returns End of the last element
func (el ExpressionList) End() int {
	if len(el) == 0 {
		return 0
	}
	return el[len(el)-1].End()
}

// TokenLiteral returns the literal of the first element
func (el ExpressionList) TokenLiteral() string {
	if len(el) == 0 {
		return ""
	}
	return el[0].TokenLiteral()
}
func (el ExpressionList) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range el {
		elements = append(elements, e.String())
	}
	out.WriteString(strings.Join(elements, ", "))
	return out.String()
}

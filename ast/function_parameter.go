package ast

import (
	"bytes"
)

// A FunctionParameter represents a parameter in a function literal
type FunctionParameter struct {
	Name    *Identifier
	Default Expression
}

func (f *FunctionParameter) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (f *FunctionParameter) Pos() int { return f.Name.Pos() }

// End returns the position of the default end if it exists, otherwise the end position of Name
func (f *FunctionParameter) End() int {
	if f.Default != nil {
		return f.Default.End()
	}
	return f.Name.End()
}

// TokenLiteral returns the token of the parameter name
func (f *FunctionParameter) TokenLiteral() string { return f.Name.TokenLiteral() }
func (f *FunctionParameter) String() string {
	var out bytes.Buffer
	out.WriteString(f.Name.String())
	if f.Default != nil {
		out.WriteString(" = ")
		out.WriteString(encloseInParensIfNeeded(f.Default))
	}
	return out.String()
}

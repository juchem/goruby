package ast

import (
	"bytes"
	"strings"
)

// MultiAssignment represents multiple variables on the lefthand side
type MultiAssignment struct {
	Variables []*Identifier
	Values    []Expression
}

func (m *MultiAssignment) String() string {
	var out bytes.Buffer
	vars := make([]string, len(m.Variables))
	for i, v := range m.Variables {
		vars[i] = v.Value
	}
	out.WriteString(strings.Join(vars, ", "))
	out.WriteString(" = ")
	values := make([]string, len(m.Values))
	for i, v := range m.Values {
		values[i] = v.String()
	}
	out.WriteString(strings.Join(values, ", "))
	return out.String()
}
func (m *MultiAssignment) literalNode() {}

// Pos returns the position of first character belonging to the node
func (m *MultiAssignment) Pos() int { return m.Variables[0].Pos() }

// End returns the position of first character immediately after the node
func (m *MultiAssignment) End() int        { return m.Values[len(m.Values)-1].End() }
func (m *MultiAssignment) expressionNode() {}

// TokenLiteral returns the literal of the first variable token
func (m *MultiAssignment) TokenLiteral() string { return m.Variables[0].Token.Literal }

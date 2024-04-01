package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// A ContextCallExpression represents a method call on a given Context
type ContextCallExpression struct {
	Token     token.Token      // The '.' token
	Context   Expression       // The lefthandside expression
	Function  *Identifier      // The function to call
	Arguments []Expression     // The function arguments
	Block     *BlockExpression // The function block
}

func (ce *ContextCallExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (ce *ContextCallExpression) Pos() int {
	if ce.Context != nil {
		return ce.Context.Pos()
	}
	return ce.Function.Pos()
}

// End returns the end position of the block if it exists. If not, it returns
// the end position of the last argument if any. Otherwise it returns the end
// of the function identifier
func (ce *ContextCallExpression) End() int {
	if ce.Block != nil {
		return ce.Block.End()
	}
	if len(ce.Arguments) == 0 {
		return ce.Function.End()
	}
	return ce.Arguments[len(ce.Arguments)-1].End()
}

// TokenLiteral returns the literal from token.DOT
func (ce *ContextCallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *ContextCallExpression) String() string {
	var out bytes.Buffer
	if ce.Context != nil {
		out.WriteString(ce.Context.String())
		out.WriteString(".")
	}
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	if ce.Block != nil {
		out.WriteString("\n")
		out.WriteString(ce.Block.String())
	}
	return out.String()
}

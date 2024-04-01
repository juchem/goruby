package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// YieldExpression represents self in the current context in the program
type YieldExpression struct {
	Token     token.Token  // the token.YIELD token
	Arguments []Expression // The arguments to yield
}

func (y *YieldExpression) String() string {
	var out bytes.Buffer
	out.WriteString(y.Token.Literal)
	if len(y.Arguments) != 0 {
		args := []string{}
		for _, a := range y.Arguments {
			args = append(args, a.String())
		}
		out.WriteString(" ")
		out.WriteString(strings.Join(args, ", "))
	}
	return out.String()
}
func (y *YieldExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (y *YieldExpression) Pos() int { return y.Token.Pos }

// End returns the position of first character immediately after the node
func (y *YieldExpression) End() int {
	if len(y.Arguments) == 0 {
		return y.Pos() + 5
	}
	return y.Arguments[len(y.Arguments)-1].End()
}

// TokenLiteral returns the literal of the token.YIELD token
func (y *YieldExpression) TokenLiteral() string { return y.Token.Literal }

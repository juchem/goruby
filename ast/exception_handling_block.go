package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// ExceptionHandlingBlock represents a begin/end block where exceptions are rescued
type ExceptionHandlingBlock struct {
	BeginToken token.Token
	EndToken   token.Token
	TryBody    *BlockStatement
	Rescues    []*RescueBlock
}

func (eh *ExceptionHandlingBlock) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (eh *ExceptionHandlingBlock) Pos() int { return eh.BeginToken.Pos }

// End returns the position of first character immediately after the node
func (eh *ExceptionHandlingBlock) End() int { return eh.EndToken.Pos }

// TokenLiteral returns the token literal from 'begin'
func (eh *ExceptionHandlingBlock) TokenLiteral() string { return eh.BeginToken.Literal }
func (eh *ExceptionHandlingBlock) String() string {
	var out bytes.Buffer
	out.WriteString(eh.BeginToken.Literal)
	out.WriteString("\n")
	out.WriteString(eh.TryBody.String())
	out.WriteString("\n")
	for _, r := range eh.Rescues {
		out.WriteString(r.String())
	}
	out.WriteString("end")
	return out.String()
}

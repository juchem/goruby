package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// A RescueBlock represents a rescue block
type RescueBlock struct {
	Token            token.Token
	ExceptionClasses []*Identifier
	Exception        *Identifier
	Body             *BlockStatement
}

func (rb *RescueBlock) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (rb *RescueBlock) Pos() int { return rb.Token.Pos }

// End returns the position of first character immediately after the node
func (rb *RescueBlock) End() int { return rb.Body.End() }

// TokenLiteral returns the token literal from 'rescue'
func (rb *RescueBlock) TokenLiteral() string { return rb.Token.Literal }
func (rb *RescueBlock) String() string {
	var out bytes.Buffer
	out.WriteString(rb.Token.Literal)
	if len(rb.ExceptionClasses) != 0 {
		classes := make([]string, len(rb.ExceptionClasses))
		for i, c := range rb.ExceptionClasses {
			classes[i] = c.String()
		}
		out.WriteString(strings.Join(classes, ", "))
	}
	if rb.Exception != nil {
		out.WriteString(" => ")
		out.WriteString(rb.Exception.String())
	}
	out.WriteString("\n")
	out.WriteString(rb.Body.String())
	out.WriteString("\n")
	return out.String()
}

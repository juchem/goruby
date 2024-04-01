package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// ConditionalExpression represents an if expression within the AST
type ConditionalExpression struct {
	Token       token.Token // The 'if' or 'unless' token
	EndToken    token.Token // The 'end' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// IsNegated indicates if the condition uses unless, i.e. is negated
func (ce *ConditionalExpression) IsNegated() bool {
	return ce.Token.Type == token.UNLESS
}

func (ce *ConditionalExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (ce *ConditionalExpression) Pos() int {
	if ce.EndToken.Type == token.ILLEGAL {
		return ce.Consequence.Pos()
	}
	return ce.Token.Pos
}

// End returns the position of first character immediately after the node
func (ce *ConditionalExpression) End() int {
	if ce.EndToken.Type == token.ILLEGAL {
		return ce.Consequence.Pos()
	}
	return ce.EndToken.Pos
}

// TokenLiteral returns the literal from token token.IF or token.UNLESS
func (ce *ConditionalExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *ConditionalExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Token.Literal)
	out.WriteString(ce.Condition.String())
	out.WriteString(" ")
	out.WriteString(ce.Consequence.String())
	if ce.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ce.Alternative.String())
	}
	out.WriteString(" end")
	return out.String()
}

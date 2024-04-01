package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// A LoopExpression represents a loop
type LoopExpression struct {
	Token     token.Token // while
	EndToken  token.Token // end
	Condition Expression
	Block     *BlockStatement
}

func (ce *LoopExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (ce *LoopExpression) Pos() int {
	return ce.Token.Pos
}

// End returns the position of first character immediately after the node
func (ce *LoopExpression) End() int {
	return ce.EndToken.Pos
}

// TokenLiteral returns the literal from token token.WHILE
func (ce *LoopExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *LoopExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Token.Literal)
	out.WriteString(ce.Condition.String())
	out.WriteString(" do ")
	out.WriteString(ce.Block.String())
	out.WriteString(" end")
	return out.String()
}

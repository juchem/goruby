package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// An IndexExpression represents an array or hash access in the AST
type IndexExpression struct {
	Token  token.Token // The [ token
	Left   Expression
	Index  Expression
	Length Expression
}

func (ie *IndexExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (ie *IndexExpression) Pos() int { return ie.Token.Pos }

// End returns the position of the last character belonging to the node
func (ie *IndexExpression) End() int { return ie.Index.End() }

// TokenLiteral returns the literal from token.LBRACKET
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	if ie.Length != nil {
		out.WriteString(", ")
		out.WriteString(ie.Index.String())
	}
	out.WriteString("])")
	return out.String()
}

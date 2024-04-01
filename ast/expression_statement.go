package ast

import (
	"github.com/goruby/goruby/token"
)

// An ExpressionStatement is a Statement wrapping an Expression
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) statementNode() {}

// Pos returns the position of first character belonging to the node
func (es *ExpressionStatement) Pos() int { return es.Expression.Pos() }

// End returns the position of first character immediately after the node
func (es *ExpressionStatement) End() int { return es.Expression.End() }

// TokenLiteral returns the first token of the Expression
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

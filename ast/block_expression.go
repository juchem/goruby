package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// A BlockExpression represents a Ruby block
type BlockExpression struct {
	Token      token.Token          // token.DO or token.LBRACE
	EndToken   token.Token          // token.END or token.RBRACE
	Parameters []*FunctionParameter // the block parameters
	Body       *BlockStatement      // the block body
}

func (b *BlockExpression) expressionNode() {}

// Pos returns the position of first character belonging to the node
func (b *BlockExpression) Pos() int { return b.Token.Pos }

// End returns the position of the end token
func (b *BlockExpression) End() int { return b.EndToken.Pos }

// TokenLiteral returns the literal from the Token
func (b *BlockExpression) TokenLiteral() string { return b.Token.Literal }

// String returns a string representation of the block statement
func (b *BlockExpression) String() string {
	var out bytes.Buffer
	out.WriteString(b.Token.Literal)
	if len(b.Parameters) != 0 {
		args := []string{}
		for _, a := range b.Parameters {
			args = append(args, a.String())
		}
		out.WriteString("|")
		out.WriteString(strings.Join(args, ", "))
		out.WriteString("|")
		out.WriteString("\n")
	}
	out.WriteString(b.Body.String())
	out.WriteString("\n")
	if b.Token.Type == token.LBRACE {
		out.WriteString("}")
	} else {
		out.WriteString("end")
	}
	return out.String()
}

package ast

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/token"
)

// A FunctionLiteral represents a function definition in the AST
type FunctionLiteral struct {
	Token         token.Token // The 'def' token
	EndToken      token.Token // the 'end' token
	Receiver      *Identifier
	Name          *Identifier
	Parameters    []*FunctionParameter
	CapturedBlock *BlockCapture
	Body          *BlockStatement
	Rescues       []*RescueBlock
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) literalNode()    {}

// Pos returns the position of the `def` keyword
func (fl *FunctionLiteral) Pos() int { return fl.Token.Pos }

// End returns the position of the `end` keyword
func (fl *FunctionLiteral) End() int { return fl.EndToken.Pos }

// TokenLiteral returns the literal from token.DEF
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	if fl.CapturedBlock != nil {
		params = append(params, fl.CapturedBlock.String())
	}
	out.WriteString("def ")
	if fl.Receiver != nil {
		out.WriteString(fl.Receiver.String())
		out.WriteString(".")
	}
	out.WriteString(fl.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	if fl.Body != nil {
		out.WriteString(fl.Body.String())
	}
	for _, r := range fl.Rescues {
		out.WriteString(r.String())
	}
	out.WriteString(" end")
	return out.String()
}

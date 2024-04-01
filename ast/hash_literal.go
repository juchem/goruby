package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/goruby/goruby/token"
)

// HashLiteral represents an Hash literal within the AST
type HashLiteral struct {
	Token  token.Token // the '{'
	Rbrace token.Token // the '{'
	Map    map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}
func (hl *HashLiteral) literalNode()    {}

// Pos returns the position of the left brace
func (hl *HashLiteral) Pos() int { return hl.Token.Pos }

// End returns the position of the right brace
func (hl *HashLiteral) End() int { return hl.Rbrace.Pos }

// TokenLiteral returns the literal of the token token.LBRACE
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for key, val := range hl.Map {
		elements = append(elements, fmt.Sprintf("%q => %q", key.String(), val.String()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("}")
	return out.String()
}

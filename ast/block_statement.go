package ast

import (
	"bytes"

	"github.com/goruby/goruby/token"
)

// BlockStatement represents a list of statements
type BlockStatement struct {
	// the { token or the first token from the first statement
	Token      token.Token
	EndToken   token.Token // the } token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// Pos returns the position of first character belonging to the node
func (bs *BlockStatement) Pos() int { return bs.Token.Pos }

// End returns the position of first character immediately after the node
func (bs *BlockStatement) End() int { return bs.EndToken.Pos }

// TokenLiteral returns '{' or the first token from the first statement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		if s != nil {
			out.WriteString(s.String())
		}
	}
	return out.String()
}

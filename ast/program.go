package ast

import (
	gotoken "go/token"
	"strings"
)

// A Program node is the root node within the AST.
type Program struct {
	pos        int
	File       *gotoken.File
	Statements []Statement
}

// Pos returns the position of first character belonging to the node
func (p *Program) Pos() int { return p.pos }

// End returns the position of first character immediately after the node
func (p *Program) End() int {
	if len(p.Statements) == 0 {
		return p.pos
	}
	return p.Statements[len(p.Statements)-1].End()
}
func (p *Program) String() string {
	stmts := make([]string, len(p.Statements))
	for i, s := range p.Statements {
		if s != nil {
			stmts[i] = s.String()
		}
	}
	return strings.Join(stmts, "\n")
}

// TokenLiteral returns the literal of the first statement and empty string if
// there is no statement.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

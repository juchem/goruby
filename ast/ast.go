package ast

import (
	"strings"
)

func encloseInParensIfNeeded(expr Expression) string {
	val := expr.String()
	hasParens := strings.HasPrefix(val, "(") && strings.HasSuffix(val, ")")
	_, isLiteral := expr.(literal)
	if !isLiteral && !hasParens {
		val = "(" + val + ")"
	}
	return val
}

package ast

// literal
type literal interface {
	Node
	literalNode()
}

// IsLiteral returns true if n is a literal node, false otherwise
func IsLiteral(n Node) bool {
	_, ok := n.(literal)
	return ok
}

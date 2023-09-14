package expressions

import "strings"

// interface for all other expressions to satisfy
type Expr interface {
	Eval() any
	Debug(b *strings.Builder)
}

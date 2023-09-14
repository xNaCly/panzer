package expressions

import "strings"

type Expr interface {
	Eval() any
	Debug(b *strings.Builder)
}

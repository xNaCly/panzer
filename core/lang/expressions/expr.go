package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

// interface for all other expressions to satisfy
type Expr interface {
	Eval() any
	Debug(b *strings.Builder)
	GetToken() tokens.Token
}

func Debug(ast []Expr, b *strings.Builder) string {
	for _, c := range ast {
		c.Debug(b)
	}
	str := b.String()
	b.Reset()
	return str
}

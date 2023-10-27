package expressions

import (
	"panzer/core/lang/tokens"
	"strings"
)

// interface for all other expressions to satisfy
type Expr interface {
	Eval() any
	Debug(b *strings.Builder)
	GetToken() tokens.Token
}

func Debug(ast []Expr, b *strings.Builder) string {
	astL := len(ast)
	for i, c := range ast {
		c.Debug(b)
		if i+1 < astL {
			b.WriteRune('\n')
		}
	}
	str := b.String()
	b.Reset()
	return str
}

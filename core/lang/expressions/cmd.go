package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

type Cmd struct {
	Token     tokens.Token
	Name      Expr
	Arguments []Expr
}

func (c *Cmd) Eval() any {
	return ""
}

func (c *Cmd) Debug(b *strings.Builder) {
	c.Name.Debug(b)
	b.WriteRune('[')
	cl := len(c.Arguments)
	for i, ch := range c.Arguments {
		ch.Debug(b)
		if i+1 < cl {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(']')
}

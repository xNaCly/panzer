package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

// everything passed to the shell which isnt a build in, string or any symbol
// computed from: ls -la -> Cmd{Name:"ls", Arguments: []string{"-la"}}
type Cmd struct {
	Token     tokens.Token
	Name      Expr
	Arguments []Expr
}

func (c *Cmd) Eval() any {
	return ""
}

func (c *Cmd) Debug(b *strings.Builder) {
	b.WriteString("CMD:")
	b.WriteString(c.Token.Raw)
	b.WriteRune('[')
	for i, ch := range c.Arguments {
		ch.Debug(b)
		if i+1 < len(c.Arguments) {
			b.WriteRune(',')
		}
	}
	b.WriteRune(']')
}

func (c *Cmd) GetToken() tokens.Token {
	return c.Token
}

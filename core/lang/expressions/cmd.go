package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

// everything passed to the shell which isnt a build in, string or any symbol
// computed from: ls -la -> Cmd{Name:"ls", Arguments: []string{"-la"}}
type Cmd struct {
	Token     tokens.Token
	Name      string
	Arguments []string
}

func (c *Cmd) Eval() any {
	return ""
}

func (c *Cmd) Debug(b *strings.Builder) {
	b.WriteString(c.Name)
	b.WriteRune('[')
	cl := len(c.Arguments)
	for i, ch := range c.Arguments {
		b.WriteString(ch)
		if i+1 < cl {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(']')
}

func (c *Cmd) GetToken() tokens.Token {
	return c.Token
}

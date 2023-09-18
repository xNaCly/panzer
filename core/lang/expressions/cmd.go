package expressions

import (
	"gopnzr/core/lang/tokens"
	"os"
	"os/exec"
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
	l := len(c.Arguments)
	args := make([]string, l)
	for i := 0; i < l; i++ {
		str := c.Arguments[i].GetToken().Raw
		args[i] = str
	}
	cmd := exec.Command(c.Name.GetToken().Raw, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return nil
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

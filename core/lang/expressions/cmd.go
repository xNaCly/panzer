package expressions

import (
	"gopnzr/core/lang/tokens"
	"gopnzr/core/shell/expansions"
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
	args := make([]string, 0)
	for _, e := range c.Arguments {
		t := e.GetToken()
		if t.ContainsPattern {
			args = append(args, expansions.MatchFiles(t.Raw)...)
		} else {
			val, ok := e.Eval().(string)
			if !ok {
				panic("failure in command prep")
			}
			args = append(args, val)
		}
	}
	val, ok := c.Name.Eval().(string)
	if !ok {
		panic("failure in command prep")
	}
	cmd := exec.Command(val, args...)
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

package expressions

import (
	"gopnzr/core/lang/tokens"
	"gopnzr/core/shell/expansions"
	"os"
	"os/exec"
	"strings"
)

// TODO: what about quitting running tasks, maybe start commands with context,
// store them in a map, then filter and kill if ^C is hit

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

	cmd := exec.Command(c.Name.GetToken().Raw, args...)

	// TODO: enable piping by replacing these with global buffers, such as Pipe?
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		succ := cmd.ProcessState.ExitCode()
		if succ == -1 {
			panic(err)
		}
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

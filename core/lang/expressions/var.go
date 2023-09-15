package expressions

import (
	"fmt"
	"gopnzr/core/lang/tokens"
	"gopnzr/core/shell/env"
	"strings"
)

type Var struct {
	Token tokens.Token
	Name  string
}

func (v *Var) Eval() any {
	val, ok := env.GetEnv(v.Name)
	if !ok {
		panic(fmt.Sprintf("Undefined variable $%s\n", val))
	}
	return val
}

func (v *Var) Debug(b *strings.Builder) {
	b.WriteRune('$')
	b.WriteString(v.Name)
}

func (v *Var) GetToken() tokens.Token {
	return v.Token
}

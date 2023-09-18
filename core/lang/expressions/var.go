package expressions

import (
	"gopnzr/core/lang/tokens"
	"gopnzr/core/shell/env"
	"strings"
)

type Var struct {
	Token tokens.Token
}

func (v *Var) Eval() any {
	val, ok := env.GetEnv(v.Token.Raw)
	if !ok {
		return ""
	}
	return val
}

func (v *Var) Debug(b *strings.Builder) {
	b.WriteRune('$')
	b.WriteString(v.Token.Raw)
}

func (v *Var) GetToken() tokens.Token {
	return v.Token
}

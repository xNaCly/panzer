package expressions

import (
	"gopnzr/core/lang/keywords"
	"gopnzr/core/lang/tokens"
	"strings"
)

type Keyword struct {
	Token     tokens.Token
	Arguments []Expr
}

func (k *Keyword) Eval() any {
	l := len(k.Arguments)
	args := make([]string, l)
	for i := 0; i < l; i++ {
		if str, ok := k.Arguments[i].Eval().(string); ok {
			args[i] = str
		}
	}
	keywords.KEYWORD_MAP[k.Token.Raw](args...)
	return nil
}

func (k *Keyword) Debug(b *strings.Builder) {
	b.WriteString(k.Token.Raw)
	b.WriteRune('[')
	for i, c := range k.Arguments {
		c.Debug(b)
		if i+1 < len(k.Arguments) {
			b.WriteRune(',')
		}
	}
	b.WriteRune(']')
}

func (k *Keyword) GetToken() tokens.Token {
	return k.Token
}

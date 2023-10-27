package expressions

import (
	"panzer/core/lang/keywords"
	"panzer/core/lang/tokens"
	"os"
	"strings"
)

type Keyword struct {
	Token     tokens.Token
	Arguments []Expr
}

func (k *Keyword) Eval() any {
	// TEST: does this make execution faster, due to skipping argument construction?
	if k.Token.Raw == "exit" {
		os.Exit(0)
	}
	l := len(k.Arguments)
	args := make([]string, l)
	for i := 0; i < l; i++ {
		str := castOrPanic[string](k.Arguments[i].Eval())
		args[i] = str
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

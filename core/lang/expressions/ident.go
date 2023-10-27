package expressions

import (
	"panzer/core/lang/tokens"
	"strings"
)

type Ident struct {
	Token tokens.Token
}

func (i *Ident) Eval() any {
	return i.Token.Raw
}

func (i *Ident) Debug(b *strings.Builder) {
	b.WriteString(i.Token.Raw)
}

func (i *Ident) GetToken() tokens.Token {
	return i.Token
}

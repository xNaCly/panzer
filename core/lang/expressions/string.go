package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

type String struct {
	Token tokens.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}

func (s *String) Debug(b *strings.Builder) {
	b.WriteString(s.Token.Raw)
}

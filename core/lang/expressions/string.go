package expressions

import (
	"gopnzr/core/lang/tokens"
	"strings"
)

// all characters between " and "
type String struct {
	Token tokens.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}

func (s *String) Debug(b *strings.Builder) {
	b.WriteString(s.Token.Raw)
}

func (s *String) GetToken() tokens.Token {
	return s.Token
}

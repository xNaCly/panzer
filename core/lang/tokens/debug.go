package tokens

import "strings"

func Debug(t []Token, b *strings.Builder) string {
	for _, c := range t {
		b.WriteString(LOOKUP[c.Type])
		b.WriteRune('{')
		b.WriteString(c.Raw)
		b.WriteRune('}')
		b.WriteRune('\n')
	}
	str := b.String()
	b.Reset()
	return str
}

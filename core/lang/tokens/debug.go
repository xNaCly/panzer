package tokens

import "strings"

func Debug(t []Token) string {
	b := strings.Builder{}
	for _, c := range t {
		b.WriteString(LOOKUP[c.Type])
		b.WriteRune('{')
		b.WriteString(c.Raw)
		b.WriteRune('}')
		b.WriteRune('\n')
	}
	return b.String()
}

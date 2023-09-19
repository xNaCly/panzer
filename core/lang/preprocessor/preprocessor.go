// responsible for replacing aliases and variables the source
package preprocessor

import (
	"gopnzr/core/shell/env"
	"strings"
	"unicode"
)

type Preprocessor struct {
	pos     int
	in      []rune
	inL     int
	cc      rune
	Builder *strings.Builder
}

var tempBuilder = strings.Builder{}

func (p *Preprocessor) NewInput(input string) {
	i := []rune(input)
	p.pos = 0
	p.in = i
	p.inL = len(input)
	p.cc = i[0]
	p.Builder.Reset()
}

func (p *Preprocessor) Process() string {
	for p.cc != 0 {
		if p.cc == '$' {
			p.advance()
			for (unicode.IsLetter(p.cc) || p.cc == '_') && p.cc != 0 {
				tempBuilder.WriteRune(p.cc)
				p.advance()
			}
			if val, ok := env.GetEnv(tempBuilder.String()); ok {
				p.Builder.WriteString(val)
			} else {
				p.Builder.WriteRune('$')
				p.Builder.WriteString(tempBuilder.String())
			}
			tempBuilder.Reset()
		}
		p.Builder.WriteRune(p.cc)
		p.advance()
	}

	return p.Builder.String()
}

func (p *Preprocessor) advance() {
	if p.pos+1 < p.inL {
		p.pos++
		p.cc = p.in[p.pos]
	} else {
		p.cc = 0
	}
}

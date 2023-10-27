// responsible for replacing aliases and variables the source
package preprocessor

import (
	"fmt"
	"panzer/core/shell/args"
	"panzer/core/shell/env"
	"panzer/core/state"
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

func matchesAny(r rune, runes ...rune) bool {
	for _, cr := range runes {
		if cr == r {
			return true
		}
	}
	return false
}

func (p *Preprocessor) Process(a *args.Arguments) string {
	for p.cc != 0 {
		// expansion of variables
		if p.cc == '$' {
			p.advance() // skip dollar
			for (unicode.IsLetter(p.cc) || p.cc == '_') && p.cc != 0 && p.cc != '\n' {
				tempBuilder.WriteRune(p.cc)
				p.advance()
			}
			res := tempBuilder.String()
			if val, ok := env.GetEnv(res); ok {
				p.Builder.WriteString(val)
				if a.Debug {
					fmt.Printf("found variable %q, replaced it with %q\n", res, val)
				}
			} else {
				if a.Debug {
					fmt.Printf("unknown variable %q\n", res)
				}
				p.Builder.WriteRune('$')
				p.Builder.WriteString(res)
			}
			tempBuilder.Reset()
		}

		// TODO: this needs a fix, only commands should be replaced, no command arguments
		// FIX: (TEMP): allow alias expansion only at the start of the expression
		if p.prev() == 0 {
			if unicode.IsLetter(p.cc) || p.cc == '_' || p.cc == '.' {
				for (unicode.IsLetter(p.cc) || p.cc == '_' || p.cc == '.') && p.cc != 0 && p.cc != '\n' {
					tempBuilder.WriteRune(p.cc)
					p.advance()
				}

				if tempBuilder.Len() == 0 {
					continue
				}

				res := tempBuilder.String()

				if val, ok := state.ALIASES[res]; ok {
					p.Builder.WriteString(val)
					if a.Debug {
						fmt.Printf("found alias %q, replaced it with %q\n", res, val)
					}
				} else {
					if a.Debug {
						fmt.Printf("unknown alias %q\n", res)
					}
					p.Builder.WriteString(res)
				}
				tempBuilder.Reset()
			}
		}

		p.Builder.WriteRune(p.cc)
		p.advance()
	}

	return p.Builder.String()
}

func (p *Preprocessor) prev() rune {
	if p.pos == 0 {
		return 0
	} else {
		return p.in[p.pos-1]
	}
}

func (p *Preprocessor) advance() {
	if p.pos+1 < p.inL {
		p.pos++
		p.cc = p.in[p.pos]
	} else {
		p.cc = 0
	}
}

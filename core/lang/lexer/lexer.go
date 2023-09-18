// transforms a stream of characters into tokens with meaning attached,
// prepares input for processing by the parser
package lexer

import (
	"fmt"
	"gopnzr/core/lang/tokens"
	"strings"
)

type Lexer struct {
	pos     int
	lPos    int // position on the current line
	l       int // line position of total buffer
	in      string
	inL     int
	cc      byte // current character
	Builder *strings.Builder
}

func (l *Lexer) NewInput(input string) {
	l.Builder.Reset()
	l.pos = 0
	l.lPos = 0
	l.in = input
	l.l = 0
	l.inL = len(input)
	l.cc = input[0]
}

func (l *Lexer) Lex() []tokens.Token {
	t := make([]tokens.Token, 0)
	for l.cc != 0 {
		tt := tokens.UNKNOWN
		switch l.cc {
		case ' ', '\t':
			l.advance()
			continue
		case '\n':
			// insert semicolon on enter, this enables the configuration
			t = append(t, tokens.Token{
				Pos:  l.pos,
				Type: tokens.SEMICOLON,
			})
			l.lPos = 0
			l.l++
			l.advance()
			continue
		case '|':
			tt = tokens.PIPE
		case '&':
			tt = tokens.AND
		case ';':
			tt = tokens.SEMICOLON
		case '$':
			tt = tokens.DOLLAR
		case '"':
			i := l.string()
			t = append(t, tokens.Token{
				Pos:  l.pos,
				Type: tokens.STRING,
				Raw:  i,
			})
			continue
		default:
			i := l.ident()
			tt := tokens.IDENT
			if _, isKeyword := tokens.KEYWORDS[i]; isKeyword {
				tt = tokens.KEYWORD
			}
			t = append(t, tokens.Token{
				Pos:  l.pos,
				Type: tt,
				Raw:  i,
			})
			continue
		}
		if tt != tokens.UNKNOWN {
			t = append(t, tokens.Token{
				Pos:  l.pos,
				Type: tt,
				Raw:  string(l.cc),
			})
		}
		l.advance()
	}
	t = append(t, tokens.Token{
		Pos:  l.lPos,
		Type: tokens.EOF,
		Raw:  "EOF",
	})
	return t
}

func (l *Lexer) ident() string {
	if l.cc == '=' {
		panic("got = at identifier, you probably wanted to define an environment variable, use 'set VARIABLE_NAME value' instead")
	}
	l.Builder.WriteByte(l.cc)
	l.advance()
	for !l.matchAny('&', '|', '=', ';', '$', 0, ' ') {
		l.Builder.WriteByte(l.cc)
		l.advance()
	}
	str := l.Builder.String()
	l.Builder.Reset()
	return str
}

func (l *Lexer) string() string {
	l.advance() // skip "
	for l.cc != 0 && l.cc != '"' {
		l.Builder.WriteByte(l.cc)
		l.advance()
	}

	str := l.Builder.String()

	if l.cc != '"' {
		panic(fmt.Sprintf("Unterminated string %q\n", str))
	}

	l.advance() // skip "
	l.Builder.Reset()
	return str
}

func (l *Lexer) matchAny(ts ...rune) bool {
	lcr := rune(l.cc)
	for _, r := range ts {
		if r == lcr {
			return true
		}
	}
	return false
}

func (l *Lexer) advance() {
	if l.pos+1 < l.inL {
		l.pos++
		l.lPos++
		l.cc = l.in[l.pos]
	} else {
		l.cc = 0
	}
}

// transforms a stream of characters into tokens with meaning attached,
// prepares input for processing by the parser
package lexer

import (
	"fmt"
	"gopnzr/core/lang/tokens"
	"os"
	"strings"
)

// we reuse the builder because we need it for every string and identifier -
// why reallocate?
var builder = strings.Builder{}

type Lexer struct {
	pos   int
	lPos  int // position on the current line
	l     int // line position of total buffer
	in    string
	inL   int
	cc    byte // current character
	error bool
}

func (l *Lexer) NewInput(input string) {
	l.pos = 0
	l.lPos = 0
	l.in = input
	l.l = 0
	l.inL = len(input)
	l.cc = input[0]
	l.error = false
}

func (l *Lexer) Lex() []tokens.Token {
	t := make([]tokens.Token, 0)
	for l.cc != 0 && !l.error {
		tt := tokens.UNKNOWN
		switch l.cc {
		case ' ', '\t':
			l.advance()
			continue
		case '\n':
			l.lPos = 0
			l.l++
			l.advance()
			continue
		case '|':
			tt = tokens.PIPE
		case '&':
			tt = tokens.AND
		case '*':
			tt = tokens.ASTERIKS
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
			if (l.cc >= 'a' && l.cc <= 'z') || (l.cc >= 'A' && l.cc <= 'Z') || l.cc == '-' || l.cc == '/' {
				i := l.ident()
				t = append(t, tokens.Token{
					Pos:  l.pos,
					Type: tokens.IDENT,
					Raw:  i,
				})
				continue
			} else {
				fmt.Fprintf(os.Stderr, "unknown character %q\n", l.cc)
			}
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
	if l.error {
		return []tokens.Token{}
	}
	t = append(t, tokens.Token{
		Pos:  l.lPos,
		Type: tokens.EOF,
		Raw:  "EOF",
	})
	return t
}

func (l *Lexer) ident() string {
	builder.WriteByte(l.cc)
	l.advance()
	for (l.cc >= 'a' && l.cc <= 'z') || (l.cc >= 'A' && l.cc <= 'Z') || l.cc == '_' || (l.cc >= '0' && l.cc <= '9') || l.cc == '-' || l.cc == '/' {
		builder.WriteByte(l.cc)
		l.advance()
	}
	str := builder.String()
	builder.Reset()
	return str
}

func (l *Lexer) string() string {
	l.advance() // skip "
	for l.cc != 0 && l.cc != '"' {
		builder.WriteByte(l.cc)
		l.advance()
	}

	str := builder.String()

	if l.cc != '"' {
		fmt.Fprintf(os.Stderr, "unterminated string %q\n", str)
		l.error = true
		return ""
	}

	l.advance() // skip "
	builder.Reset()
	return str
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

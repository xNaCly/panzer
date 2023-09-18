// parses a stream of tokens into an abstract syntax tree:
//
//		str      -> ".+"
//		ident    -> [A-Za-z\-]+[_0-9/]*
//		const    -> str | ident
//		var      -> 'set' ident const
//		cmd      -> const+
//	    cmdSep   -> '|' | ';' | '&&' | '>>' | '>'
//		cmd      -> cmd cmdSep cmd
//		cmd      -> cmd &
//		cmd      -> cmd+
package parser

import (
	"fmt"
	"gopnzr/core/lang/expressions"
	"gopnzr/core/lang/tokens"
)

type Parser struct {
	pos int
	in  []tokens.Token
}

func (p *Parser) NewInput(input []tokens.Token) {
	p.pos = 0
	p.in = input
}

func (p *Parser) Parser() []expressions.Expr {
	r := make([]expressions.Expr, 0)
	for p.peek().Type != tokens.EOF {
		e := p.parse()
		r = append(r, e)
		p.advance()
	}
	return r
}

func (p *Parser) parse() expressions.Expr {
	var expr expressions.Expr
	if p.peekEquals(tokens.KEYWORD) {
		expr = p.keywords()
	} else if p.matchAny(tokens.IDENT) {
		expr = p.command()
	} else {
		expr = p.consts()
	}
	return expr
}

func (p *Parser) consts() expressions.Expr {
	var expr expressions.Expr
	if p.peekEquals(tokens.STRING) {
		expr = &expressions.String{Token: p.peek()}
	} else if p.peekEquals(tokens.IDENT) {
		expr = &expressions.Ident{Token: p.peek()}
	} else if p.peekEquals(tokens.DOLLAR) {
		expr = p.variable()
	}

	return expr
}

func (p *Parser) variable() expressions.Expr {
	p.advance() // skip dollar
	expr := &expressions.Var{
		Token: p.peek(),
	}
	p.expect(tokens.IDENT)
	return expr
}

func (p *Parser) command() expressions.Expr {
	expr := &expressions.Cmd{
		Token:     p.peek(),
		Arguments: make([]expressions.Expr, 0),
	}
	expr.Name = p.consts()
	p.advance()

	for !p.matchAny(tokens.PIPE, tokens.SEMICOLON, tokens.AND, tokens.EOF) {
		expr.Arguments = append(expr.Arguments, p.consts())
		p.advance()
	}

	return expr
}

func (p *Parser) keywords() expressions.Expr {
	expr := &expressions.Keyword{
		Token:     p.peek(),
		Arguments: make([]expressions.Expr, 0),
	}
	p.advance()
	for !p.matchAny(tokens.PIPE, tokens.SEMICOLON, tokens.AND, tokens.EOF) {
		expr.Arguments = append(expr.Arguments, p.consts())
		p.advance()
	}
	return expr
}

func (p *Parser) expect(t tokens.TokenType) {
	if p.peekEquals(t) {
		p.advance()
	} else {
		panic(fmt.Sprintf("Unexpected %s, wanted %s", tokens.LOOKUP[p.peek().Type], tokens.LOOKUP[t]))
	}
}

func (p *Parser) matchAny(t ...tokens.TokenType) bool {
	ctt := p.peek().Type
	for _, tt := range t {
		if ctt == tt {
			return true
		}
	}
	return false
}

func (p *Parser) peekEquals(t tokens.TokenType) bool {
	return p.in[p.pos].Type == t
}

func (p *Parser) peek() tokens.Token {
	return p.in[p.pos]
}

func (p *Parser) advance() {
	if p.peek().Type != tokens.EOF {
		p.pos++
	}
}

// parses a stream of tokens into an abstract syntax tree:
//
//		str      -> ".+"
//		ident    -> [A-Za-z\-]+[_0-9/]*
//		const    -> str | ident
//		varDef   -> 'set' ident const
//		varUse   -> '$' ident
//		cmd      -> const+
//	    cmdSep   -> '|' | ';' | '&&' | '>>' | '>'
//		cmd      -> cmd cmdSep cmd
//		cmd      -> cmd &
//		cmd      -> cmd+
package parser

import (
	"panzer/core/lang/expressions"
	"panzer/core/lang/tokens"
	"log"
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
	if p.matchAny(tokens.STRING, tokens.KEYWORD) {
		expr = &expressions.String{Token: p.peek()}
		p.advance()
	} else if p.peekEquals(tokens.IDENT) {
		expr = &expressions.Ident{Token: p.peek()}
		p.advance()
	} else {
		log.Printf("Unexpected %s", tokens.LOOKUP[p.peek().Type])
	}
	return expr
}

func (p *Parser) command() expressions.Expr {
	expr := &expressions.Cmd{
		Token:     p.peek(),
		Arguments: make([]expressions.Expr, 0),
	}
	expr.Name = p.consts()

	for !p.matchAny(tokens.PIPE, tokens.SEMICOLON, tokens.AND, tokens.EOF) {
		expr.Arguments = append(expr.Arguments, p.consts())
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
	}
	return expr
}

func (p *Parser) expect(t tokens.TokenType) {
	if p.peekEquals(t) {
		p.advance()
	} else {
		log.Printf("Unexpected %s, wanted %s", tokens.LOOKUP[p.peek().Type], tokens.LOOKUP[t])
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

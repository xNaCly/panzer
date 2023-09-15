package lang

import (
	"gopnzr/core/lang/expressions"
	"gopnzr/core/lang/lexer"
	"gopnzr/core/lang/parser"
	"strings"
)

// we reuse the builder because we need it for every string and identifier in
// the lexer and for debugging purposes
var b = strings.Builder{}

// we reuse the lexer to skip a lexer allocation for each input the shell gets
var lex = lexer.Lexer{
	Builder: &b,
}

// we do the same for the parser
var par = parser.Parser{}

func Compile(input string) {
	lex.NewInput(input)
	token := lex.Lex()
	par.NewInput(token)
	ast := par.Parser()
	exec(ast)
}

func exec(ast []expressions.Expr) {
	for _, e := range ast {
		e.Eval()
	}
}

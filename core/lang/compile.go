package lang

import (
	"fmt"
	"panzer/core/lang/expressions"
	"panzer/core/lang/lexer"
	"panzer/core/lang/parser"
	"panzer/core/lang/preprocessor"
	"panzer/core/lang/tokens"
	"panzer/core/shell/args"
	"strings"
)

// we reuse the builder because we need it for every string and identifier in
// the lexer and for debugging purposes
var b = strings.Builder{}

// we reuse the lexer to skip a lexer allocation for each input the shell gets
var prep = preprocessor.Preprocessor{
	Builder: &b,
}

// we reuse the lexer to skip a lexer allocation for each input the shell gets
var lex = lexer.Lexer{
	Builder: &b,
}

// we do the same for the parser
var par = parser.Parser{}

func Compile(input string, a *args.Arguments) {
	prep.NewInput(input)
	input = prep.Process(a)

	lex.NewInput(input)
	token := lex.Lex()

	if a.Debug {
		fmt.Println(tokens.Debug(token, &b))
	}

	par.NewInput(token)
	ast := par.Parser()

	if a.Debug {
		fmt.Println(expressions.Debug(ast, &b))
	}

	exec(ast)
}

func exec(ast []expressions.Expr) {
	for _, e := range ast {
		e.Eval()
	}
}

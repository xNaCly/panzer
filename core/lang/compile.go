package lang

import (
	"fmt"
	"gopnzr/core/lang/lexer"
	"gopnzr/core/lang/tokens"
)

// we reuse the lexer to skip a lexer allocation for each input the shell gets
var lex = lexer.Lexer{}

func Compile(input string) {
	lex.NewInput(input)
	token := lex.Lex()
	fmt.Print(tokens.Debug(token))
}

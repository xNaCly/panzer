package tokens

const (
	UNKNOWN TokenType = iota // used for error handling and default value of a token in the lexer
	EOF                      // END OF INPUT

	DOT       // .
	PIPE      // |
	AND       // &
	ASTERIKS  // *
	SEMICOLON // ;
	DOLLAR    // $

	STRING // ".*"
	IDENT  // [A-Za-z\-]+[_0-9]*
)

var LOOKUP = map[TokenType]string{
	UNKNOWN:   "UNKNOWN",
	EOF:       "EOF",
	DOT:       "DOT",
	PIPE:      "PIPE",
	AND:       "AND",
	ASTERIKS:  "ASTERIKS",
	SEMICOLON: "SEMICOLON",
	DOLLAR:    "DOLLAR",
	STRING:    "STRING",
	IDENT:     "IDENT",
}

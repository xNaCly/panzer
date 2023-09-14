package tokens

const (
	UNKNOWN TokenType = iota // used for error handling and default value of a token in the lexer

	// symbols

	PIPE      // |
	AND       // &
	ASTERIKS  // *
	SEMICOLON // ;
	DOLLAR    // $

	STRING // ".*"
	IDENT  // [A-Za-z\-]+[_0-9]*
)

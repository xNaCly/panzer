package tokens

var KEYWORDS = map[string]struct{}{
	"cd":     {},
	"set":    {},
	"alias":  {},
	"dstack": {},
	"help":   {},
	"exit":   {},
}

const (
	UNKNOWN   TokenType = iota // used for error handling and default value of a token in the lexer
	EOF                        // END OF INPUT
	PIPE                       // |
	AND                        // &
	SEMICOLON                  // ;
	STRING                     // ".*"
	IDENT                      // [A-Za-z\-]+[_0-9]*
	KEYWORD                    // build in stuff
)

var LOOKUP = map[TokenType]string{
	UNKNOWN:   "UNKNOWN",
	EOF:       "EOF",
	PIPE:      "PIPE",
	AND:       "AND",
	SEMICOLON: "SEMICOLON",
	STRING:    "STRING",
	IDENT:     "IDENT",
	KEYWORD:   "KEYWORD",
}

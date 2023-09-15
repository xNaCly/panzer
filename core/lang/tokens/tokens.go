package tokens

var KEYWORDS = map[string]struct{}{
	"cd":  {},
	"set": {},
	"env": {},
}

const (
	UNKNOWN   TokenType = iota // used for error handling and default value of a token in the lexer
	EOF                        // END OF INPUT
	PIPE                       // |
	AND                        // &
	ASTERIKS                   // *
	SEMICOLON                  // ;
	DOLLAR                     // $
	STRING                     // ".*"
	IDENT                      // [A-Za-z\-]+[_0-9]*
	KEYWORD                    // build in stuff
)

var LOOKUP = map[TokenType]string{
	UNKNOWN:   "UNKNOWN",
	EOF:       "EOF",
	PIPE:      "PIPE",
	AND:       "AND",
	ASTERIKS:  "ASTERIKS",
	SEMICOLON: "SEMICOLON",
	DOLLAR:    "DOLLAR",
	STRING:    "STRING",
	IDENT:     "IDENT",
	KEYWORD:   "KEYWORD",
}

package tokens

type TokenType uint

type Token struct {
	Pos  int
	Type TokenType
	Raw  string
}

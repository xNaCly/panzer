package keywords

// maps keywords to their functions
var KEYWORD_MAP = map[string]func(args ...string){
	"set":    Set,
	"cd":     Cd,
	"dstack": Dstack,
	"alias":  Alias,
	"help":   Help,
}

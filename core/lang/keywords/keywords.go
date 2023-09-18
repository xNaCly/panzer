package keywords

// TODO: add source <file>
// TODO: add cd <path>
var KEYWORD_MAP = map[string]func(args ...string){
	"env": Env,
	"set": Set,
	"cd":  Cd,
}

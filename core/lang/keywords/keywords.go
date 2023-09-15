package keywords

var KEYWORD_MAP = map[string]func(args ...string){
	"env": Env,
	"set": Set,
}

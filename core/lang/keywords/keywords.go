package keywords

import (
	"os"
)

var KEYWORD_MAP = map[string]func(args ...string){
	"env": Env,
	"set": Set,
	"exit": func(_ ...string) {
		os.Exit(0)
	},
}

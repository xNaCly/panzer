package keywords

import (
	"panzer/core/shell/env"
	"log"
	"strings"
)

// sets an env variable
func Set(args ...string) {
	if len(args) < 2 {
		log.Panicf("set: expected two arguments: one for variable name and one for the variable value: 'set NAME value', got %d", len(args))
	}
	env.SetEnv(args[0], strings.Join(args[1:], ":"))
}

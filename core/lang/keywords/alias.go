package keywords

import (
	"fmt"
	"panzer/core/state"
)

// Creates a new alias
func Alias(args ...string) {
	al := len(args)
	if al == 0 {
		for k, v := range state.ALIASES {
			fmt.Print("alias ", k, " \"", v, "\"\n")
		}
		return
	}
	if al != 2 {
		panic("alias: expected 2 arguments, one for alias name, one for alias content")
	}

	state.ALIASES[args[0]] = args[1]
}

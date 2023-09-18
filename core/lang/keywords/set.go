package keywords

import (
	"gopnzr/core/shell/env"
	"strings"
)

// sets an env variable
func Set(args ...string) {
	if len(args) < 2 {
		panic("set requires one argument for variable name and one for the variable value")
	}
	env.SetEnv(args[0], strings.Join(args[1:], ":"))
}

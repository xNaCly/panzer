package keywords

import (
	"os"
	"strings"
)

func Set(args ...string) {
	if len(args) < 2 {
		panic("set requires one argument for variable name and one for the variable value")
	}
	os.Setenv(args[0], strings.Join(args[1:], ":"))
}

package keywords

import (
	"log"
	"os"
	"strings"
)

// sets an env variable
func Set(args ...string) {
	if len(args) < 2 {
		log.Panicf("set: expected two arguments: one for variable name and one for the variable value: 'set NAME value', got %d", len(args))
	}
	os.Setenv(args[0], strings.Join(args[1:], ":"))
}

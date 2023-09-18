package keywords

import (
	"fmt"
	"os"
)

// Iterates over env variables and prints them
func Env(args ...string) {
	if len(args) > 0 {
		panic("env: too many arguments")
	}
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

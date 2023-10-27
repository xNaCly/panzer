package keywords

import (
	"fmt"
	"panzer/core/state"
)

// prints elements in dir stack
func Dstack(args ...string) {
	for i, d := range state.DIR_STACK.Stack {
		fmt.Printf("[%d]: %q\n", i, d)
	}
}

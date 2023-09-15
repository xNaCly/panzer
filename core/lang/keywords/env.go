package keywords

import (
	"fmt"
	"os"
)

func Env(args ...string) {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

package core

import (
	"bufio"
	"fmt"
	"gopnzr/core/shell/prompt"
	"os"
)

func Loop() {
	err := prompt.PreComputePlaceholders()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while computing prompt placeholders: ", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt.ComputePrompt())
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(input)
	}
}

package keywords

// TODO: history
// TODO: last dir via -

import (
	"fmt"
	"gopnzr/core/shell/prompt"
	"os"
)

func Cd(args ...string) {
	al := len(args)
	if al > 1 {
		panic("Too many arguments for cd")
	} else if al == 0 {
		panic("no path given for cd")
	}

	dir := args[0]
	resolvedDir := ""

	switch dir {
	case "~":
		h, err := os.UserHomeDir()
		if err != nil {
			h = "/"
		}
		resolvedDir = h
	case ".":
		return
	default:
		resolvedDir = dir
	}

	data, err := os.Stat(resolvedDir)
	if err != nil {
		panic(fmt.Sprintf("cd: the directory %q does not exist", resolvedDir))
	}

	if !data.IsDir() {
		panic(fmt.Sprintf("cd: %q is not a directory", resolvedDir))
	}

	err = os.Chdir(resolvedDir)
	if err != nil {
		panic(err)
	}

	prompt.UpdatePrompt()
}

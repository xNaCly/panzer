package keywords

// TODO: history
// TODO: last dir via -

import (
	"gopnzr/core/shell/prompt"
	"os"
)

func Cd(args ...string) {
	if len(args) > 1 {
		panic("Too many arguments for cd")
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

	_, err := os.Stat(resolvedDir)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(resolvedDir)
	if err != nil {
		panic(err)
	}

	prompt.UpdatePrompt()
}

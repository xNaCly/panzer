package keywords

// TODO: history
// TODO: last dir via -

import (
	"fmt"
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/prompt"
	"gopnzr/core/shell/system"
	"os"
	"path/filepath"
)

// resolves the dir from the given argument, checks if the target is a dir,
// changes process working directory to the target
func Cd(args ...string) {
	al := len(args)
	dir := ""
	if al > 1 {
		// cd can not accept more than 1 argument
		panic("Too many arguments for cd")
	} else if al == 0 {
		// no arguments, means we want to go $HOME
		dir = "~"
	} else {
		dir = args[0]
	}

	// we skip cd if we are already at the desired path
	if dir == system.Getwd() {
		return
	}

	switch dir {
	case "~":
		h, err := os.UserHomeDir()
		if err != nil {
			h = "/"
		}
		dir = h
	case ".":
		// we skip changing dir due to us already being at .
		return
	}

	data, err := os.Stat(dir)
	if err != nil {
		panic(fmt.Sprintf("cd: the directory %q does not exist", dir))
	}

	if !data.IsDir() {
		panic(fmt.Sprintf("cd: %q is not a directory", dir))
	}

	cleansedPath, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(cleansedPath)
	if err != nil {
		panic(err)
	}

	env.SetEnv("PWD", cleansedPath)

	// we changed the directory, we need to update the prompt
	prompt.UpdatePrompt()
}

package keywords

import (
	"panzer/core/shell/env"
	"panzer/core/shell/prompt"
	"panzer/core/shell/system"
	"panzer/core/state"
	"log"
	"os"
	"path/filepath"
)

// resolves the dir from the given argument, checks if the target is a dir,
// changes process working directory to the target
func Cd(args ...string) {
	dir := ""

	if len(args) > 1 {
		panic("cd: too many arguments for")
	} else if len(args) == 0 {
		// no arguments, means we want to go $HOME
		dir = "~"
	} else {
		dir = args[0]
	}

	addToHist := true

	switch dir {
	case "-":
		dir = state.LAST_DIR
	case "^":
		dir = state.DIR_STACK.Pop()
		addToHist = false
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

	// we skip cd if we are already at the desired path
	if dir == system.Getwd() {
		return
	}

	data, err := os.Stat(dir)
	if err != nil {
		log.Panicf("cd: the directory %q does not exist", dir)
	}

	if !data.IsDir() {
		log.Panicf("cd: %q is not a directory", dir)
	}

	cleansedPath, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}

	prevDir := system.Getwd()

	err = os.Chdir(cleansedPath)
	if err != nil {
		panic(err)
	}

	state.LAST_DIR = prevDir
	if addToHist {
		state.DIR_STACK.Add(prevDir)
	}

	env.SetEnv("PWD", cleansedPath)

	// we changed the directory, we need to update the prompt
	prompt.UpdatePrompt()
}

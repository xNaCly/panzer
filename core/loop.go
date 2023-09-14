package core

import (
	"errors"
	"fmt"
	"github.com/chzyer/readline"
	"gopnzr/core/lang"
	a "gopnzr/core/shell/args"
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/prompt"
	"gopnzr/core/shell/system"
	"io"
	"os"
)

// main entry point for the shell
// performs the following actions:
//
// 1. computes the value for $PWD
//
// 2. registers notifier for syscalls, such as SIGINT, SIGTERM, etc
//
// 3. computation of at startup known prompt placeholders
//
// 4. starting a go routine for signal handling
//
// 5. main loop
//   - computing the prompt
//   - waiting for input
//   - waiting for input
//   - exiting on EOF (Ctrl+D)
func Shell() {
	args := a.Get()
	env.SetEnv("PWD", system.Getwd())

	err := prompt.PreComputePlaceholders()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while computing prompt placeholders: ", err)
	}

	if args.Command != "" {
		lang.Compile(args.Command)
		return
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt: prompt.ComputePrompt(),
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		input, err := rl.Readline()

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("bye bye, see you :^)")
				os.Exit(0)
			} else if errors.Is(err, readline.ErrInterrupt) {
				continue
			}

			fmt.Fprintln(os.Stderr, err)
		}

		if input == "" {
			continue
		}
		lang.Compile(input)
	}
}

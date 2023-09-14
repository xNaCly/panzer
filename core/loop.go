package core

import (
	"bufio"
	"errors"
	"fmt"
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/prompt"
	"gopnzr/core/shell/system"
	"io"
	"os"
	"os/signal"
	"syscall"
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
	env.SetEnv("PWD", system.Getwd())

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	err := prompt.PreComputePlaceholders()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error while computing prompt placeholders: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			s := <-cancelChan
			switch s {
			case syscall.SIGINT:
				fmt.Print("\n", prompt.ComputePrompt())
			case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Printf("\nerr: [%s]\n", s.String())
				os.Exit(1)
			}
		}
	}()

	for {
		fmt.Print(prompt.ComputePrompt())
		input, err := reader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("\nexit")
				os.Exit(0)
			}

			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Print(input)
	}
}

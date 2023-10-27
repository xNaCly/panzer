package keywords

import (
	"gopnzr/core/state"
	"log"
	"os"
	"path/filepath"
)

func Help(args ...string) {
	var requestedPage string
	if len(args) == 0 {
		requestedPage = "main"
	} else if len(args) > 1 {
		log.Panicf("help: expected 1 argument, got %d", len(args))
	} else {
		requestedPage = args[0]
	}
	cleansedPath := filepath.Clean(filepath.Join("help", requestedPage+".txt"))
	main, err := state.HELP_PAGES.ReadFile(cleansedPath)
	if err != nil {
		log.Panicf("help: failed to access help page: %q - not found", requestedPage)
	}
	os.Stdout.Write(main)
}

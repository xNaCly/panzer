// command line arguments handling, registers and returns arguments parsed by flag.Parse
package args

import "flag"

// wrapper for structuring arguments
type Arguments struct {
	Command string // command to execute, can be specified via "-c"
	Debug   bool   // whether to enable debug logging
}

// registers cli flags, invokes flag for parsing, structures and returns the
// parsed flags into args.Arguments
func Get() Arguments {
	a := Arguments{}

	flag.StringVar(&a.Command, "c", "", "command to execute, executes, exits")
	flag.BoolVar(&a.Debug, "d", false, "whether to enable debug logging")

	flag.Parse()
	return a
}

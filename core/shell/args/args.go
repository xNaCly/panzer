// command line arguments handling, registers and returns arguments parsed by flag.Parse
package args

import "flag"

// wrapper for structuring arguments
type Arguments struct {
	// command to execute, can be specified via "-c"
	Command string
}

// registers cli flags, invokes flag for parsing, structures and returns the
// parsed flags into args.Arguments
func Get() Arguments {
	a := Arguments{}

	flag.StringVar(&a.Command, "c", "", "command to execute, executes, exits")

	flag.Parse()
	return a
}

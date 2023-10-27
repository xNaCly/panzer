package complete

import (
	"panzer/core/shell/system"

	"github.com/chzyer/readline"
)

// TODO: this works like ass, completion only works for the first word :(
func BuildCompleter() readline.AutoCompleter {
	return readline.NewPrefixCompleter(
		readline.PcItemDynamic(system.GetwdFiles(".")),
	)
}

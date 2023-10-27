package state

import (
	"embed"
)

const VERSION = "0.0.1"
const VERSION_SUFFIX = "dev"

//go:embed help/*.txt
var HELP_PAGES embed.FS

var ALIASES = make(map[string]string, 0)
var DIR_STACK = DirStack{Stack: make([]string, 0), max: 9}
var LAST_DIR string

type DirStack struct {
	Stack []string
	max   int
}

func (d *DirStack) Add(s string) {
	if d.max == len(d.Stack) {
		d.Stack = d.Stack[1:]
	} else {
		d.Stack = append(d.Stack, s)
	}
}

func (d *DirStack) Pop() string {
	l := d.Stack[len(d.Stack)-1]
	d.Stack = d.Stack[:1]
	return l
}

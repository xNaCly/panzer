// prompt generation, prompt placeholder computation
package prompt

import (
	"os"
	"os/user"
	"panzer/core/git"
	"panzer/core/shell/system"
	"strconv"
	"strings"
	"time"
)

// home directory
var HOME = "/"

// default prompt: USERNAME@HOSTNAME WORKINGDIRECTORY >
const DEFAULT_PROMPT = `\7\u\0@\6\h\0 \8\w\0 \2>\0 `

const remove_last_character = "\x08"

// contains all possible placeholders a prompt could contain
var prompt_placeholders = map[rune]string{
	'u': "", // username
	'h': "", // hostname
	'w': "", // pwd
	'd': "", // name of directory
	'D': "", // current date
	't': "", // current time (24h)
	'T': "", // current time (12h)
	'U': "", // current time (unixepoch)
	'b': "", // git branch
	'S': "", // git status (M for modified)
	's': "panzer",
	'0': "\033[0m",
	'1': "\033[31m",
	'2': "\033[32m",
	'3': "\033[33m",
	'4': "\033[91m",
	'5': "\033[92m",
	'6': "\033[93m",
	'7': "\033[94m",
	'8': "\033[95m",
	'9': "\033[96m",
}

// computes placeholder values that are known at startup, this decreases load
// on the main loop prompt computation
func PreComputePlaceholders() (e error) {
	HOME, _ = os.UserHomeDir()
	u, e := user.Current()

	prompt_placeholders['u'] = u.Username

	// date only needs to be computed at startup, who keeps their shell active
	// more than a day?
	prompt_placeholders['D'] = time.Now().Format(time.DateOnly)

	h, e := os.Hostname()
	prompt_placeholders['h'] = h

	return
}

// updates all values in the placeholders map that are either unknown at
// startup or require recomputation on a prompt redraw, such as cwd and time
func UpdatePrompt() {
	t := time.Now()
	prompt_placeholders['t'] = t.Format(time.TimeOnly)
	prompt_placeholders['T'] = t.Format("03:04:05PM")
	prompt_placeholders['w'] = system.Getwd()
	prompt_placeholders['U'] = strconv.FormatInt(t.UnixMilli(), 10)
	// FIXME: bug here
	// prompt_placeholders['b'] = git.Branch()
	// if git.Status() {
	// 	prompt_placeholders['S'] = "M"
	// }
	dir := system.Getdir()
	if dir == prompt_placeholders['u'] {
		dir = "~"
	}
	prompt_placeholders['d'] = dir
}

// checks if custom prompt is set, returns either that prompt or the default
// prompt with placeholders replaced
func ComputePrompt() string {
	UpdatePrompt()
	prompt := DEFAULT_PROMPT
	if val, ok := os.LookupEnv("PROMPT"); ok {
		prompt = val
	}
	return replacePlaceholders(prompt)
}

// formats the working directory according to the configuration
func formatWd(path string) string {
	if path == HOME {
		return "~"
	}

	b := strings.Builder{}

	for _, c := range path {
		str := b.String()
		if str == HOME {
			b.Reset()
			b.WriteRune('~')
		}
		b.WriteRune(c)
	}

	return b.String()
}

// replaces placeholders in the given prompt with the values in
// 'prompt_placeholders', works by detecting slashes and writing the
// 'prompt_placeholders' value of the placeholder into a string builder, which
// gets returned, this should be incredibly faster than calling strings.Replace
// on each placeholder
func replacePlaceholders(prompt string) string {
	prompt_placeholders['w'] = formatWd(prompt_placeholders['w'])
	b := strings.Builder{}
	placeHolderMode := false
	for _, c := range prompt {
		if c == '\\' && placeHolderMode {
			b.WriteRune('\\')
		} else if c == '\\' {
			placeHolderMode = true
		} else if placeHolderMode {
			if t, ok := prompt_placeholders[c]; ok {
				if t != "" {
					b.WriteString(t)
				} else {
					// INFO: this removes the uncessecary whitespace before
					// empty prompts placeholders, such as git
					b.WriteString(remove_last_character)
				}
			}
			// TODO: log not found escapes
			placeHolderMode = false
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

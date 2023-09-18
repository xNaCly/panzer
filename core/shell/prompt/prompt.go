package prompt

import (
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/system"
	"os"
	"os/user"
	"strings"
	"time"
)

var HOME = "/"

const DEFAULT_PROMPT = `\7\u\0@\6\h\0 \8\w\0 \2>\0 `

// contains all possible placeholders a prompt could contain
var prompt_placeholders = map[rune]string{
	'u': "",
	'h': "",
	'w': "",
	'd': "",
	'D': "",
	't': "",
	'T': "",
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
	// TODO: shell name (S)
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
	// TODO: support git-branch (b) (either nothing or the branch name, see 'git branch')
	// TODO: support git-status (s) (either nothing or M for modified, see 'git status --short')
	t := time.Now()
	prompt_placeholders['t'] = t.Format(time.TimeOnly)
	prompt_placeholders['T'] = t.Format("03:04:05PM")
	prompt_placeholders['w'] = system.Getwd()
	prompt_placeholders['d'] = system.Getdir()
}

// checks if custom prompt is set, returns either that prompt or the default
// prompt with placeholders replaced
func ComputePrompt() string {
	UpdatePrompt()
	prompt := DEFAULT_PROMPT
	if val, ok := env.GetEnv("PROMPT"); ok {
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
				b.WriteString(t)
			}
			placeHolderMode = false
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

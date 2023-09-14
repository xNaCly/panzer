package prompt

import (
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/system"
	"os"
	"os/user"
	"strings"
)

const DEFAULT_PROMPT = `\u@\h \d :: `

// contains all possible placeholders a prompt could contain
var prompt_placeholders = map[rune]string{
	'u': "",
	'h': "",
	'w': "",
	'd': "",
}

// computes placeholder values that are known at startup, this decreases load
// on the main loop prompt computation
func PreComputePlaceholders() (e error) {
	u, e := user.Current()

	prompt_placeholders['u'] = u.Username
	pwd, _ := env.GetEnv("PWD")
	prompt_placeholders['w'] = pwd
	prompt_placeholders['d'] = system.Getdir()

	h, e := os.Hostname()
	prompt_placeholders['h'] = h
	return
}

// checks if custom prompt is set, returns either that prompt or the default
// prompt with placeholders replaced
func ComputePrompt() string {
	prompt := DEFAULT_PROMPT
	if val, ok := env.GetEnv("GPNZR_PROMPT"); ok {
		prompt = val
	}
	return replacePlaceholders(prompt)
}

// formats the working directory according to the configuration
func formatWd(path string) string {
	if !env.GetEnvBool("GPNZR_PROMPT_SHORT_PWD") {
		return path
	}
	b := strings.Builder{}
	var lc rune
	for _, c := range path {
		if lc == '/' {
			b.WriteRune(c)
			b.WriteRune('/')
		}
		lc = c
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
		if c == '\\' {
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

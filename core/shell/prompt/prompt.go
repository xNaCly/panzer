package prompt

import (
	"gopnzr/core/shell/env"
	"gopnzr/core/shell/system"
	"os"
	"os/user"
	"strings"
)

const DEFAULT_PROMPT = `\u@\h \d :: `

var prompt_placeholders = map[rune]string{
	'u': "",
	'h': "",
	'w': "",
	'd': "",
}

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

func ComputePrompt() string {
	prompt := DEFAULT_PROMPT
	if val, ok := env.GetEnv("GPNZR_PROMPT"); ok {
		prompt = val
	}
	return replacePlaceholders(prompt)
}

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

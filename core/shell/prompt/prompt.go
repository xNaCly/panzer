package prompt

import (
	"gopnzr/core/shell/env"
	"os"
	"os/user"
	"strings"
)

const DEFAULT_PROMPT = `\u@\h \w :: `

var HOME = "/"

var prompt_placeholders = map[rune]string{
	'u': "",
	'h': "",
	'w': "",
}

func PreComputePlaceholders() (e error) {
	u, e := user.Current()
	HOME = u.HomeDir
	prompt_placeholders['u'] = u.Username
	h, e := os.Hostname()
	prompt_placeholders['h'] = h
	prompt_placeholders['w'] = u.HomeDir
	return
}

func ComputePrompt() string {
	prompt := DEFAULT_PROMPT
	if val, ok := env.GetEnv("GPNZR_PROMPT"); ok {
		prompt = val
	}
	return replacePlaceholders(prompt)
}

func replacePlaceholders(prompt string) string {
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

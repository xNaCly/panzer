package git

import (
	"bytes"
	"os/exec"
)

// looks git branch up via git cli
func Branch() string {
	if !isRepo() {
		return ""
	}
	out, err := branch()
	if err != nil {
		return ""
	}
	return out
}

// looks git branch up via git cli
func branch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	out = bytes.TrimSpace(out)
	return string(out), nil
}

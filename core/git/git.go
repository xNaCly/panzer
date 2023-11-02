// provides wrappers for displaying the current git branch, wether or not
// modifications were made, etc - currently depends on the 'git' executable being in path
package git

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
)

// checks if git exists on the operating system and in the path
func hasGit() bool {
	_, found := exec.LookPath("git")
	return found != nil
}

// checks if cwd is a git repo
func isRepo() bool {
	if !hasGit() {
		return false
	}
	if val, err := os.Stat(".git"); errors.Is(err, fs.ErrNotExist) || !val.IsDir() {
		return false
	}
	return true
}

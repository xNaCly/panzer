package git

import "os/exec"

// returns true if changes in the current repo were made
func Status() bool {
	if !isRepo() {
		return false
	}
	cmd := exec.Command("git", "status", "--short")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return len(out) > 0
}

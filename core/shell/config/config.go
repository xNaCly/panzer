// reads and interprets config files from several locations in the following
// order
//
//   - $HOME/.panzerc
//   - $XDG_CONFIG_HOME/.panzerc
//   - /etc/.panzerc
//
// therefore the first valid and existing configuration at any of the paths
// above is used
package config

import (
	"panzer/core/lang"
	"panzer/core/shell/args"
	"log"
	"os"
)

var paths = []string{
	"$HOME",
	"$XDG_CONFIG_HOME",
	"/etc",
}

// checks which of the config location contains a file, returns first found,
// returns false if none found
func lookup() (string, bool) {
	confHome, err := os.UserConfigDir()
	skipConfHome := err != nil
	for _, p := range paths {
		switch p {
		case "$HOME":
			home, err := os.UserHomeDir()
			if err != nil {
				p = "/"
			} else {
				p = home
			}
		case "$XDG_CONFIG_HOME":
			if skipConfHome {
				continue
			}
			p = confHome
		}
		p += "/.panzerc"
		stat, err := os.Stat(p)
		if err != nil {
			continue
		}

		if !stat.IsDir() {
			return p, true
		}
	}
	return "", false
}

// attempts to load configuration from the first found path, attempts to
// interpret the configuration
func Load(a *args.Arguments) {
	path, found := lookup()
	if !found {
		return
	}

	out, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read configuration")
	}

	defer func() {
		if err := recover(); err != nil {
			log.Print("error in configuration: ", err)
		}
	}()

	lang.Compile(string(out), a)
}

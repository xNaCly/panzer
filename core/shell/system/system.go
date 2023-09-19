// abstractions for operating system interactions
package system

import (
	"os"
	"path"
)

// returns the current working directory or /
func Getwd() (wd string) {
	wd, err := os.Getwd()

	if err != nil {
		wd = "/"
	}

	return
}

// returns only the name of the current directory, replaces ~ with home
func Getdir() (dir string) {
	wd := path.Base(Getwd())
	if wd == "." {
		return
	}
	return wd
}

func GetwdFiles(wd string) func(string) []string {
	return func(l string) []string {
		names := make([]string, 0)
		files, _ := os.ReadDir(wd)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

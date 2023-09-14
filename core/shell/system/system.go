package system

import (
	"gopnzr/core/shell/env"
	"os"
	"path"
)

func Getwd() (wd string) {
	wd, err := os.Getwd()

	if err != nil {
		wd = "/"
	}

	return
}

func Getdir() (dir string) {
	pwd, ok := env.GetEnv("PWD")
	dir = "/"
	if !ok {
		return
	}
	wd := path.Base(pwd)
	if wd == "." {
		return
	}
	return wd
}

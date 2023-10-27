// preforms expansions, such as:
//
// - *.c -> all files in current dir with .c postfix
//
// - libary.? -> matches all files with the libary suffix and a single
// character after the ., such as library.h, library.c
//
// - $(cat test) -> execute command between $()
package expansions

import (
	"panzer/core/shell/system"
	"os"
	"path"
)

// accepts a pattern, matches that pattern agains all files in the current
// directory and returns all matching filenames
func MatchFiles(pattern string) []string {
	res := make([]string, 0)

	wd := system.Getwd()

	files, err := os.ReadDir(wd)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		matched, err := path.Match(pattern, file.Name())
		if matched && err == nil {
			res = append(res, file.Name())
		}
	}

	return res
}

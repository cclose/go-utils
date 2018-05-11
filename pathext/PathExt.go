package pathext

import (
	"fmt"
	"os"
	"strconv"
)

// Splits the specified filepath string into an array of strings based on element, in order from root to leaf
// If the path begins at the root "/", this is concatenated onto the first directory.
// Trailing / is not preserved
//
// EX:
//   path = /path/to/a/place/foo.txt
//   returns = ["/path", "to", "a", "place", "foo.txt"]
//
//   path = path/to/a/place/
//   returns = ["path", "to", "a", "place"]
//
func SplitAll(path string) (dirs []string, err error) {
	dirStart := 0
	newDir := false
	for i := 0; i < len(path); i++ {
		if os.IsPathSeparator(path[i]) {
			if i > 0 { //if we're at the system root, preserve that as part of the first dir
				newDir = true
				fmt.Println("append " + strconv.Itoa(dirStart) + ":" + strconv.Itoa(i))
				dirs = append(dirs, path[dirStart:i])
			}
		} else if newDir {
			newDir = false
			dirStart = i
		}
	}
	//append last portion
	if !newDir {
		dirs = append(dirs, path[dirStart:])
	}

	return
}

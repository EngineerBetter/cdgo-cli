package main

import "errors"
import "os"
import "path/filepath"

type DirectoryFinder interface {
	Find(directory string, in string) (string, error)
}

type RecursiveFinder struct{}

func (*RecursiveFinder) Find(needle string, haystack string) (result string, errOut error) {
	fi, err := os.Lstat(haystack)

	if err != nil {
		return result, errors.New(haystack + " did not exist")
	} else if !fi.IsDir() {
		return result, errors.New(haystack + " is not a directory")
	}

	filepath.Walk(haystack, func(path string, fi os.FileInfo, errIn error) (errOut error) {
		if fi.Name() == needle {
			result = path
			return filepath.SkipDir
		}

		return
	})

	if result == "" {
		errOut = errors.New(needle + "not found in " + haystack)
	}

	return
}

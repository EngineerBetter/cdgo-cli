package main

import "errors"
import "os"

type DirectoryFinder interface {
	Find(directory string, in string) (string, error)
}

type RecursiveFinder struct{}

func (*RecursiveFinder) Find(needle string, haystack string) (result string, err error) {
	info, err := os.Stat(haystack)

	if err != nil {
		return result, err
	} else if !info.IsDir() {
		return result, errors.New("Path was not a directory")
	}

	return "foo", nil
}
